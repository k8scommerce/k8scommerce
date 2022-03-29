package config

import (
	eventsconfig "github.com/k8scommerce/k8scommerce/internal/events/config"
	"github.com/k8scommerce/k8scommerce/internal/repos"
	storageconfig "github.com/k8scommerce/k8scommerce/internal/storage/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig repos.PostgresConfig
	UploadConfig   storageconfig.UploadConfig
	EventsConfig   eventsconfig.EventsConfig
	MaxBytes       int
}
