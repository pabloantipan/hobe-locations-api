package bussines

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/services"
	"github.com/pabloantipan/hobe-locations-api/utils"
)

type locationsBusiness struct {
	pictureService  services.PicturesService
	locationService services.LocationsService
}

func NewLocationsBusiness(
	pictureService services.PicturesService,
	locationService services.LocationsService,
) LocationsBusiness {
	return &locationsBusiness{
		pictureService:  pictureService,
		locationService: locationService,
	}
}

type LocationsBusiness interface {
	Add(locationRequest models.LocationRequest) (*models.Location, error)
	GetThemByEmail(email string) (*[]models.Location, error)
	GetThemByMapSquare(authorEmail string, request models.LocationMarkersRequest) (*[]models.Location, error)
}

func (s *locationsBusiness) GetThemByEmail(email string) (*[]models.Location, error) {
	return s.locationService.GetThemByEmail(email)
}

func (s *locationsBusiness) Add(request models.LocationRequest) (*models.Location, error) {
	locationID := uuid.New().String()

	pictures, errs := s.uploadPictures(locationID, request.Pictures)
	log.Printf("Errors uploading files: %v", errs)

	location := models.Location{
		ID:             locationID,
		UserID:         request.UserID,
		UserEmail:      request.UserEmail,
		UserFirebaseID: request.UserFirebaseID,
		Name:           request.Name,
		Comment:        request.Comment,
		Latitude:       request.Latitude,
		Longitude:      request.Longitude,
		Accuracy:       request.Accuracy,
		Address:        request.Address,
		Pictures:       pictures,
		CreatedOn:      time.Now(),
	}

	location, err := s.locationService.Add(location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (s *locationsBusiness) uploadPictures(locationID string, pictures []*multipart.FileHeader) ([]models.BucketPicture, []error) {
	var pictureURLs = make([]models.BucketPicture, 0)

	var errors = make([]error, 0)

	for _, picture := range pictures {
		uploadResult, err := s.pictureService.Upload(picture, locationID)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		pictureURLs = append(pictureURLs, models.BucketPicture{
			URL:        uploadResult.URL,
			Name:       uploadResult.Name,
			UploadedAt: uploadResult.UploadedAt,
		})
	}

	return pictureURLs, errors
}

func (s *locationsBusiness) GetThemByMapSquare(authorEmail string, request models.LocationMarkersRequest) (*[]models.Location, error) {
	allLocations, err := s.locationService.GetThemByMapSquare(request.Square)
	if err != nil {
		return nil, err
	}

	var locations = make([]models.Location, 0)
	for _, location := range *allLocations {
		if !utils.Contains(request.GottenLocationIDs, location.ID) {
			locations = append(locations, location)
		}
	}

	return &locations, nil
}
