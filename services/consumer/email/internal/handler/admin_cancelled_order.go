package handler

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/consumer/email/internal/logic"
	"github.com/k8scommerce/k8scommerce/services/consumer/email/internal/svc"

	"github.com/k8scommerce/k8scommerce/internal/events"
	"github.com/k8scommerce/k8scommerce/internal/events/eventkey"

	"github.com/wagslane/go-rabbitmq"
	"github.com/zeromicro/go-zero/core/logx"
)

func adminCancelledOrder(ev events.EventManager, svcCtx *svc.ServiceContext) error {
	// consume AdminCancelledOrder events
	return ev.Consume(eventkey.AdminCancelledOrder.AsKey(), "AdminCancelledOrder", func(d rabbitmq.Delivery) rabbitmq.Action {
		go func() {
			data, err := eventkey.AdminCancelledOrder.Unmarshal(d.Body)
			if err != nil {
				logx.Error("error unmarshalling d.Body : %s", err.Error())
			}

			ctx := context.Background()
			l := logic.NewAdminCancelledOrderLogic(ctx, svcCtx)
			if err := l.Send(data); err != nil {
				logx.Error("error sending email NewAdminCancelledOrderLogic: %s", err.Error())
			} else {
				logx.Infof("done sending email for customer id: %d", data.Customer.Id)
			}
		}()

		return rabbitmq.Ack
	})
}
