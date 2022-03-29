package middleware

import (
	"context"
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/config"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
)

type LocaleMiddleware struct {
	Config *config.Config
}

func NewLocaleMiddleware() *LocaleMiddleware {
	return &LocaleMiddleware{}
}

func (m *LocaleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, types.Locale, "en")

		if langs, ok := r.Header["Accept-Language"]; ok {
			ctx = context.WithValue(ctx, types.Locale, langs[0])
		}
		next(w, r.WithContext(ctx))
	}
}
