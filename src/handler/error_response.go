package handler

import (
	"encoding/json"
	"net/http"
)

func HandleErrorResponse(writer http.ResponseWriter, responseError *ErrorResponse) {
	writer.WriteHeader(responseError.GetCode())
	json.NewEncoder(writer).Encode(
		ErrorResponse{
			Code:    responseError.GetCode(),
			Message: responseError.Error(),
		})
}

func HandleUnexpectedErrorResponse(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(
		ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Unexpected error",
		})
}
