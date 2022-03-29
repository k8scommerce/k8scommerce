package main

import (
	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/icrowley/fake"
)

func createUsers() {
	// create super user
	err := repo.User().Create(&models.User{
		FirstName: "Admin",
		LastName:  "User",
		Email:     "admin@k8scommerce.com",
		Password:  "admin",
	})
	if err != nil {
		panic(err)
	}

	// create other users
	for i := 0; i < 30; i++ {

		err = repo.User().Create(&models.User{
			FirstName: fake.FirstName(),
			LastName:  fake.LastName(),
			Email:     fake.EmailAddress(),
			Password:  fake.Password(6, 14, true, true, true),
		})
		if err != nil {
			panic(err)
		}
	}
}
