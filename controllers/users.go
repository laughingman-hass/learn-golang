package controllers

import (
	"fmt"
	"learn-golang/models"
	"learn-golang/views"
	"net/http"
)

func NewUsers(us *models.UserService) *UsersController {
	return &UsersController{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type UsersController struct {
	NewView *views.View
	us      *models.UserService
}

type Signupform struct {
	Name     string `schema:"name"`
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

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, form)
}
