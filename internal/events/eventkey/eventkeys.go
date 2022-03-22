package eventkey

type EventKey string

// emails
type AdminCancelledOrderKey EventKey
type AdminFailedOrderKey EventKey
type AdminNewOrderKey EventKey
type CustomerAccountConfirmedEmailKey EventKey
type CustomerCompletedOrderKey EventKey
type CustomerConfirmationEmailKey EventKey
type CustomerNewAccountKey EventKey
type CustomerNoteKey EventKey
type CustomerOnHoldOrderKey EventKey
type CustomerPasswordChangedKey EventKey
type CustomerProcessingOrderKey EventKey
type CustomerRefundedOrderKey EventKey
type CustomerResetPasswordKey EventKey
type CustomerSaleKey EventKey

var AdminCancelledOrder AdminCancelledOrderKey = "AdminCancelledOrder"
var AdminFailedOrder AdminFailedOrderKey = "AdminFailedOrder"
var AdminNewOrder AdminNewOrderKey = "AdminNewOrder"
var CustomerAccountConfirmedEmail CustomerAccountConfirmedEmailKey = "CustomerAccountConfirmedEmail" // unsure about this one
var CustomerCompletedOrder CustomerCompletedOrderKey = "CustomerCompletedOrder"
var CustomerConfirmationEmail CustomerConfirmationEmailKey = "CustomerConfirmationEmail"
var CustomerNewAccount CustomerNewAccountKey = "CustomerNewAccount"
var CustomerNote CustomerNoteKey = "CustomerNote"
var CustomerOnHoldOrder CustomerOnHoldOrderKey = "CustomerOnHoldOrder"
var CustomerPasswordChanged CustomerPasswordChangedKey = "CustomerPasswordChanged"
var CustomerProcessingOrder CustomerProcessingOrderKey = "CustomerProcessingOrder"
var CustomerRefundedOrder CustomerRefundedOrderKey = "CustomerRefundedOrder"
var CustomerResetPassword CustomerResetPasswordKey = "CustomerResetPassword"
var CustomerSale CustomerSaleKey = "CustomerSale"

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

var CustomerCreated CustomerCreatedKey = "CustomerCreated"
var CustomerUpdated CustomerUpdatedKey = "CustomerUpdated"
var CustomerForgotPassword CustomerForgotPasswordKey = "CustomerForgotPassword"
var CustomerChangedPassword CustomerChangedPasswordKey = "CustomerChangedPassword"
var CustomerLoginSuccess CustomerLoginSuccessKey = "CustomerLoginSuccess"
var CustomerLoginFailed CustomerLoginFailedKey = "CustomerLoginFailed"
var CustomerAddedBillingAddress CustomerAddedBillingAddressKey = "CustomerAddedBillingAddress"
var CustomerUpdatedBillingAddress CustomerUpdatedBillingAddressKey = "CustomerUpdatedBillingAddress"
var CustomerDeletedBillingAddress CustomerDeletedBillingAddressKey = "CustomerDeletedBillingAddress"
var CustomerAddedShippingAddress CustomerAddedShippingAddressKey = "CustomerAddedShippingAddress"
var CustomerUpdatedShippingAddress CustomerUpdatedShippingAddressKey = "CustomerUpdatedShippingAddress"
var CustomerDeletedShippingAddress CustomerDeletedShippingAddressKey = "CustomerDeletedShippingAddress"

// catalog
type CatalogImageUploadedKey EventKey
type CatalogProductAddedKey EventKey
type CatalogProductUpdatedKey EventKey
type CatalogProductDeletedKey EventKey
type CatalogCategoryAddedKey EventKey
type CatalogCategoryUpdatedKey EventKey
type CatalogCategoryDeletedKey EventKey

var CatalogImageUploaded CatalogImageUploadedKey = "CatalogImageUploaded"
var CatalogProductAdded CatalogProductAddedKey = "CatalogProductAdded"
var CatalogProductUpdated CatalogProductUpdatedKey = "CatalogProductUpdated"
var CatalogProductDeleted CatalogProductDeletedKey = "CatalogProductDeleted"
var CatalogCategoryAdded CatalogCategoryAddedKey = "CatalogCategoryAdded"
var CatalogCategoryUpdated CatalogCategoryUpdatedKey = "CatalogCategoryUpdated"
var CatalogCategoryDeleted CatalogCategoryDeletedKey = "CatalogCategoryDeleted"

// cart
type CartProductAddedKey EventKey
type CartProductUpdatedKey EventKey
type CartProductDeletedKey EventKey

var CartProductAdded CartProductAddedKey = "CartProductAdded"
var CartProductUpdated CartProductUpdatedKey = "CartProductUpdated"
var CartProductDeleted CartProductDeletedKey = "CartProductDeleted"

// user
type UserCreatedKey EventKey
type UserUpdatedKey EventKey
type UserForgotPasswordKey EventKey
type UserChangedPasswordKey EventKey
type UserLoginSuccessKey EventKey
type UserLoginFailedKey EventKey

var UserCreated UserCreatedKey = "UserCreated"
var UserUpdated UserUpdatedKey = "UserUpdated"
var UserForgotPassword UserForgotPasswordKey = "UserForgotPassword"
var UserChangedPassword UserChangedPasswordKey = "UserChangedPassword"
var UserLoginSuccess UserLoginSuccessKey = "UserLoginSuccess"
var UserLoginFailed UserLoginFailedKey = "UserLoginFailed"

// debug/testing
type DebugTestingKey EventKey

var DebugTesting DebugTestingKey = "DebugTesting"
