// Code generated by goctl. DO NOT EDIT.
package types

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"`
}

type ResponseStatus struct {
	StatusCode    int64  `json:"statusCode"`              // RFC http status code, ie. 204, etc - https://go.dev/src/net/http/status.go
	StatusMessage string `json:"statusMessage,omitempty"` // status message
}

type Category struct {
	Id              int64  `json:"id"`                        // category id
	ParentId        int64  `json:"parentId"`                  // parent category id. references Category.Id
	Slug            string `json:"slug"`                      // slug name of the category
	Name            string `json:"name"`                      // name of category
	Description     string `json:"description"`               // description of category
	MetaTitle       string `json:"metaTitle,omitempty"`       // metatag title for SEO
	MetaDescription string `json:"metaDescription,omitempty"` // metatag description for SEO
	MetaKeywords    string `json:"metaKeywords,omitempty"`    // metatag keywords for SEO
	SortOrder       int32  `json:"sortOrder"`                 // sort order of menu items on the same level and same parent id
}

type Product struct {
	Id               int64     `json:"id"`                        // product id
	Slug             string    `json:"slug"`                      // product slug
	Name             string    `json:"name"`                      // product name
	ShortDescription string    `json:"shortDescription"`          // product short description. used in category pages
	Description      string    `json:"description"`               // category description
	MetaTitle        string    `json:"metaTitle,omitempty"`       // metatag title for SEO
	MetaDescription  string    `json:"metaDescription,omitempty"` // metatag description for SEO
	MetaKeywords     string    `json:"metaKeywords,omitempty"`    // metatag keywords for SEO
	Variants         []Variant `json:"variants,omitempty"`        // collection of Variant objects
}

type Variant struct {
	Id        int64   `json:"id,omitempty"`        // variant id
	IsDefault bool    `json:"isDefault,omitempty"` // is default variant. each product must have exactly 1 default variant
	Sku       string  `json:"sku,omitempty"`       // variant sku. usually the product sku with appended identification tags
	Weight    float64 `json:"weight,omitempty"`    // variant weight. used in calculating shipping
	Height    float64 `json:"height,omitempty"`    // variant height. used in calculating shipping
	Width     float64 `json:"width,omitempty"`     // variant width. used in calculating shipping
	Depth     float64 `json:"depth,omitempty"`     // variant depth. used in calculating shipping
	Price     Price   `json:"price,omitempty"`     // variant Price object
}

type Price struct {
	Id                     int64   `json:"id,omitempty"`                     // price id
	Amount                 float64 `json:"amount,omitempty"`                 // price amount
	DisplayAmount          string  `json:"displayAmount,omitempty"`          // price display amount
	CompareAtAmount        float64 `json:"compareAtAmount,omitempty"`        // price compare amount
	DisplayCompareAtAmount string  `json:"displayCompareAtAmount,omitempty"` // price display compare amount
	Currency               string  `json:"currency,omitempty"`               // price currency. example: USD, CAN, etc.
}

type GetAllCategoriesResponse struct {
	Categories     []Category     `json:"categories,omitempty"` // a collection of Category
	ResponseStatus ResponseStatus `json:"status"`               // a ResponseStatus object
}

type GetCategoryBySlugRequest struct {
	Slug string `path:"slug"` // slug name of the category
}

