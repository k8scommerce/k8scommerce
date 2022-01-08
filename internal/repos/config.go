package repos

type Config struct {
	Connection                   string
	MaxOpenConnections           int
	MaxIdleConnections           int
	MaxConnectionLifetimeMinutes int
}
