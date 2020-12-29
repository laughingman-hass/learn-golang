package controllers

import (
	"fmt"
	"learn-golang/context"
	"learn-golang/models"
	"learn-golang/views"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	GalleryPath = "gallery"
)

func NewGalleries(gs models.GalleryService, r *mux.Router) *GalleriesController {
	return &GalleriesController{
		New:      views.NewView("bootstrap", "galleries/new"),
		ShowView: views.NewView("bootstrap", "galleries/show"),
		gs:       gs,
		r:        r,
	}
}

type GalleriesController struct {
	New      *views.View
	ShowView *views.View
	gs       models.GalleryService
	r        *mux.Router
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
	url, err := gc.r.Get(GalleryPath).URL("id", fmt.Sprintf("%v", gallery.ID))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.Redirect(w, r, url.Path, http.StatusFound)
}

func (gc *GalleriesController) Show(w http.ResponseWriter, r *http.Request) {
	var vd views.Data

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return
	}

	gallery, err := gc.gs.ByID(uint(id))
	if err != nil {
		switch err {
		case models.ErrNotFound:
			http.Error(w, "Gallery Not Found", http.StatusNotFound)
			return
		default:
			http.Error(w, "Opps! Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	vd.Yield = gallery
	gc.ShowView.Render(w, vd)
}
