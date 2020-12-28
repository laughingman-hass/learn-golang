package controllers

import (
	"learn-golang/models"
	"learn-golang/views"
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
