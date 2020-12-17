package controllers

import (
	"fmt"
	"learn-golang/views"
	"net/http"
)

func NewUsers() *UsersController {
	return &UsersController{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type UsersController struct {
	NewView *views.View
}

func (u UsersController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["password"])
	fmt.Fprintln(w, "This is a temporary response.")
}
