package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	CartRpc            zrpc.RpcClientConf
	InventoryRpc       zrpc.RpcClientConf
	OthersBoughtRpc    zrpc.RpcClientConf
	ProductRpc         zrpc.RpcClientConf
	SimilarProductsRpc zrpc.RpcClientConf
	UserRpc            zrpc.RpcClientConf
}
