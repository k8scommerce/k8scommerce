package main

import (
	"log"
	"os"

	"github.com/k8scommerce/k8scommerce/internal/repos"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var (
	db   *sqlx.DB
	repo repos.Repo
)

func init() {
	err := godotenv.Load("./../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repo = repos.NewRepo(&repos.PostgresConfig{
		DataSourceName:               os.Getenv("POSTGRES_DSN"),
		MaxOpenConnections:           100,
		MaxIdleConnections:           2,
		MaxConnectionLifetimeMinutes: 5,
	})
	db = repo.GetRawDB()
}
