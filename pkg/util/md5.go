package util

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func EncodeMD5(value string) string {
	m := md5.New()
	// 将当前的string数据写到哈希中
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
