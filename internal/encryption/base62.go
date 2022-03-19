package encryption

import (
	"encoding/hex"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// great article on creating base-62 strings
// https://medium.com/@anabhishek.jha/base-62-text-encoding-decoding-b43921c7a954

func NewBase62() Base62 {
	return &base62{
		encodingChunkSize: 2,
		base:              62,
		characterSet:      "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		decodingChunkSize: 3,
	}
}

type Base62 interface {
	Encode(str string) string
	Decode(encoded string) (string, error)
}

type base62 struct {
	encodingChunkSize int
	base              uint64
	characterSet      string
	decodingChunkSize int
}

func (e *base62) Encode(str string) string {
	var encoded strings.Builder

	inBytes := []byte(str)
	byteLength := len(inBytes)

	for i := 0; i < byteLength; i += e.encodingChunkSize {
		chunk := inBytes[i:e.minOf(i+e.encodingChunkSize, byteLength)]
		s := hex.EncodeToString(chunk)
		val, _ := strconv.ParseUint(s, 16, 64)
		w := e.padLeft(e.toBase62(val), "0", e.decodingChunkSize)
		encoded.WriteString(w)
	}
	return encoded.String()
}

func (e *base62) Decode(encoded string) (string, error) {
	decodedBytes := []byte{}
	for i := 0; i < len(encoded); i += e.decodingChunkSize {
		chunk := encoded[i:e.minOf(i+e.decodingChunkSize, len(encoded))]
		val, err := e.fromBase62(chunk)
		if err != nil {
			return "", err
		}
		chunkHex := strconv.FormatUint(val, 16)
		dst := make([]byte, hex.DecodedLen(len([]byte(chunkHex))))
		_, err = hex.Decode(dst, []byte(chunkHex))
		if err != nil {
			return "", errors.Wrap(err, "malformed input")
		}
		decodedBytes = append(decodedBytes, dst...)
	}
	s := string(decodedBytes)
	return s, nil
}

func (e *base62) minOf(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (e *base62) padLeft(str, pad string, length int) string {
	for len(str) < length {
		str = pad + str
	}
	return str
}

func (e *base62) toBase62(num uint64) string {
	encoded := ""
	for num > 0 {
		r := num % e.base
		num /= e.base
		encoded = string(e.characterSet[r]) + encoded

	}
	return encoded
}

func (e *base62) fromBase62(encoded string) (uint64, error) {
	var val uint64
	for index, char := range encoded {
		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(e.characterSet, char)
		if pos == -1 {
			return 0, errors.New("invalid character: " + string(char))
		}

		val += uint64(pos) * uint64(math.Pow(float64(e.base), float64(pow)))
	}

	return val, nil
}
