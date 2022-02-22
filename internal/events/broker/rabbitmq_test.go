package broker_test

import (
	"k8scommerce/internal/events/broker"
	"k8scommerce/internal/events/config"
	"k8scommerce/internal/events/eventkey"
	"log"
	"os"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/wagslane/go-rabbitmq"
)

var _ = Describe("RabbitMQBroker", func() {
	defer GinkgoRecover()

	var err error
	var mq broker.Broker

	err = godotenv.Load("../../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getConfig := func() *config.RabbitMQConfig {
		return &config.RabbitMQConfig{
			Url: os.Getenv("RABBITMQ_URL"),
		}
	}
	Expect(getConfig().Url).ToNot(BeEmpty())

	BeforeEach(func() {
		mq, err = broker.MustNewRabbitMQBroker(getConfig())
		Expect(err).To(BeNil())
		Expect(mq).ToNot(BeNil())
	})

	It("should connect to rabbitmq", func() {
		err = mq.Connect()
		Expect(err).To(BeNil())
	})

	Describe("Pub/Sub", func() {

		var message = "hello, world"
		var consumerName = "ginkgo-tester"

		BeforeEach(func() {
			err = mq.Connect()
			Expect(err).To(BeNil())
		})

		It("should publish a message", func() {
			content := []byte(message)
			err = mq.Publish(eventkey.DebugTesting.AsKey(), content)
			Expect(err).To(BeNil())
		})

		It("should consume a message", func() {
			err = mq.Consume(eventkey.DebugTesting.AsKey(), consumerName, func(d rabbitmq.Delivery) rabbitmq.Action {
				Expect(string(d.Body)).ToNot(BeNil())
				Expect(string(d.Body)).To(Equal(message))
				return rabbitmq.Ack
			})
			Expect(err).To(BeNil())
		})
	})
})
