package broker

import (
	"k8scommerce/internal/events/eventkey"

	"github.com/wagslane/go-rabbitmq"
)

type Broker interface {
	Connect() error
	Publish(key eventkey.EventKey, content []byte) error
	Consume(key eventkey.EventKey, consumerName string, handler rabbitmq.Handler) error
	Disconnect()
}
