package controllers

import (
	"learn-golang/models"
	"learn-golang/views"
	"log"
	"net/http"
)

func NewUsers(us models.UserService) *UsersController {
	return &UsersController{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type UsersController struct {
	NewView *views.View
	us      models.UserService
}

type Signupform struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func (uc UsersController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	uc.NewView.Render(w, r, nil)
}

func (uc *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		form Signupform
		vd   views.Data
	)

	if err := parseForm(&form, r); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		uc.NewView.Render(w, r, vd)
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := uc.us.Create(&user); err != nil {
		vd.SetAlert(err)
		uc.NewView.Render(w, r, vd)
		return
	}

	if err := signIn(w, &user, uc.us); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/galleries", http.StatusFound)
}
