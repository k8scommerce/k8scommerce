package svc

import (
	"k8scommerce/internal/gcache"
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Repo   repos.Repo
	Cache  gcache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repos.NewRepo(&c.PostgresConfig),
	}
}
