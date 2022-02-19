package buildsql_test

import (
	"fmt"
	"k8scommerce/internal/buildsql"
	"k8scommerce/internal/models"
	"log"
	"net/url"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	defer GinkgoRecover()

	expectNilErr := func(err error) {
		Expect(err).To(BeNil())
	}

	expectErr := func(err error) {
		Expect(err).To(Not(BeNil()))
	}

	Describe("BuildWhere", func() {

		// getFields := func() map[string]string {
		// 	return map[string]string{
		// 		"id":     "p",  // product alias
		// 		"name":   "p",  // product alias
		// 		"slug":   "p",  // product alias
		// 		"sku":    "v",  // product alias
		// 		"amount": "pr", // price alias
		// 	}
		// }

		// on := "filter=p-name-Practical Cotton Gloves&filter=v-sku-practical-cotton-gloves&sortOn=p-name&sortOn=-pr-amount"

		var builder buildsql.QueryBuilder

		// var removeItemToCartLogic *logic.RemoveItemInCartLogic
		BeforeEach(func() {
			builder = buildsql.QueryBuilder{}
		})

		It("should return valid AllowedFilterFields from a string map", func() {
			// Expect(addItemToCartLogic).To(Not(BeNil()))
			builder.AllowedFilterFields = map[string]string{
				"id":     "p",  // product alias
				"name":   "p",  // product alias
				"slug":   "p",  // product alias
				"sku":    "v",  // product alias
				"amount": "pr", // price alias
			}

			// where, err := buildsql.BuildWhere(on)
			// expectNilErr(err)
			// fmt.Println("where:", where)
		})

		It("should correctly parse a param string", func() {

			on := "filter=p-name-Practical Cotton Gloves&filter=v-sku-practical-cotton-gloves&sortOn=p-name&sortOn=-pr-amount"

			err := builder.ParseParamString(on)
			expectNilErr(err)

			Expect(len(builder.Filters)).To(Equal(2))
			Expect(builder.Filters[0].TableAlias).To(Equal("p"))
			Expect(builder.Filters[0].FieldName).To(Equal("name"))
			Expect(builder.Filters[0].Value).To(Equal("Practical Cotton Gloves"))

			Expect(builder.Filters[1].TableAlias).To(Equal("v"))
			Expect(builder.Filters[1].FieldName).To(Equal("sku"))
			Expect(builder.Filters[1].Value).To(Equal("practical-cotton-gloves"))

			Expect(len(builder.Sorts)).To(Equal(2))
			Expect(builder.Sorts[0].TableAlias).To(Equal("p"))
			Expect(builder.Sorts[0].FieldName).To(Equal("name"))

			Expect(builder.Sorts[1].TableAlias).To(Equal("pr"))
			Expect(builder.Sorts[1].FieldName).To(Equal("amount"))
		})

		It("should error on parsing an invalid param string", func() {

			on := "filter=p_name_Practical Cotton Gloves&filter=v_sku_practical-cotton-gloves&sortOn=p_name&sortOn=-pr_amount"

			err := builder.ParseParamString(on)
			expectErr(err)
		})

		It("should correctly build ", func() {

			on := "filter=p-name-like-Practical&filter=p-name-nlike-Cotton&filter=v-sku-eq-practical-cotton-gloves&sortOn=p-name&sortOn=-pr-amount"

			where, orderBy, namedParamMap, err := builder.Build(on, map[string]interface{}{
				"p":  models.Product{}, // product alias
				"v":  models.Variant{}, // product alias
				"pr": models.Price{},   // product alias
			})
			fmt.Println("")
			fmt.Println("where:", where)
			fmt.Println("namedParamMap:", namedParamMap)
			fmt.Println("orderBy:", orderBy)
			// Expect(len(builder.Filters)).To(Equal(2))

			// Expect(addItemToCartLogic).To(Not(BeNil()))

			// where, err := buildsql.BuildWhere(on)
			expectNilErr(err)
			// fmt.Println("where:", where)
		})

		It("should parse raw url ", func() {

			// on := "/v1/products/0/20?filter=p-name-like-practical&sortOn=p-id"
			on := "http://api.local.k8sly.com:8888/v1/products/0/20?filter=p-name-like-cotton&sortOn=p-id"

			on = strings.Replace(on, "\\u0026", "&", -1)

			decodedValue, err := url.QueryUnescape(on)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(decodedValue)

			where, orderBy, namedParamMap, err := builder.Build(on, map[string]interface{}{
				"p":  models.Product{}, // product alias
				"v":  models.Variant{}, // product alias
				"pr": models.Price{},   // product alias
			})
			fmt.Println("")
			fmt.Println("where:", where)
			fmt.Println("namedParamMap:", namedParamMap)
			fmt.Println("orderBy:", orderBy)
			// Expect(len(builder.Filters)).To(Equal(2))

			// Expect(addItemToCartLogic).To(Not(BeNil()))

			// where, err := buildsql.BuildWhere(on)
			expectNilErr(err)
			// fmt.Println("where:", where)
		})
	})
})
