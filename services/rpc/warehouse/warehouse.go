package main

import (
	"flag"
	"fmt"
	"log"

	"k8scommerce/services/rpc/warehouse/internal/config"
	"k8scommerce/services/rpc/warehouse/internal/server"
	"k8scommerce/services/rpc/warehouse/internal/svc"
	"k8scommerce/services/rpc/warehouse/pb/warehouse"

	"github.com/joho/godotenv"
	"github.com/localrivet/gcache"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/warehouse.yaml", "the config file")
var envFile = flag.String("e", "./.env", "the .env file")

func main() {
	flag.Parse()

	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)
	universe := gcache.NewUniverse(c.ListenOn)
	srv := server.NewWarehouseClientServer(ctx, universe)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		warehouse.RegisterWarehouseClientServer(grpcServer, srv)

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
