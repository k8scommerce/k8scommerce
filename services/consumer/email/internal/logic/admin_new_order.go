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

type AdminNewOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminNewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminNewOrderLogic {
	return &AdminNewOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminNewOrderLogic) Send(in *eventtype.AdminNewOrder) error {
	var body bytes.Buffer

	ADMIN_NEW_ORDER.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.AdminNewOrder.Name
	fromEmail := in.StoreSetting.Config.Emails.AdminNewOrder.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "AdminNewOrderSubject",
		DefaultMessage: &i18n.Message{
			ID:          "AdminNewOrderSubject",
			Description: "The subject of the admin new order email",
			Other:       "New order #{{ .Order }}",
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
