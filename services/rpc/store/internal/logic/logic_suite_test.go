package logic_test

import (
	"log"
	"os"
	"testing"

	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	repo repos.Repo
)

func TestLogic(t *testing.T) {
	defer GinkgoRecover()
	RegisterFailHandler(Fail)
	dbConnect()
	RunSpecs(t, "Logic Suite")
}

func dbConnect() {
	err := godotenv.Load("../../../../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getPostgresConfig := func() *repos.PostgresConfig {
		return &repos.PostgresConfig{
			DataSourceName: os.Getenv("POSTGRES_DSN"),
		}
	}

	repo = repos.NewRepo(getPostgresConfig())
}
