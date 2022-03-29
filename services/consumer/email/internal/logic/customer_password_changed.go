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

type CustomerPasswordChangedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerPasswordChangedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerPasswordChangedLogic {
	return &CustomerPasswordChangedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerPasswordChangedLogic) Send(in *eventtype.CustomerPasswordChanged) error {
	var body bytes.Buffer

	CUSTOMER_PASSWORD_CHANGED.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerPasswordChanged.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerPasswordChanged.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerPasswordChangedSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerPasswordChangedSubject",
			Description: "The subject of the customer password changed email",
			Other:       "Your password has changed",
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
