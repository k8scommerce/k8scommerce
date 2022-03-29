package main

import (
	"context"

	"github.com/k8scommerce/k8scommerce/internal/models"
)

func createStore() {

	// state := faker.Address().StateAbbr()

	// address := fmt.Sprintf("%s\n%s, %s %s",
	// 	faker.Address().StreetAddress(),
	// 	faker.Address().City(),
	// 	state,
	// 	faker.Address().ZipCodeByState(state),
	// )

	store := &models.Store{
		Name:        "Demo Store",
		Description: toNullString("My Demo Store"),
		URL:         "http://localhost:4200/",
		IsDefault:   true,
	}
	if err := store.Insert(context.Background(), db); err != nil {
		panic(err)
	}
	storeID = store.ID
}

func createStoreSetting() {
	config := `{"currency":{"default_currency":"USD","supported_currencies":["USD","CAN"]},"locale":{"default_locale":"en-US","supported_locales":["en-US","en-CA"]},"contact":{"phone":{"corportate":"800-555-1212","customer_support":"800-555-1212","custom":{"Sales":"800-555-1212 Ext 1234"}},"addresses":[{"name":"Warehouse 1","street":"123 Any Street","apt_suite":"100","city":"Big Springs","state_province":"VA","country":"US","postal_code":"12345","is_default":true}]},"emails":{"default":{"name":"K8sCommerce","email":"test@k8scommerce.com"},"customer_support":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_completed_order":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_confirmation_email":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_new_account":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_note":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_on_hold_order":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_password_changed":{"name":"K8sCommerce","email":"no-reply@k8scommerce.com"},"customer_processing_order":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_refunded_order":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"customer_reset_password":{"name":"K8sCommerce","email":"no-reply@k8scommerce.com"},"customer_sale":{"name":"K8sCommerce","email":"support@k8scommerce.com"},"admin_cancelled_order":{"name":"K8sCommerce","email":"no-reply@k8scommerce.com"},"admin_failed_order":{"name":"K8sCommerce","email":"no-reply@k8scommerce.com"},"admin_new_order":{"name":"K8sCommerce","email":"no-reply@k8scommerce.com"}}}`
	storeSetting := &models.StoreSetting{
		StoreID: storeID,
		Config:  []byte(config),
	}
	if err := storeSetting.Insert(context.Background(), db); err != nil {
		panic(err)
	}
}
