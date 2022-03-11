package main

import (
	"k8scommerce/internal/repos"
	"log"
	"os"

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
		MaxOpenConnections:           10,
		MaxIdleConnections:           2,
		MaxConnectionLifetimeMinutes: 5,
	})
	db = repo.GetRawDB()
}
