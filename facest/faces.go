package facest

import "time"

type FacesRes struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    FacesList `json:"data"`
}
type FacesList struct {
	Count      int    `json:"count"`
	PagesCount int    `json:"pages_count"`
	Faces      []Face `json:"faces"`
}

type Face struct {
	FaceToken  string      `json:"face_token"`
	FaceID     string      `json:"face_id"`
	FaceImages []FaceImage `json:"face_images"`
	CreatedAt  time.Time   `json:"created_at"`
}

type FaceImage struct {
	ImageToken string    `json:"image_token"`
	ImageURL   string    `json:"image_url"`
	CreatedAt  time.Time `json:"created_at"`
}
