package models

type Player struct {
	ID       string  `json:"id" example:"1"`
	Name     string  `json:"name" binding:"required" example:"John Doe"`
	Age      int     `json:"age" binding:"required" example:"25"`
	Position string  `json:"position" binding:"required" example:"Forward"`
	Rating   float64 `json:"rating" binding:"required" example:"7.5"`
}
