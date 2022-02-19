package config

import (
	"k8scommerce/internal/repos"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig  repos.PostgresConfig
	InventoryRpc    zrpc.RpcClientConf
	OthersBoughtRpc zrpc.RpcClientConf
	// SimilarProductsRpc zrpc.RpcClientConf
}
