package handler

import (
	"net/http"
	"strconv"
)

func HandleQueryParameters(request *http.Request) (int, int, error) {
	requests, requestsParamError := strconv.Atoi(request.URL.Query().Get("requests"))
	numberOfIntegers, lengthParamError := strconv.Atoi(request.URL.Query().Get("length"))

	if requestsParamError != nil || lengthParamError != nil {
		return 0, 0, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Parameters are invalid, pass integers for both params",
		}
	}

	if requests < 0 || numberOfIntegers < 0 {
		return 0, 0, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Parameters cannot be negative",
		}
	}

	return requests, numberOfIntegers, nil
}
