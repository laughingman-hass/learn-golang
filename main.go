package main

import (
	"fmt"
	"learn-golang/controllers"
	"learn-golang/models"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host     = "db"
	port     = "5432"
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "lenslocked_dev"
)

func main() {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserServices(connectionInfo)
	must(err)
	defer us.Close()
	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(us)
	sessionsController := controllers.NewSession(us)

	r := mux.NewRouter()
	r.NotFoundHandler = http.Handler(staticController.NotFound)

	r.Handle("/", staticController.Home).Methods("GET")
	r.Handle("/contact", staticController.Contact).Methods("GET")
	r.Handle("/faq", staticController.FAQ).Methods("GET")

	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	r.HandleFunc("/cookietest", usersController.CookieTest).Methods("GET")
	r.HandleFunc("/login", sessionsController.New).Methods("GET")
	r.HandleFunc("/login", sessionsController.Create).Methods("POST")

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
