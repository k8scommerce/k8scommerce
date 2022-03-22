package customers

import (
	"context"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/customer/customerclient"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.CustomerLoginRequest) (resp *types.CustomerLoginResponse, err error) {
	resp = &types.CustomerLoginResponse{
		Success: false,
	}

	found, err := l.svcCtx.CustomerRpc.Login(l.ctx, &customerclient.LoginRequest{
		StoreId:  l.ctx.Value(types.StoreKey).(int64),
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	if found.Customer == nil || found.Customer.Id == 0 {
		return resp, nil
	}

	// create the token
	jwtToken, err := getJwt(
		l.svcCtx.Config.Auth.AccessExpire,
		l.svcCtx.Config.Auth.AccessSecret,
		map[string]interface{}{
			"customerId": found.Customer.Id,
		},
	)
	if err != nil {
		return nil, err
	}

	customer := types.Customer{
		FirstName: found.Customer.FirstName,
		LastName:  found.Customer.LastName,
		Email:     found.Customer.Email,
	}

	resp.JwtToken = *jwtToken
	resp.Customer = customer
	resp.Success = true

	return resp, nil
}

func (l *LoginLogic) getJwt(accessExpire int64, accessSecret string, payload map[string]interface{}) (*types.JwtToken, error) {

	now := time.Now().Unix()
	accessToken, err := l.genToken(now, l.svcCtx.Config.Auth.AccessSecret, payload, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtToken{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *LoginLogic) genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
