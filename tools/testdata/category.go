package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/gosimple/slug"
	"github.com/iancoleman/strcase"
)

func createCategory() {
	parentPath := func(parentId int64) string {
		// Query the DB
		rows, err := repo.GetRawDB().Query(`select 
				parent.name as name 
			from category as node, 
				category as parent
			where node.lft between parent.lft and parent.rgt
			and node.id = $1
			and parent.depth > 1
			order by parent.lft
		`, parentId)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var paths []string

		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				log.Fatal(err)
			}
			paths = append(paths, slug.Make(name))
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		// fmt.Println(parentId, len(paths), strings.Join(paths, "/"))

		return strings.Join(paths, "/")
	}

	createOne := func(parentID interface{}, name string, sortOrder int64) int64 {
		var slugStr = ""
		var rootPath = ""
		var pID sql.NullInt64
		if parentID != nil {
			id := parentID.(int64)
			pID = toNullInt64(id)

			rootPath = parentPath(id)
		}

		if storeID == 0 {
			storeID = 1
		}

		if rootPath != "" {
			slugStr = fmt.Sprintf("%s/%s", rootPath, slug.Make(name))
		} else {
			slugStr = slug.Make(name)
		}

		category := &models.Category{
			ParentID:    pID,
			StoreID:     storeID,
			Slug:        slugStr,
			Name:        name,
			Description: toNullString(name),
			SortOrder:   toNullInt64(sortOrder),
		}

		if err := category.Insert(context.Background(), db); err != nil {
			fmt.Println(strcase.ToKebab(name))
			panic(err)
		}

		db.MustExec(fmt.Sprintf("CALL rebuild_category_nested_set(%d);", storeID))

		return category.ID
	}

	categories["root"] = createOne(nil, "Root Category", 1)

	categories["women"] = createOne(categories["root"], "Women", 1)

	categories["What's New"] = createOne(categories["women"], "What's New", 1)
	categories["New Arrivals"] = createOne(categories["What's New"], "New Arrivals", 1)
	categories["Natural Ease"] = createOne(categories["What's New"], "Natural Ease", 2)
	categories["Accessories Shop"] = createOne(categories["What's New"], "Accessories Shop", 3)
	categories["Dress Shop"] = createOne(categories["What's New"], "Dress Shop", 4)
	categories["Welington Collection"] = createOne(categories["What's New"], "Welington Collection", 5)
	categories["Team USA Collection"] = createOne(categories["What's New"], "Team USA Collection", 6)
	categories["Vintage"] = createOne(categories["What's New"], "Vintage", 7)
	categories["The Bear Shop"] = createOne(categories["What's New"], "The Bear Shop", 8)
	categories["Pin of Solidarity"] = createOne(categories["What's New"], "Pin of Solidarity", 9)

	categories["Clothing"] = createOne(categories["women"], "Clothing", 2)
	categories["Sweaters"] = createOne(categories["Clothing"], "Sweaters", 1)
	categories["Dresses & Jumpsuits"] = createOne(categories["Clothing"], "Dresses & Jumpsuits", 2)
	categories["T-Shirts & Sweatshirts"] = createOne(categories["Clothing"], "T-Shirts & Sweatshirts", 3)
	categories["Polo Shirts"] = createOne(categories["Clothing"], "Polo Shirts", 4)
	categories["Outerwear & Vests"] = createOne(categories["Clothing"], "Outerwear & Vests", 5)
	categories["Jackets & Blazers"] = createOne(categories["Clothing"], "Jackets & Blazers", 6)
	categories["Pants"] = createOne(categories["Clothing"], "Pants", 7)
	categories["Jeans & Denim"] = createOne(categories["Clothing"], "Jeans & Denim", 8)
	categories["Skirts"] = createOne(categories["Clothing"], "Skirts", 9)
	categories["Shorts"] = createOne(categories["Clothing"], "Shorts", 10)
	categories["Activewear"] = createOne(categories["Clothing"], "Activewear", 11)
	categories["Swimsuits & Cover-Ups"] = createOne(categories["Clothing"], "Swimsuits & Cover-Ups", 12)
	categories["Sleepwear & Loungewear"] = createOne(categories["Clothing"], "Sleepwear & Loungewear", 13)

	categories["Petite (Sizes 2–14)"] = createOne(categories["women"], "Petite (Sizes 2–14)", 3)
	categories["Woman (Sizes 14–22)"] = createOne(categories["women"], "Woman (Sizes 14–22)", 4)

	categories["Shoes"] = createOne(categories["women"], "Shoes", 5)
	categories["Boots"] = createOne(categories["Shoes"], "Boots", 1)
	categories["Flats & Sneakers"] = createOne(categories["Shoes"], "Flats & Sneakers", 2)
	categories["Heels & Pumps"] = createOne(categories["Shoes"], "Heels & Pumps", 3)
	categories["Sandals"] = createOne(categories["Shoes"], "Sandals", 4)

	categories["Accessories"] = createOne(categories["women"], "Accessories", 6)
	categories["Handbags"] = createOne(categories["Accessories"], "Handbags", 1)
	categories["Hats, Scarves & Gloves"] = createOne(categories["Accessories"], "Hats, Scarves & Gloves", 2)
	categories["Masks & Bandannas"] = createOne(categories["Accessories"], "Masks & Bandannas", 3)
	categories["Wallets & Accessories"] = createOne(categories["Accessories"], "Wallets & Accessories", 4)
	categories["Belts"] = createOne(categories["Accessories"], "Belts", 5)
	categories["Socks"] = createOne(categories["Accessories"], "Socks", 6)
	categories["Sunglasses"] = createOne(categories["Accessories"], "Sunglasses", 7)
	categories["Fashion Jewelry"] = createOne(categories["Accessories"], "Fashion Jewelry", 8)
	categories["Fragrance"] = createOne(categories["Accessories"], "Fragrance", 9)

	categories["Watches & Fine Jewelry"] = createOne(categories["women"], "Watches & Fine Jewelry", 7)

	categories["Create Your Own"] = createOne(categories["women"], "Create Your Own", 8)
	categories["Clothing"] = createOne(categories["Create Your Own"], "Clothing", 1)
	categories["Accessories"] = createOne(categories["Create Your Own"], "Accessories", 2)
	categories["Custom Outerwear"] = createOne(categories["Create Your Own"], "Custom Outerwear", 3)
	categories["Made to Order"] = createOne(categories["Create Your Own"], "Made to Order", 4)

	categories["Our Brands"] = createOne(categories["women"], "Our Brands", 9)
	categories["Polo Ralph Lauren"] = createOne(categories["Our Brands"], "Polo Ralph Lauren", 1)
	categories["RLX"] = createOne(categories["Our Brands"], "RLX", 2)
	categories["Collection"] = createOne(categories["Our Brands"], "Collection", 3)
	categories["Double RL"] = createOne(categories["Our Brands"], "Double RL", 4)
	categories["Lauren"] = createOne(categories["Our Brands"], "Lauren", 5)
	categories["Pink Pony"] = createOne(categories["Our Brands"], "Pink Pony", 6)
	categories["Golf"] = createOne(categories["Our Brands"], "Golf", 7)

	categories["Sale"] = createOne(categories["women"], "Sale", 10)
	categories["Clothing"] = createOne(categories["Sale"], "Clothing", 1)
	categories["Shoes"] = createOne(categories["Sale"], "Shoes", 1)
	categories["Accessories"] = createOne(categories["Sale"], "Accessories", 1)
	categories["Collection: Enjoy Up to 30% Off"] = createOne(categories["Sale"], "Collection: Enjoy Up to 30% Off", 1)
	categories["Double RL: Enjoy Up to 30% Off"] = createOne(categories["Sale"], "Double RL: Enjoy Up to 30% Off", 1)

	//
	//
	//
	//
	//
	categories["men"] = createOne(categories["root"], "Men", 1)

	categories["What's New"] = createOne(categories["men"], "What's New", 1)
	categories["New Arrivals"] = createOne(categories["What's New"], "New Arrivals", 1)
	categories["Natural Ease"] = createOne(categories["What's New"], "Natural Ease", 2)
	categories["Accessories Shop"] = createOne(categories["What's New"], "Accessories Shop", 3)
	categories["Dress Shop"] = createOne(categories["What's New"], "Dress Shop", 4)
	categories["Welington Collection"] = createOne(categories["What's New"], "Welington Collection", 5)
	categories["Team USA Collection"] = createOne(categories["What's New"], "Team USA Collection", 6)
	categories["Vintage"] = createOne(categories["What's New"], "Vintage", 7)
	categories["The Bear Shop"] = createOne(categories["What's New"], "The Bear Shop", 8)
	categories["Pin of Solidarity"] = createOne(categories["What's New"], "Pin of Solidarity", 9)

	categories["Clothing"] = createOne(categories["men"], "Clothing", 2)
	categories["Sweaters"] = createOne(categories["Clothing"], "Sweaters", 1)
	categories["Dresses & Jumpsuits"] = createOne(categories["Clothing"], "Dresses & Jumpsuits", 2)
	categories["T-Shirts & Sweatshirts"] = createOne(categories["Clothing"], "T-Shirts & Sweatshirts", 3)
	categories["Polo Shirts"] = createOne(categories["Clothing"], "Polo Shirts", 4)
	categories["Outerwear & Vests"] = createOne(categories["Clothing"], "Outerwear & Vests", 5)
	categories["Jackets & Blazers"] = createOne(categories["Clothing"], "Jackets & Blazers", 6)
	categories["Pants"] = createOne(categories["Clothing"], "Pants", 7)
	categories["Jeans & Denim"] = createOne(categories["Clothing"], "Jeans & Denim", 8)
	categories["Skirts"] = createOne(categories["Clothing"], "Skirts", 9)
	categories["Shorts"] = createOne(categories["Clothing"], "Shorts", 10)
	categories["Activewear"] = createOne(categories["Clothing"], "Activewear", 11)
	categories["Swimsuits & Cover-Ups"] = createOne(categories["Clothing"], "Swimsuits & Cover-Ups", 12)
	categories["Sleepwear & Loungewear"] = createOne(categories["Clothing"], "Sleepwear & Loungewear", 13)

	categories["Petite (Sizes 2–14)"] = createOne(categories["men"], "Petite (Sizes 2–14)", 3)
	categories["Woman (Sizes 14–22)"] = createOne(categories["men"], "Woman (Sizes 14–22)", 4)

	categories["Shoes"] = createOne(categories["men"], "Shoes", 5)
	categories["Boots"] = createOne(categories["Shoes"], "Boots", 1)
	categories["Flats & Sneakers"] = createOne(categories["Shoes"], "Flats & Sneakers", 2)
	categories["Heels & Pumps"] = createOne(categories["Shoes"], "Heels & Pumps", 3)
	categories["Sandals"] = createOne(categories["Shoes"], "Sandals", 4)

	categories["Accessories"] = createOne(categories["men"], "Accessories", 6)
	categories["Handbags"] = createOne(categories["Accessories"], "Handbags", 1)
	categories["Hats, Scarves & Gloves"] = createOne(categories["Accessories"], "Hats, Scarves & Gloves", 2)
	categories["Masks & Bandannas"] = createOne(categories["Accessories"], "Masks & Bandannas", 3)
	categories["Wallets & Accessories"] = createOne(categories["Accessories"], "Wallets & Accessories", 4)
	categories["Belts"] = createOne(categories["Accessories"], "Belts", 5)
	categories["Socks"] = createOne(categories["Accessories"], "Socks", 6)
	categories["Sunglasses"] = createOne(categories["Accessories"], "Sunglasses", 7)
	categories["Fashion Jewelry"] = createOne(categories["Accessories"], "Fashion Jewelry", 8)
	categories["Fragrance"] = createOne(categories["Accessories"], "Fragrance", 9)

	categories["Watches & Fine Jewelry"] = createOne(categories["men"], "Watches & Fine Jewelry", 7)

	categories["Create Your Own"] = createOne(categories["men"], "Create Your Own", 8)
	categories["Clothing"] = createOne(categories["Create Your Own"], "Clothing", 1)
	categories["Accessories"] = createOne(categories["Create Your Own"], "Accessories", 2)
	categories["Custom Outerwear"] = createOne(categories["Create Your Own"], "Custom Outerwear", 3)
	categories["Made to Order"] = createOne(categories["Create Your Own"], "Made to Order", 4)

	categories["Our Brands"] = createOne(categories["men"], "Our Brands", 9)
	categories["Polo Ralph Lauren"] = createOne(categories["Our Brands"], "Polo Ralph Lauren", 1)
	categories["RLX"] = createOne(categories["Our Brands"], "RLX", 2)
	categories["Collection"] = createOne(categories["Our Brands"], "Collection", 3)
	categories["Double RL"] = createOne(categories["Our Brands"], "Double RL", 4)
	categories["Lauren"] = createOne(categories["Our Brands"], "Lauren", 5)
	categories["Pink Pony"] = createOne(categories["Our Brands"], "Pink Pony", 6)
	categories["Golf"] = createOne(categories["Our Brands"], "Golf", 7)

	categories["Sale"] = createOne(categories["men"], "Sale", 10)
	categories["Clothing"] = createOne(categories["Sale"], "Clothing", 1)
	categories["Shoes"] = createOne(categories["Sale"], "Shoes", 1)
	categories["Accessories"] = createOne(categories["Sale"], "Accessories", 1)
	categories["Collection: Enjoy Up to 30% Off"] = createOne(categories["Sale"], "Collection: Enjoy Up to 30% Off", 1)
	categories["Double RL: Enjoy Up to 30% Off"] = createOne(categories["Sale"], "Double RL: Enjoy Up to 30% Off", 1)

	// categories["kids"] = createOne(categories["root"], "Kids", 3)
	// categories["kids pants"] = createOne(categories["kids"], "Kids Pants", 1)
	// categories["kids shirts"] = createOne(categories["kids"], "Kids Shirts", 2)
}
