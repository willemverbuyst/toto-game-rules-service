package routes

import (
	"toto-game-rules-service/api/controllers"

	"github.com/gin-gonic/gin"
)

func RulesRoute(router *gin.Engine) {
	router.GET("/rules", controllers.GetAllRules())
	router.GET("/rules/:ruleId", controllers.GetARule())
}