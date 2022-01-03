package logic

import (
	"context"
	"time"

	"k8scommerce/internal/utils"
	"k8scommerce/services/api/client/internal/svc"
	"k8scommerce/services/api/client/internal/types"
	"k8scommerce/services/rpc/user/userclient"

	"github.com/golang-jwt/jwt"

	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LoginLogic) Login(req types.CustomerLoginRequest) (*types.CustomerLoginResponse, error) {
	// logx.Info("RECEIVED ", req.Email)

	res, err := l.svcCtx.UserRpc.Login(l.ctx, &userclient.LoginRequest{
		Username: req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// create the token
	jwtToken, err := l.getJwt(map[string]interface{}{
		"userId": res.User.Id,
	})
	if err != nil {
		return nil, err
	}

	customer := types.Customer{}
	utils.TransformObj(res.User, &customer)

	return &types.CustomerLoginResponse{
		JwtToken:      *jwtToken,
		Customer:      customer,
		StatusCode:    res.StatusCode,
		StatusMessage: res.StatusMessage,
	}, nil
}

func (l *LoginLogic) getJwt(payload map[string]interface{}) (*types.JwtToken, error) {
	var accessExpire = l.svcCtx.Config.Auth.AccessExpire

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
