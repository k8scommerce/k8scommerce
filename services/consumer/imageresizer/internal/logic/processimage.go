package logic

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"k8scommerce/internal/models"
	"k8scommerce/internal/storage/asset"
	storagetypes "k8scommerce/internal/storage/types"
	"k8scommerce/internal/workerpool"
	"k8scommerce/services/consumer/imageresizer/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	totalWorkers int = 7
	wp           workerpool.WorkerPool
)

type resizedImage struct {
	Tag string `json:"tag"`
	URL string `json:"url"`
}

func init() {
	wp = workerpool.NewWorkerPool(totalWorkers)
	wp.StartWorkers()
}

type ProcessImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessImageLogic {
	return &ProcessImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// http://localhost:8088/resize?nocrop=true&type=jpeg&url=https://k8scommerce.s3.us-west-1.amazonaws.com/uploads/d/b/e/Pizigani_1367_Chart_10MB.jpeg&width=800&height=800&quality=80&extend=white
func (l *ProcessImageLogic) ProcessImage(in *models.Asset) error {
	var baseUrl = fmt.Sprintf("%s/resize?", l.svcCtx.Config.ImageResizeConfig.ImaginaryURL)

	assetModel, err := l.svcCtx.Repo.Asset().GetAssetById(in.ID)
	if err != nil {
		return status.Errorf(codes.Internal, "asset with id of %d is not found: %s", in.ID, err.Error())
	}

	totalTasks := len(l.svcCtx.Config.ImageResizeConfig.ImageSettings)
	resizedImageC := make(chan resizedImage, totalTasks)

	for _, setting := range l.svcCtx.Config.ImageResizeConfig.ImageSettings {
		setting := setting // https://stackoverflow.com/a/10117080

		wp.AddTask(func() {
			var err error

			url := fmt.Sprintf("%snocrop=true&type=%s&url=%s&width=%d&height=%d&quality=%d&extend=white", baseUrl, setting.OutputType, url.QueryEscape(in.URL), setting.Width, setting.Height, setting.Quality)
			fileName := strings.Join([]string{
				strings.TrimSuffix(in.Name, filepath.Ext(in.Name)),
				"-",
				setting.SizeTag,
				".",
				setting.OutputType,
			}, "")

			logx.Infof("processing image id: %d, size: %s", assetModel.ID, setting.SizeTag)

			resizedImage, err := l.resizeAndSave(url, setting.SizeTag, fileName, *in)
			if err != nil {
				logx.Error(err)
			}

			resizedImageC <- *resizedImage
		})
	}

	imageSizes := make(storagetypes.ImageSizes)
	for i := 0; i < totalTasks; i++ {
		resizedImage := <-resizedImageC
		if resizedImage.Tag != "" {
			imageSizes[resizedImage.Tag] = resizedImage.URL
		}
	}

	if sizes, err := imageSizes.Marshal(); err != nil {
		return status.Errorf(codes.Internal, "failed to marshal sizes: %s", err.Error())
	} else {
		assetModel.Sizes = sizes
		if err = l.svcCtx.Repo.Asset().Update(assetModel); err != nil {
			return status.Errorf(codes.Internal, "failed to updated database with sizes: %s", err.Error())
		}
	}

	return nil
}

func (l *ProcessImageLogic) resizeAndSave(URL, tag, fileName string, in models.Asset) (*resizedImage, error) {
	resizedImage := &resizedImage{}

	file, err := asset.MustNewFile(fileName, l.svcCtx.Config.UploadConfig)
	if err != nil {
		return resizedImage, status.Errorf(codes.Internal, "file creation error: %s", err.Error())
	}

	file.Kind = asset.Kind(in.Kind)

	if err := file.Open(in.ContentType); err != nil {
		return resizedImage, err
	}

	response, err := http.Get(URL)
	if err != nil {
		return resizedImage, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return resizedImage, errors.New("received non 200 response code")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return resizedImage, fmt.Errorf("could not read image body: %s", err.Error())
	}

	if err := file.Write(body, 1); err != nil {
		return resizedImage, fmt.Errorf("chunk write error: %s", err.Error())
	}

	if err := file.Close(); err != nil {
		return resizedImage, fmt.Errorf("error closing file: %s", err.Error())
	}

	resizedImage.Tag = tag
	resizedImage.URL = file.URL

	return resizedImage, nil
}
