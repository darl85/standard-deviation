package main

import (
    "net/http"
    "standard-deviation/src/random_api"
)


func hello(writer http.ResponseWriter, request *http.Request) {
    random_api.getRandomIntegers(5, 1, 100)
}

func main() {

    http.HandleFunc("/hello", hello)

    http.ListenAndServe(":8080", nil)
}
