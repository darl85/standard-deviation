package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func HandleErrorResponse(writer http.ResponseWriter, responseError interface{ApiResponseErrorInterface}) {
	writer.WriteHeader(responseError.GetCode())
	json.NewEncoder(writer).Encode(
		ErrorResponse{
			Code: responseError.GetCode(),
			Message: responseError.Error(),
		})
}
