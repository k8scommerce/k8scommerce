package svc

import (
	"k8scommerce/services/rpc/customer/internal/config"
	"k8scommerce/services/rpc/customer/internal/repos"
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
