package handler


type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (errorResponse *ResponseError) Error() string {
	return errorResponse.Message
}
func (errorResponse *ResponseError) GetCode() int {
	return errorResponse.Code
}
