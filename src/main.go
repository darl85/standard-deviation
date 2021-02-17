package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "standard-deviation/src/calculations"
    "standard-deviation/src/handler"
    "standard-deviation/src/numbers"
    "standard-deviation/src/random_api"
)

func meanHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    requests, numberOfIntegers, paramsValidationError := handler.HandleQueryParameters(request)

    if paramsValidationError != nil {
        if assertErr, ok := paramsValidationError.(*handler.ApiResponseError); ok {
            handler.HandleErrorResponse(writer, &handler.ErrorResponse{
                Code:    assertErr.GetCode(),
                Message: assertErr.Error(),
            })
        }
        return
    }

    numberSetsCollection, collectNumberError := numbers.CollectNumberSets(requests, numberOfIntegers, random_api.RandomApiClient)
    if collectNumberError != nil {
        if assertcollectNumberError, ok := collectNumberError.(*random_api.ClientError); ok {
            handler.HandleErrorResponse(writer, &handler.ErrorResponse{
                Code:    assertcollectNumberError.GetCode(),
                Message: assertcollectNumberError.Error(),
            })
        }
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

