// Code generated by goctl. DO NOT EDIT.
package types

type JwtToken struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type ResponseStatus struct {
	StatusCode    int64  `json:"statusCode"`              // RFC http status code, ie. 204, etc - https://go.dev/src/net/http/status.go
	StatusMessage string `json:"statusMessage,omitempty"` // status message
}

type PingResponse struct {
	Ping string `json:"ping"`
}

type Cart struct {
	Item       []Item  `json:"items"`      // a collection of Item
	TotalPrice float64 `json:"totalPrice"` // the sum total of the cart
}

type Item struct {
	Sku       string  `json:"sku"`       // an item's variant sku number
	Quantity  int32   `json:"quantity"`  // how many of identical items
	Price     float64 `json:"price"`     // the item's price
	ExpiresAt string  `json:"expiresAt"` // when this item expires in the cart
}

type GetCartRequest struct {
	CustomerId int64 `path:"customerId"` // a customer's id
}

type GetCartResponse struct {
	Cart Cart `json:"cart"` // a Cart object
}

type AddItemToCartRequest struct {
	CustomerId int64 `path:"customerId"` // a customer's id
	Item       Item  `json:"item"`       // an Item object
}

type AddItemToCartResponse struct {
	Cart Cart `json:"cart"` // a Cart object
}

type UpdateCartItemQuantityRequest struct {
	CustomerId int64  `path:"customerId"` // a customer's id
	Sku        string `path:"sku"`        // an item's variant sku number
	Quanity    int32  `json:"quanity"`    // a new quantity
}

type UpdateCartItemQuantityResponse struct {
	Cart Cart `json:"cart"` // a Cart object
}

type RemoveCartItemRequest struct {
	CustomerId int64  `path:"customerId"` // a customer's id
	Sku        string `path:"sku"`        // an item's variant sku number
	Quanity    int32  `json:"quanity"`    // a new quantity
}

type RemoveCartItemResponse struct {
	Cart Cart `json:"cart"` // a Cart object
}

type ClearCartRequest struct {
	CustomerId int64 `path:"customerId"` // a customer's id
}

type ClearCartResponse struct {
	Deleted bool `json:"deleted"` // a boolean true/false if successful
}

type Category struct {
	Id              int64  `json:"id"`                                 // category id
	ParentId        int64  `json:"parentId"`                           // parent category id. references Category.Id
	Slug            string `json:"slug"`                               // slug name of the category
	Name            string `json:"name"`                               // name of category
	Description     string `json:"description,optional,omitempty"`     // description of category
	MetaTitle       string `json:"metaTitle,optional,omitempty"`       // metatag title for SEO
	MetaDescription string `json:"metaDescription,optional,omitempty"` // metatag description for SEO
	MetaKeywords    string `json:"metaKeywords,optional,omitempty"`    // metatag keywords for SEO
	Depth           int32  `json:"depth,optional,omitempty"`           // category level depth
	SortOrder       int32  `json:"sortOrder,optional,omitempty"`       // sort order of menu items on the same level and same parent id
}

type CategoryPair struct {
	Slug string `json:"slug"` // slug name of the category
	Name string `json:"name"` // name of category
}

type Product struct {
	Id               int64          `json:"id"`                                  // product id
	Slug             string         `json:"slug"`                                // product slug
	Name             string         `json:"name"`                                // product name
	ShortDescription string         `json:"shortDescription,optional,omitempty"` // product short description. used in category pages
	Description      string         `json:"description,optional,omitempty"`      // category description
	MetaTitle        string         `json:"metaTitle,optional,omitempty"`        // metatag title for SEO
	MetaDescription  string         `json:"metaDescription,optional,omitempty"`  // metatag description for SEO
	MetaKeywords     string         `json:"metaKeywords,optional,omitempty"`     // metatag keywords for SEO
	Variants         []Variant      `json:"variants,optional,omitempty"`         // collection of Variant objects
	DefaultImage     Asset          `json:"defaultImage,optional,omitempty"`     // default Asset object of image type
	Images           []Asset        `json:"images,optional,omitempty"`           // array of Asset objects of image type
	Categories       []CategoryPair `json:"categories,optional,omitempty"`       // array of Asset objects of image type
}

