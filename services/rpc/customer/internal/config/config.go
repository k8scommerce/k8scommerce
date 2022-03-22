package config

import (
	encryptionconfig "k8scommerce/internal/encryption/config"
	eventsconfig "k8scommerce/internal/events/config"
	"k8scommerce/internal/repos"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	PostgresConfig   repos.PostgresConfig
	EventsConfig     eventsconfig.EventsConfig
	EncryptionConfig encryptionconfig.EncryptionConfig
}
