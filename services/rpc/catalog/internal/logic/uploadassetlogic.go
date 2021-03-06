package logic

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/events/eventkey"
	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/storage/asset"
	"github.com/k8scommerce/k8scommerce/internal/utils/humanizer"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UploadAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	// mu       sync.Mutex
}

func NewUploadAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadAssetLogic {
	return &UploadAssetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

///
/// The problem seems to be the file losing it's handle
///
///
///
// grpc.MaxCallSendMsgSize(maxFileSize)
func (l *UploadAssetLogic) UploadAsset(stream catalog.CatalogClient_UploadAssetServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive file info")
	}

	file, err := asset.MustNewFile(req.GetAsset().Name, l.svcCtx.Config.UploadConfig)
	if err != nil {
		return status.Errorf(codes.Internal, "file creation error: %s", err.Error())
	}

	kind, err := l.getAssetKind(req.GetAsset().GetKind())
	if err != nil {
		return err
	}
	file.Kind = kind

	storeId := req.GetAsset().GetStoreId()
	productId := req.GetAsset().GetProductId()
	variantId := req.GetAsset().GetVariantId()

	// let's check for an existing file of the same name/variant id
	if exists, err := l.svcCtx.Repo.Asset().AssetExists(storeId, req.GetAsset().Name); err != nil || exists {
		if err != nil {
			return status.Errorf(codes.Internal, "error checking for existing asset: %s", err.Error())
		}
		return status.Errorf(codes.Internal, "asset already exists for this store: %s: ", req.GetAsset().Name)
	}

	maxUploadSize, err := l.getMaxUploadFilesize(req.GetAsset().GetKind())
	if err != nil {
		return err
	}

	// start db insert transaction
	tx, err := l.svcCtx.Repo.Begin()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to start a transaction: %s", err.Error())
	}

	var uploadSize int64 = 0
	partNumber := 1
	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			if err := file.Close(); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			modelAsset := &models.Asset{
				StoreID:     storeId,
				ProductID:   productId,
				VariantID:   variantId,
				Kind:        file.Kind.Int(),
				Name:        file.BaseName,
				DisplayName: sql.NullString{Valid: true, String: file.BaseName},
				ContentType: file.ContentType,
				URL:         file.URL,
				SortOrder:   sql.NullInt64{Valid: true, Int64: 100},
				Sizes:       []byte("{}"),
			}

			// save the asset to the database
			if err := l.svcCtx.Repo.Asset().CreateTx(modelAsset, tx); err != nil {
				l.svcCtx.Repo.Rollback(tx)
				return status.Errorf(codes.Internal, "saving asset to database failed: %s", err.Error())
			}

			// commit the transaction
			if err := l.svcCtx.Repo.Commit(tx); err != nil {
				return status.Errorf(codes.Internal, "saving asset db commit failed: %s", err.Error())
			}

			if bytes, err := eventkey.CatalogImageUploaded.Marshal(modelAsset); err != nil {
				logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CatalogImageUploaded, err.Error())
			} else {
				// publish event
				err = l.svcCtx.EventManager.Publish(eventkey.CatalogImageUploaded.AsKey(), bytes)
				if err != nil {
					logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CatalogImageUploaded, err.Error())
				}
			}

			// return it
			return stream.SendAndClose(&catalog.Asset{
				Id:          modelAsset.ID,
				StoreId:     modelAsset.StoreID,
				ProductId:   modelAsset.ProductID,
				VariantId:   modelAsset.VariantID,
				Kind:        catalog.AssetKind(file.Kind.Int32()), // catelog kind is zero based
				DisplayName: modelAsset.DisplayName.String,
				Name:        modelAsset.Name,
				Url:         modelAsset.URL,
				ContentType: modelAsset.ContentType,
				SortOrder:   modelAsset.SortOrder.Int64,
			})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		chunk := req.GetChunk()
		size := len(chunk)

		if partNumber == 1 {
			contentType, err := l.getContentType(chunk)
			if err != nil {
				return err
			}

			if err := file.Open(contentType); err != nil {
				return err
			}
		}

		uploadSize += int64(size)
		if uploadSize > humanizer.HumanToSize(maxUploadSize) {
			return status.Errorf(codes.InvalidArgument, "file is too large: %d > %s. kind: %s", uploadSize, maxUploadSize, file.Kind.String())
		}

		// we stream the file to all transports (filesystem, aws, azure, gcp, etc. )
		// this way we can handle large files if needed
		if err := file.Write(chunk, partNumber); err != nil {
			return fmt.Errorf("chunk write error: %s", err.Error())
		}
		partNumber++
	}

	// the image is uploaded

}

func (l *UploadAssetLogic) getContentType(chunk []byte) (string, error) {
	if len(chunk) > 1 {
		fileType := http.DetectContentType(chunk)
		return fileType, nil

	}
	return "", status.Error(codes.Internal, "mime type cannot be detected. file buffer length is zero")
}

func (l *UploadAssetLogic) getAssetKind(assetKind catalog.AssetKind) (asset.Kind, error) {
	var kind asset.Kind
	var isSet = false
	switch assetKind {
	case catalog.AssetKind_unknown:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_unknown.String())
	case catalog.AssetKind_image:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_image.String())
	case catalog.AssetKind_document:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_document.String())
	case catalog.AssetKind_audio:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_audio.String())
	case catalog.AssetKind_video:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_video.String())
	case catalog.AssetKind_archive:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_archive.String())
	}

	if !isSet {
		return asset.Kind(catalog.AssetKind_unknown.String()), status.Error(codes.FailedPrecondition, "unknown asset kind. must be one of image, document, audio or video")
	}
	return kind, nil
}

func (l *UploadAssetLogic) getMaxUploadFilesize(assetKind catalog.AssetKind) (string, error) {
	var maxUploadSize string

	switch assetKind {
	case catalog.AssetKind_image:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Images.MaxUploadSize
	case catalog.AssetKind_document:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Documents.MaxUploadSize
	case catalog.AssetKind_audio:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Audio.MaxUploadSize
	case catalog.AssetKind_video:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Video.MaxUploadSize
	case catalog.AssetKind_archive:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Archive.MaxUploadSize
	}

	if maxUploadSize == "" {
		return "", status.Error(codes.FailedPrecondition, "upload allowed type is unknown")
	}
	return maxUploadSize, nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
