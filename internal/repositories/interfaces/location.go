package interfaces

import "github.com/pabloantipan/hobe-locations-api/internal/models"

type LocationRepository interface {
	Add(location models.Location) (models.Location, error)
	GetByID(id string) (models.Location, error)
	GetAll() ([]models.Location, error)
	Update(location models.Location) (models.Location, error)
	Delete(id string) error
}
