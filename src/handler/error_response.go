package handler

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

func HandleErrorResponse(writer http.ResponseWriter, responseError *ApiResponseError) {
	writer.WriteHeader(responseError.Code)
	json.NewEncoder(writer).Encode(
		ErrorResponse{
			Code: responseError.Code,
			Message: responseError.Error(),
		})
}
