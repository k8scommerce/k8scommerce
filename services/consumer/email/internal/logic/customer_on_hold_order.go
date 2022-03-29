package logic

import (
	"bytes"
	"context"
	"fmt"

	"github.com/k8scommerce/k8scommerce/services/consumer/email/internal/svc"

	"github.com/k8scommerce/k8scommerce/internal/events/eventkey/eventtype"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	mail "github.com/xhit/go-simple-mail/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerOnHoldOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerOnHoldOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerOnHoldOrderLogic {
	return &CustomerOnHoldOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerOnHoldOrderLogic) Send(in *eventtype.CustomerOnHoldOrder) error {
	var body bytes.Buffer

	CUSTOMER_ON_HOLD_ORDER.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerOnHoldOrder.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerOnHoldOrder.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerOnHoldOrderSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerOnHoldOrderSubject",
			Description: "The subject of the customer order on hold email",
			Other:       "Your order #{{ .Order }} has been placed on hold",
		},
		TemplateData: map[string]interface{}{
			"Order": "TBD",
		},
	})
	if err != nil {
		logx.Info(err)
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
		logx.Info(err)
	}

	return nil
}
