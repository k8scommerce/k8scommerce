package main

import (
	"flag"
	"fmt"

	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/config"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/server"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/product/pb/product"

	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/joho/godotenv/autoload"
)

var configFile = flag.String("f", "etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	universe := gcache.NewUniverse(c.ListenOn)
	srv := server.NewProductClientServer(ctx, universe)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		product.RegisterProductClientServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}

		sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		update := func() {
			universe.Set(sub.Values()...)
			fmt.Printf("universe.Set: %#v\n", sub.Values())
		}
		sub.AddListener(update)
		update()
	})
	defer s.Stop()
	defer universe.Shutdown()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
