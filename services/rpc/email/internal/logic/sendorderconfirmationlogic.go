package logic

import (
	"context"
	"sync"

	"k8scommerce/services/rpc/email/internal/svc"
	"k8scommerce/services/rpc/email/pb/email"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendOrderConfirmationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	mu sync.Mutex
}

func NewSendOrderConfirmationLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *SendOrderConfirmationLogic {
	return &SendOrderConfirmationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendOrderConfirmationLogic) SendOrderConfirmation(in *email.SendOrderConfirmationRequest) (*email.Empty, error) {
	res := &email.Empty{}
	return res, nil
}
