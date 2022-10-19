package controllers

import (
	"net/http"
	"toto-game-rules-service/api/responses"

	"github.com/gin-gonic/gin"
)

// CheckRoot godoc
// @Summary Test root
// @Description Responds with "Hello world" message.
// @Tags root
// @Produce json
// @Success 200 {object} responses.RootResponse
// @Router / [get]
func CheckRoot() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, responses.RootResponse{Status: http.StatusOK, Message: "Hello world"})
	}
}
