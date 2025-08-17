package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /user/view/{id}", app.userView)
	mux.HandleFunc("GET /user/create", app.userCreate)
	mux.HandleFunc("POST /user/create", app.userCreatePost)

	mux.HandleFunc("GET /excercise/view/{id}", app.exerciseView)
	mux.HandleFunc("GET /excercise/create", app.exerciseCreate)
	mux.HandleFunc("POST /excercise/create", app.exerciseCreatePost)

	return mux
}
