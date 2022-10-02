package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"toto-game/go-service/configs"

	"github.com/gin-contrib/cors"
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

func testRoot(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	configs.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", testRoot)
	router.GET("/rules", getRules)

	err := router.Run(":9090")

	if err != nil {
		log.Fatal("error when running the server: ", err)
	}
}
