package bussines

import (
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/services"
)

type LocationService struct {
	pictureService services.PictureServiceInterface
	// locationService services.LocationServiceInterface
}

func NewLocationService(pictureService services.PictureServiceInterface) LocationServiceInterface {
	return &LocationService{pictureService: pictureService}
}

type LocationServiceInterface interface {
	Add(locationRequest models.LocationRequest) (*models.Location, error)
}

func (s *LocationService) Add(locationRequest models.LocationRequest) (*models.Location, error) {
	return nil, nil
}
