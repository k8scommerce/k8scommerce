package session

import (
	"encoding/base64"
	"encoding/binary"

	"k8scommerce/internal/encryption"
)

func NewSession(encrypter encryption.Encrypter, sessionId string) Session {
	return &session{
		encrypter: encrypter,
		sessionId: sessionId,
	}
}

type Session interface {
	GenSessionId(customerId int64) string
	GetCustomerId() int64
}

type session struct {
	encrypter encryption.Encrypter
	sessionId string
}

func (s *session) GenSessionId(customerId int64) string {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(customerId))
	encoded, _ := s.encrypter.Encrypt(base64.StdEncoding.EncodeToString(bytes))
	return encoded
}

func (s *session) GetCustomerId() int64 {
	bytes, _ := base64.StdEncoding.DecodeString(s.sessionId)
	decoded, _ := s.encrypter.Decrypt(string(bytes))
	customerId := int64(binary.LittleEndian.Uint64([]byte(decoded)))
	return customerId
}
