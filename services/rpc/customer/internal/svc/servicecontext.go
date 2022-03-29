package svc

import (
	"k8scommerce/internal/encryption"
	"k8scommerce/internal/events"
	"k8scommerce/internal/gcache"
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/customer/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	EventManager events.EventManager
	Encrypter    encryption.Encrypter
	Cache        gcache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         repos.NewRepo(&c.PostgresConfig),
		EventManager: events.NewEventManager(&c.EventsConfig),
		Encrypter:    encryption.NewEncrypter(&c.EncryptionConfig),
	}
}
