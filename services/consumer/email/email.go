package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"k8scommerce/internal/events"
	"k8scommerce/services/consumer/email/internal/config"
	"k8scommerce/services/consumer/email/internal/handler"
	"k8scommerce/services/consumer/email/internal/svc"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "etc/email.yaml", "the config file")
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
			logx.Error("internal consumer.email error: %v", p)
		}
	}()

	// start the event manager
	ev := events.NewEventManager(&svcCtx.Config.EventsConfig)
	handler.MustHandle(ev, svcCtx)

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
