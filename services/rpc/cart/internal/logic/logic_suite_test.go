package logic_test

import (
	"log"
	"os"
	"testing"

	"k8scommerce/internal/gcache"
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/cart/internal/config"
	"k8scommerce/services/rpc/cart/internal/server"
	"k8scommerce/services/rpc/cart/internal/svc"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zeromicro/go-zero/core/conf"
)

var (
	repo repos.Repo
	srv  *server.CartClientServer
)

func TestRepos(t *testing.T) {
	defer GinkgoRecover()
	RegisterFailHandler(Fail)
	loadEnv()
	dbConnect()
	var c config.Config
	conf.MustLoad("../../etc/cart.yaml", &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)
	ctx.Cache = gcache.NewGCache()
	srv = server.NewCartClientServer(ctx)

	RunSpecs(t, "CartLogic Suite")
}

func loadEnv() {
	err := godotenv.Load("../../../../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func dbConnect() {
	getPostgresConfig := func() *repos.PostgresConfig {
		return &repos.PostgresConfig{
			DataSourceName: os.Getenv("POSTGRES_DSN"),
		}
	}

	repo = repos.NewRepo(getPostgresConfig())
}

var truncateCart = func() {
	_, err := repo.GetRawDB().Exec(`TRUNCATE cart RESTART IDENTITY CASCADE;`)
	Expect(err).To(BeNil())
}
var truncateCartItem = func() {
	_, err := repo.GetRawDB().Exec(`TRUNCATE cart_item RESTART IDENTITY CASCADE;`)
	Expect(err).To(BeNil())
}
