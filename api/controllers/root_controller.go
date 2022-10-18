package controllers

import (
	"net/http"
	"toto-game-rules-service/api/responses"

	"github.com/gin-gonic/gin"
)

func CheckRoot() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, responses.RootResponse{Status: http.StatusOK, Message: "Hello world"})
	}
}
