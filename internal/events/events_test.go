package events_test

import (
	"k8scommerce/internal/events"
	"k8scommerce/internal/events/config"
	"k8scommerce/internal/events/eventkey"
	"log"
	"os"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/wagslane/go-rabbitmq"
)

var _ = Describe("Events", func() {
	defer GinkgoRecover()

	var err error
	var em events.EventManager

	err = godotenv.Load("../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getConfig := func() *config.EventsConfig {
		return &config.EventsConfig{
			Enabled:  true,
			RabbitMQ: true,
			RabbitMQConfig: config.RabbitMQConfig{
				Url: os.Getenv("RABBITMQ_URL"),
			},
		}
	}
	Expect(getConfig()).ToNot(BeEmpty())

	BeforeEach(func() {
		em = events.NewEventManager(getConfig())
		Expect(em).ToNot(BeNil())
	})

	Describe("Pub/Sub", func() {

		var message = "hello, world"
		var consumerName = "ginkgo-tester"

		It("should publish a message", func() {
			content := []byte(message)
			err = em.Publish(eventkey.DebugTesting.AsKey(), content)
			Expect(err).To(BeNil())
		})

		It("should consume a message", func() {
			err = em.Consume(eventkey.DebugTesting.AsKey(), consumerName, func(d rabbitmq.Delivery) rabbitmq.Action {
				Expect(string(d.Body)).ToNot(BeNil())
				Expect(string(d.Body)).To(Equal(message))
				return rabbitmq.Ack
			})
			Expect(err).To(BeNil())
		})
	})
})
