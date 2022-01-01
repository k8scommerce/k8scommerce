package main

import (
	"context"
	"database/sql"
	"fmt"

	"ecomm/internal/repos/models"
)

func createCategory() {
	createOne := func(parentID interface{}, name string, sortOrder int64) int64 {
		var pID sql.NullInt64
		if parentID != nil {
			id := parentID.(int64)
			pID = toNullInt64(id)
		}

		category := &models.Category{
			ParentID:    pID,
			StoreID:     storeID,
			Name:        name,
			Description: toNullString(name),
			SortOrder:   toNullInt64(sortOrder),
		}
		if err := category.Insert(context.Background(), db); err != nil {
			panic(err)
		}
		return category.ID
	}

	categories["root"] = createOne(nil, "Root Category", 1)

	categories["women"] = createOne(categories["root"], "Women", 1)
	categories["womens pants"] = createOne(categories["women"], "Pants", 1)
	categories["womens shirts"] = createOne(categories["women"], "Shirts", 2)

	categories["men"] = createOne(categories["root"], "Men", 2)
	categories["mens pants"] = createOne(categories["men"], "Pants", 1)
	categories["mens shirts"] = createOne(categories["men"], "Shirts", 2)

	categories["kids"] = createOne(categories["root"], "Kids", 3)
	categories["kids pants"] = createOne(categories["kids"], "Pants", 1)
	categories["kids shirts"] = createOne(categories["kids"], "Shirts", 2)

	db.MustExec(fmt.Sprintf("CALL rebuild_category_nested_set(%d);", storeID))
}
