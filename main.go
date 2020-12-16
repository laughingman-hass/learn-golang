package main

import (
	"html/template"
	"learn-golang/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeView         *views.View
	contactView      *views.View
	faqTempate       *template.Template
	notFoundTemplate *template.Template
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)

	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.Execute(w, nil)

	if err != nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqTempate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	if err := notFoundTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")

	var err error
	faqTempate, err = template.ParseFiles(
		"views/faq.gohtml",
		"views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	notFoundTemplate, err = template.ParseFiles(
		"views/not_found.gohtml",
		"views/layouts/footer.gohtml",
	)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	http.ListenAndServe(":3000", r)
}
