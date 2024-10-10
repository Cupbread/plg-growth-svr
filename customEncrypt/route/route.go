package route

import (
	"customEncrypt/controller"
	"customEncrypt/dao/model"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(router *gin.Engine, ce *model.CustomEncrypt) {

	encryptController := controller.NewEncryptController(ce)
	router.POST("/encrypt/string", encryptController.EncryptString)
	router.POST("/encrypt/mapstring", encryptController.EncryptMapString)

	errorController := controller.NewErrorController(ce)
	router.POST("/encrypt/error", errorController.GetFailureInfo)

}
