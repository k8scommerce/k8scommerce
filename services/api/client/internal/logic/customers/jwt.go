package customers

import (
	"time"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/svc"
	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

	"github.com/dgrijalva/jwt-go"
)

func genJwtToken(svcCtx *svc.ServiceContext, payloads map[string]interface{}) (*types.JwtToken, error) {
	now := time.Now().UnixMilli()
	accessExpire := svcCtx.Config.Auth.AccessExpire

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	accessToken, err := token.SignedString([]byte(svcCtx.Config.Auth.AccessSecret))

	return &types.JwtToken{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, err
}
