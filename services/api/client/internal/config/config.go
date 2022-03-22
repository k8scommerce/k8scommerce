package config

import (
	encryptionconfig "k8scommerce/internal/encryption/config"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	EncryptionConfig encryptionconfig.EncryptionConfig
	Auth             struct {
		AccessSecret string
		AccessExpire int64
	}
	CartRpc            zrpc.RpcClientConf
	CatalogRpc         zrpc.RpcClientConf
	CustomerRpc        zrpc.RpcClientConf
	InventoryRpc       zrpc.RpcClientConf
	OthersBoughtRpc    zrpc.RpcClientConf
	PaymentRpc         zrpc.RpcClientConf
	ShippingRpc        zrpc.RpcClientConf
	SimilarProductsRpc zrpc.RpcClientConf
	StoreRpc           zrpc.RpcClientConf
	UserRpc            zrpc.RpcClientConf
	WarehouseRpc       zrpc.RpcClientConf
}
