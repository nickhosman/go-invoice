package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
    mux.HandleFunc("GET /invoice/view/{id}", app.invoiceView)
    mux.HandleFunc("GET /invoice/create", app.invoiceCreate)
    mux.HandleFunc("POST /invoice/create", app.invoiceCreatePost)

	return mux
}
