package util

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHashOfService(serviceString string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(serviceString))
	return hex.EncodeToString(algorithm.Sum(nil))
}
