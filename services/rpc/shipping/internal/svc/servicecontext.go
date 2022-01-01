package svc

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/shipping/internal/config"
	"github.com/k8scommerce/k8scommerce/services/rpc/shipping/internal/repos"
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
