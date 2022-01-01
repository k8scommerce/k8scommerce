package logic

// import (
// 	"context"

// 	"ecomm/services/api/client/internal/svc"
// 	"ecomm/services/api/client/internal/types"
// 	"ecomm/services/rpc/cart/cartclient"

// 	"github.com/tal-tech/go-zero/core/logx"
// )

// type GetCartSessionLogic struct {
// 	logx.Logger
// 	ctx    context.Context
// 	svcCtx *svc.ServiceContext
// }

// func NewGetCartSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCartSessionLogic {
// 	return GetCartSessionLogic{
// 		Logger: logx.WithContext(ctx),
// 		ctx:    ctx,
// 		svcCtx: svcCtx,
// 	}
// }

// func (l *GetCartSessionLogic) GetCartSession(req types.GetSessionRequest) (*types.GetSessionResponse, error) {
// 	response, err := l.svcCtx.CartRpc.GetSession(l.ctx, &cartclient.GetSessionRequest{
// 		UserId: 0,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := &types.GetSessionResponse{
// 		SessionId: response.SessionId,
// 	}
// 	return res, nil
// }
