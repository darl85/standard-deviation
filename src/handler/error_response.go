package handler

import (
	"encoding/json"
	"net/http"
)

func HandleErrorResponse(writer http.ResponseWriter, responseError *ResponseError) {
	writer.WriteHeader(responseError.GetCode())
	json.NewEncoder(writer).Encode(
		ResponseError{
			Code:    responseError.GetCode(),
			Message: responseError.Error(),
		})
}

func HandleUnexpectedErrorResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(
		ResponseError{
			Code:    http.StatusInternalServerError,
			Message: "Unexpected error",
		})
}
