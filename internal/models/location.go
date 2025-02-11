package models

import (
	"mime/multipart"
	"time"
)

type LocationRequest struct {
	ID        string                  `form:"id" example:"1"`
	Name      string                  `form:"name" binding:"required" example:"John Doe"`
	Comment   string                  `form:"comment" binding:"required" example:"This is a description"`
	Latitude  float64                 `form:"latitude" binding:"required" example:"-34.603722"`
	Longitude float64                 `form:"longitude" binding:"required" example:"-58.381592"`
	Pictures  []*multipart.FileHeader `form:"pictures" binding:"required"`
	Address   string                  `form:"address" binding:"required" example:"Av. Corrientes 1234"`
}

type BucketPicture struct {
	Name       string    `json:"name" binding:"required" example:"1"`
	URL        string    `json:"url" binding:"required" example:"https://storage.googleapis.com/bucket-name/picture.jpg"`
	UploadedAt time.Time `json:"uploadedAt" binding:"required" example:"2021-01-01T00:00:00Z"`
}

type Location struct {
	ID        string          `json:"id" binding:"required" example:"1"`
	Name      string          `json:"name" binding:"required" example:"John Doe"`
	Comment   string          `json:"comment" binding:"required" example:"This is a description"`
	Latitude  float64         `json:"latitude" binding:"required" example:"-34.603722"`
	Longitude float64         `json:"longitude" binding:"required" example:"-58.381592"`
	Pictures  []BucketPicture `json:"pictures" binding:"optional"`
	Address   string          `json:"address" binding:"required" example:"Av. Corrientes 1234"`
	CreatedOn time.Time       `json:"created_on" binding:"required" example:"2021-01-01T00:00:00Z"`
}
