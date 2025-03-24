package constants

import "github.com/pabloantipan/hobe-locations-api/internal/models"

var LOCATION_POINT_TYPES = []models.PointType{
	{
		ID:    "1",
		Value: "tent",
		Label: "Ocupación en carpa o ruco armado",
	},
	{
		ID:    "2",
		Value: "makeshift",
		Label: "Ocupación en refugio callejero",
	},
	{
		ID:    "4",
		Value: "unsheltered",
		Label: "Ocupación al aire libre",
	},
}

var LOCATION_ORDER_KEY_OPTIONS = []models.LocationObjectKey{
	{
		ID:        "1",
		Value:     "created_on",
		Label:     "Más reciente",
		Direction: "desc",
	},
	{
		ID:        "2",
		Value:     "created_on",
		Label:     "Más antiguo",
		Direction: "asc",
	},
	{
		ID:        "3",
		Value:     "name",
		Label:     "Alfabético",
		Direction: "asc",
	},
	// {
	// 	ID:        "1",
	// 	Value:     "name",
	// 	Label:     "Nombre Punto",
	// 	Direction: "asc",
	// },
	// {
	// 	ID:        "2",
	// 	Value:     "address",
	// 	Label:     "Dirección",
	// 	Direction: "asc",
	// },
	// {
	// 	ID:        "3",
	// 	Value:     "pointType",
	// 	Label:     "Tipo de punto",
	// 	Direction: "asc",
	// },
	// {
	// 	ID:        "4",
	// 	Value:     "menCount",
	// 	Label:     "Hombres",
	// 	Direction: "desc",
	// },
	// {
	// 	ID:        "5",
	// 	Value:     "womenCount",
	// 	Label:     "Mujeres",
	// 	Direction: "desc",
	// },
	// {
	// 	ID:        "6",
	// 	Value:     "hasMigrants",
	// 	Label:     "Migrantes",
	// 	Direction: "desc",
	// },
	// {
	// 	ID:        "7",
	// 	Value:     "canSurvey",
	// 	Label:     "Fue Encuestado",
	// 	Direction: "desc",
	// },
	// {
	// 	ID:        "8",
	// 	Value:     "created_on",
	// 	Label:     "Fecha de creación",
	// 	Direction: "desc",
	// },
}
