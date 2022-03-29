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

func customerSale(ev events.EventManager, svcCtx *svc.ServiceContext) error {
	// consume CustomerSale events
	return ev.Consume(eventkey.CustomerSale.AsKey(), "CustomerSale", func(d rabbitmq.Delivery) rabbitmq.Action {
		go func() {
			data, err := eventkey.CustomerSale.Unmarshal(d.Body)
			if err != nil {
				logx.Error("error unmarshalling d.Body : %s", err.Error())
			}

			ctx := context.Background()
			l := logic.NewCustomerSaleLogic(ctx, svcCtx)
			if err := l.Send(data); err != nil {
				logx.Error("error sending email NewCustomerSaleLogic: %s", err.Error())
			} else {
				logx.Infof("done sending email for customer id: %d", data.Customer.Id)
			}
		}()

		return rabbitmq.Ack
	})
}
