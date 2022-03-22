package logic

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/consumer/email/internal/svc"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	mail "github.com/xhit/go-simple-mail/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerResetPasswordLogic {
	return &CustomerResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerResetPasswordLogic) Send(in *eventtype.CustomerResetPassword) error {
	var body bytes.Buffer

	// url encode the code
	in.Code = url.QueryEscape(in.Code)

	CUSTOMER_RESET_PASSWORD.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerResetPassword.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerResetPassword.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerResetPasswordSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerResetPasswordSubject",
			Description: "The subject of the customer password reset email",
			Other:       "Your password reset request",
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
