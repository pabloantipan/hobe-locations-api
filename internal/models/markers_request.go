package models

type MapSquareBounds struct {
	NorthLat float64 `json:"northLat" binding:"required" example:"-33.321" format:"float"`
	SouthLat float64 `json:"southLat" binding:"required" example:"-33.321" format:"float"`
	EastLng  float64 `json:"eastLng" binding:"required" example:"-33.321" format:"float"`
	WestLng  float64 `json:"westLng" binding:"required" example:"-33.321" format:"float"`
}

type LocationMarkersRequest struct {
	Square            MapSquareBounds `json:"square" binding:"required"`
	GottenLocationIDs []string        `json:"gottenLocationIDs" binding:"required"`
}
