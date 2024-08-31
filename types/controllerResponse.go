package controller_response

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
