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

type CustomerNoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerNoteLogic {
	return &CustomerNoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerNoteLogic) Send(in *eventtype.CustomerNote) error {
	var body bytes.Buffer

	CUSTOMER_NOTE.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.CustomerNote.Name
	fromEmail := in.StoreSetting.Config.Emails.CustomerNote.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "CustomerNoteSubject",
		DefaultMessage: &i18n.Message{
			ID:          "CustomerNoteSubject",
			Description: "The subject of the customer note email",
			Other:       "A note has been added to your account",
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
