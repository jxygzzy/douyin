package md5util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func MD5WithSalt(s string, salt string) string {
	sum := md5.Sum([]byte(s + salt))
	return hex.EncodeToString(sum[:])
}
