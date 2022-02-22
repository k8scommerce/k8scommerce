package eventkey

type EventKey string

// customer events
type CustomerCreatedKey EventKey
type CustomerUpdatedKey EventKey
type CustomerForgotPasswordKey EventKey
type CustomerChangedPasswordKey EventKey
type CustomerLoginSuccessKey EventKey
type CustomerLoginFailedKey EventKey
type CustomerAddedBillingAddressKey EventKey
type CustomerUpdatedBillingAddressKey EventKey
type CustomerDeletedBillingAddressKey EventKey
type CustomerAddedShippingAddressKey EventKey
type CustomerUpdatedShippingAddressKey EventKey
type CustomerDeletedShippingAddressKey EventKey

var CustomerCreated CustomerCreatedKey = "customer-created"
var CustomerUpdated CustomerUpdatedKey = "customer-updated"
var CustomerForgotPassword CustomerForgotPasswordKey = "customer-forgot-password"
var CustomerChangedPassword CustomerChangedPasswordKey = "customer-changed-password"
var CustomerLoginSuccess CustomerLoginSuccessKey = "customer-login-success"
var CustomerLoginFailed CustomerLoginFailedKey = "customer-login-failed"
var CustomerAddedBillingAddress CustomerAddedBillingAddressKey = "customer-billing-address-added"
var CustomerUpdatedBillingAddress CustomerUpdatedBillingAddressKey = "customer-billing-address-updated"
var CustomerDeletedBillingAddress CustomerDeletedBillingAddressKey = "customer-billing-address-deleted"
var CustomerAddedShippingAddress CustomerAddedShippingAddressKey = "customer-shipping-address-added"
var CustomerUpdatedShippingAddress CustomerUpdatedShippingAddressKey = "customer-shipping-address-updated"
var CustomerDeletedShippingAddress CustomerDeletedShippingAddressKey = "customer-shipping-address-deleted"

// catalog
type CatalogImageUploadedKey EventKey
type CatalogProductAddedKey EventKey
type CatalogProductUpdatedKey EventKey
type CatalogProductDeletedKey EventKey
type CatalogCategoryAddedKey EventKey
type CatalogCategoryUpdatedKey EventKey
type CatalogCategoryDeletedKey EventKey

var CatalogImageUploaded CatalogImageUploadedKey = "catalog-image-uploaded"
var CatalogProductAdded CatalogProductAddedKey = "catalog-product-added"
var CatalogProductUpdated CatalogProductUpdatedKey = "catalog-product-updated"
var CatalogProductDeleted CatalogProductDeletedKey = "catalog-product-deleted"
var CatalogCategoryAdded CatalogCategoryAddedKey = "catalog-category-added"
var CatalogCategoryUpdated CatalogCategoryUpdatedKey = "catalog-category-updated"
var CatalogCategoryDeleted CatalogCategoryDeletedKey = "catalog-category-deleted"

// cart
type CartProductAddedKey EventKey
type CartProductUpdatedKey EventKey
type CartProductDeletedKey EventKey

var CartProductAdded CartProductAddedKey = "cart-product-added"
var CartProductUpdated CartProductUpdatedKey = "cart-product-updated"
var CartProductDeleted CartProductDeletedKey = "cart-product-deleted"

// user
type UserCreatedKey EventKey
type UserUpdatedKey EventKey
type UserForgotPasswordKey EventKey
type UserChangedPasswordKey EventKey
type UserLoginSuccessKey EventKey
type UserLoginFailedKey EventKey

var UserCreated UserCreatedKey = "user-created"
var UserUpdated UserUpdatedKey = "user-updated"
var UserForgotPassword UserForgotPasswordKey = "user-forgot-password"
var UserChangedPassword UserChangedPasswordKey = "user-changed-password"
var UserLoginSuccess UserLoginSuccessKey = "user-login-success"
var UserLoginFailed UserLoginFailedKey = "user-login-failed"

// debug/testing
type DebugTestingKey EventKey

var DebugTesting DebugTestingKey = "ginkgo-test-event"
