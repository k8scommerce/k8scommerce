package types

type Config struct {
	DbConnection                 string
	MaxOpenConnections           int
	MaxIdleConnections           int
	MaxConnectionLifetimeMinutes int
	Debug                        bool
}
