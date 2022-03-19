package middleware

import (
	"context"
	"fmt"
	"net/http"

	"k8scommerce/services/api/client/internal/config"
	"k8scommerce/services/api/client/internal/types"

	"github.com/golang-jwt/jwt/v4"
)

type StoreKeyMiddleware struct {
	// hashCoder utils.HashCoder
	hashSalt string
	config   config.Config
}

func NewStoreKeyMiddleware(c config.Config) *StoreKeyMiddleware {
	// hashCoder := utils.NewHashCoder(c.hashSalt, utils.Store)

	return &StoreKeyMiddleware{
		// hashCoder: hashCoder,
		hashSalt: c.HashSalt,
		config:   c,
	}
}

func (m *StoreKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if result, ok := r.Header["Store-Key"]; ok {
			if len(result) == 0 {
				http.Error(w, "error: invalid Store-Key value", http.StatusUnauthorized)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, types.StoreKey, int64(0))

			token, err := jwt.ParseWithClaims(result[0], &types.StoreKeyClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(m.hashSalt), nil
			})

			var storeId int64 = 0

			if claims, ok := token.Claims.(*types.StoreKeyClaims); ok && token.Valid {
				storeId = claims.StoreId
				// fmt.Printf("%v %v", claims.StoreId, claims.RegisteredClaims.Issuer)
			} else {

				http.Error(w, fmt.Sprintf("error: invalid Store-Key header: %s", err.Error()), http.StatusUnauthorized)
				return
			}

			if storeId != 0 {
				ctx = context.WithValue(ctx, types.StoreKey, storeId)
				next(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "error: missing Store-Key header", http.StatusUnauthorized)
	}
}
