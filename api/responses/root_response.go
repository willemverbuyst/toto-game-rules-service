package responses

type RootResponse struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"success"`
}
