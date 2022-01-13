package middleware

import (
	"context"
	"k8scommerce/internal/utils"
	"k8scommerce/services/api/client/internal/types"
	"net/http"
)

type StoreKeyMiddleware struct {
	hashCoder utils.HashCoder
}

func NewStoreKeyMiddleware(hashSalt string) *StoreKeyMiddleware {
	hashCoder := utils.NewHashCoder(hashSalt, utils.Store)

	return &StoreKeyMiddleware{
		hashCoder: hashCoder,
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
			ctx = context.WithValue(ctx, types.StoreKey, int64(1))

			// decode the hashed ID
			id := m.hashCoder.Decode(result[0])
			if id != 0 {
				ctx = context.WithValue(ctx, types.StoreKey, id)
				next(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "error: missing Store-Key header :: "+m.hashCoder.Encode(1), http.StatusUnauthorized)
	}
}
