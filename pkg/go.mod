module github.com/k8scommerce/k8scommerce/pkg

go 1.17

require (
	github.com/google/uuid v1.3.0
	github.com/rabbitmq/amqp091-go v1.2.0
	github.com/wagslane/go-rabbitmq v0.7.0
)

replace github.com/k8scommerce/k8scommerce/pkg/hooks => ./hooks

replace github.com/k8scommerce/k8scommerce/pkg/models => ./models

replace github.com/k8scommerce/k8scommerce/pkg/utils => ./utils
