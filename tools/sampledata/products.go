package main

import (
	"fmt"
	"strings"
)

var (
	products       Products = Products{}
	productRootSet bool     = false
)

type Product struct {
	ID       int64
	ParentID int64
	Name     string
	Parent   string
	Path     string
}

type Products []Product

func (c *Products) ToSql() {
	for _, category := range products {
		sql := fmt.Sprintf(`
		insert into products (id, parent_id, name) values (%d, %d, '%s')
		`, category.ID, category.ParentID, category.Name)

		fmt.Println(sql)
	}
}

func (c *Products) Add(name, parent, path string) {
	// pop the last element off the path
	// parts := strings.Split(path, ">")
	// path = strings.Join(parts[:len(parts)-1], ">")
	if parent == "" {
		parent = "root"
	}

	if parentObj := c.exists(name, parent); parentObj == nil {
		if parent == "root" && !productRootSet {
			products = append(products, Product{
				ID:       ids.NextCategoryID(),
				ParentID: 1,
				Name:     parent,
				Parent:   "",
				Path:     "",
			})
			productRootSet = true
		}

		products = append(products, Product{
			ID:       ids.NextCategoryID(),
			ParentID: products.getParentID(parent, path),
			Name:     name,
			Parent:   parent,
			Path:     path,
		})
	}
}

func (c *Products) getParentID(name, path string) int64 {
	for _, obj := range products {
		if obj.Name == name {
			return obj.ID
		}
	}
	return -1
}

func (c *Products) exists(name, parent string) *Product {
	for _, obj := range products {
		if obj.Name == name && obj.Parent == parent {
			return &obj
		}
	}
	return nil
}

func (c *Products) Parse(cat string) {
	paths := strings.Split(cat, "|")
	for _, path := range paths {
		parts := strings.Split(path, ">")
		parent := ""
		for _, category := range parts {
			products.Add(category, parent, path)
			parent = category
		}
	}
}
