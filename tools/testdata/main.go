package main

import (
	"fmt"
	"k8scommerce/internal/repos"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var (
	totalProducts int              = 100
	categories    map[string]int64 = make(map[string]int64)
	storeID       int64
	db            *sqlx.DB
	repo          repos.Repo
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

func main() {
	truncateAll()
	createStore()
	createCategory()
	createArchetypes()
	createProducts()
	createUsers()
}

func truncateAll() {
	tables := []string{
		"store",
		"category",
		"archetype",
		"option",
		"option_item",
		"archetype_option",
		"property",
		"archetype_property",
		"product",
		"users",
	}
	for _, table := range tables {
		db.MustExec(fmt.Sprintf("truncate %s RESTART IDENTITY CASCADE;", table))
	}
}
