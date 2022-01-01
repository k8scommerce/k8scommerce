package main

import (
	"ecomm/internal/env"
	"ecomm/internal/repos"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/joho/godotenv/autoload"
)

var (
	totalProducts int              = 100
	categories    map[string]int64 = make(map[string]int64)
	storeID       int64
	db            *sqlx.DB
)

func init() {
	env.Load()
	repo := repos.MustNewRepo()
	db = repo.GetRawDB()
}

func main() {
	truncateAll()
	createStore()
	createCategory()
	createArchetypes()
	createProducts()
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
	}
	for _, table := range tables {
		db.MustExec(fmt.Sprintf("truncate %s RESTART IDENTITY CASCADE;", table))
	}
}
