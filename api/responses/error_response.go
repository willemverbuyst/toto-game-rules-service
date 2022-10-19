package responses

type ErrorResponse struct {
	Status  int    `json:"status" example:"500"`
	Message string `json:"message" example:"fail"`
	Error   string `json:"error" example:"the error"`
}
