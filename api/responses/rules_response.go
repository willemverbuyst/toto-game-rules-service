package responses

type RuleResponse struct {
	Status  int                    `json:"status" example:"200"`
	Message string                 `json:"message" example:"success"`
	Data    map[string]interface{} `json:"data"`
}
