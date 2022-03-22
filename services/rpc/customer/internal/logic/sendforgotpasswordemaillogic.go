package logic

import (
	"context"
	"k8scommerce/internal/events/eventkey"
	"k8scommerce/internal/utils"
	"k8scommerce/services/rpc/customer/internal/svc"
	"k8scommerce/services/rpc/customer/pb/customer"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
)

type SendForgotPasswordEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendForgotPasswordEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *SendForgotPasswordEmailLogic {
	return &SendForgotPasswordEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendForgotPasswordEmailLogic) SendForgotPasswordEmail(in *customer.SendForgotPasswordEmailRequest) (*customer.SendForgotPasswordEmailResponse, error) {
	found, err := l.svcCtx.Repo.Customer().GetCustomerByEmail(in.StoreId, in.Email)
	if err != nil {
		return nil, err
	}

	out := &customer.Customer{}
	utils.TransformObj(found, &out)

	res := &customer.SendForgotPasswordEmailResponse{
		Success: true,
	}

	if bytes, err := eventkey.CustomerForgotPassword.Marshal(out); err != nil {
		logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CustomerForgotPassword, err.Error())
		res.Success = false
	} else {
		// publish event
		err = l.svcCtx.EventManager.Publish(eventkey.CustomerForgotPassword.AsKey(), bytes)
		if err != nil {
			logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CustomerForgotPassword, err.Error())
			res.Success = false
		}
	}

	return res, err
}
