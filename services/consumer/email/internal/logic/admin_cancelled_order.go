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

type AdminCancelledOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminCancelledOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminCancelledOrderLogic {
	return &AdminCancelledOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminCancelledOrderLogic) Send(in *eventtype.AdminCancelledOrder) error {
	var body bytes.Buffer

	ADMIN_CANCELLED_ORDER.Execute(&body, in)

	fromName := in.StoreSetting.Config.Emails.AdminCancelledOrder.Name
	fromEmail := in.StoreSetting.Config.Emails.AdminCancelledOrder.Email
	if fromName == "" {
		fromName = in.StoreSetting.Config.Emails.Default.Name
	}
	if fromEmail == "" {
		fromEmail = in.StoreSetting.Config.Emails.Default.Email
	}

	subject, err := l.svcCtx.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "AdminCancelledOrderSubject",
		DefaultMessage: &i18n.Message{
			ID:          "AdminCancelledOrderSubject",
			Description: "The subject of the admin order cancelation email",
			Other:       "Order #{{ .Order }} has been canceled",
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
