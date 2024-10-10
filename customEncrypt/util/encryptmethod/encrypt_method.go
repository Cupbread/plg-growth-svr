package encryptmethod

import (
	"crypto/sha256"
	"customEncrypt/util"
	"encoding/base64"
	"encoding/hex"
)

func Sha256Encrypt(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func Md5Encrypt(data string) (string, error) {
	hash, err := util.Md5([]byte(data))
	return hex.EncodeToString(hash[:]), err
}

func Base64Encrypt(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
