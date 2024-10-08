package main

import (
	"crypto/sha256"
	"customEncrypt/util"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
)

type CustomEncrypt struct {
	failureCount int
	failedData   map[string]error
}

func New() *CustomEncrypt {
	return &CustomEncrypt{
		failureCount: 0,
		failedData:   make(map[string]error),
	}
}

func (ce *CustomEncrypt) EncryptStrings(strings []string, method string) []string {
	var encryptedStrings []string
	for _, str := range strings {
		encrypted, err := ce.encrypt(str, method)
		if err != nil {
			ce.failureCount++
			ce.failedData[str] = err
			continue
		}
		encryptedStrings = append(encryptedStrings, encrypted)
	}
	return encryptedStrings
}

func (ce *CustomEncrypt) EncryptMap(data map[string]string, method string) map[string]string {
	encryptedMap := make(map[string]string)
	for key, value := range data {
		encrypted, err := ce.encrypt(value, method)
		if err != nil {
			ce.failureCount++
			ce.failedData[key] = err
			continue
		}
		encryptedMap[key] = encrypted
	}
	return encryptedMap
}

func (ce *CustomEncrypt) GetFailureInfo() (int, map[string]error) {
	return ce.failureCount, ce.failedData
}

func (ce *CustomEncrypt) encrypt(data, method string) (string, error) {
	switch method {
	case "sha256":
		return sha256Encrypt(data), nil
	case "md5":
		hash, err := md5Encrypt(data)
		return hash, err
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

func md5Encrypt(data string) (string, error) {
	hash, err := util.Md5([]byte(data))
	return hex.EncodeToString(hash[:]), err
}

func base64Encrypt(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func main() {
	r := gin.Default()
	ce := New()
	r.POST("/encrypt/mapstring", func(c *gin.Context) {
		var requestData struct {
			DataToEncrypt  map[string]string `json:"dataToEncrypt"`
			EncryptionType string            `json:"encryptionType"`
		}

		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		encryptedMap := ce.EncryptMap(requestData.DataToEncrypt, requestData.EncryptionType)

		c.JSON(200, gin.H{"encryptedMap": encryptedMap})
	})

	r.POST("/encrypt/string", func(c *gin.Context) {
		var requestData struct {
			Strings        []string `json:"strings"`
			EncryptionType string   `json:"encryptionType"`
		}

		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}

		encryptedStrings := ce.EncryptStrings(requestData.Strings, requestData.EncryptionType)

		c.JSON(200, gin.H{"encryptedStrings": encryptedStrings})
	})

	r.POST("/encrypt/error", func(c *gin.Context) {
		failureCount, failedData := ce.GetFailureInfo()

		failedDataString := make(map[string]string)
		for key, value := range failedData {
			failedDataString[key] = value.Error()
		}

		c.JSON(200, gin.H{
			"failureCount": failureCount,
			"failedData":   failedDataString,
		})
	})

	r.Run(":8080")
}
