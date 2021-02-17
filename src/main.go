package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "standard-deviation/src/calculations"
    "standard-deviation/src/handler"
    "standard-deviation/src/numbers"
    "standard-deviation/src/random_api"
)

type ApiResponseError struct {
    code    int
    message string
}

func (responseError *ApiResponseError) Error() string {
    return responseError.message
}
func (responseError *ApiResponseError) GetCode() int {
    return responseError.code
}

func meanHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    requests, numberOfIntegers, paramsValidationError := handler.HandleQueryParameters(request)

    if paramsValidationError != nil {
        handler.HandleErrorResponse(writer, paramsValidationError)
        return
    }

    numberSetsCollection, collectNumberError := numbers.CollectNumberSets(requests, numberOfIntegers, random_api.RandomApiClient)
    if collectNumberError != nil {
        handler.HandleErrorResponse(writer, &ApiResponseError{
            code:    collectNumberError.GetCode(),
            message: collectNumberError.Error(),
        })
        return
    }

    deviations := calculations.GetStandardDeviations(numberSetsCollection, &calculations.StandardDeviationCalculator)
    handler.HandleSuccessResponse(writer, deviations)
}

func main() {
    apiRouter := mux.NewRouter().StrictSlash(true)
    apiRouter.HandleFunc("/random/mean", meanHandler).
        Methods("GET")
    http.ListenAndServe(":8080", apiRouter)
}

