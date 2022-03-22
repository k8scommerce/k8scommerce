package types

import "github.com/golang-jwt/jwt/v4"

type StoreKeyClaims struct {
	StoreId int64  `json:"storeKey"`
	Url     string `json:"url"`
	jwt.RegisteredClaims
}
