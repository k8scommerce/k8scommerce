package svc

import (
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/config"

	"github.com/k8scommerce/k8scommerce/internal/events"
	"github.com/k8scommerce/k8scommerce/internal/gcache"
	"github.com/k8scommerce/k8scommerce/internal/repos"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	EventManager events.EventManager
	Cache        gcache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         repos.NewRepo(&c.PostgresConfig),
		EventManager: events.NewEventManager(&c.EventsConfig),
	}
}
