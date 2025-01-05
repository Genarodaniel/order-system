package response

type ResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}
