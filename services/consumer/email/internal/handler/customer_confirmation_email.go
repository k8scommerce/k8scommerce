package handler

import (
	"context"
	"k8scommerce/internal/events"
	"k8scommerce/internal/events/eventkey"
	"k8scommerce/services/consumer/email/internal/logic"
	"k8scommerce/services/consumer/email/internal/svc"

	"github.com/wagslane/go-rabbitmq"
	"github.com/zeromicro/go-zero/core/logx"
)

func customerConfirmationEmail(ev events.EventManager, svcCtx *svc.ServiceContext) error {
	// consume CustomerConfirmationEmail events
	return ev.Consume(eventkey.CustomerConfirmationEmail.AsKey(), "CustomerConfirmationEmail", func(d rabbitmq.Delivery) rabbitmq.Action {
		go func() {
			data, err := eventkey.CustomerConfirmationEmail.Unmarshal(d.Body)
			if err != nil {
				logx.Error("error unmarshalling d.Body : %s", err.Error())
			}

			ctx := context.Background()
			l := logic.NewCustomerConfirmationEmailLogic(ctx, svcCtx)
			if err := l.Send(data); err != nil {
				logx.Error("error sending email NewCustomerConfirmationEmailLogic: %s", err.Error())
			} else {
				logx.Infof("done sending email for customer id: %d", data.Customer.Id)
			}
		}()

		return rabbitmq.Ack
	})
}