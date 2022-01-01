module github.com/k8s-commerce/k8s-commerce/pkg

go 1.17

require (
	github.com/google/uuid v1.3.0
	github.com/rabbitmq/amqp091-go v1.2.0
	github.com/wagslane/go-rabbitmq v0.7.0
)

replace github.com/k8s-commerce/k8s-commerce/pkg/hooks => ./hooks

replace github.com/k8s-commerce/k8s-commerce/pkg/models => ./models

replace github.com/k8s-commerce/k8s-commerce/pkg/utils => ./utils
