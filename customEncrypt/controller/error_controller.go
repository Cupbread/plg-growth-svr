package controller

import (
	"customEncrypt/dao/model"
	"github.com/gin-gonic/gin"
)

type ErrorController struct {
	customEncrypt *model.CustomEncrypt
}

func NewErrorController(ce *model.CustomEncrypt) *ErrorController {
	return &ErrorController{
		customEncrypt: ce,
	}
}

func (ec *ErrorController) GetFailureInfo(c *gin.Context) {
	failureNum := ec.customEncrypt.GetFailureNum()

	c.JSON(200, gin.H{
		"failureNum": failureNum,
	})
}
