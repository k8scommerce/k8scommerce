package store

import (
	"context"
	"time"

	"k8scommerce/services/api/admin/internal/svc"
	"k8scommerce/services/api/admin/internal/types"
	"k8scommerce/services/rpc/store/storeclient"

	"github.com/golang-jwt/jwt/v4"
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
	resp = &types.GenerateStoreKeyTokenResponse{}

	result, err := l.svcCtx.StoreRpc.GetStoreById(l.ctx, &storeclient.GetStoreByIdRequest{
		Id: req.StoreId,
	})
	if err != nil {
		return nil, err
	}

	claims := types.StoreKeyClaims{
		StoreId: result.Store.Id,
		Url:     result.Store.Url,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(((24 * time.Hour) * 365) * 100)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "K8sCommerce",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(l.svcCtx.Config.HashSalt))
	resp.Token = ss
	return resp, err
}
