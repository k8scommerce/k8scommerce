package logic

import (
	"context"

	"k8scommerce/services/rpc/store/internal/svc"
	"k8scommerce/services/rpc/store/pb/store"

	"github.com/localrivet/galaxycache"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext, universe *galaxycache.Universe) *CreateStoreLogic {
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
