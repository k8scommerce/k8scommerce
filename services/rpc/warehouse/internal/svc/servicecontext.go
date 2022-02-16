package svc

import (
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/warehouse/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Repo   repos.Repo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repos.MustNewRepo(&c.PostgresConfig),
	}
}
