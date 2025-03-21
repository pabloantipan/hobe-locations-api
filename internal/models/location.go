package models

import (
	"time"
)

type BucketPicture struct {
	Name       string    `json:"name" binding:"required" example:"1"`
	URL        string    `json:"url" binding:"required" example:"https://storage.googleapis.com/bucket-name/picture.jpg"`
	UploadedAt time.Time `json:"uploadedAt" binding:"required" example:"2021-01-01T00:00:00Z"`
}

type Location struct {
	ID             string          `json:"id" binding:"required" example:"1"`
	UserID         string          `form:"userId" binding:"required" example:"1"`
	UserEmail      string          `form:"userEmail" binding:"required" example:"john@doe.com"`
	UserFirebaseID string          `form:"userFirebaseId" binding:"required" example:"123456"`
	Name           string          `json:"name" binding:"required" example:"John Doe"`
	Address        string          `json:"address" binding:"required" example:"Av. Corrientes 1234"`
	Comment        string          `json:"comment" binding:"required" example:"This is a description"`
	Latitude       float64         `json:"latitude" binding:"required" example:"-34.603722"`
	Longitude      float64         `json:"longitude" binding:"required" example:"-58.381592"`
	Accuracy       float64         `json:"accuracy" binding:"required" example:"0.0001"`
	PointType      string          `json:"pointType" binding:"required" example:"ruco"`
	MenCount       int             `json:"menCount" binding:"required" example:"2"`
	WomenCount     int             `json:"womenCount" binding:"required" example:"2"`
	HasMigrants    bool            `json:"hasMigrants" binding:"required" example:"true"`
	CanSurvey      bool            `json:"canSurvey" binding:"required" example:"true"`
	Pictures       []BucketPicture `json:"pictures" binding:"optional"`
	CreatedOn      time.Time       `json:"created_on" binding:"required" example:"2021-01-01T00:00:00Z"`
}
