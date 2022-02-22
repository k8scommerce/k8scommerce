package events

import (
	"fmt"
	"k8scommerce/internal/events/broker"
	"k8scommerce/internal/events/config"
	"k8scommerce/internal/events/eventkey"

	"github.com/wagslane/go-rabbitmq"
	"github.com/zeromicro/go-zero/core/logx"
)

type EventManager interface {
	Publish(key eventkey.EventKey, content []byte) error
	Consume(key eventkey.EventKey, consumerName string, handler rabbitmq.Handler) error
	Disconnect()
}

func NewEventManager(cfg *config.EventsConfig) EventManager {
	// catch any panics
	defer func() {
		if rec := recover(); rec != nil {
			logx.Infof("Panic Recovered in NewEventManager: %#v", rec)
		}
	}()

	em := &eventManager{
		cfg: cfg,
	}

	if err := em.setBroker(); err != nil {
		logx.Infof("error setting event broker: %s", err.Error())
	}

	return em
}

type eventManager struct {
	cfg    *config.EventsConfig
	broker broker.Broker
}

func (e *eventManager) Publish(key eventkey.EventKey, content []byte) error {
	return e.broker.Publish(key, content)
}

func (e *eventManager) Consume(key eventkey.EventKey, consumerName string, handler rabbitmq.Handler) error {
	return e.broker.Consume(key, consumerName, handler)
}

func (e *eventManager) Disconnect() {
	e.broker.Disconnect()
}

func (e *eventManager) setBroker() error {
	var eventBroker broker.Broker
	if e.cfg.RabbitMQ {
		mq, err := broker.MustNewRabbitMQBroker(&e.cfg.RabbitMQConfig)
		if err != nil {
			return fmt.Errorf("couldn't create rabbitMQ broker: %s", err)
		}
		eventBroker = mq
	}
	if eventBroker != nil {
		if err := eventBroker.Connect(); err != nil {
			return fmt.Errorf("even broker could not connect: %s", err)
		}

		e.broker = eventBroker
		return nil
	}

	return fmt.Errorf("no broker enabled in config")
}
