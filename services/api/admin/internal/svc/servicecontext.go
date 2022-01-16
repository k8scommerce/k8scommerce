package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"k8scommerce/services/api/admin/internal/config"
	"k8scommerce/services/api/admin/internal/middleware"
)

type ServiceContext struct {
	Config   config.Config
	Locale   rest.Middleware
	StoreKey rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Locale:   middleware.NewLocaleMiddleware().Handle,
		StoreKey: middleware.NewStoreKeyMiddleware().Handle,
	}
}
