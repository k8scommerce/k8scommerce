package svc

import (
	"k8scommerce/internal/events"
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/catalog/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	EventManager events.EventManager
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         repos.NewRepo(&c.PostgresConfig),
		EventManager: events.NewEventManager(&c.EventsConfig),
	}
}
