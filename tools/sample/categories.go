package main

import (
	"context"
	"database/sql"
	"fmt"
	"k8scommerce/internal/models"
	"os"

	"github.com/gocarina/gocsv"
)

var (
	categories Categories = Categories{}
)

type Categories []Category
type Category struct {
	ID          int64          `csv:"id"`
	ParentID    int64          `csv:"parent_id"`
	Slug        string         `csv:"slug"`
	Name        string         `csv:"name"`
	SortOrder   int64          `csv:"sort_order"`
	StoreID     int64          `csv:"store_id"`
	Description sql.NullString `csv:"description"`
}

func contains(s []int64, eval int64) bool {
	for _, v := range s {
		if v == eval {
			return true
		}
	}

	return false
}

func (c *Categories) ToMap() map[string]int64 {
	res := make(map[string]int64)
	for _, category := range *c {
		res[category.Name] = category.ID
	}
	return res
}

func (c *Categories) Save() {
	saved := []int64{}
	for _, category := range *c {
		model := &models.Category{
			ID:       category.ID,
			ParentID: toNullInt64(category.ParentID),
			StoreID:  storeID,
			Slug:     category.Slug,
			Name:     category.Name,
		}

		if category.ParentID == 0 {
			model.ParentID = sql.NullInt64{}
		}

		if !contains(saved, category.ID) {
			if err := model.Insert(context.Background(), db); err != nil {
				fmt.Printf("ERROR: %s, %d\n", model.Name, model.ID)
				panic(err)
			}
			saved = append(saved, model.ID)
		}
	}

	c.rebuildNestedSet()
}

func (c *Categories) rebuildNestedSet() {
	db.MustExec(fmt.Sprintf("CALL rebuild_category_nested_set(%d);", storeID))
}

func (c *Categories) Import(filename string) *Categories {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, c); err != nil {
		panic(err)
	}

	*c = append(*c, categories...)
	// for _, category := range categories {
	// 	// fmt.Printf("%#v\n\n\n", category)
	// 	*c = append(*c, category)
	// }

	return c
}
