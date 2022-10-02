package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testRoot(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func RootRoute(router *gin.Engine) {
	router.GET("/", testRoot)
}
