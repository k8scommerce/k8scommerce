package config

import (
	"k8scommerce/internal/repos"

	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig repos.PostgresConfig
}
