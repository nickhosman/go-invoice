package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /invoice/view/{id}", invoiceView)
    mux.HandleFunc("GET /invoice/create", invoiceCreate)
    mux.HandleFunc("POST /invoice/create", invoiceCreatePost)

    log.Print("starting server on :4000")

    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
