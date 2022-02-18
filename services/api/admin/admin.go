package main

import (
	"flag"
	"fmt"
	"log"

	"k8scommerce/services/api/admin/internal/config"
	"k8scommerce/services/api/admin/internal/handler"
	"k8scommerce/services/api/admin/internal/svc"

	"github.com/joho/godotenv"
	middleware "github.com/muhfajar/go-zero-cors-middleware"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")
var envFile = flag.String("e", "./.env", "the .env file")

func main() {
	flag.Parse()

	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	c.Timeout = 0
	ctx := svc.NewServiceContext(c)

	cors := middleware.NewCORSMiddleware(&middleware.Options{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token", "Store-Key"},
	})

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors("*"),
		rest.WithNotAllowedHandler(cors.Handler()),
	)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
