package svc

import (
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/config"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/middleware"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/cartclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/catalogclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/customer/customerclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/inventory/inventoryclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/othersbought/othersboughtclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/payment/paymentclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/shipping/shippingclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/similarproducts/similarproductsclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/storeclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/user/userclient"
	"github.com/k8scommerce/k8scommerce/services/rpc/warehouse/warehouseclient"

	"github.com/k8scommerce/k8scommerce/internal/encryption"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	Encrypter          encryption.Encrypter
	Locale             rest.Middleware
	Filter             rest.Middleware
	StoreKey           rest.Middleware
	CartRpc            cartclient.CartClient
	CatalogRpc         catalogclient.CatalogClient
	CustomerRpc        customerclient.CustomerClient
	InventoryRpc       inventoryclient.InventoryClient
	OthersBoughtRpc    othersboughtclient.OthersBoughtClient
	PaymentRpc         paymentclient.PaymentClient
	ShippingRpc        shippingclient.ShippingClient
	SimilarProductsRpc similarproductsclient.SimilarProductsClient
	StoreRpc           storeclient.StoreClient
	UserRpc            userclient.UserClient
	WarehouseRpc       warehouseclient.WarehouseClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		Encrypter:          encryption.NewEncrypter(&c.EncryptionConfig),
		Locale:             middleware.NewLocaleMiddleware().Handle,
		Filter:             middleware.NewFilterMiddleware().Handle,
		StoreKey:           middleware.NewStoreKeyMiddleware(c.EncryptionConfig).Handle,
		CartRpc:            cartclient.NewCartClient(zrpc.MustNewClient(c.CartRpc)),
		CatalogRpc:         catalogclient.NewCatalogClient(zrpc.MustNewClient(c.CatalogRpc)),
		CustomerRpc:        customerclient.NewCustomerClient(zrpc.MustNewClient(c.CustomerRpc)),
		InventoryRpc:       inventoryclient.NewInventoryClient(zrpc.MustNewClient(c.InventoryRpc)),
		OthersBoughtRpc:    othersboughtclient.NewOthersBoughtClient(zrpc.MustNewClient(c.OthersBoughtRpc)),
		PaymentRpc:         paymentclient.NewPaymentClient(zrpc.MustNewClient(c.PaymentRpc)),
		ShippingRpc:        shippingclient.NewShippingClient(zrpc.MustNewClient(c.ShippingRpc)),
		SimilarProductsRpc: similarproductsclient.NewSimilarProductsClient(zrpc.MustNewClient(c.SimilarProductsRpc)),
		StoreRpc:           storeclient.NewStoreClient(zrpc.MustNewClient(c.StoreRpc)),
		WarehouseRpc:       warehouseclient.NewWarehouseClient(zrpc.MustNewClient(c.WarehouseRpc)),
		UserRpc:            userclient.NewUserClient(zrpc.MustNewClient(c.UserRpc)),
	}
}
