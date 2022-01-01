package svc

import (
	"github.com/k8s-commerce/k8s-commerce/services/rpc/customer/internal/config"
	"github.com/k8s-commerce/k8s-commerce/services/rpc/customer/internal/repos"
)

type ServiceContext struct {
	Config config.Config
	Repo   repos.Repo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repos.MustNewRepo(&c),
	}
}
