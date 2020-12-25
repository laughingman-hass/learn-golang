package controllers

import (
	"fmt"
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
	uc.NewView.Render(w, nil)
}

func (uc *UsersController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		form Signupform
		vd   views.Data
	)

	if err := parseForm(&form, r); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		uc.NewView.Render(w, vd)
		return
	}

	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	if err := uc.us.Create(&user); err != nil {
		vd.SetAlert(err)
		uc.NewView.Render(w, vd)
		return
	}

	if err := signIn(w, &user, uc.us); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

func (uc *UsersController) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := uc.us.BySession(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintln(w, user)
}
