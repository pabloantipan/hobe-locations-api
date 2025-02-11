package models

type Picture struct {
	File      []byte `json:"file" binding:"required"`
	Timestamp string `json:"timestamp" example:"2021-01-01T00:00:00Z"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude" binding:"required" example:"-34.603722"`
	Longitude float64 `json:"longitude" binding:"required" example:"-58.381592"`
}

type LocationRequest struct {
	ID          string `json:"id" example:"1"`
	Name        string `json:"name" binding:"required" example:"John Doe"`
	Comment     string `json:"comment" binding:"required" example:"This is a description"`
	Coordinates `json:"coordinates" binding:"required"`
	Pictures    []Picture `json:"pictures" binding:"optional"`
	Address     string    `json:"address" binding:"required" example:"Av. Corrientes 1234"`
}

type BucketPicture struct {
	ID        string `json:"id" binding:"required" example:"1"`
	URL       string `json:"url" binding:"required" example:"https://storage.googleapis.com/bucket-name/picture.jpg"`
	Timestamp string `json:"timestamp" binding:"required" example:"2021-01-01T00:00:00Z"`
}

type Location struct {
	ID          string          `json:"id" binding:"required" example:"1"`
	Name        string          `json:"name" binding:"required" example:"John Doe"`
	Comment     string          `json:"comment" binding:"required" example:"This is a description"`
	Latitude    float64         `json:"latitude" binding:"required" example:"-34.603722"`
	Longitude   float64         `json:"longitude" binding:"required" example:"-58.381592"`
	PictureURLs []BucketPicture `json:"picture_urls" binding:"optional"`
	Address     string          `json:"address" binding:"required" example:"Av. Corrientes 1234"`
	CreatedOn   string          `json:"created_on" binding:"required" example:"2021-01-01T00:00:00Z"`
}
