package bussines

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/services"
)

type LocationBusiness struct {
	pictureService  services.PictureServiceInterface
	locationService services.LocationServiceInterface
}

func NewLocationBusiness(
	pictureService services.PictureServiceInterface,
	locationService services.LocationServiceInterface,
) LocationBusinessInterface {
	return &LocationBusiness{
		pictureService:  pictureService,
		locationService: locationService,
	}
}

type LocationBusinessInterface interface {
	Add(locationRequest models.LocationRequest) (*models.Location, error)
}

func (s *LocationBusiness) Add(request models.LocationRequest) (*models.Location, error) {
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

func (s *LocationBusiness) uploadPictures(locationID string, pictures []*multipart.FileHeader) ([]models.BucketPicture, []error) {
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
