package svc

import (
	"k8scommerce/internal/repos"
	storageconfig "k8scommerce/internal/storage/config"
	"k8scommerce/services/rpc/catalog/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	UploadConfig storageconfig.UploadConfig
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repos.MustNewRepo(&c.PostgresConfig),
	}
}
