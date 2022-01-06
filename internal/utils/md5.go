package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// StringToMD5 - Converts a string to md5
func StringToMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
