package config

import (
	eventsconfig "github.com/k8scommerce/k8scommerce/internal/events/config"
	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig repos.PostgresConfig
	EventsConfig   eventsconfig.EventsConfig
	EmailConfig    EmailConfig
}

type EmailConfig struct {
	Host        string // = "smtp.example.com"
	Port        int    // = 587
	Username    string // "test@example.com"
	Password    string // "examplepass"
	Encryption  string // "STARTTLS", "SSL/TLS", "STARTTLS"
	DKIMOptions DKIMOptions
}

type DKIMOptions struct {
	PrivateKey            string //= []byte(privateKey)
	Domain                string // "example.com"
	SignatureExpireIn     int    //= 3600
	AddSignatureTimestamp bool   // = true
}
