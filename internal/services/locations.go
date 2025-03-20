package services

import (
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/datastore"
)

type locationsService struct {
	repo datastore.LocationRepository
}

func NewLocationService(repo datastore.LocationRepository) LocationsService {
	return &locationsService{repo: repo}
}

type LocationsService interface {
	Add(location models.Location) (models.Location, error)
	GetByID(id string) (models.Location, error)
	GetAll() ([]models.Location, error)
	GetThemByEmail(email string) (*[]models.Location, error)
	GetThemByMapSquare(square models.MapSquareBounds) (*[]models.Location, error)
	Update(location models.Location) (models.Location, error)
	Delete(id string) error
}

func (s *locationsService) Add(location models.Location) (models.Location, error) {
	return s.repo.Add(location)
}

func (s *locationsService) GetByID(id string) (models.Location, error) {
	return s.repo.GetByID(id)
}

func (s *locationsService) GetThemByEmail(email string) (*[]models.Location, error) {
	return s.repo.GetThemByEmail(email)
}

func (s *locationsService) GetThemByMapSquare(square models.MapSquareBounds) (*[]models.Location, error) {
	return s.repo.GetThemByMapSquare(square)
}

func (s *locationsService) GetAll() ([]models.Location, error) {
	return s.repo.GetAll()
}

func (s *locationsService) Update(location models.Location) (models.Location, error) {
	return s.repo.Update(location)
}

func (s *locationsService) Delete(id string) error {
	return s.repo.Delete(id)
}
