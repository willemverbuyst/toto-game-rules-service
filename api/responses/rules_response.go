package responses

import "toto-game-rules-service/api/models"

type RuleResponse struct {
	Status  int                    `json:"status" example:"200"`
	Message string                 `json:"message" example:"success"`
	Data    map[string]interface{} `json:"data"`
}

type RulesResponse struct {
	Status  int           `json:"status" example:"200"`
	Message string        `json:"message" example:"success"`
	Data    []models.Rule `json:"data"`
	Results int           `json:"results" example:"1"`
}
