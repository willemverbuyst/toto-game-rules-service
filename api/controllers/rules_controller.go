package controllers

import (
	"context"
	"fmt"
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

// GetRule godoc
// @Summary Get rule by ID
// @Description Responds with a rule as JSON
// @Tags rules
// @Accept json
// @Produce json
// @Param id path string true "Rule ID"
// @Success 200 {object} responses.RuleResponse
// @Failure 404 {object} responses.ErrorResponse
// @Router /rules/{id} [get]
func GetRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("id")
		var rule models.Rule
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(ruleId)

		fmt.Print(objId)

		err := rulesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&rule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: rule})
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
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleRule models.Rule
			if err = results.Decode(&singleRule); err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			}

			rules = append(rules, singleRule)
		}

		c.JSON(http.StatusOK,
			responses.RulesResponse{Status: http.StatusOK, Message: "success", Data: rules, Results: len(rules)})
	}
}

// AddRule godoc
// @Summary      Add Rule
// @Description  Responds with the rule created of as JSON.
// @Tags         rules
// @Accept       json
// @Produce      json
// @Param        user body models.Rule true "Add rule"
// @Success      201 {object} responses.RuleResponse
// @Router       /rules [post]
func AddRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var rule models.Rule
		defer cancel()

		// validate the request body
		if err := c.BindJSON(&rule); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Error: err.Error()})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&rule); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Error: validationErr.Error()})
			return
		}

		newRule := models.Rule{
			Id:       primitive.NewObjectID(),
			Question: rule.Question,
			Answers:  rule.Answers,
		}

		result, err := rulesCollection.InsertOne(ctx, newRule)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			return
		}

		var ruleInserted models.Rule
		error := rulesCollection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&ruleInserted)

		if error != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.RuleResponse{Status: http.StatusCreated, Message: "success", Data: ruleInserted})
	}
}

func DeleteRule() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		ruleId := c.Param("id")
		defer cancel()

		result, err := rulesCollection.DeleteOne(ctx, bson.M{"id": ruleId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.ErrorResponse{Status: http.StatusNotFound, Message: "fail", Error: "Rule with specified ID not found!"},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.RuleGeneralResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Rule successfully deleted!"}},
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
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Error: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&rule); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "fail", Error: validationErr.Error()})
			return
		}

		update := bson.M{"order": rule.Order, "question": rule.Question, "answers": rule.Answers}
		result, err := rulesCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
			return
		}

		//get updated rule details
		var updatedRule models.Rule
		if result.MatchedCount == 1 {
			err := rulesCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedRule)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: "fail", Error: err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, responses.RuleResponse{Status: http.StatusOK, Message: "success", Data: updatedRule})
	}
}
