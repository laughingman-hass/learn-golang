package controllers

import (
	"fmt"
	"learn-golang/views"
	"net/http"
)

func NewUsers() *UsersController {
	return &UsersController{
		NewView: views.NewView("bootstrap", "users/new"),
	}
}

type UsersController struct {
	NewView *views.View
}

type Signupform struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (u UsersController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var form Signupform

	if err := parseForm(&form, r); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
}
