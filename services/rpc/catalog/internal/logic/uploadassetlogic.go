package logic

import (
	"context"
	"io"
	"k8scommerce/internal/storage"
	"k8scommerce/internal/utils/humanizer"
	"k8scommerce/services/rpc/catalog/internal/svc"
	"k8scommerce/services/rpc/catalog/pb/catalog"

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

	file := storage.NewFile()
	file.Name = req.GetAsset().Name
	file.Mime = req.GetAsset().Mime

	kind, err := l.getAssetKind(req.GetKind())
	if err != nil {
		return err
	}
	file.AssetKind = kind

	maxUploadSize, err := l.getMaxUploadFilesize(req.GetKind())
	if err != nil {
		return err
	}

	var uploadSize int64 = 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		req, err := stream.Recv()
		if err == io.EOF {
			if err := storage.SaveFile(file, l.svcCtx.Config.UploadConfig); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&catalog.UploadAssetResponse{})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		chunk := req.GetChunk()
		size := len(chunk)

		uploadSize += int64(size)
		if uploadSize > humanizer.HumanToSize(maxUploadSize) {
			return status.Errorf(codes.InvalidArgument, "file is too large: %d > %d", uploadSize, maxUploadSize)
		}

		if err := file.Write(chunk); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	// the image is uploaded

}

func (l *UploadAssetLogic) getAssetKind(assetKind catalog.UploadAssetRequest_AssetKind) (storage.AssetKind, error) {

	var kind storage.AssetKind
	var isSet = false
	switch assetKind {
	case catalog.UploadAssetRequest_Image:
		isSet = true
		kind = storage.AssetKind(catalog.UploadAssetRequest_Image.Number())
	case catalog.UploadAssetRequest_Document:
		isSet = true
		kind = storage.AssetKind(catalog.UploadAssetRequest_Document.Number())
	case catalog.UploadAssetRequest_Audio:
		isSet = true
		kind = storage.AssetKind(catalog.UploadAssetRequest_Audio.Number())
	case catalog.UploadAssetRequest_Video:
		isSet = true
		kind = storage.AssetKind(catalog.UploadAssetRequest_Video.Number())
	}

	if !isSet {
		return -1, status.Error(codes.FailedPrecondition, "unknown asset kind. must be one of image, document, audio or video")
	}
	return kind, nil
}

func (l *UploadAssetLogic) getMaxUploadFilesize(assetKind catalog.UploadAssetRequest_AssetKind) (string, error) {
	var maxUploadSize string

	switch assetKind {
	case catalog.UploadAssetRequest_Image:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Images.MaxUploadSize
	case catalog.UploadAssetRequest_Document:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Documents.MaxUploadSize
	case catalog.UploadAssetRequest_Audio:
		maxUploadSize = l.svcCtx.Config.UploadConfig.AllowedTypes.Audio.MaxUploadSize
	case catalog.UploadAssetRequest_Video:
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
