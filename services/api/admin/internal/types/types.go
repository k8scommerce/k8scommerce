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

type Category struct {
	Id              int64  `json:"id"`                       // category id
	ParentId        int64  `json:"parentId,optional"`        // parent category id. references Category.Id
	Slug            string `json:"slug"`                     // slug name of the category
	Name            string `json:"name"`                     // name of category
	Description     string `json:"description,optional"`     // description of category
	MetaTitle       string `json:"metaTitle,optional"`       // metatag title for SEO
	MetaDescription string `json:"metaDescription,optional"` // metatag description for SEO
	MetaKeywords    string `json:"metaKeywords,optional"`    // metatag keywords for SEO
	Depth           int32  `json:"depth,optional"`           // category level depth
	SortOrder       int32  `json:"sortOrder,optional"`       // sort order of menu items on the same level and same parent id
}

type Product struct {
	Id               int64     `json:"id"`                        // product id
	Slug             string    `json:"slug"`                      // product slug
	Name             string    `json:"name"`                      // product name
	ShortDescription string    `json:"shortDescription,optional"` // product short description. used in category pages
	Description      string    `json:"description,optional"`      // category description
	MetaTitle        string    `json:"metaTitle,optional"`        // metatag title for SEO
	MetaDescription  string    `json:"metaDescription,optional"`  // metatag description for SEO
	MetaKeywords     string    `json:"metaKeywords,optional"`     // metatag keywords for SEO
	Variants         []Variant `json:"variants,optional"`         // collection of Variant objects
}

type Variant struct {
	Id        int64   `json:"id"`              // variant id
	IsDefault bool    `json:"isDefault"`       // is default variant. each product must have exactly 1 default variant
	Sku       string  `json:"sku"`             // variant sku. usually the product sku with appended identification tags
	Weight    float64 `json:"weight,optional"` // variant weight. used in calculating shipping
	Height    float64 `json:"height,optional"` // variant height. used in calculating shipping
	Width     float64 `json:"width,optional"`  // variant width. used in calculating shipping
	Depth     float64 `json:"depth,optional"`  // variant depth. used in calculating shipping
	Price     Price   `json:"price,optional"`  // variant Price object
}

type Price struct {
	Id                     int64   `json:"id,optional"`                     // price id
	Amount                 float64 `json:"amount"`                          // price amount
	DisplayAmount          string  `json:"displayAmount"`                   // price display amount
	CompareAtAmount        float64 `json:"compareAtAmount,optional"`        // price compare amount
	DisplayCompareAtAmount string  `json:"displayCompareAtAmount,optional"` // price display compare amount
	Currency               string  `json:"currency,optional"`               // price currency. example: USD, CAN, etc.
}

type GetAllCategoriesRequest struct {
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	SortOn      string `form:"sortOn,optional"`
}

type GetAllCategoriesResponse struct {
	Categories   []Category `json:"categories"` // a collection of Category
	TotalRecords int64      `json:"totalRecords"`
	TotalPages   int64      `json:"totalPages"`
}

type GetCategoryBySlugRequest struct {
	Slug string `path:"slug"` // slug name of the category
}

type GetCategoryByIdRequest struct {
	Id int64 `path:"id"`
}

type CreateCategoryRequest struct {
	Category Category `json:"category"`
}

type UpdateCategoryRequest struct {
	Id       int64    `json:"id"`
	Category Category `json:"category"`
}

type DeleteCategoryRequest struct {
	Id int64 `path:"id"`
}

type DeleteCategoryResponse struct {
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
	SortOn      string `form:"sortOn,optional"`
}

type GetProductsByCategoryIdResponse struct {
	Products     []Product `json:"products"`
	TotalRecords int64     `json:"totalRecords"`
	TotalPages   int64     `json:"totalPages"`
}

type GetAllProductsRequest struct {
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	SortOn      string `form:"sortOn,optional"`
}

type GetAllProductsResponse struct {
	Products     []Product `json:"products"`
	TotalRecords int64     `json:"totalRecords"`
	TotalPages   int64     `json:"totalPages"`
}

type CreateProductRequest struct {
	Product Product `json:"product"`
}

type UpdateProductRequest struct {
	Id      int64   `json:"path"`
	Product Product `json:"product"`
}

type DeleteProductRequest struct {
	Id int64 `path:"id"`
}

type DeleteProductResponse struct {
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
	Id        int64  `json:"id"`                // user id
	FirstName string `json:"firstName"`         // first name
	LastName  string `json:"lastName"`          // last name
	Email     string `json:"email"`             // email address
	Password  string `json:"password,optional"` // password
}

type PermissionGroup struct {
	Id        int64  `json:"id"`        // permission group id
	GroupName string `json:"groupName"` // groupName
}

type UsersPermissionGroups struct {
	UserId            int64 `json:"userId"`            // user id
	PermissionGroupId int64 `json:"permissionGroupId"` // permission group id
}

type UserLoginRequest struct {
	Email    string `json:"email"`    // email address
	Password string `json:"password"` // password
}

type UserLoginResponse struct {
	JwtToken JwtToken `json:"jwt"`  // JwtToken object
	User     User     `json:"user"` // User object
}

type GetAllUsersRequest struct {
	CurrentPage int64  `path:"currentPage"`
	PageSize    int64  `path:"pageSize"`
	SortOn      string `form:"sortOn,optional"`
}

type GetAllUsersResponse struct {
	Users        []User `json:"users"`
	TotalRecords int64  `json:"totalRecords"`
	TotalPages   int64  `json:"totalPages"`
}
