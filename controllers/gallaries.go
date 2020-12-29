package controllers

import (
	"fmt"
	"learn-golang/context"
	"learn-golang/models"
	"learn-golang/views"
	"log"
	"net/http"
)

func NewGalleries(gs models.GalleryService) *GalleriesController {
	return &GalleriesController{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

type GalleriesController struct {
	New *views.View
	gs  models.GalleryService
}

type GalleryForm struct {
	Title string `schema:"title"`
}

func (gc *GalleriesController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		form GalleryForm
		vd   views.Data
	)

	if err := parseForm(&form, r); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		gc.New.Render(w, vd)
		return
	}

	user := context.User(r.Context())

	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}

	if err := gc.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		gc.New.Render(w, vd)
		return
	}

	fmt.Fprintln(w, gallery)
}
