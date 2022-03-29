package config

import (
	eventsconfig "github.com/k8scommerce/k8scommerce/internal/events/config"
	imagesconfig "github.com/k8scommerce/k8scommerce/internal/images/config"
	"github.com/k8scommerce/k8scommerce/internal/repos"
	storageconfig "github.com/k8scommerce/k8scommerce/internal/storage/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig    repos.PostgresConfig
	EventsConfig      eventsconfig.EventsConfig
	ImageResizeConfig imagesconfig.ImageResizeConfig
	UploadConfig      storageconfig.UploadConfig
}
