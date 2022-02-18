package main

import (
	"flag"
	"fmt"

	"k8scommerce/services/rpc/inventory/internal/config"
	"k8scommerce/services/rpc/inventory/internal/server"
	"k8scommerce/services/rpc/inventory/internal/svc"
	"k8scommerce/services/rpc/inventory/pb/inventory"

	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/inventory.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	universe := gcache.NewUniverse(c.ListenOn)
	srv := server.NewInventoryClientServer(ctx, universe)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		inventory.RegisterInventoryClientServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}

		// sub, err := discov.NewSubscriber(c.Etcd.Hosts, c.Etcd.Key)
		// if err != nil {
		// 	fmt.Println("ERROR:", err)
		// }

		// update := func() {
		// 	universe.Set(sub.Values()...)
		// 	fmt.Printf("universe.Set: %#v\n", sub.Values())
		// }
		// sub.AddListener(update)
		// update()
	})
	defer s.Stop()
	defer universe.Shutdown()

	fmt.Printf("Starting %s.rpc server at %s...\n", "inventory", c.ListenOn)
	s.Start()
}
