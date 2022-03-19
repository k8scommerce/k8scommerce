package logic

import (
	"bytes"
	"context"
	"log"

	"k8scommerce/internal/events/eventkey/eventtype"
	"k8scommerce/services/consumer/email/internal/svc"

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
