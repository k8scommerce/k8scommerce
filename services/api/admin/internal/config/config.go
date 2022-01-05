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
	CatalogRpc         zrpc.RpcClientConf
	CustomerRpc        zrpc.RpcClientConf
	EmailRpc           zrpc.RpcClientConf
	InventoryRpc       zrpc.RpcClientConf
	OthersBoughtRpc    zrpc.RpcClientConf
	PaymentRpc         zrpc.RpcClientConf
	ShippingRpc        zrpc.RpcClientConf
	SimilarProductsRpc zrpc.RpcClientConf
	StoreRpc           zrpc.RpcClientConf
	UserRpc            zrpc.RpcClientConf
	WarehouseRpc       zrpc.RpcClientConf
}
