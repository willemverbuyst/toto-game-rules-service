package main

import (
	"log"
	"toto-game-rules-service/api/configs"
	"toto-game-rules-service/api/routes"
	_ "toto-game-rules-service/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Toto Game Rule Service API
// @version 1.0
// @description Toto Game Rule Service API in Go using Gin framework.
// @host localhost:9090
func main() {
	configs.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())
	routes.RootRoute(router)
	routes.RulesRoute(router)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":9090")

	if err != nil {
		log.Fatal("error when running the server: ", err)
	}
}
