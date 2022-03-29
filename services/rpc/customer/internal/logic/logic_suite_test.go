package logic_test

import (
	"log"
	"testing"

	"k8scommerce/services/rpc/customer/internal/config"
	"k8scommerce/services/rpc/customer/internal/server"
	"k8scommerce/services/rpc/customer/internal/svc"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zeromicro/go-zero/core/conf"
)

var (
	srv *server.CustomerClientServer
)

func TestLogic(t *testing.T) {
	defer GinkgoRecover()
	RegisterFailHandler(Fail)
	loadEnv()
	var c config.Config
	conf.MustLoad("../../etc/customer.yaml", &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)
	srv = server.NewCustomerClientServer(ctx)

	RunSpecs(t, "Logic Suite")
}

func loadEnv() {
	err := godotenv.Load("../../../../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
