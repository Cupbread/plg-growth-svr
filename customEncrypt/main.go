package main

import (
	"customEncrypt/dao/model"
	"customEncrypt/route"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	ce := model.New()

	route.RegisterWebRoutes(r, ce)

	r.Run(":8080")
}
