package repos

type PostgresConfig struct {
	DataSourceName               string
	MaxOpenConnections           int
	MaxIdleConnections           int
	MaxConnectionLifetimeMinutes int
}
