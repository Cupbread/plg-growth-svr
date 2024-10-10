package controller

import (
	"customEncrypt/dao/model"
	"github.com/gin-gonic/gin"
)

type EncryptController struct {
	customEncrypt *model.CustomEncrypt
}

func NewEncryptController(ce *model.CustomEncrypt) *EncryptController {
	return &EncryptController{
		customEncrypt: ce,
	}
}

func (ec *EncryptController) EncryptString(c *gin.Context) {
	var requestData struct {
		Strings        []string `json:"strings"`
		EncryptionType string   `json:"encryptionType"`
		UserID         string   `json:"userId"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if requestData.UserID == "" {
		c.JSON(400, gin.H{"error": "UserID is empty"})
		return
	}

	encryptedStrings := ec.customEncrypt.EncryptStrings(requestData.Strings, requestData.EncryptionType, requestData.UserID)

	c.JSON(200, gin.H{"encryptedStrings": encryptedStrings})
}

func (ec *EncryptController) EncryptMapString(c *gin.Context) {
	var requestData struct {
		DataToEncrypt  map[string]string `json:"dataToEncrypt"`
		EncryptionType string            `json:"encryptionType"`
		UserID         string            `json:"userId"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	if requestData.UserID == "" {
		c.JSON(400, gin.H{"error": "UserID is empty"})
		return
	}

	encryptedMap := ec.customEncrypt.EncryptMap(requestData.DataToEncrypt, requestData.EncryptionType, requestData.UserID)

	c.JSON(200, gin.H{"encryptedMap": encryptedMap})
}
