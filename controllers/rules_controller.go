package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"toto-game/go-service/configs"
	"toto-game/go-service/models"
	"toto-game/go-service/responses"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var rulesCollection *mongo.Collection = configs.GetCollection(configs.DB, "rules")

func GetARule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("ruleId")
		var rule models.Rule
		defer cancel()

		objId, _ := strconv.Atoi(ruleId)

		err := rulesCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&rule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RuleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": rule}})
	}
}

func GetAllRules() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var rules []models.Rule
		defer cancel()

		results, err := rulesCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RuleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleRule models.Rule
			if err = results.Decode(&singleRule); err != nil {
				c.JSON(http.StatusInternalServerError, responses.RuleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			rules = append(rules, singleRule)
		}

		c.JSON(http.StatusOK,
			responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": rules, "number": len(rules)}},
		)
	}
}
