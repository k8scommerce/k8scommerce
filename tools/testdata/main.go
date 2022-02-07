package main

import (
	"fmt"
	"k8scommerce/internal/repos"

	"github.com/jmoiron/sqlx"

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
	// env.Load()
	repo = repos.MustNewRepo(&repos.Config{
		Connection:                   "postgres://postgres:postgres@localhost:5432/k8scommerce?connect_timeout=180&sslmode=disable",
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
