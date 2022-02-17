package logic

import (
	"context"
	"io"
	"k8scommerce/internal/storage/asset"
	"k8scommerce/internal/utils/humanizer"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"
	"net/http"

	"github.com/localrivet/galaxycache"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UploadAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	universe *galaxycache.Universe
	// mu       sync.Mutex
}

func NewUploadAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *UploadAssetLogic {
	return &UploadAssetLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		universe: universe,
	}
}

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

	maxUploadSize, err := l.getMaxUploadFilesize(req.GetAsset().GetKind())
	if err != nil {
		return err
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

			return stream.SendAndClose(&catalog.UploadAssetResponse{})
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

			file.Open(contentType)
		}

		uploadSize += int64(size)
		if uploadSize > humanizer.HumanToSize(maxUploadSize) {
			return status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", uploadSize, maxUploadSize)
		}

		// we stream the file to all transports (filesystem, aws, azure, gcp, etc. )
		// this way we can handle large files if needed
		if err := file.Write(chunk, partNumber); err != nil {
			return status.Error(codes.Internal, err.Error())
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
	case catalog.AssetKind_Image:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_Image.Number())
	case catalog.AssetKind_Document:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_Document.Number())
	case catalog.AssetKind_Audio:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_Audio.Number())
	case catalog.AssetKind_Video:
		isSet = true
		kind = asset.Kind(catalog.AssetKind_Video.Number())
	}

	if !isSet {
		return -1, status.Error(codes.FailedPrecondition, "unknown asset kind. must be one of image, document, audio or video")
	}
	return kind, nil
}

func (l *UploadAssetLogic) getMaxUploadFilesize(assetKind catalog.AssetKind) (string, error) {
	var maxUploadSize string

	switch assetKind {
	case catalog.AssetKind_Image:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Images.MaxUploadSize
	case catalog.AssetKind_Document:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Documents.MaxUploadSize
	case catalog.AssetKind_Audio:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Audio.MaxUploadSize
	case catalog.AssetKind_Video:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Video.MaxUploadSize
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
