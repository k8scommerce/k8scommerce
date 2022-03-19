package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"k8scommerce/internal/encryption/config"
)

// Encode and decode text
// this is for sensitive data

func NewEncrypter(c *config.EncryptionConfig) Encrypter {
	return &encrypter{
		Secret: c.Secret,
		Token:  c.Token,
		base62: NewBase62(),
	}
}

type Encrypter interface {
	Encrypt(text string) (string, error)
	Decrypt(text string) (string, error)
}

type encrypter struct {
	Secret string
	Token  string
	base62 Base62
}

func (e *encrypter) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(e.Secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(e.Token))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return e.encode(cipherText), nil
}

func (e *encrypter) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(e.Secret))
	if err != nil {
		return "", err
	}
	cipherText := e.decode(text)
	cfb := cipher.NewCFBDecrypter(block, []byte(e.Token))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func (e *encrypter) encode(b []byte) string {
	buf := base64.StdEncoding.EncodeToString(b)
	return string(e.base62.Encode(buf))
}

func (e *encrypter) decode(s string) []byte {
	decoded, err := e.base62.Decode(s)
	if err != nil {
		panic(err)
	}

	data, err := base64.StdEncoding.DecodeString(decoded)
	if err != nil {
		panic(err)
	}

	return data
}
