package svc

import (
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/catalog/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Repo   repos.Repo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo: repos.MustNewRepo(&repos.Config{
			Connection:                   c.Postgres.Connection,
			MaxOpenConnections:           c.Postgres.MaxOpenConnections,
			MaxIdleConnections:           c.Postgres.MaxIdleConnections,
			MaxConnectionLifetimeMinutes: c.Postgres.MaxConnectionLifetimeMinutes,
		}),
	}
}
