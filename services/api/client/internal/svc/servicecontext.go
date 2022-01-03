package svc

import (
	"k8scommerce/services/api/client/internal/config"
	"k8scommerce/services/api/client/internal/middleware"
	"k8scommerce/services/rpc/cart/cartclient"
	"k8scommerce/services/rpc/inventory/inventoryclient"
	"k8scommerce/services/rpc/othersbought/othersboughtclient"
	"k8scommerce/services/rpc/product/productclient"
	"k8scommerce/services/rpc/similarproducts/similarproductsclient"
	"k8scommerce/services/rpc/user/userclient"

	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config             config.Config
	CartRpc            cartclient.CartClient
	InventoryRpc       inventoryclient.InventoryClient
	OtherBoughtRpc     othersboughtclient.OthersBoughtClient
	ProductRpc         productclient.ProductClient
	SimilarProductsRpc similarproductsclient.SimilarProductsClient
	UserRpc            userclient.UserClient
	Locale             rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		Locale:             middleware.NewLocaleMiddleware().Handle,
		CartRpc:            cartclient.NewCartClient(zrpc.MustNewClient(c.CartRpc)),
		InventoryRpc:       inventoryclient.NewInventoryClient(zrpc.MustNewClient(c.InventoryRpc)),
		OtherBoughtRpc:     othersboughtclient.NewOthersBoughtClient(zrpc.MustNewClient(c.OthersBoughtRpc)),
		ProductRpc:         productclient.NewProductClient(zrpc.MustNewClient(c.ProductRpc)),
		SimilarProductsRpc: similarproductsclient.NewSimilarProductsClient(zrpc.MustNewClient(c.SimilarProductsRpc)),
		UserRpc:            userclient.NewUserClient(zrpc.MustNewClient(c.UserRpc)),
	}
}
