package handler

import (
	"net/http"
	"strconv"
)

type ApiResponseError struct {
	code    int
	message string
}

func (error *ApiResponseError) Error() string {
	return error.message
}
func (error *ApiResponseError) GetCode() int {
	return error.code
}

func HandleQueryParameters(request *http.Request) (int, int, error) {
	requests, requestsParamError := strconv.Atoi(request.URL.Query().Get("requests"))
	numberOfIntegers, lengthParamError := strconv.Atoi(request.URL.Query().Get("length"))

	if requestsParamError != nil || lengthParamError != nil {
		return 0, 0, &ApiResponseError{
			code:    http.StatusBadRequest,
			message: "Parameters are invalid, pass integers for both params",
		}
	}

	if requests < 0 || numberOfIntegers < 0 {
		return 0, 0, &ApiResponseError{
			code:    http.StatusBadRequest,
			message: "Parameters cannot be negative",
		}
	}

	return requests, numberOfIntegers, nil
}
