package logic

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/consumer/email/internal/svc"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	mail "github.com/xhit/go-simple-mail/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerRefundedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerRefundedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerRefundedOrderLogic {
	return &CustomerRefundedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerRefundedOrderLogic) Send(in *eventtype.CustomerRefundedOrder) error {
	var body bytes.Buffer

	CUSTOMER_REFUNDED_ORDER.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerRefundedOrder.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerRefundedOrder.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerRefundedOrderSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerRefundedOrderSubject",
			Description: "The subject of the customer order refund email",
			Other:       "Your order #{{ .Order }} has been refunded",
		},
		TemplateData: map[string]interface{}{
			"Order": "TBD",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	msg := mail.NewMSG()
	msg.SetFrom(fmt.Sprintf("%s <%s>", fromName, fromEmail))
	msg.AddTo(in.Customer.Email)
	msg.SetSubject(subject)
	msg.SetBody(mail.TextHTML, body.String())

	// Send msg
	err = msg.Send(l.svcCtx.EmailClient.GetSMTPClient())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
