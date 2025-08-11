package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	addr := flag.String("addr", ":4000", "HTTP routing address")
	flag.Parse()

	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on port: %v", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
