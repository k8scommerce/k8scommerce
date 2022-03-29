package middleware

import (
	"context"
	"net/http"

	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/config"
	"github.com/k8scommerce/k8scommerce/services/api/admin/internal/types"
)

type FilterMiddleware struct {
	Config *config.Config
}

func NewFilterMiddleware() *FilterMiddleware {
	return &FilterMiddleware{}
}

func (m *FilterMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, types.Filter, r.RequestURI)
		next(w, r.WithContext(ctx))
	}
}
