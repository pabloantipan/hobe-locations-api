package models

type PointType struct {
	ID    string `json:"id" binding:"required" example:"1" format:"string"`
	Value string `json:"value" binding:"required" example:"1" format:"string"`
	Label string `json:"label" binding:"required" example:"1" format:"string"`
}

type PointTypesResponse struct {
	Types []PointType `json:"types" binding:"required"`
}

type LocationObjectKey struct {
	ID        string `json:"id" binding:"required" example:"1" format:"string"`
	Value     string `json:"value" binding:"required" example:"1" format:"string"`
	Label     string `json:"label" binding:"required" example:"1" format:"string"`
	Direction string `json:"direction" binding:"required" example:"1" format:"string"`
}

type LocationOrderKeysResponse struct {
	Keys []LocationObjectKey `json:"keys" binding:"required"`
}
