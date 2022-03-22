package customers

import (
	"k8scommerce/services/api/client/internal/types"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJwt(accessExpire int64, accessSecret string, payload map[string]interface{}) (*types.JwtToken, error) {

	now := time.Now().Unix()
	accessToken, err := genToken(now, accessSecret, payload, accessExpire)
	if err != nil {
		return nil, err
	}

	return &types.JwtToken{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func genToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
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
