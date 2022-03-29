package svc

import (
	"github.com/k8scommerce/k8scommerce/services/consumer/email/internal/config"
	"github.com/k8scommerce/k8scommerce/services/consumer/email/internal/email"

	"github.com/k8scommerce/k8scommerce/internal/events"
	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type ServiceContext struct {
	Config       config.Config
	Repo         repos.Repo
	EventManager events.EventManager
	EmailClient  email.EmailClient
	Localizer    *i18n.Localizer
}

func NewServiceContext(c config.Config) *ServiceContext {

	// create the language localizer
	bundle := i18n.NewBundle(language.English)
	localizer := i18n.NewLocalizer(bundle, language.English.String())

	return &ServiceContext{
		Config:       c,
		Repo:         repos.NewRepo(&c.PostgresConfig),
		EventManager: events.NewEventManager(&c.EventsConfig),
		EmailClient:  email.NewEmailClient(c.EmailConfig).Connect(),
		Localizer:    localizer,
	}
}
