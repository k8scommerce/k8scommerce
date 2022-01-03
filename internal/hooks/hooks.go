package hooks

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type RabbitmqConfig struct {
	Username string
	Password string
	Hostname string
	Port     int32
}

type Hooks interface {
	Publish(key string, content interface{})
	Subscribe(key string) interface{}
}

func NewPublisher(config RabbitmqConfig, routingKeys []string) *rabbitmq.Publisher {
	publisher, err := rabbitmq.NewPublisher(
		fmt.Sprintf("amqp://%s:%s@%s:%d",
			config.Username,
			config.Password,
			config.Hostname,
			config.Port,
		), amqp.Config{},
		rabbitmq.WithPublisherOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}

	return publisher
}

func NewSubscriber(config RabbitmqConfig) rabbitmq.Consumer {
	consumer, err := rabbitmq.NewConsumer(
		fmt.Sprintf("amqp://%s:%s@%s:%d",
			config.Username,
			config.Password,
			config.Hostname,
			config.Port,
		), amqp.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	return consumer
}

// handler rabbitmq.Handler, queue string, routingKeys []string
// err = consumer.StartConsuming(handler,
// 	queue,
// 	[]string{"routing_key", "routing_key_2"},
// 	rabbitmq.WithConsumeOptionsConcurrency(10),
// 	rabbitmq.WithConsumeOptionsQueueDurable,
// 	rabbitmq.WithConsumeOptionsQuorum,
// 	rabbitmq.WithConsumeOptionsBindingExchangeName("events"),
// 	rabbitmq.WithConsumeOptionsBindingExchangeKind("topic"),
// 	rabbitmq.WithConsumeOptionsBindingExchangeDurable,
// 	rabbitmq.WithConsumeOptionsConsumerName(consumerName),
// )
// if err != nil {
// 	log.Fatal(err)
// }
