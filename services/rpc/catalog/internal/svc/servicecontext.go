package svc

import (
	"k8scommerce/internal/repos"
	"k8scommerce/internal/storage"
	"k8scommerce/services/rpc/catalog/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	UploadConfig storage.UploadConfig
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repos.MustNewRepo(&c.PostgresConfig),
	}
}
