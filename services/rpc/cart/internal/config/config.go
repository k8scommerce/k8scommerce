package config

import "github.com/tal-tech/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Postgres           PostgresConfig
	InventoryRpc       zrpc.RpcClientConf
	OthersBoughtRpc    zrpc.RpcClientConf
	SimilarProductsRpc zrpc.RpcClientConf
}

type PostgresConfig struct {
	Connection                   string
	MaxOpenConnections           int
	MaxIdleConnections           int
	MaxConnectionLifetimeMinutes int
}
