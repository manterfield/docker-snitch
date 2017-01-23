package main

import (
    "net/http"
    "fmt"
    "github.com/manterfield/docker-snitch"
)

func index() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world!")
    })
}

func images() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, dockersnitch.Images())
    })
}

func main() {
    http.Handle("/", index())
    http.Handle("/images", images())
    // http.Handle("/ping", index()) -- current date and "pong"
    // http.Handle("/images", apiHandler())
    // http.Handle("/containers", apiHandler())
    http.ListenAndServe(":8080", nil)
}