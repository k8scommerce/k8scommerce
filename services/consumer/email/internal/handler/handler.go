package handler

import (
	"k8scommerce/internal/events"
	"k8scommerce/services/consumer/email/internal/svc"
	"log"
)

func MustHandle(ev events.EventManager, svcCtx *svc.ServiceContext) {
	// Customer Emails
	if err := customerCompletedOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerCompletedOrder")
	}
	if err := customerConfirmationEmail(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerConfirmationEmail")
	}
	if err := customerNewAccount(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerNewAccount")
	}
	if err := customerNote(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerNote")
	}
	if err := customerOnHoldOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerOnHoldOrder")
	}
	if err := customerPasswordChanged(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerPasswordChanged")
	}
	if err := customerProcessingOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerProcessingOrder")
	}
	if err := customerRefundedOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerRefundedOrder")
	}
	if err := customerResetPassword(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerResetPassword")
	}
	if err := customerSale(ev, svcCtx); err != nil {
		log.Fatal("Error loading customerSale")
	}

	// Admin Emails
	if err := adminCancelledOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading adminCancelledOrder")
	}
	if err := adminFailedOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading adminFailedOrder")
	}
	if err := adminNewOrder(ev, svcCtx); err != nil {
		log.Fatal("Error loading adminNewOrder")
	}

}
