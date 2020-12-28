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
	services, err := models.NewServices(connectionInfo)
	must(err)
	defer services.Close()

	services.AutoMigrate()

	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(services.User)
	sessionsController := controllers.NewSession(services.User)
	galleriesController := controllers.NewGalleries(services.Gallery)

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

	// Gallery routes
	r.Handle("/galleries/new", galleriesController.New).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
