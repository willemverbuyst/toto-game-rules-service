package responses

type ErrorResponse struct {
	Status  int    `json:"status" example:"500"`
	Message string `json:"message" example:"fail"`
}

type ValidationErrorResponse struct {
	Status  int                    `json:"status" example:"500"`
	Message string                 `json:"message" example:"fail"`
	Data    map[string]interface{} `json:"data"`
}
