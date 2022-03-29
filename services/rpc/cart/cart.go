package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/config"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/server"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/cart/pb/cart"

	"github.com/k8scommerce/k8scommerce/internal/gcache"

	"github.com/joho/godotenv"
	"github.com/mailgun/groupcache/v2"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/cart.yaml", "the config file")
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
	srv := server.NewCartClientServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		cart.RegisterCartClientServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}

		// gcache peer listener
		gcache.PeerListener(pool, c.ListenOn, c.Etcd)
	})

	// gcache server
	server := gcache.Serve(pool, c.ListenOn)
	defer server.Shutdown(context.Background())

	defer s.Stop()

	fmt.Printf("Starting %s.rpc server at %s...\n", "cart", c.ListenOn)
	s.Start()
}
