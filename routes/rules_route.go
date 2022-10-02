package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Answer struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type Rule struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []Answer `json:"answers"`
}

func getRules(context *gin.Context) {
	content, err := os.ReadFile("./rules.json")
	if err != nil {
		log.Fatal("error when opening file: ", err)
	}

	var payload []Rule
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("error during Unmarshall(): ", err)
	}

	context.IndentedJSON(http.StatusOK, payload)

}

func RulesRoute(router *gin.Engine) {
	router.GET("/rules", getRules)
}
