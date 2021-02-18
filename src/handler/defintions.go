package handler


type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (error *ErrorResponse) Error() string {
	return error.Message
}
func (error *ErrorResponse) GetCode() int {
	return error.Code
}
