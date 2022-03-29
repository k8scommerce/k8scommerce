package config

import (
	encryptionconfig "github.com/k8scommerce/k8scommerce/internal/encryption/config"
	eventsconfig "github.com/k8scommerce/k8scommerce/internal/events/config"
	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig   repos.PostgresConfig
	EventsConfig     eventsconfig.EventsConfig
	EncryptionConfig encryptionconfig.EncryptionConfig
}
