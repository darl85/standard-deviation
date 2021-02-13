package main

import (
    "fmt"
    "net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {

    fmt.Fprintf(writer, "hello\n")
}

func headers(writer http.ResponseWriter, request *http.Request) {

    for name, headers := range request.Header {
        for _, h := range headers {
            fmt.Fprintf(writer, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8080", nil)
}
