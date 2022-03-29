package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/k8scommerce/k8scommerce/services/consumer/imageresizer/internal/config"
	"github.com/k8scommerce/k8scommerce/services/consumer/imageresizer/internal/logic"
	"github.com/k8scommerce/k8scommerce/services/consumer/imageresizer/internal/svc"

	"github.com/k8scommerce/k8scommerce/internal/events"
	"github.com/k8scommerce/k8scommerce/internal/events/eventkey"

	"github.com/joho/godotenv"
	rabbitmq "github.com/wagslane/go-rabbitmq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var consumerName = "ImageResizer"

var configFile = flag.String("f", "etc/imageresizer.yaml", "the config file")
var envFile = flag.String("e", "./.env", "the .env file")

func main() {
	flag.Parse()

	err := godotenv.Load(*envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	svcCtx := svc.NewServiceContext(c)

	defer func() {
		if p := recover(); p != nil {
			logx.Error("internal error: %v", p)
		}
	}()

	// start the event manager
	ev := events.NewEventManager(&svcCtx.Config.EventsConfig)

	// consume image upload streams
	err = ev.Consume(eventkey.CatalogImageUploaded.AsKey(), consumerName, func(d rabbitmq.Delivery) rabbitmq.Action {
		go func() {
			asset, err := eventkey.CatalogImageUploaded.Unmarshal(d.Body)
			if err != nil {
				logx.Error("error unmarshalling d.Body : %s", err.Error())
			}

			ctx := context.Background()
			l := logic.NewProcessImageLogic(ctx, svcCtx)
			if err := l.ProcessImage(asset); err != nil {
				logx.Error("error accessing NewProcessImageLogic: %s", err.Error())
			} else {
				logx.Infof("done processing images for image id: %d", asset.ID)
			}
		}()

		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal("Error loading Consumer")
	}

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logx.Info(sig)
		done <- true
	}()

	logx.Info("awaiting signal")
	<-done
	logx.Info("stopping consumer")

	// wait for server to acknowledge the cancel
	ev.Disconnect()
}
