package logic

import (
	"bytes"
	"context"
	"fmt"

	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/consumer/email/internal/svc"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	mail "github.com/xhit/go-simple-mail/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminFailedOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminFailedOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminFailedOrderLogic {
	return &AdminFailedOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminFailedOrderLogic) Send(in *eventtype.AdminFailedOrder) error {
	var body bytes.Buffer

	ADMIN_FAILED_ORDER.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.AdminFailedOrder.Name
	fromEmail := in.StoreSetting.Config.Emails.AdminFailedOrder.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "AdminFailedOrderSubject",
		DefaultMessage: &i18n.Message{
			ID:          "AdminFailedOrderSubject",
			Description: "The subject of the admin failed order email",
			Other:       "Order #{{ .Order }} has failed",
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
