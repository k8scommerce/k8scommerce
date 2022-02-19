package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	HashSalt string
	Auth     struct {
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
