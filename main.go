package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from invoicer!"))
}

func invoiceView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific invoice with ID %d...", id)
	w.Write([]byte(msg))
}

func invoiceCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new invoice..."))
}

func invoiceCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new invoice..."))
}
