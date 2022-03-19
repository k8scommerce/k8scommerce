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