type GetCategoryBySlugResponse struct {
	Category       Category       `json:"category,omitempty"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type GetCategoryByIdRequest struct {
	Id int64 `json:"id,omitempty"`
}

type GetCategoryByIdResponse struct {
	Category       Category       `json:"category,omitempty"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type CreateCategoryRequest struct {
	Category Category `json:"category,omitempty"`
}

type CreateCategoryResponse struct {
	Category       Category       `json:"category,omitempty"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type UpdateCategoryRequest struct {
	Id       int64    `json:"id,omitempty"`
	Category Category `json:"category,omitempty"`
}

type UpdateCategoryResponse struct {
	Category       Category       `json:"category,omitempty"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type DeleteCategoryRequest struct {
	Id int64 `json:"id,omitempty"`
}

type DeleteCategoryResponse struct {
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type GetProductBySkuRequest struct {
	Sku string `path:"sku"`
}

type GetProductBySlugRequest struct {
	Slug string `path:"slug"` // slug name of the category
}

type GetProductResponse struct {
	Product        Product        `json:"product"` // slug name of the category
	ResponseStatus ResponseStatus `json:"status"`  // a ResponseStatus object
}

type GetProductByIdRequest struct {
	Id int64 `path:"id"`
}

type GetProductsByCategoryIdRequest struct {
	CategoryId  int64  `path:"categoryId"`
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	SortOn      string `form:"sortOn,optional"`
}

type GetProductsByCategoryIdResponse struct {
	Products       []Product      `json:"products"`
	TotalRecords   int64          `json:"totalRecords"`
	TotalPages     int64          `json:"totalPages"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type GetAllProductsRequest struct {
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	SortOn      string `form:"sortOn"`
}

type GetAllProductsResponse struct {
	Products       []Product      `json:"products"`
	TotalRecords   int64          `json:"totalRecords"`
	TotalPages     int64          `json:"totalPages"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type CreateProductRequest struct {
	Product Product `json:"product:omitempty"`
}

type CreateProductResponse struct {
	Product        Product        `json:"product:omitempty"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type UpdateProductRequest struct {
	Id      int64   `json:"path:omitempty"`
	Product Product `json:"product:omitempty"`
}

type UpdateProductResponse struct {
	Product        Product        `json:"product:omitempty"`
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type DeleteProductRequest struct {
	Id int64 `json:"path:omitempty"`
}

type DeleteProductResponse struct {
	ResponseStatus ResponseStatus `json:"status"` // a ResponseStatus object
}

type Customer struct {
	Id        int64  `json:"id"`                 // customer id
	FirstName string `json:"firstName"`          // first name
	LastName  string `json:"lastName"`           // last or given name
	Email     string `json:"email,required"`     // email address
	Password  string `json:"password,omitempty"` // password
}

type NewCustomer struct {
	FirstName       string  `json:"firstName,required"`         // first name
	LastName        string  `json:"lastName,required"`          // last or given name
	Email           string  `json:"email,required"`             // email address, unique per store id
	Password        string  `json:"password,required"`          // password
	BillingAddress  Address `json:"billingAddress,optional"`    // Address object
	ShippingAddress Address `json:"shippingAddresses,optional"` // Address object
}

type CustomerAccount struct {
	CustomerId        int64     `json:"id"`                // customer id
	BillingAddress    Address   `json:"billingAddress"`    // Address object
	ShippingAddresses []Address `json:"shippingAddresses"` // collection of Address objects
}

type Address struct {
	Street        string `json:"street"`             // street name, ie: 1723 NW 23rd Ave.
	City          string `json:"city"`               // city name
	StateProvince string `json:"stateProvince"`      // state or province name
	Country       string `json:"country"`            // IISO 3166-1 alpha-2 country code. https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
	PostalCode    string `json:"postalCode"`         // postal or zip code
	IsDefault     bool   `json:"isDefault,required"` // indicates if this is a default address
}

type CustomerLoginRequest struct {
	Email    string `json:"email,required"`    // email address, unique to each store id
	Password string `json:"password,required"` // password
}

type CustomerLoginResponse struct {
	JwtToken       JwtToken       `json:"jwt"`      // jwt token
	Customer       Customer       `json:"customer"` // Customer object
	ResponseStatus ResponseStatus `json:"status"`   // a ResponseStatus object
}

type CreateCustomerRequest struct {
	Customer NewCustomer `json:"customer"` // NewCustomer object
}

type CreateCustomerResponse struct {
	Customer       Customer       `json:"customer"` // Customer object
	ResponseStatus ResponseStatus `json:"status"`   // a ResponseStatus object
}

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	JwtToken      JwtToken `json:"jwt"`
	User          User     `json:"user"`
	StatusCode    int64    `json:"statusCode"`
	StatusMessage string   `json:"statusMessage"`
}
