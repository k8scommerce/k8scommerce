package main

import (
	"context"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func createStore() {

	// state := faker.Address().StateAbbr()

	// address := fmt.Sprintf("%s\n%s, %s %s",
	// 	faker.Address().StreetAddress(),
	// 	faker.Address().City(),
	// 	state,
	// 	faker.Address().ZipCodeByState(state),
	// )

	store := &models.Store{
		Name:        "Demo Store",
		Description: toNullString("My Demo Store"),
		URL:         "http://localhost:4200/",
		IsDefault:   true,
	}
	if err := store.Insert(context.Background(), db); err != nil {
		panic(err)
	}
	storeID = store.ID
}
