package svc

import (
	"ecomm/services/api/admin/internal/config"
	"ecomm/services/rpc/product/productclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productclient.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProductClient(zrpc.MustNewClient(c.ProductRpc)),
	}
}
