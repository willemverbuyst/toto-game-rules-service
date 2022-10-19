package routes

import (
	"toto-game-rules-service/api/controllers"

	"github.com/gin-gonic/gin"
)

func RulesRoute(router *gin.Engine) {
	router.GET("/rules", controllers.GetAllRules())
	router.GET("/rules/:id", controllers.GetRule())
	router.POST("/rules", controllers.AddRule())
	router.DELETE("rules/:id", controllers.DeleteRule())
	router.PUT("rules/:id", controllers.UpdateRule())
}
