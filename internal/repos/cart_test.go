package repos_test

import (
	"github.com/k8scommerce/k8scommerce/internal/models"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cart", func() {
	defer GinkgoRecover()

	var cartId uuid.UUID
	// var cartSessionId
	// cartSessionId = session.NewSessionId()
	testCart := &models.Cart{
		StoreID: 1,
	}
	Expect(repo).ToNot(BeNil())

	truncateCart := func() {
		_, err := repo.GetRawDB().Exec(`TRUNCATE cart RESTART IDENTITY CASCADE;`)
		Expect(err).To(BeNil())
	}
	truncateCartItem := func() {
		_, err := repo.GetRawDB().Exec(`TRUNCATE cart_item RESTART IDENTITY CASCADE;`)
		Expect(err).To(BeNil())
	}

	// BeforeEach(func() {
	// 	truncateCartItem()
	// 	truncateCart()
	// })

	Describe("Build Cart With Customer & Items", func() {
		truncateCartItem()
		truncateCart()

		Describe("Create", func() {
			It("should create a cart", func() {
				err := repo.Cart().Create(testCart)
				Expect(err).To(BeNil())
				Expect(testCart).ToNot(BeNil())
				Expect(testCart.ID).ToNot(BeNil())
				cartId = testCart.ID
				Expect(cartId).ToNot(BeNil())

			})
		})

		Describe("GetById", func() {
			It("should get the cart by id", func() {
				// get the cart
				crt, err := repo.Cart().GetByCartId(cartId)
				Expect(err).To(BeNil())
				Expect(crt).ToNot(BeNil())

				// // get the cart items
				// items, err := repo.CartItem().GetByCartID(cartId)
				// Expect(err).To(BeNil())
				// Expect(items).ToNot(BeNil())
			})
		})

		Describe("Attach Customer", func() {
			It("should attach a customer to the cart by customer id", func() {
				crt, err := repo.Cart().GetByCartId(cartId)
				Expect(err).To(BeNil())
				Expect(crt).ToNot(BeNil())

			})
		})
	})
})
