package config

import (
	eventsconfig "k8scommerce/internal/events/config"
	"k8scommerce/internal/repos"
	storageconfig "k8scommerce/internal/storage/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig repos.PostgresConfig
	UploadConfig   storageconfig.UploadConfig
	EventsConfig   eventsconfig.EventsConfig
	MaxBytes       int
}
