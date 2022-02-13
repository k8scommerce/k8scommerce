package svc

import (
	"k8scommerce/services/api/admin/internal/config"
	"k8scommerce/services/api/admin/internal/middleware"
	"k8scommerce/services/rpc/cart/cartclient"
	"k8scommerce/services/rpc/catalog/catalogclient"
	"k8scommerce/services/rpc/customer/customerclient"
	"k8scommerce/services/rpc/email/emailclient"
	"k8scommerce/services/rpc/inventory/inventoryclient"
	"k8scommerce/services/rpc/othersbought/othersboughtclient"
	"k8scommerce/services/rpc/payment/paymentclient"
	"k8scommerce/services/rpc/shipping/shippingclient"
	"k8scommerce/services/rpc/similarproducts/similarproductsclient"
	"k8scommerce/services/rpc/store/storeclient"
	"k8scommerce/services/rpc/user/userclient"
	"k8scommerce/services/rpc/warehouse/warehouseclient"

	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	Locale             rest.Middleware
	Filter             rest.Middleware
	StoreKey           rest.Middleware
	CartRpc            cartclient.CartClient
	CatalogRpc         catalogclient.CatalogClient
	CustomerRpc        customerclient.CustomerClient
	EmailRpc           emailclient.EmailClient
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
		Locale:             middleware.NewLocaleMiddleware().Handle,
		Filter:             middleware.NewFilterMiddleware().Handle,
		StoreKey:           middleware.NewStoreKeyMiddleware(c.HashSalt).Handle,
		CartRpc:            cartclient.NewCartClient(zrpc.MustNewClient(c.CartRpc)),
		CatalogRpc:         catalogclient.NewCatalogClient(zrpc.MustNewClient(c.CatalogRpc)),
		CustomerRpc:        customerclient.NewCustomerClient(zrpc.MustNewClient(c.CustomerRpc)),
		EmailRpc:           emailclient.NewEmailClient(zrpc.MustNewClient(c.EmailRpc)),
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
