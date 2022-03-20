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

type CustomerProcessingOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerProcessingOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerProcessingOrderLogic {
	return &CustomerProcessingOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerProcessingOrderLogic) Send(in *eventtype.CustomerProcessingOrder) error {
	var body bytes.Buffer

	CUSTOMER_PROCESSING_ORDER.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerProcessingOrder.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerProcessingOrder.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerProcessingOrderSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerProcessingOrderSubject",
			Description: "The subject of the customer order processing email",
			Other:       "Your order #{{ .Order }} is being processed",
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
