package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

type CustomEncrypt struct {
	failureCount int
}

func New() *CustomEncrypt {
	return &CustomEncrypt{failureCount: 0}
}

func (ce *CustomEncrypt) EncryptStrings(strings []string, method string) ([]string, error) {
	var encryptedStrings []string
	for _, str := range strings {
		encrypted, err := ce.encrypt(str, method)
		if err != nil {
			ce.failureCount++
			continue
		}
		encryptedStrings = append(encryptedStrings, encrypted)
	}
	return encryptedStrings, nil
}

func (ce *CustomEncrypt) EncryptMap(data map[string]string, method string) (map[string]string, error) {
	encryptedMap := make(map[string]string)
	for key, value := range data {
		encrypted, err := ce.encrypt(value, method)
		if err != nil {
			ce.failureCount++
			continue
		}
		encryptedMap[key] = encrypted
	}
	return encryptedMap, nil
}

func (ce *CustomEncrypt) GetFailureCount() int {
	return ce.failureCount
}

func (ce *CustomEncrypt) encrypt(data, method string) (string, error) {
	switch method {
	case "sha256":
		return sha256Encrypt(data), nil
	case "md5":
		return md5Encrypt(data), nil
	case "base64":
		return base64Encrypt(data), nil
	default:
		return "", errors.New("unsupported encryption method")
	}
}

func sha256Encrypt(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func md5Encrypt(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func base64Encrypt(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func main() {
	ce := New()

	stringsToEncrypt := []string{"test1", "test1"}
	encryptedStrings, _ := ce.EncryptStrings(stringsToEncrypt, "base64")
	fmt.Println("Encrypted Strings:", encryptedStrings)

	dataToEncrypt := map[string]string{"key1": "value1", "key2": "value2"}
	encryptedMap, _ := ce.EncryptMap(dataToEncrypt, "md5")
	fmt.Println("Encrypted Map:", encryptedMap)

	fmt.Println("Failure Count:", ce.GetFailureCount())
}
