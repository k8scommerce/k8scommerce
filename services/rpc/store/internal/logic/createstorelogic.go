package logic

import (
	"context"

	"github.com/k8scommerce/k8scommerce/services/rpc/store/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/pb/store"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStoreLogic {
	return &CreateStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStoreLogic) CreateStore(in *store.CreateStoreRequest) (*store.CreateStoreResponse, error) {
	res := &store.CreateStoreResponse{}
	return res, nil
}
