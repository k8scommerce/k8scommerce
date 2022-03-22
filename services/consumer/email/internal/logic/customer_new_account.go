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

type CustomerNewAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerNewAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerNewAccountLogic {
	return &CustomerNewAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerNewAccountLogic) Send(in *eventtype.CustomerNewAccount) error {
	var body bytes.Buffer

	CUSTOMER_NEW_ACCOUNT.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerNewAccount.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerNewAccount.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerNewAccountSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerNewAccountSubject",
			Description: "The subject of the customer new account email",
			Other:       "Your account has been created",
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
