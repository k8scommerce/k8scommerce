package svc

import (
	"k8scommerce/internal/repos"
	"k8scommerce/services/rpc/cart/internal/config"
	"k8scommerce/services/rpc/inventory/inventoryclient"
	"k8scommerce/services/rpc/othersbought/othersboughtclient"

	"github.com/tal-tech/go-zero/zrpc"

	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type ServiceContext struct {
	Config         config.Config
	Repo           repos.Repo
	Publisher      *rabbitmq.Publisher
	InventoryRpc   inventoryclient.InventoryClient
	OtherBoughtRpc othersboughtclient.OthersBoughtClient
	// SimilarProductsRpc similarproductsclient.SimilarProductsClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		Repo: repos.MustNewRepo(&repos.Config{
			Connection:                   c.Postgres.Connection,
			MaxOpenConnections:           c.Postgres.MaxOpenConnections,
			MaxIdleConnections:           c.Postgres.MaxIdleConnections,
			MaxConnectionLifetimeMinutes: c.Postgres.MaxConnectionLifetimeMinutes,
		}),
		// Publisher:          InitRabbitMQPublisher(&c),
		InventoryRpc:   inventoryclient.NewInventoryClient(zrpc.MustNewClient(c.InventoryRpc)),
		OtherBoughtRpc: othersboughtclient.NewOthersBoughtClient(zrpc.MustNewClient(c.OthersBoughtRpc)),
		// SimilarProductsRpc: similarproductsclient.NewSimilarProductsClient(zrpc.MustNewClient(c.SimilarProductsRpc)),
	}
}
