package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"

	"github.com/k8scommerce/k8scommerce/internal/encryption"
	encryptionconfig "github.com/k8scommerce/k8scommerce/internal/encryption/config"
)

type StoreKeyMiddleware struct {
	encrypter encryption.Encrypter
}

func NewStoreKeyMiddleware(config encryptionconfig.EncryptionConfig) *StoreKeyMiddleware {
	return &StoreKeyMiddleware{
		encrypter: encryption.NewEncrypter(&config),
	}
}

func (m *StoreKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if result, ok := r.Header["Store-Key"]; ok {
			if len(result) == 0 {
				http.Error(w, "error: invalid Store-Key value", http.StatusUnauthorized)
				return
			}

			var err error
			var id int64 = 0
			var decryptedStr string

			ctx := r.Context()
			ctx = context.WithValue(ctx, types.StoreKey, id) // add a default value

			decryptedStr, err = m.encrypter.Decrypt(result[0])
			if err == nil {
				id, err = strconv.ParseInt(decryptedStr, 10, 64)
				if err == nil {
					ctx = context.WithValue(ctx, types.StoreKey, id)
					next(w, r.WithContext(ctx))
					return
				}
			}
		}

		http.Error(w, "error: missing Store-Key header", http.StatusUnauthorized)
	}
}
