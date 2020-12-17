package controllers

import "learn-golang/views"

func NewStatic() *Static {
	return &Static{
		Home:     views.NewView("bootstrap", "static/home"),
		Contact:  views.NewView("bootstrap", "static/contact"),
		FAQ:      views.NewView("bootstrap", "static/faq"),
		NotFound: views.NewView("bootstrap", "static/not_found"),
	}
}

type Static struct {
	Home     *views.View
	Contact  *views.View
	FAQ      *views.View
	NotFound *views.View
}
