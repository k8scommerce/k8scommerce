package main

import (
	"context"
	"fmt"
	"k8scommerce/internal/models"

	"github.com/iancoleman/strcase"
)

var options = make(map[string]int64)
var optionItemNames = make(map[int64]string)
var optionItems = make(map[int64][]int64)
var archetypes = make(map[string]int64)
var properties = make(map[string]int64)

func createArchetypes() {
	a := archetype{}

	options["color"] = a.newOption("Color", "Color")
	a.addBulkOptionItems(options["color"], []string{
		"Black",
		"Blue",
		"Brown",
		"Cream",
		"Gold",
		"Green",
		"Grey",
		"Multi",
		"Orange",
		"Pink",
		"Purple",
		"Red",
		"Tan",
		"White",
		"Yellow",
	})

	options["clothing sizes"] = a.newOption("Size", "Size")
	a.addBulkOptionItems(options["clothing sizes"], []string{
		`XXS`,
		`XS`,
		`S`,
		`M`,
		`L`,
		`XL`,
		`XXL`,
		`00`,
		`0`,
		`2`,
		`2T`,
		`3T`,
		`3`,
		`4T`,
		`4`,
		`5`,
		`5.5`,
		`6`,
		`6X`,
		`6.5`,
		`7`,
		`7.5`,
		`8`,
		`8.5`,
		`9`,
		`9.5`,
		`10`,
		`10.5`,
		`11`,
		`11.5`,
		`12`,
		`13`,
		`14`,
		`14.5`,
		`15`,
		`15.5`,
		`16`,
		`16.5`,
		`17`,
		`17.5`,
		`18`,
		`18.5`,
		`24`,
		`25`,
		`26`,
		`27`,
		`28`,
		`29`,
		`30`,
		`31`,
		`32`,
		`33`,
		`34`,
		`35`,
		`36`,
		`38`,
		`40`,
		`42`,
		`44`,
		`46`,
		`48`,
		`1X Big`,
		`2X Big`,
		`3X Big`,
		`4X Big`,
		`5X Big`,
		`6X Big`,
		`L Tall`,
		`XL Tall`,
		`2XL Tall`,
		`3XL Tall`,
		`4XL Tall`,
		`2.5 oz`,
		`4.2 oz`,
		`6.0 oz`,
		`20"x20"`,
		`50"x70"`,
		`1X`,
		`2X`,
		`3X`,
		`One Size`,
		`24M`,
	})

	// properties
	properties["long-sleeve"] = a.newProperty("Long Sleeve")
	properties["short-sleeve"] = a.newProperty("Short Sleeve")
	properties["hooded"] = a.newProperty("Hooded")
	properties["capris"] = a.newProperty("Capris")
	properties["pleated"] = a.newProperty("Pleated")
	properties["comfort-fit"] = a.newProperty("Comfort Fit")

	// Classic Fit.
	// All Big sizes have an 8½" inseam. Size Big 2X has 44" waist that expands to 56" and an 11¾" rise.
	// All Tall sizes have a 9½" inseam. Size Tall 2L has 44" waist that expands to 56" and an 11¾" rise.
	// Elastic waistband with an interior drawstring. Zip fly with a buttoned closure.
	// Side on-seam pockets. Two back buttoned pockets.
	// Signature embroidered Pony at the front left hem.
	// 97% cotton, 3% elastane.
	// Machine washable.
	// Imported.

	// make the archetypes
	archetypes["shirts"] = a.newArchetype("Shirts")
	a.addArchetypeOptionRelation(archetypes["shirts"], options["clothing sizes"])
	a.addArchetypeOptionRelation(archetypes["shirts"], options["color"])

	a.addArchetypePropertyRelation(archetypes["shirts"], properties["long-sleeve"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["short-sleeve"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["hooded"])

	archetypes["pants"] = a.newArchetype("Pants")
	a.addArchetypeOptionRelation(archetypes["pants"], options["clothing sizes"])
	a.addArchetypeOptionRelation(archetypes["pants"], options["color"])

	a.addArchetypePropertyRelation(archetypes["shirts"], properties["capris"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["pleated"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["comfort-fit"])

	// a.addArchetypeCategoryRelation(archetypes["pants"], categories["mens pants"])
	// a.addArchetypeCategoryRelation(archetypes["pants"], categories["womens pants"])
	// a.addArchetypeCategoryRelation(archetypes["pants"], categories["kids pants"])

	// a.addArchetypeCategoryRelation(archetypes["shirts"], categories["mens shirts"])
	// a.addArchetypeCategoryRelation(archetypes["shirts"], categories["womens shirts"])
	// a.addArchetypeCategoryRelation(archetypes["shirts"], categories["kids shirts"])
}

type archetype struct{}

func (a *archetype) newArchetype(name string) int64 {
	record := &models.Archetype{
		Name: name,
	}
	if err := record.Insert(context.Background(), db); err != nil {
		panic(err)
	}
	return record.ID
}

func (a *archetype) newOption(name, displayName string) int64 {
	record := &models.Option{
		Name:        name,
		DisplayName: displayName,
	}
	if err := record.Insert(context.Background(), db); err != nil {
		fmt.Println("ERROR:", name, displayName)
		panic(err)
	}
	return record.ID
}

func (a *archetype) newOptionItem(optionID int64, name string, displayName string, sortOrder int) int64 {
	record := &models.OptionItem{
		OptionID:    optionID,
		Name:        name,
		DisplayName: displayName,
		SortOrder:   sortOrder,
	}
	if err := record.Insert(context.Background(), db); err != nil {
		panic(err)
	}
	return record.ID
}

func (a *archetype) addBulkOptionItems(optionID int64, items []string) {
	for sortOrder, item := range items {
		added := a.newOptionItem(optionID, item, item, sortOrder+1)
		optionItems[optionID] = append(optionItems[optionID], added)
		optionItemNames[added] = item
	}
}

func (a *archetype) addArchetypeOptionRelation(archetypeID, optionID int64) {
	record := &models.ArchetypeOption{
		ArchetypeID: archetypeID,
		OptionID:    optionID,
	}
	if err := record.Insert(context.Background(), db); err != nil {
		panic(err)
	}
}

func (a *archetype) newProperty(name string) int64 {
	record := &models.Property{
		Name:        name,
		DisplayName: name,
		FilterParam: strcase.ToSnake(name),
		Fiterable:   true,
	}
	if err := record.Insert(context.Background(), db); err != nil {
		panic(err)
	}
	return record.ID
}

func (a *archetype) addArchetypePropertyRelation(archetypeID, propertyID int64) {
	record := &models.ArchetypeProperty{
		ArchetypeID: archetypeID,
		PropertyID:  propertyID,
	}
	if err := record.Insert(context.Background(), db); err != nil {
		panic(err)
	}
}

func (a *archetype) addArchetypeCategoryRelation(archetypeID, categoryID int64) {
	if archetypeID != 0 && categoryID != 0 {
		record := &models.ArchetypeCategory{
			ArchetypeID: archetypeID,
			CategoryID:  categoryID,
		}
		if err := record.Insert(context.Background(), db); err != nil {
			panic(err)
		}
	}
}
