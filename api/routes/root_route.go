package routes

import (
	"toto-game-rules-service/api/controllers"

	"github.com/gin-gonic/gin"
)

func RootRoute(router *gin.Engine) {
	router.GET("/", controllers.CheckRoot())
}
