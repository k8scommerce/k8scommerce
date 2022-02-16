package config

import (
	"k8scommerce/internal/repos"
	"k8scommerce/internal/storage"

	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig repos.PostgresConfig
	UploadConfig   storage.UploadConfig
}

// type PostgresConfig struct {
// 	Connection                   string
// 	MaxOpenConnections           int
// 	MaxIdleConnections           int
// 	MaxConnectionLifetimeMinutes int
// }
