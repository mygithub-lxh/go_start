package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算 md5 hash
func Md5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
