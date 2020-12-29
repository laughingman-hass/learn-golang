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
	GalleryPath       = "gallery"
	EditGalleryPath   = "editGallery"
	UpdateGalleryPath = "updateGallery"
	DeleteGalleryPath = "deleteGallery"
)

func NewGalleries(gs models.GalleryService, r *mux.Router) *GalleriesController {
	return &GalleriesController{
		New:       views.NewView("bootstrap", "galleries/new"),
		EditView:  views.NewView("bootstrap", "galleries/edit"),
		IndexView: views.NewView("bootstrap", "galleries/index"),
		ShowView:  views.NewView("bootstrap", "galleries/show"),
		gs:        gs,
		r:         r,
	}
}

type GalleriesController struct {
	New       *views.View
	IndexView *views.View
	ShowView  *views.View
	EditView  *views.View
	gs        models.GalleryService
	r         *mux.Router
}

type GalleryForm struct {
	Title string `schema:"title"`
}

func (gc *GalleriesController) Index(w http.ResponseWriter, r *http.Request) {
	var (
		vd views.Data
	)

	user := context.User(r.Context())
	galleries, err := gc.gs.ByUserID(user.ID)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	vd.Yield = galleries
	gc.IndexView.Render(w, r, vd)
}

func (gc *GalleriesController) Create(w http.ResponseWriter, r *http.Request) {
	var (
		form GalleryForm
		vd   views.Data
	)

	if err := parseForm(&form, r); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		gc.New.Render(w, r, vd)
		return
	}

	user := context.User(r.Context())

	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}

	if err := gc.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		gc.New.Render(w, r, vd)
		return
	}
	url, err := gc.r.Get(EditGalleryPath).URL("id", fmt.Sprintf("%v", gallery.ID))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	http.Redirect(w, r, url.Path, http.StatusFound)
}

func (gc *GalleriesController) Edit(w http.ResponseWriter, r *http.Request) {
	var vd views.Data

	gallery, err := gc.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "Gallery not found", http.StatusNotFound)
		return
	}

	vd.Yield = gallery
	gc.EditView.Render(w, r, vd)
}

func (gc *GalleriesController) Show(w http.ResponseWriter, r *http.Request) {
	var vd views.Data

	gallery, err := gc.galleryByID(w, r)
	if err != nil {
		return
	}

	vd.Yield = gallery
	gc.ShowView.Render(w, r, vd)
}

func (gc *GalleriesController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		vd   views.Data
		form GalleryForm
	)

	gallery, err := gc.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "Gallery not found", http.StatusNotFound)
		return
	}

	vd.Yield = gallery

	if err := parseForm(&form, r); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		gc.EditView.Render(w, r, vd)
		return
	}

	gallery.Title = form.Title
	err = gc.gs.Update(gallery)
	if err != nil {
		vd.SetAlert(err)
		gc.EditView.Render(w, r, vd)
		return
	}
	vd.Alert = &views.Alert{
		Level:   views.AlertLevelSuccess,
		Message: "Gallery successfully updated!",
	}
	gc.ShowView.Render(w, r, vd)
}

func (gc *GalleriesController) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		vd views.Data
	)
	gallery, err := gc.galleryByID(w, r)
	if err != nil {
		return
	}

	user := context.User(r.Context())
	if gallery.UserID != user.ID {
		http.Error(w, "Gallery not found", http.StatusNotFound)
		return
	}

	err = gc.gs.Delete(gallery.ID)
	if err != nil {
		vd.SetAlert(err)
		vd.Yield = gallery
		gc.EditView.Render(w, r, vd)
		return
	}

	http.Redirect(w, r, "/galleries", http.StatusFound)
}

func (gc *GalleriesController) galleryByID(w http.ResponseWriter, r *http.Request) (*models.Gallery, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return nil, err
	}

	gallery, err := gc.gs.ByID(uint(id))
	if err != nil {
		switch err {
		case models.ErrNotFound:
			http.Error(w, "Gallery Not Found", http.StatusNotFound)
			return nil, err
		default:
			http.Error(w, "Opps! Something went wrong", http.StatusInternalServerError)
			return nil, err
		}
	}
	return gallery, nil
}
