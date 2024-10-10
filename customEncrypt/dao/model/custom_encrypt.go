package model

import (
	"customEncrypt/util/encryptmethod"
	"errors"
)

type CustomEncrypt struct {
	failureNum map[string]int
}

func New() *CustomEncrypt {
	return &CustomEncrypt{
		failureNum: make(map[string]int),
	}
}

func (ce *CustomEncrypt) EncryptStrings(strings []string, method string, userid string) []string {
	var encryptedStrings []string
	for _, str := range strings {
		if len(str) == 0 {
			ce.failureNum[userid]++
			continue
		}
		encrypted, err := ce.encrypt(str, method)
		if err != nil {
			ce.failureNum[userid]++
			continue
		}
		encryptedStrings = append(encryptedStrings, encrypted)
	}
	return encryptedStrings
}

func (ce *CustomEncrypt) EncryptMap(data map[string]string, method string, userid string) map[string]string {
	encryptedMap := make(map[string]string)
	for key, value := range data {
		encrypted, err := ce.encrypt(value, method)
		if err != nil {
			ce.failureNum[userid]++
			continue
		}
		encryptedMap[key] = encrypted
	}
	return encryptedMap
}

func (ce *CustomEncrypt) GetFailureNum() map[string]int {
	return ce.failureNum
}

func (ce *CustomEncrypt) encrypt(data, method string) (string, error) {
	switch method {
	case "sha256":
		//return sha256Encrypt(data), nil
		return encryptmethod.Sha256Encrypt(data), nil
	case "md5":
		hash, err := encryptmethod.Md5Encrypt(data)
		return hash, err
	case "base64":
		return encryptmethod.Base64Encrypt(data), nil
	default:
		return "", errors.New("unsupported encryption method")
	}
}
