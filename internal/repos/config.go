package repos

type PostgresConfig struct {
	Connection                   string
	MaxOpenConnections           int
	MaxIdleConnections           int
	MaxConnectionLifetimeMinutes int
}
