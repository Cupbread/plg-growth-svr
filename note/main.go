package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"note/configs"
	"note/route"
	"note/storge/mysql"
)

func main() {
	configs.Setup("./configs/config.yaml")
	configs.SetApplicationIsInit()
	mysql.Setup()

	if mysql.DBGorm == nil {
		log.Fatal("Database connection failed")
	}

	router := gin.Default()
	route.RegisterWebRoutes(router)

	router.Run(":8080")
}
