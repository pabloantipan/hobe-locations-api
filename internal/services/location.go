package services

import (
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/interfaces"
)

type LocationService struct {
	repo interfaces.LocationRepository
}

func NewLocationService(repo interfaces.LocationRepository) LocationServiceInterface {
	return &LocationService{repo: repo}
}

type LocationServiceInterface interface {
	Add(location models.Location) (models.Location, error)
	GetByID(id string) (models.Location, error)
	GetAll() ([]models.Location, error)
	Update(location models.Location) (models.Location, error)
	Delete(id string) error
}

func (s *LocationService) Add(location models.Location) (models.Location, error) {
	return s.repo.Add(location)
}

func (s *LocationService) GetByID(id string) (models.Location, error) {
	return s.repo.GetByID(id)
}

func (s *LocationService) GetAll() ([]models.Location, error) {
	return s.repo.GetAll()
}

func (s *LocationService) Update(location models.Location) (models.Location, error) {
	return s.repo.Update(location)
}

func (s *LocationService) Delete(id string) error {
	return s.repo.Delete(id)
}
