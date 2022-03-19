package svc

import (
	"k8scommerce/internal/events"
	"k8scommerce/internal/repos"
	"k8scommerce/services/consumer/email/internal/config"
	"k8scommerce/services/consumer/email/internal/email"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	EventManager events.EventManager
	EmailClient  email.EmailClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         repos.NewRepo(&c.PostgresConfig),
		EventManager: events.NewEventManager(&c.EventsConfig),
		EmailClient:  email.NewEmailClient(c.EmailConfig).Connect(),
	}
}
