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

var (
	categories      Categories = Categories{}
	categoryRootSet bool       = false
)

type Category struct {
	ID        int64
	ParentID  int64
	Name      string
	Parent    string
	Path      string
	SortOrder int64
}

type Categories []Category

func (c *Categories) parentPath(parentId int64) string {
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

	return strings.Join(paths, "/")
}

func (c *Categories) create(parentID interface{}, name string, sortOrder int64) int64 {
	var slugStr = ""
	var rootPath = ""
	var pID sql.NullInt64

	if parentID != nil {
		id := parentID.(int64)
		if id != -1 {
			pID = toNullInt64(id)
			rootPath = c.parentPath(id)
		}
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

	return category.ID
}

func (c *Categories) rebuildNestedSet() {
	db.MustExec(fmt.Sprintf("CALL rebuild_category_nested_set(%d);", storeID))
}

func (c *Categories) Save() {
	for _, category := range *c {
		c.create(category.ParentID, category.Name, category.SortOrder)
		c.rebuildNestedSet()
	}
}

func (c *Categories) Add(name, parent, path string) {
	if parent == "" {
		parent = "root"
	}

	if parentObj := c.exists(name, parent); parentObj == nil {
		if parent == "root" && !categoryRootSet {
			*c = append(*c, Category{
				ID:        ids.NextCategoryID(),
				ParentID:  -1,
				Name:      parent,
				Parent:    "",
				Path:      "",
				SortOrder: 1,
			})
			categoryRootSet = true
		}

		*c = append(*c, Category{
			ID:        ids.NextCategoryID(),
			ParentID:  c.getParentID(parent, path),
			Name:      name,
			Parent:    parent,
			Path:      path,
			SortOrder: sos.NextCategorySortOrder(parent),
		})
	}
}

func (c *Categories) getParentID(name, path string) int64 {
	for _, obj := range *c {
		if obj.Name == name {
			return obj.ID
		}
	}
	return -1
}

func (c *Categories) exists(name, parent string) *Category {
	for _, obj := range *c {
		if obj.Name == name && obj.Parent == parent {
			return &obj
		}
	}
	return nil
}

func (c *Categories) Parse(cat string) {
	paths := strings.Split(cat, "|")
	for _, path := range paths {
		parts := strings.Split(path, ">")
		parent := ""
		for _, category := range parts {
			c.Add(category, parent, path)
			parent = category
		}
	}
}
