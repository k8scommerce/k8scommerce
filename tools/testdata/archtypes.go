package main

import (
	"context"
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
		"Antique Cherry Red",
		"Antique Heliconia",
		"Antique Irish Green",
		"Antique Jade Dome",
		"Antique Orange",
		"Antique Royal",
		"Antique Sapphire",
		"Ash",
		"Ash Grey",
		"Azalea",
		"Berry",
		"Black",
		"Blackberry",
		"Blue Dusk",
		"Brown Savana",
		"Cardinal Red",
		"Carolina Blue",
		"Charcoal",
		"Cherry Red",
		"Chestnut",
		"Cobalt",
		"Coral Silk",
		"Cornsilk",
		"Daisy",
		"Dark Chocolate",
		"Dark Heather",
		"Electric Green",
		"Forest Green",
		"Galapagos Blue",
		"Garnet",
		"Gold",
		"Gravel",
		"Heather Cardinal",
		"Heather Indigo",
		"Heather Irish Green",
		"Heather Military Green",
		"Heather Navy",
		"Heather Orange",
		"Heather Purple",
		"Heather Red",
		"Heather Royal",
		"Heather Sapphire",
		"Heliconia",
		"Honey",
		"Ice Grey",
		"Indigo Blue",
		"Iris",
		"Irish Green",
		"Jade Dome",
		"Kelly Green",
		"Kiwi",
		"Light Blue",
		"Lilac",
		"Lime",
		"Marbled Charcoal",
		"Marbled Galapagos Blue",
		"Marbled Heliconia",
		"Marbled Navy",
		"Marbled Royal",
		"Maroon",
		"Meadow",
		"Metro Blue",
		"Midnight",
		"Military Green",
		"Mint Green",
		"Moss",
		"Natural",
		"Navy",
		"Neon Blue",
		"Neon Green",
		"Old Gold",
		"Olive",
		"Orange",
		"Orchid",
		"PFD (Prepared for dye)",
		"Pistachio",
		"Prairie Dust",
		"Purple",
		"Red",
		"Royal",
		"RS Sport Grey",
		"Russet",
		"Rusty Bronze",
		"Safety Green",
		"Safety Orange",
		"Safety Pink",
		"Sand",
		"Sapphire",
		"Sky",
		"Sport Grey",
		"Stone Blue",
		"Sunset",
		"Tan",
		"Tangerine",
		"Tennessee Orange",
		"Texas Orange",
		"Tropical Blue",
		"Turf Green",
		"Tweed",
		"Vegas Gold",
		"Violet",
		"White",
		"Yellow Haze",
	})

	options["clothing sizes"] = a.newOption("Clothing Sizes", "Clothing Sizes")
	a.addBulkOptionItems(options["clothing sizes"], []string{
		"Petite",
		"Small",
		"Medium",
		"Large",
		"X-Large",
		"XX-Large",
		"XXX-Large",
	})

	options["waist size"] = a.newOption("Waist Size", "Waist Size")
	a.addBulkOptionItems(options["waist size"], []string{
		"US 26 / Euro42",
		"US 27 / Euro 43",
		"US 28 / Euro 44",
		"US 29 / Euro 45",
		"US 30 / Euro 46",
		"US 31 / Euro 47",
		"US 32 / Euro 48",
		"US 33 / Euro 49",
		"US 34 / Euro 50",
		"US 35 / Euro 51",
		"US 36 / Euro 52",
		"US 38 / Euro 54",
		"US 40 / Euro 56",
		"US 42 / Euro 58",
		"US 44 / Euro 60",
	})

	// properties
	properties["long-sleeve"] = a.newProperty("Long Sleeve")
	properties["short-sleeve"] = a.newProperty("Short Sleeve")
	properties["hooded"] = a.newProperty("Hooded")
	properties["capris"] = a.newProperty("Capris")
	properties["pleated"] = a.newProperty("Pleated")
	properties["comfort-fit"] = a.newProperty("Comfort Fit")

	// make the archetypes
	archetypes["shirts"] = a.newArchetype("Shirts")
	a.addArchetypeOptionRelation(archetypes["shirts"], options["clothing sizes"])
	a.addArchetypeOptionRelation(archetypes["shirts"], options["color"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["long-sleeve"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["short-sleeve"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["hooded"])

	archetypes["pants"] = a.newArchetype("Pants")
	a.addArchetypeOptionRelation(archetypes["pants"], options["waist size"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["capris"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["pleated"])
	a.addArchetypePropertyRelation(archetypes["shirts"], properties["comfort-fit"])

	a.addArchetypeCategoryRelation(archetypes["pants"], categories["mens pants"])
	a.addArchetypeCategoryRelation(archetypes["pants"], categories["womens pants"])
	a.addArchetypeCategoryRelation(archetypes["pants"], categories["kids pants"])

	a.addArchetypeCategoryRelation(archetypes["shirts"], categories["mens shirts"])
	a.addArchetypeCategoryRelation(archetypes["shirts"], categories["womens shirts"])
	a.addArchetypeCategoryRelation(archetypes["shirts"], categories["kids shirts"])
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
