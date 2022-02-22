package svc

import (
	"k8scommerce/internal/events"
	"k8scommerce/internal/images"
	"k8scommerce/internal/repos"
	"k8scommerce/services/consumer/imageresizer/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	EventManager events.EventManager
	ImageResizer images.ImageResizer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Repo:         repos.NewRepo(&c.PostgresConfig),
		EventManager: events.NewEventManager(&c.EventsConfig),
		ImageResizer: images.NewImageResizer(&c.ImageResizeConfig),
	}
}
