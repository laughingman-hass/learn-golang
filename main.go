package main

import (
	"fmt"
	"learn-golang/controllers"
	"learn-golang/middleware"
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

	r := mux.NewRouter()
	requireSessionMW := middleware.RequireSession{
		UserService: services.User,
	}

	staticController := controllers.NewStatic()
	usersController := controllers.NewUsers(services.User)
	sessionsController := controllers.NewSession(services.User)
	galleriesController := controllers.NewGalleries(services.Gallery, r)

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
	r.Handle("/galleries/new", requireSessionMW.Apply(galleriesController.New)).Methods("GET")
	r.HandleFunc("/galleries", requireSessionMW.ApplyFn(galleriesController.Create)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", requireSessionMW.ApplyFn(galleriesController.Show)).Methods("GET").Name("gallery")

	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
