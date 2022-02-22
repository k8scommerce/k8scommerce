package config

import (
	eventsconfig "k8scommerce/internal/events/config"
	imagesconfig "k8scommerce/internal/images/config"
	"k8scommerce/internal/repos"
	storageconfig "k8scommerce/internal/storage/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig    repos.PostgresConfig
	EventsConfig      eventsconfig.EventsConfig
	ImageResizeConfig imagesconfig.ImageResizeConfig
	UploadConfig      storageconfig.UploadConfig
}
