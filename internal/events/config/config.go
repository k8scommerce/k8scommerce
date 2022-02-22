package config

type EventsConfig struct {
	Enabled        bool
	RabbitMQ       bool
	RabbitMQConfig RabbitMQConfig
}

type RabbitMQConfig struct {
	Url string
}
