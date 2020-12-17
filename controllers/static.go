package controllers

import "learn-golang/views"

func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact:  views.NewView("bootstrap", "views/static/contact.gohtml"),
		FAQ:      views.NewView("bootstrap", "views/static/faq.gohtml"),
		NotFound: views.NewView("bootstrap", "views/static/not_found.gohtml"),
	}
}

type Static struct {
	Home     *views.View
	Contact  *views.View
	FAQ      *views.View
	NotFound *views.View
}
