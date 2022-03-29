/*
Copyright © 2022 K8sCommerce
*/
package builder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"gopkg.in/yaml.v2"
)

type productBuilder struct {
	apiURL    string
	storeKey  string
	outputDir string
	baseName  string
}

func NewProductBuilder(apiURL, storeKey, outputDir, baseName string) Builder {
	return &productBuilder{
		apiURL:    apiURL,
		storeKey:  storeKey,
		outputDir: outputDir,
		baseName:  baseName,
	}
}

func (b *productBuilder) Build() {
	for _, slug := range b.getProductSlugs() {
		dir := path.Clean(fmt.Sprintf("%s/%s", b.outputDir, b.baseName))
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
			os.Exit(1)
		}

		prod := b.getProduct(slug)
		if prod != nil {
			// write the file
			b.write(dir, fmt.Sprintf("%s.md", slug), prod)
		} else {
			log.Printf("Could not write product: %s", slug)
		}
	}
}

func (b *productBuilder) write(dir, filename string, product *Product) {

	mdFile := path.Clean(fmt.Sprintf("%s/%s", dir, filename))

	if _, err := os.Stat(mdFile); !errors.Is(err, os.ErrNotExist) {
		if err := os.Truncate(mdFile, 0); err != nil {
			log.Printf("Failed to truncate page: %s %v", mdFile, err)
		}
	}

	// Create a file for writing
	f, err := os.Create(mdFile)
	if err != nil {
		// failed to create/open the file
		fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
		os.Exit(1)
	}
	f.WriteString("---\n")

	enc := yaml.NewEncoder(f)
	// enc.SetIndent("", "    ")
	if err := enc.Encode(product); err != nil {
		// failed to encode
		fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
		os.Exit(1)
	}
	f.WriteString("---")
	if err := f.Close(); err != nil {
		// failed to close the file
		fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
		os.Exit(1)
	}
}

func (b *productBuilder) getProduct(slug string) *Product {
	var product *Product
	url := fmt.Sprintf("%s/v1/product/slug/%s", b.apiURL, slug)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic("ERROR:" + err.Error())
	}
	req.Header.Set("Store-Key", b.storeKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
		return nil
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
		os.Exit(1)
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(responseBody, &product)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
			os.Exit(1)
		}

		product.Title = product.Name

		t, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
		product.Date = t
	} else {
		fmt.Fprintf(os.Stderr, "\n\nerror: %s\n%s\n", resp.Status, string(responseBody))
		os.Exit(1)
	}
	return product
}

func (b *productBuilder) getProductSlugs() []string {
	var slugs []string
	baseUrl := b.apiURL + "/v1/products/%d/%d"

	client := &http.Client{}
	totalPages := int64(1)
	productsPerPage := int64(1000)
	for i := int64(0); i < totalPages; i++ {
		url := fmt.Sprintf(baseUrl, i, productsPerPage)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
			os.Exit(1)
		}
		req.Header.Set("Store-Key", b.storeKey)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error when sending request to the server")
			return nil
		}
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
			os.Exit(1)
		}

		if resp.StatusCode == http.StatusOK {
			var getAllProductsResponse catalog.GetAllProductsResponse
			err = json.Unmarshal(responseBody, &getAllProductsResponse)
			if err != nil {
				fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
				os.Exit(1)
			}

			// adjust the totalPages based on the response
			totalPages = getAllProductsResponse.TotalPages

			// append the products
			for _, p := range getAllProductsResponse.Products {
				slugs = append(slugs, p.Slug)
			}
		} else {
			fmt.Fprintf(os.Stderr, "\n\nerror: %s\n%s\n", resp.Status, string(responseBody))
			os.Exit(1)
		}
	}

	return slugs
}

// type GetAllProductsResponse struct {
// 	Products     []Product `json:"products" yaml:"products"`
// 	TotalRecords int64     `json:"total_records" yaml:"total_records"`
// 	TotalPages   int64     `json:"total_pages" yaml:"total_pages"`
// }

// type CategoryPair struct {
// 	Slug string `json:"slug" yaml:"slug"` // product slug
// 	Name string `json:"name" yaml:"name"` // product name
// }

type Product struct {
	catalog.Product
	Title string    `json:"title" yaml:"title"` // product name
	Date  time.Time `json:"date" yaml:"date"`
	Tags  []string  `json:"tags" yaml:"tags"`
}

// type Variant struct {
// 	// Id        int64   `json:"id" yaml:"id"`               // variant id
// 	IsDefault bool    `json:"is_default" yaml:"is_default"` // is default variant. each product must have exactly 1 default variant
// 	Sku       string  `json:"sku" yaml:"sku"`               // variant sku. usually the product sku with appended identification tags
// 	Weight    float64 `json:"weight" yaml:"weight"`         // variant weight. used in calculating shipping
// 	Height    float64 `json:"height" yaml:"height"`         // variant height. used in calculating shipping
// 	Width     float64 `json:"width" yaml:"width"`           // variant width. used in calculating shipping
// 	Depth     float64 `json:"depth" yaml:"depth"`           // variant depth. used in calculating shipping
// 	Price     Price   `json:"price" yaml:"price"`           // variant Price object
// }

// type Price struct {
// 	// Id                     int64   `json:"id" yaml:"id"`                                         // price id
// 	Amount               float64 `json:"amount" yaml:"amount"`                                 // price amount
// 	SalePrice            float64 `json:"sale_price" yaml:"sale_price"`                         // sale price
// 	FormattedSalePrice   string  `json:"formatted_sale_price" yaml:"formatted_sale_price"`     // formatted sale price
// 	RetailPrice          float64 `json:"retail_price" yaml:"retail_price"`                     // retail price
// 	FormattedRetailPrice string  `json:"formatted_retail_price" yaml:"formatted_retail_price"` // formatted retail price
// 	Currency             string  `json:"currency" yaml:"currency"`                             // price currency. example: USD, CAN, etc.
// }

// type Asset struct {
// 	// Id int64 `json:"id" yaml:"id"` // asset id
// 	// ProductId   int64             `json:"product_id" yaml:"product_id"`     // product id
// 	VariantId   int64             `json:"variant_id" yaml:"variant_id"`     // variant id
// 	Name        string            `json:"name" yaml:"name"`                 // asset name
// 	DisplayName string            `json:"display_name" yaml:"display_name"` // display name
// 	Url         string            `json:"url" yaml:"url"`                   // full, public access url
// 	Kind        int               `json:"kind" yaml:"kind"`                 // asset kind (0=unknown|1=image|2=document|3=audio|4=video|5=archive)
// 	ContentType string            `json:"content_type" yaml:"content_type"` // content type (mime type)
// 	SortOrder   int64             `json:"sort_order" yaml:"sort_order"`     // sort order
// 	Sizes       map[string]string `json:"sizes" yaml:"sizes"`               // map[tag:string]url:string
// }
