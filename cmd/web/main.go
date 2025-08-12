package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	addr := flag.String("addr", ":4000", "HTTP routing address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := application{
		logger: logger,
	}
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	log.Printf("Starting server on port: %v", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
