package repos_test

// import (
// 	"ecomm/internal/env"

// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("Products", func() {

// 	adapter := nativeadapter.NewNativeAdapter()

// 	expectNilErr := func(err error) {
// 		Expect(err).To(BeNil())
// 	}

// 	connect := func() {
// 		env.Load()
// 		adapter.LoadConfig()
// 		_, err := adapter.Connect()
// 		expectNilErr(err)
// 	}

// 	BeforeEach(func() {
// 		connect()
// 	})

// 	Describe("DB Connection", func() {

// 		Describe("GetProductBySku", func() {
// 			It("should return a valid product by sku", func() {
// 				validSku := "gorgeous-granite-chair"
// 				product, err := adapter.GetProductBySku(validSku)
// 				Expect(err).To(BeNil())
// 				Expect(product).To(Not(BeNil()))
// 				defaults := 0
// 				for _, variant := range product.Variants {
// 					if variant.IsDefault {
// 						defaults++
// 						Expect(variant.Sku).To(BeIdenticalTo(validSku))
// 					}
// 				}
// 				Expect(defaults).To(Equal(1))
// 			})
// 		})

// 		Describe("GetProductById", func() {
// 			It("should return a valid product by id", func() {
// 				validId := int64(1)
// 				product, err := adapter.GetProductById(validId)
// 				Expect(err).To(BeNil())
// 				Expect(product).To(Not(BeNil()))
// 				Expect(product.Id).To(BeIdenticalTo(validId))
// 				defaults := 0
// 				for _, variant := range product.Variants {
// 					if variant.IsDefault {
// 						defaults++
// 					}
// 				}
// 				Expect(defaults).To(Equal(1))
// 			})
// 		})

// 		Describe("GetProductsByCategoryId", func() {
// 			It("should return all valid products by category_id", func() {
// 				validCategoryId := int64(2)
// 				_, err := adapter.GetProductsByCategoryId(validCategoryId)
// 				Expect(err).To(BeNil())
// 				// Expect(product).To(Not(BeNil()))
// 				// Expect(product).To(BeIdenticalTo(validCategoryId))
// 			})
// 		})

// 		Describe("GetAllProducts", func() {

// 		})

// 		Describe("CreateProduct", func() {

// 		})

// 		Describe("UpdateProduct", func() {

// 		})

// 		Describe("DeleteProduct", func() {

// 		})
// 	})

// })
