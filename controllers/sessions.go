package controllers

import (
	"learn-golang/models"
	"learn-golang/views"
	"log"
	"net/http"

	"learn-golang/rand"
)

func NewSession(us models.UserService) *SessionsController {
	return &SessionsController{
		NewView: views.NewView("bootstrap", "sessions/new"),
		us:      us,
	}
}

type SessionsController struct {
	NewView *views.View
	us      models.UserService
}

func (sc *SessionsController) New(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	sc.NewView.Render(w, nil)
}

func (sc *SessionsController) Create(w http.ResponseWriter, r *http.Request) {
	vd := views.Data{}
	form := SessionParams{}

	if err := parseForm(&form, r); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		sc.NewView.Render(w, vd)
		return
	}

	user, err := sc.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			vd.AlertError("Invalid email address")
		default:
			vd.SetAlert(err)
		}
		sc.NewView.Render(w, vd)
		return
	}

	if err = signIn(w, user, sc.us); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		return
	}

	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

type SessionParams struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func signIn(w http.ResponseWriter, user *models.User, us models.UserService) error {
	if user.SessionToken == "" {
		token, err := rand.NewSessionToken()
		if err != nil {
			return err
		}
		user.SessionToken = token
		err = us.Update(user)
		if err != nil {
			return err
		}
	}

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    user.SessionToken,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	return nil
}
