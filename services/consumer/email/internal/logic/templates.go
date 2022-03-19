package logic

import "html/template"

var (
	// Customer Emails
	CUSTOMER_COMPLETED_ORDER    *template.Template
	CUSTOMER_CONFIRMATION_EMAIL *template.Template
	CUSTOMER_NEW_ACCOUNT        *template.Template
	CUSTOMER_NOTE               *template.Template
	CUSTOMER_ON_HOLD_ORDER      *template.Template
	CUSTOMER_PASSWORD_CHANGED   *template.Template
	CUSTOMER_PROCESSING_ORDER   *template.Template
	CUSTOMER_REFUNDED_ORDER     *template.Template
	CUSTOMER_RESET_PASSWORD     *template.Template
	CUSTOMER_SALE               *template.Template

	// Admin Emails
	ADMIN_CANCELLED_ORDER *template.Template
	ADMIN_FAILED_ORDER    *template.Template
	ADMIN_NEW_ORDER       *template.Template
)

func init() {
	// Customer Emails
	CUSTOMER_COMPLETED_ORDER = template.Must(template.ParseFiles("./templates/customer/completed_order.html"))
	CUSTOMER_CONFIRMATION_EMAIL = template.Must(template.ParseFiles("./templates/customer/customer_confirmation_email.html"))
	CUSTOMER_NEW_ACCOUNT = template.Must(template.ParseFiles("./templates/customer/new_account.html"))
	CUSTOMER_NOTE = template.Must(template.ParseFiles("./templates/customer/note.html"))
	CUSTOMER_ON_HOLD_ORDER = template.Must(template.ParseFiles("./templates/customer/on_hold_order.html"))
	CUSTOMER_PASSWORD_CHANGED = template.Must(template.ParseFiles("./templates/customer/password_changed.html"))
	CUSTOMER_PROCESSING_ORDER = template.Must(template.ParseFiles("./templates/customer/processing_order.html"))
	CUSTOMER_REFUNDED_ORDER = template.Must(template.ParseFiles("./templates/customer/refunded_order.html"))
	CUSTOMER_RESET_PASSWORD = template.Must(template.ParseFiles("./templates/customer/reset_password.html"))
	CUSTOMER_SALE = template.Must(template.ParseFiles("./templates/customer/sale.html"))

	// Admin Emails
	ADMIN_CANCELLED_ORDER = template.Must(template.ParseFiles("./templates/admin/cancelled_order.html"))
	ADMIN_FAILED_ORDER = template.Must(template.ParseFiles("./templates/admin/failed_order.html"))
	ADMIN_NEW_ORDER = template.Must(template.ParseFiles("./templates/admin/new_order.html"))
}
