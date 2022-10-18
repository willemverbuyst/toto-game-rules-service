package responses

import "toto-game-rules-service/api/models"

type RuleResponse struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    models.Rule `json:"data"`
}

type RulesResponse struct {
	Status  int           `json:"status" example:"200"`
	Message string        `json:"message" example:"success"`
	Data    []models.Rule `json:"data"`
	Number  int           `json:"number" example:"10"`
}
