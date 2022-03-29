package svc

import (
	"github.com/k8scommerce/k8scommerce/services/consumer/imageresizer/internal/config"

	"github.com/k8scommerce/k8scommerce/internal/events"
	"github.com/k8scommerce/k8scommerce/internal/repos"
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
