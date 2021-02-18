package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "standard-deviation/src/calculations"
    "standard-deviation/src/handler"
    "standard-deviation/src/numbers"
    "standard-deviation/src/random_api"
    "time"
)

const timeout = 2 * time.Second

func meanHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    requests, numberOfIntegers, paramsValidationError := handler.HandleQueryParameters(request)

    if paramsValidationError != nil {
        if assertErr, ok := paramsValidationError.(*handler.ResponseError); ok {
            handler.HandleErrorResponse(writer, &handler.ResponseError{
                Code:    assertErr.GetCode(),
                Message: assertErr.Error(),
            })
        } else {
            handler.HandleUnexpectedErrorResponse(writer)
        }
        return
    }

    numberSetsCollection, collectNumberError := numbers.CollectNumberSets(
        requests,
        numberOfIntegers,
        random_api.RandomApiClient,
        timeout,
    )
    if collectNumberError != nil {
        if assertCollectNumberError, ok := collectNumberError.(*numbers.CollectingNumbersError); ok {
            handler.HandleErrorResponse(writer, &handler.ResponseError{
                Code:    assertCollectNumberError.GetCode(),
                Message: assertCollectNumberError.Error(),
            })
        } else {
            handler.HandleUnexpectedErrorResponse(writer)
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

