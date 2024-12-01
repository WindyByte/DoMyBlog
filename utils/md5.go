package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(buff string) string {
	hash := md5.New()
	hash.Write([]byte(buff))
	rawData := hash.Sum(nil)
	return hex.EncodeToString(rawData)
}
