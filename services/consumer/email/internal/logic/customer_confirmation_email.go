package logic

import (
	"bytes"
	"context"
	"log"
	"net/url"

	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/consumer/email/internal/svc"

	mail "github.com/xhit/go-simple-mail/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerConfirmationEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerConfirmationEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerConfirmationEmailLogic {
	return &CustomerConfirmationEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerConfirmationEmailLogic) Send(in *eventtype.CustomerConfirmationEmail) error {
	var body bytes.Buffer

	// url encode the code
	in.Code = url.QueryEscape(in.Code)

	CUSTOMER_CONFIRMATION_EMAIL.Execute(&body, in)

	// Create email
	msg := mail.NewMSG()
	msg.SetFrom("K8sCommerce <alma.tuck@k8scommerce.com>")
	msg.AddTo("alma.tuck@gmail.com")
	msg.SetSubject("Confirm your email address")
	msg.SetBody(mail.TextHTML, body.String())

	// Send msg
	err := msg.Send(l.svcCtx.EmailClient.GetSMTPClient())
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
