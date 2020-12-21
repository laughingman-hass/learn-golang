package controllers

import (
	"fmt"
	"learn-golang/models"
	"learn-golang/views"
	"net/http"
)

func NewSession(us *models.UserService) *SessionsController {
	return &SessionsController{
		NewView: views.NewView("bootstrap", "sessions/new"),
		us:      us,
	}
}

type SessionsController struct {
	NewView *views.View
	us      *models.UserService
}

func (sc *SessionsController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := sc.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (sc *SessionsController) Create(w http.ResponseWriter, r *http.Request) {
	form := SessionParams{}
	if err := parseForm(&form, r); err != nil {
		panic(err)
	}

	user, err := sc.us.Authenticate(form.Email, form.Password)
	switch err {
	case models.ErrNotFound:
		fmt.Fprintln(w, "Invalid email address")
	case models.ErrInvalidPassword:
		fmt.Fprintln(w, "Invalid password provided")
	case nil:
		fmt.Fprintln(w, user)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type SessionParams struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
