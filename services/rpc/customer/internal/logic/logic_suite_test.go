package logic_test

import (
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/customer/internal/config"
	"k8scommerce/services/rpc/customer/internal/server"
	"k8scommerce/services/rpc/customer/internal/svc"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/localrivet/gcache"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zeromicro/go-zero/core/conf"
)

var (
	repo repos.Repo
	srv  *server.CustomerClientServer
)

func TestLogic(t *testing.T) {
	defer GinkgoRecover()
	RegisterFailHandler(Fail)
	loadEnv()
	var c config.Config
	conf.MustLoad("../../etc/customer.yaml", &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)
	universe := gcache.NewUniverse(c.ListenOn)
	srv = server.NewCustomerClientServer(ctx, universe)

	RunSpecs(t, "Logic Suite")
}

func loadEnv() {
	err := godotenv.Load("../../../../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
