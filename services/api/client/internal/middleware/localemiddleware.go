package middleware

import (
	"context"
	"net/http"

	"k8scommerce/services/api/client/internal/config"
	"k8scommerce/services/api/client/internal/types"
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
		ctx = context.WithValue(ctx, types.ClientLocale, "en")

		if langs, ok := r.Header["Accept-Language"]; ok {
			// m.Config.Locale = langs[0]
			ctx = context.WithValue(ctx, types.ClientLocale, langs[0])
		}
		next(w, r.WithContext(ctx))
	}
}