type Variant struct {
	Id        int64   `json:"id"`                        // variant id
	IsDefault bool    `json:"isDefault"`                 // is default variant. each product must have exactly 1 default variant
	Sku       string  `json:"sku"`                       // variant sku. usually the product sku with appended identification tags
	Weight    float64 `json:"weight,optional,omitempty"` // variant weight. used in calculating shipping
	Height    float64 `json:"height,optional,omitempty"` // variant height. used in calculating shipping
	Width     float64 `json:"width,optional,omitempty"`  // variant width. used in calculating shipping
	Depth     float64 `json:"depth,optional,omitempty"`  // variant depth. used in calculating shipping
	Price     Price   `json:"price,optional,omitempty"`  // variant Price object
}

type Price struct {
	Id                   int64   `json:"id,optional,omitempty"`                   // price id
	SalePrice            float64 `json:"salePrice"`                               // sale price
	FormattedSalePrice   string  `json:"formattedSalePrice"`                      // formatted sale price
	RetailPrice          float64 `json:"retailPrice,optional,omitempty"`          // retail price
	FormattedRetailPrice string  `json:"formattedRetailPrice,optional,omitempty"` // formatted retail price
	Currency             string  `json:"currency,optional,omitempty"`             // currency. example: USD, CAN, etc.
}

type Asset struct {
	Id          int64             `json:"id,optional,omitempty"`                       // asset id
	ProductId   int64             `json:"productId,optional,omitempty"`                // product id
	VariantId   int64             `json:"variantId,optional,omitempty"`                // variant id
	Name        string            `json:"name,optional,omitempty"`                     // asset name
	DisplayName string            `json:"displayName,optional,omitempty"`              // display name
	Url         string            `json:"url,optional,omitempty"`                      // full, public access url
	Kind        int               `json:"kind,optional,omitempty,options=0|1|2|3|4|5"` // asset kind (0=unknown|1=image|2=document|3=audio|4=video|5=archive)
	ContentType string            `json:"contentType,optional,omitempty"`              // content type (mime type)
	SortOrder   int64             `json:"sortOrder,optional,omitempty"`                // sort order
	Sizes       map[string]string `json:"sizes,optional,omitempty"`                    // map[tag:string]url:string
}

type GetAllCategoriesResponse struct {
	Categories []Category `json:"categories"` // a collection of Category
}

type GetCategoryBySlugRequest struct {
	Slug string `path:"slug"` // slug name of the category
}

type GetCategoryByIdRequest struct {
	Id int64 `path:"id"`
}

type GetProductBySkuRequest struct {
	Sku string `path:"sku"`
}

type GetProductBySlugRequest struct {
	Slug string `path:"slug"` // slug name of the category
}

type GetProductByIdRequest struct {
	Id int64 `path:"id"`
}

type GetProductsByCategoryIdRequest struct {
	CategoryId  int64  `path:"categoryId"`
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	Filter      string `form:"filter,optional,omitempty"`
	SortOn      string `form:"sortOn,optional,omitempty"`
}

type GetProductsByCategoryIdResponse struct {
	Products     []Product `json:"products"`
	TotalRecords int64     `json:"totalRecords"`
	TotalPages   int64     `json:"totalPages"`
}

type GetAllProductsRequest struct {
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	Filter      string `form:"filter,optional,omitempty"`
	SortOn      string `form:"sortOn,optional,omitempty"`
}

type GetAllProductsResponse struct {
	Products     []Product `json:"products"`
	TotalRecords int64     `json:"totalRecords"`
	TotalPages   int64     `json:"totalPages"`
}

type Customer struct {
	Id        int64  `json:"id"`                // customer id
	FirstName string `json:"firstName"`         // first name
	LastName  string `json:"lastName"`          // last or given name
	Email     string `json:"email,required"`    // email address
	Password  string `json:"password,optional"` // password
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
	JwtToken JwtToken `json:"jwt"`      // jwt token
	Customer Customer `json:"customer"` // Customer object
	Success  bool     `json:"success"`  // success bool
}

type CreateCustomerRequest struct {
	Customer NewCustomer `json:"customer"` // NewCustomer object
}

type CreateCustomerResponse struct {
	Customer Customer `json:"customer"` // Customer object
}

type CheckForExistingEmailRequest struct {
	Email string `json:"email"` // Customer object
}

type CheckForExistingEmailResponse struct {
	Exists bool `json:"exists"` // boolean true/false if email exists or not
}

type GetCustomerResponse struct {
	Customer Customer `json:"customer"` // Customer object
}
