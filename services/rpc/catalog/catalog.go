package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/config"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/server"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"github.com/k8scommerce/k8scommerce/internal/gcache"

	"github.com/joho/godotenv"
	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/catalog.yaml", "the config file")
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
	pool := groupcache.NewHTTPPoolOpts(c.ListenOn, &groupcache.HTTPPoolOptions{})
	ctx.Cache = gcache.NewGCache()
	srv := server.NewCatalogClientServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		catalog.RegisterCatalogClientServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}

		// gcache peer listener
		gcache.PeerListener(pool, c.ListenOn, c.Etcd)
	})

	// gcache server
	server := gcache.Serve(pool, c.ListenOn)
	defer server.Shutdown(context.Background())

	s.AddOptions(
		grpc.MaxSendMsgSize(c.MaxBytes),
		grpc.MaxRecvMsgSize(c.MaxBytes),
	)

	defer s.Stop()
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
