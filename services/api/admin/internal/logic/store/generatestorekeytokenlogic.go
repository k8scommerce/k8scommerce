package store

import (
	"context"
	"strconv"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
	"github.com/k8scommerce/k8scommerce/services/rpc/store/storeclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateStoreKeyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateStoreKeyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) GenerateStoreKeyTokenLogic {
	return GenerateStoreKeyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateStoreKeyTokenLogic) GenerateStoreKeyToken(req types.GenerateStoreKeyTokenRequest) (resp *types.GenerateStoreKeyTokenResponse, err error) {

	result, err := l.svcCtx.StoreRpc.GetStoreById(l.ctx, &storeclient.GetStoreByIdRequest{
		Id: req.StoreId,
	})
	if err != nil {
		return nil, err
	}

	token, err := l.svcCtx.Encrypter.Encrypt(strconv.Itoa(int(result.Store.Id)))
	resp = &types.GenerateStoreKeyTokenResponse{
		Token: token,
	}
	return resp, err
}
