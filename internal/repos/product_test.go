package repos_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product", func() {
	defer GinkgoRecover()

	Describe("GetProductBySlug", func() {

		It("should get a product by a slug", func() {
			storeId := int64(1)
			slug := "antelopealpine-wizardarrow-rabbitmalachite"
			Expect(repo).ToNot(BeNil())
			result, err := repo.Product().GetProductBySlug(storeId, slug)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
	})

	Describe("GetProductById", func() {

		It("should get a product by id", func() {
			id := int64(1)
			Expect(repo).ToNot(BeNil())
			result, err := repo.Product().GetProductById(id)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			fmt.Printf("%#v", result)
		})
	})

	Describe("GetProductBySku", func() {

		It("should get a product by sku", func() {
			storeId := int64(1)
			sku := "antelopealpine-wizardarrow-rabbitmalachite"
			Expect(repo).ToNot(BeNil())
			result, err := repo.Product().GetProductBySku(storeId, sku)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
	})

	Describe("GetAllProducts", func() {

		It("should get all products", func() {
			storeId := int64(1)
			currentPage := int64(0)
			pageSize := int64(1000)
			filter := ""

			Expect(repo).ToNot(BeNil())
			result, err := repo.Product().GetAllProducts(storeId, currentPage, pageSize, filter)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
	})
})
