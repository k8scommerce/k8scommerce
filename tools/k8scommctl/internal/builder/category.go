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

	"github.com/k8scommerce/k8scommerce/services/rpc/catalog/pb/catalog"

	"gopkg.in/yaml.v2"
)

type categoryBuilder struct {
	apiURL    string
	storeKey  string
	outputDir string
	baseName  string
}

func NewCategoryBuilder(apiURL, storeKey, outputDir, baseName string) Builder {
	return &categoryBuilder{
		apiURL:    apiURL,
		storeKey:  storeKey,
		outputDir: outputDir,
		baseName:  baseName,
	}
}

func (b *categoryBuilder) Build() {
	categories := b.getCategories()
	for _, category := range categories {
		dir := path.Clean(fmt.Sprintf("%s/%s/%s", b.outputDir, b.baseName, category.Slug))
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
			os.Exit(1)
		}
		// write the file
		b.write(dir, "_index.md", category)
	}
}

func (b *categoryBuilder) write(dir, filename string, category *catalog.Category) {
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
	if err := enc.Encode(category); err != nil {
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

func (b *categoryBuilder) getCategories() []*catalog.Category {
	var categories []*catalog.Category
	url := b.apiURL + "/v1/categories"

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
		os.Exit(1)
		return nil
	}
	req.Header.Set("Store-Key", b.storeKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nerror: Error when sending request to the server: %s\n\n", err.Error())
		os.Exit(1)
		return nil
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
		os.Exit(1)
	}

	if resp.StatusCode == http.StatusOK {
		var getAllCategoriesResponse catalog.GetAllCategoriesResponse
		err = json.Unmarshal(responseBody, &getAllCategoriesResponse)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nerror: %s\n\n", err.Error())
			os.Exit(1)
		}

		categories = append(categories, getAllCategoriesResponse.Categories...)

	} else {
		fmt.Fprintf(os.Stderr, "\n\nerror: %s\n%s\n", resp.Status, string(responseBody))
		os.Exit(1)
	}
	return categories
}
