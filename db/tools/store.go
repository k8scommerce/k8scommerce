package main

import (
	"context"
	"fmt"

	"ecomm/internal/repos/models"

	"syreclabs.com/go/faker"
)

func createStore() {

	state := faker.Address().StateAbbr()

	address := fmt.Sprintf("%s\n%s, %s %s",
		faker.Address().StreetAddress(),
		faker.Address().City(),
		state,
		faker.Address().ZipCodeByState(state),
	)

	store := &models.Store{
		Name:                       "Demo Store",
		Description:                toNullString("My Demo Store"),
		Address:                    toNullString(address),
		MailFromAddress:            toNullString(address),
		URL:                        "http://localhost:4500/",
		IsDefault:                  true,
		DefaultCountryID:           int64(1),
		DefaultLocale:              "America/Denver",
		DefaultCurrency:            "USD",
		ContactPhone:               toNullString(faker.PhoneNumber().PhoneNumber()),
		CustomerSupportEmail:       toNullString(faker.Internet().Email()),
		NewOrderNotificationsEmail: toNullString(faker.Internet().Email()),
	}
	if err := store.Insert(context.Background(), db); err != nil {
		panic(err)
	}
	storeID = store.ID
}
