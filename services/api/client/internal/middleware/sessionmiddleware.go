package middleware

import (
	"context"
	"net/http"

	"k8scommerce/internal/encryption"
	"k8scommerce/internal/session"
	"k8scommerce/services/api/client/internal/types"
)

type SessionMiddleware struct {
	encrypter encryption.Encrypter
}

func NewSessionMiddleware(encrypter encryption.Encrypter) *SessionMiddleware {
	return &SessionMiddleware{
		encrypter: encrypter,
	}
}

func (m *SessionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if result, ok := r.Header["Session"]; ok {
			if len(result) == 0 {
				http.Error(w, "error: invalid session value", http.StatusUnauthorized)
				return
			}

			var err error
			var sessionId string

			ctx := r.Context()
			ctx = context.WithValue(ctx, types.Session, session.NewSession(m.encrypter, "")) // add a default value
			sessionId, err = m.encrypter.Decrypt(result[0])
			if err == nil {
				sess := session.NewSession(m.encrypter, sessionId)
				ctx = context.WithValue(ctx, types.Session, sess)
				next(w, r.WithContext(ctx))
				return
			}
		}

		http.Error(w, "error: missing Session header", http.StatusUnauthorized)
	}
}
