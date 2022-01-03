package repos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8scommerce/services/rpc/similarproducts/internal/repos"
)

var _ = Describe("Products", func() {
	expectNilErr := func(err error) {
		Expect(err).To(BeNil())
	}

	Describe("Helpers", func() {

		Describe("BuildOrderBy", func() {
			It("should return a valid order by string", func() {
				expected := "ORDER BY product.name ASC, product.color DESC, price.amount DESC"

				on := "name,-color,-amount"

				orderBy, err := repos.BuildOrderBy(on, map[string]string{
					"name":   "product",
					"color":  "product",
					"amount": "price",
				})

				expectNilErr(err)
				Expect(orderBy).To(Equal(expected))
			})
		})
	})
})
