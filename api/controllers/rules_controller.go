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
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": rule}})
	}
}

// GetRules godoc
// @Summary Get rules array
// @Description Responds with the list of all rules as JSON.
// @Tags rules
// @Produce json
// @Success 200 {object} responses.RulesResponse
// @Router /rules [get]
func GetRules() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var rules []models.Rule
		defer cancel()

		results, err := rulesCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleRule models.Rule
			if err = results.Decode(&singleRule); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			}

			rules = append(rules, singleRule)
		}

		c.JSON(http.StatusOK,
			responses.RulesResponse{Status: http.StatusOK, Message: "success", Data: rules, Results: len(rules)})
	}
}

func AddRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var rule models.Rule
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&rule); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&rule); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Data: map[string]interface{}{"error": validationErr.Error()}})
			return
		}

		newRule := models.Rule{
			Id:       primitive.NewObjectID(),
			Question: rule.Question,
			Answers:  rule.Answers,
		}

		result, err := rulesCollection.InsertOne(ctx, newRule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
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
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.ErrorResponse{Status: http.StatusNotFound, Message: "fail", Data: map[string]interface{}{"error": "Rule with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Rule successfully deleted!"}},
		)
	}
}

func UpdateRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("id")
		var rule models.Rule
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ruleId)

		//validate the request body
		if err := c.BindJSON(&rule); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&rule); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Data: map[string]interface{}{"error": validationErr.Error()}})
			return
		}

		update := bson.M{"order": rule.Order, "question": rule.Question, "answers": rule.Answers}
		result, err := rulesCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
			return
		}

		//get updated rule details
		var updatedRule models.Rule
		if result.MatchedCount == 1 {
			err := rulesCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedRule)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Data: map[string]interface{}{"error": err.Error()}})
				return
			}
		}

		c.JSON(http.StatusOK, responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedRule}})
	}
}
