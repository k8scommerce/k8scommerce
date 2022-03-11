package main

import "fmt"

var tables Tables = []Table{
	"store",
	"category",
	"archetype",
	"option",
	"option_item",
	"archetype_option",
	"property",
	"product_category",
	"archetype_property",
	"product",
	"users",
	"asset",
}

type Tables []Table

func (t *Tables) Truncate() {
	for _, table := range *t {
		table.Truncate()
	}
}

type Table string

func (t *Table) Truncate() {
	db.MustExec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE;", *t))
}
