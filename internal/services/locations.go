package services

import (
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/datastore"
)

type LocationsService struct {
	repo datastore.LocationRepository
}

func NewLocationService(repo datastore.LocationRepository) LocationsServiceInterface {
	return &LocationsService{repo: repo}
}

type LocationsServiceInterface interface {
	Add(location models.Location) (models.Location, error)
	GetByID(id string) (models.Location, error)
	GetAll() ([]models.Location, error)
	GetThemByEmail(email string) (*[]models.Location, error)
	Update(location models.Location) (models.Location, error)
	Delete(id string) error
}

func (s *LocationsService) Add(location models.Location) (models.Location, error) {
	return s.repo.Add(location)
}

func (s *LocationsService) GetByID(id string) (models.Location, error) {
	return s.repo.GetByID(id)
}

func (s *LocationsService) GetThemByEmail(email string) (*[]models.Location, error) {
	return s.repo.GetThemByEmail(email)
}

func (s *LocationsService) GetAll() ([]models.Location, error) {
	return s.repo.GetAll()
}

func (s *LocationsService) Update(location models.Location) (models.Location, error) {
	return s.repo.Update(location)
}

func (s *LocationsService) Delete(id string) error {
	return s.repo.Delete(id)
}
