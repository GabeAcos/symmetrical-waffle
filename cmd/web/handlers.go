package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"snippetboxsolo/internal/models"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	users, err := app.users.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, user := range users {
		fmt.Fprintf(w, "%+v\n", user)
	}
	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	http.Error(w, "Internal Server error", http.StatusInternalServerError)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
}

func (app *application) userView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	user, err := app.users.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	fmt.Fprintf(w, "%+v", user)

}

func (app *application) userCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Create"))
}

func (app *application) userCreatePost(w http.ResponseWriter, r *http.Request) {
	name := "sophia"
	age := 2
	height := 24
	weight := 30

	id, err := app.users.Insert(name, age, height, weight)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/view/%d", id), http.StatusSeeOther)
}

func (app *application) exerciseView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}

	excercise, err := app.exercises.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord){
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	
	fmt.Fprintf(w, "%+v", excercise)
}

func (app *application) exerciseCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excersise Create"))
}

func (app *application) exerciseCreatePost(w http.ResponseWriter, r *http.Request) {

}