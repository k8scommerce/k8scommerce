package config

import (
	"k8scommerce/internal/repos"
	"k8scommerce/internal/storage/config"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig repos.PostgresConfig
	UploadConfig   config.UploadConfig
	MaxBytes       int
}
