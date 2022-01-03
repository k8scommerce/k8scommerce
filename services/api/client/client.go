package main

import (
	"flag"
	"fmt"

	"k8scommerce/services/api/client/internal/config"
	"k8scommerce/services/api/client/internal/handler"
	"k8scommerce/services/api/client/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/client.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors("*"),
	)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
