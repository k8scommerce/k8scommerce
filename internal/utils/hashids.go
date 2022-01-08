package utils

import (
	"strconv"

	"github.com/speps/go-hashids"
)

// ModelKind ...
type ModelKind int

const (
	Category ModelKind = iota
	Customer
	Product
	User
	Variant
	Store
)

func (t ModelKind) String() string {
	return [...]string{
		"Category",
		"Customer",
		"Product",
		"User",
		"Variant",
		"Store",
	}[t]
}

func NewHashCoder(hashSalt string, modelKind ModelKind) HashCoder {
	data := hashids.NewData()
	data.Alphabet = "abcdefghijklmnopqrstuvzyx"
	data.Salt = strconv.Itoa(int(modelKind)) + StringToMD5(hashSalt)
	data.MinLength = 8
	hashID, _ := hashids.NewWithData(data)

	return &hashCoder{
		data:   hashids.NewData(),
		hashID: hashID,
	}
}

type HashCoder interface {
	Encode(id int64) string
	Decode(encodedHash string) int64
}

// hashCoder ...
type hashCoder struct {
	data   *hashids.HashIDData
	hashID *hashids.HashID
}

// Encode ...
func (h *hashCoder) Encode(id int64) string {
	encoded, _ := h.hashID.EncodeInt64([]int64{id})
	return encoded
}

// Decode ...
func (h *hashCoder) Decode(encodedHash string) int64 {
	decoded, _ := h.hashID.DecodeInt64WithError(encodedHash)
	if len(decoded) > 0 {
		return decoded[0]
	}
	return 0
}
