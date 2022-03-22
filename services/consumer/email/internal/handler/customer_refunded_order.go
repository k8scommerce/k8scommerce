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

func customerRefundedOrder(ev events.EventManager, svcCtx *svc.ServiceContext) error {
	// consume CustomerRefundedOrder events
	return ev.Consume(eventkey.CustomerRefundedOrder.AsKey(), "CustomerRefundedOrder", func(d rabbitmq.Delivery) rabbitmq.Action {
		go func() {
			data, err := eventkey.CustomerRefundedOrder.Unmarshal(d.Body)
			if err != nil {
				logx.Error("error unmarshalling d.Body : %s", err.Error())
			}

			ctx := context.Background()
			l := logic.NewCustomerRefundedOrderLogic(ctx, svcCtx)
			if err := l.Send(data); err != nil {
				logx.Error("error sending email NewCustomerRefundedOrderLogic: %s", err.Error())
			} else {
				logx.Infof("done sending email for customer id: %d", data.Customer.Id)
			}
		}()

		return rabbitmq.Ack
	})
}
