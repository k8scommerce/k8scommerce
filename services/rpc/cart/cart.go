package main

import (
	"flag"
	"fmt"

	"k8scommerce/services/rpc/cart/internal/config"
	"k8scommerce/services/rpc/cart/internal/server"
	"k8scommerce/services/rpc/cart/internal/svc"
	"k8scommerce/services/rpc/cart/pb/cart"

	"github.com/localrivet/gcache"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cart.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	universe := gcache.NewUniverse(c.ListenOn)
	srv := server.NewCartClientServer(ctx, universe)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		cart.RegisterCartClientServer(grpcServer, srv)

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
