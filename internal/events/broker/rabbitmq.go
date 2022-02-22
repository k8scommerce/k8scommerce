package broker

import (
	"k8scommerce/internal/events/config"
	"k8scommerce/internal/events/eventkey"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wagslane/go-rabbitmq"
)

func MustNewRabbitMQBroker(cfg *config.RabbitMQConfig) (Broker, error) {
	return &rabbitMQBroker{
		cfg: cfg,
	}, nil
}

type rabbitMQBroker struct {
	cfg          *config.RabbitMQConfig
	publisher    *rabbitmq.Publisher
	consumer     rabbitmq.Consumer
	consumerName string
}

func (e *rabbitMQBroker) Connect() error {
	if err := e.newPublisher(); err != nil {
		return err
	}

	if err := e.newConsumer(); err != nil {
		return err
	}

	return nil
}

func (e *rabbitMQBroker) Disconnect() {
	noWait := false
	e.consumer.StopConsuming(e.consumerName, noWait)
	e.consumer.Disconnect()
}

func (e *rabbitMQBroker) Publish(key eventkey.EventKey, content []byte) error {
	err := e.publisher.Publish(
		content,
		[]string{string(key)},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("k8scommerce-events"),
	)
	return err
}

func (e *rabbitMQBroker) Consume(key eventkey.EventKey, consumerName string, handler rabbitmq.Handler) error {
	e.consumerName = consumerName
	err := e.consumer.StartConsuming(
		handler,
		e.consumerName,
		[]string{string(key)},
		rabbitmq.WithConsumeOptionsConcurrency(10),
		rabbitmq.WithConsumeOptionsQueueDurable,
		rabbitmq.WithConsumeOptionsQuorum,
		rabbitmq.WithConsumeOptionsBindingExchangeName("k8scommerce-events"),
		rabbitmq.WithConsumeOptionsBindingExchangeKind("topic"),
		rabbitmq.WithConsumeOptionsBindingExchangeDurable,
		rabbitmq.WithConsumeOptionsConsumerName(consumerName),
	)
	if err != nil {
		return err
	}

	return nil
}

func (e *rabbitMQBroker) newPublisher() error {
	publisher, err := rabbitmq.NewPublisher(
		e.cfg.Url,
		amqp.Config{},
		rabbitmq.WithPublisherOptionsLogging,
	)
	if err != nil {
		return err
	}

	e.publisher = publisher
	return nil
}

func (e *rabbitMQBroker) newConsumer() error {
	consumer, err := rabbitmq.NewConsumer(
		e.cfg.Url,
		amqp.Config{},
		rabbitmq.WithConsumerOptionsLogging,
	)
	if err != nil {
		return err
	}
	e.consumer = consumer
	return nil
}
