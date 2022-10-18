package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"toto-game-rules-service/api/configs"
	"toto-game-rules-service/api/models"
	"toto-game-rules-service/api/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var rulesCollection *mongo.Collection = configs.GetCollection(configs.DB, "rules")

func GetARule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("id")
		var rule models.Rule
		defer cancel()

		objId, _ := strconv.Atoi(ruleId)

		err := rulesCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&rule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail"})
			return
		}

		c.JSON(http.StatusOK, responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: rule})
	}
}

func GetAllRules() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var rules []models.Rule
		defer cancel()

		results, err := rulesCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail"})
			return
		}

		fmt.Print(results)

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleRule models.Rule
			if err = results.Decode(&singleRule); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail"})
			}

			fmt.Print(singleRule)

			rules = append(rules, singleRule)
		}

		c.JSON(http.StatusOK,
			responses.RulesResponse{Status: http.StatusOK, Message: "success", Data: rules, Number: len(rules)},
		)
	}
}
