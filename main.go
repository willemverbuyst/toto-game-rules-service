package main

import (
	"log"
	"toto-game-rules-service/api/configs"
	"toto-game-rules-service/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())
	routes.RootRoute(router)
	routes.RulesRoute(router)

	err := router.Run(":9090")

	if err != nil {
		log.Fatal("error when running the server: ", err)
	}
}
