package repos_test

import (
	"database/sql"
	"fmt"

	"github.com/k8scommerce/k8scommerce/internal/models"
	"github.com/k8scommerce/k8scommerce/internal/repos"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Customer", func() {
	defer GinkgoRecover()

	var (
		storeId  = int64(1)
		email    = "test@k8scommerce.com"
		password = "test123"
	)

	deleteCustomer := func() {
		_, err := repo.GetRawDB().Exec(`delete from customer where email = $1`, email)
		Expect(err).To(BeNil())
	}

	createCustomer := func() (models.Customer, error) {
		customer := models.Customer{
			StoreID:   storeId,
			FirstName: "Test",
			LastName:  "Customer",
			Email:     email,
			Password:  sql.NullString{String: password, Valid: true},
		}

		err := repo.Customer().Create(&customer)
		Expect(err).To(BeNil())
		return customer, err
	}

	Describe("CreateCustomer", func() {
		BeforeEach(func() {
			deleteCustomer()
		})

		It("should create a customer", func() {
			customer, err := createCustomer()
			Expect(err).To(BeNil())
			Expect(customer).ToNot(BeNil())
			Expect(customer.Email).ToNot(BeNil())
		})
	})

	Describe("GetCustomerByEmail", func() {
		BeforeEach(func() {
			deleteCustomer()
			createCustomer()
		})

		It("should get a customer by email", func() {
			email := email
			Expect(repo).ToNot(BeNil())
			result, err := repo.Customer().GetCustomerByEmail(storeId, email)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})

		It("should fail geting a customer by email", func() {
			email := "fail@example.com"
			Expect(repo).ToNot(BeNil())
			result, err := repo.Customer().GetCustomerByEmail(storeId, email)
			Expect(err).ToNot(BeNil())
			Expect(result).To(BeNil())
		})
	})

	Describe("Login", func() {
		BeforeEach(func() {
			deleteCustomer()
			createCustomer()
		})

		It("should authenticate a valid user", func() {
			Expect(repo).ToNot(BeNil())
			result, err := repo.Customer().Login(storeId, email, password)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})

		It("should not authenticate an invalid user", func() {
			Expect(repo).ToNot(BeNil())
			password = "abc@example.com"
			result, err := repo.Customer().Login(storeId, email, password)
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(Equal(fmt.Sprintf("status %d: error %v", repos.CustomerLoginErrorCode, repos.CustomerLoginError)))
			Expect(result).To(BeNil())
		})
	})
})
