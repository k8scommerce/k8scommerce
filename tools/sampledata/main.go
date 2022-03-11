package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

var (
	storeID int64 = 1

	ids = &IDs{}
	sos = &SortOrders{}
)

type CSVProductRecord struct {
	Id                   string `csv:"ID"`
	Type                 string `csv:"Type"`
	Sku                  string `csv:"SKU"`
	Name                 string `csv:"Name"`
	Published            string `csv:"Published"`
	IsFeatured           string `csv:"Is featured?"`
	VisibilityInCatalog  string `csv:"Visibility in catalog"`
	ShortDescription     string `csv:"Short description"`
	Description          string `csv:"description"`
	DateSalePriceStarts  string `csv:"Date sale price starts"`
	DateSalePriceEnds    string `csv:"Date sale price ends"`
	TaxStatus            string `csv:"Tax status"`
	TaxClass             string `csv:"Tax class"`
	InStock              string `csv:"In stock?"`
	Stock                string `csv:"Stock"`
	BackordersAllowed    string `csv:"Backorders allowed?"`
	SoldIndividually     string `csv:"Sold individually?"`
	WeightLbs            string `csv:"Weight (lbs)"`
	LengthIn             string `csv:"Length (in)"`
	WidthIn              string `csv:"Width (in)"`
	HeightIn             string `csv:"Height (in)"`
	AllowCustomerReviews string `csv:"Allow customer reviews?"`
	PurchaseNote         string `csv:"Purchase note"`
	SalePrice            string `csv:"Sale price"`
	RegularPrice         string `csv:"Regular price"`
	Categories           string `csv:"Categories"`
	Tags                 string `csv:"Tags"`
	ShippingClass        string `csv:"Shipping class"`
	Images               string `csv:"Images"`
	DownloadLimit        string `csv:"Download limit"`
	DownloadExpiryDays   string `csv:"Download expiry days"`
	Parent               string `csv:"Parent"`
	GroupedProducts      string `csv:"Grouped products"`
	Upsells              string `csv:"Upsells"`
	CrossSells           string `csv:"Cross-sells"`
	ExternalUrl          string `csv:"External URL"`
	ButtonText           string `csv:"Button text"`
	Position             string `csv:"Position"`
	Attribute_1Name      string `csv:"Attribute 1 name"`
	Attribute_1ValueS    string `csv:"Attribute 1 value(s)"`
	Attribute_2Name      string `csv:"Attribute 2 name"`
	Attribute_2ValueS    string `csv:"Attribute 2 value(s)"`
	Attribute_3Name      string `csv:"Attribute 3 name"`
	Attribute_3ValueS    string `csv:"Attribute 3 value(s)"`
	Attribute_4Name      string `csv:"Attribute 4 name"`
	Attribute_4ValueS    string `csv:"Attribute 4 value(s)"`
	Attribute_5Name      string `csv:"Attribute 5 name"`
	Attribute_5ValueS    string `csv:"Attribute 5 value(s)"`
	MetaWpcomIsMarkdown  string `csv:"Meta: _wpcom_is_markdown"`
	Download_1Name       string `csv:"Download 1 name"`
	Download_1Url        string `csv:"Download 1 URL"`
	Download_2Name       string `csv:"Download 2 name"`
	Download_2Url        string `csv:"Download 2 URL"`
}

func main() {

	tables.Truncate()

	productsFile, err := os.OpenFile("products.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer productsFile.Close()

	products := []*CSVProductRecord{}

	if err := gocsv.UnmarshalFile(productsFile, &products); err != nil { // Load products from file
		panic(err)
	}

	for _, p := range products {

		// categories = append(categories, p.Categories)
		categories.Parse(p.Categories)

		// fmt.Println(
		// 	p.Id,
		// 	p.Type,
		// 	p.Sku,
		// 	p.Name,
		// 	p.Published,
		// 	p.IsFeatured,
		// 	p.VisibilityInCatalog,
		// 	p.ShortDescription,
		// 	p.Description,
		// 	p.DateSalePriceStarts,
		// 	p.DateSalePriceEnds,
		// 	p.TaxStatus,
		// 	p.TaxClass,
		// 	p.InStock,
		// 	p.Stock,
		// 	p.BackordersAllowed,
		// 	p.SoldIndividually,
		// 	p.WeightLbs,
		// 	p.LengthIn,
		// 	p.WidthIn,
		// 	p.HeightIn,
		// 	p.AllowCustomerReviews,
		// 	p.PurchaseNote,
		// 	p.SalePrice,
		// 	p.RegularPrice,
		// 	p.Categories,
		// 	p.Tags,
		// 	p.ShippingClass,
		// 	p.Images,
		// 	p.DownloadLimit,
		// 	p.DownloadExpiryDays,
		// 	p.Parent,
		// 	p.GroupedProducts,
		// 	p.Upsells,
		// 	p.CrossSells,
		// 	p.ExternalUrl,
		// 	p.ButtonText,
		// 	p.Position,
		// 	p.Attribute_1Name,
		// 	p.Attribute_1ValueS,
		// 	p.Attribute_2Name,
		// 	p.Attribute_2ValueS,
		// 	p.Attribute_3Name,
		// 	p.Attribute_3ValueS,
		// 	p.Attribute_4Name,
		// 	p.Attribute_4ValueS,
		// 	p.Attribute_5Name,
		// 	p.Attribute_5ValueS,
		// 	p.MetaWpcomIsMarkdown,
		// 	p.Download_1Name,
		// 	p.Download_1Url,
		// 	p.Download_2Name,
		// 	p.Download_2Url,
		// )
	}

	// fmt.Println(categories)

	// fmt.Println(catmap)

	for path, total := range sos.Category {
		fmt.Println(path, total)
	}

	categories.Save()

}
