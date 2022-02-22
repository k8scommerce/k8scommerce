package config

type EventKey string

// customer events
const (
	Event_CustomerCreated                EventKey = "customer-created"
	Event_CustomerUpdated                EventKey = "customer-updated"
	Event_CustomerForgotPassword         EventKey = "customer-forgot-password"
	Event_CustomerChangedPassword        EventKey = "customer-changed-password"
	Event_CustomerLoginSuccess           EventKey = "customer-login-success"
	Event_CustomerLoginFailed            EventKey = "customer-login-failed"
	Event_CustomerAddedBillingAddress    EventKey = "customer-billing-address-added"
	Event_CustomerUpdatedBillingAddress  EventKey = "customer-billing-address-updated"
	Event_CustomerDeletedBillingAddress  EventKey = "customer-billing-address-deleted"
	Event_CustomerAddedShippingAddress   EventKey = "customer-shipping-address-added"
	Event_CustomerUpdatedShippingAddress EventKey = "customer-shipping-address-updated"
	Event_CustomerDeletedShippingAddress EventKey = "customer-shipping-address-deleted"

	// catalog
	Event_CatalogImageUploaded   EventKey = "catalog-image-uploaded"
	Event_CatalogProductAdded    EventKey = "catalog-product-added"
	Event_CatalogProductUpdated  EventKey = "catalog-product-updated"
	Event_CatalogProductDeleted  EventKey = "catalog-product-deleted"
	Event_CatalogCategoryAdded   EventKey = "catalog-category-added"
	Event_CatalogCategoryUpdated EventKey = "catalog-category-updated"
	Event_CatalogCategoryDeleted EventKey = "catalog-category-deleted"

	// cart
	Event_CartProductAdded   EventKey = "cart-product-added"
	Event_CartProductUpdated EventKey = "cart-product-updated"
	Event_CartProductDeleted EventKey = "cart-product-deleted"

	// user
	Event_UserCreated         EventKey = "user-created"
	Event_UserUpdated         EventKey = "user-updated"
	Event_UserForgotPassword  EventKey = "user-forgot-password"
	Event_UserChangedPassword EventKey = "user-changed-password"
	Event_UserLoginSuccess    EventKey = "user-login-success"
	Event_UserLoginFailed     EventKey = "user-login-failed"
)
