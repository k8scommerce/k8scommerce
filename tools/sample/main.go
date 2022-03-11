package main

import (
	"math/rand"
	"time"
)

var (
	storeID     int64 = 1
	minProducts       = 1
	maxProducts       = 3
)

func main() {
	rand.Seed(time.Now().UnixNano())
	tables.Truncate()

	createStore()
	createArchetypes()

	categories.Import("categories.csv").Save()
	products.Generate(categories)

	associateProductCategories(categories)

	createUsers()
	createImages()
}
