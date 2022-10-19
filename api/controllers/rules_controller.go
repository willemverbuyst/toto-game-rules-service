package controllers

import (
	"context"
	"net/http"
	"time"
	"toto-game-rules-service/api/configs"
	"toto-game-rules-service/api/models"
	"toto-game-rules-service/api/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var rulesCollection *mongo.Collection = configs.GetCollection(configs.DB, "rules")
var validate = validator.New()

func GetRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("id")
		var rule models.Rule
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ruleId)

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
			responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": rules, "results": len(rules)}})
	}
}

func AddRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var rule models.Rule
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&rule); err != nil {
			c.JSON(http.StatusBadRequest, responses.RuleResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&rule); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.RuleResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newRule := models.Rule{
			Id:       primitive.NewObjectID(),
			Question: rule.Question,
			Answers:  rule.Answers,
		}

		result, err := rulesCollection.InsertOne(ctx, newRule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RuleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.RuleResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DeleteRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("id")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ruleId)

		result, err := rulesCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.RuleResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.RuleResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Rule with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Rule successfully deleted!"}},
		)
	}
}
