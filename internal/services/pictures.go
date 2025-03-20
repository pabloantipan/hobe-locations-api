package services

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/pabloantipan/hobe-locations-api/internal/exceptions"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/storage"
	"github.com/pabloantipan/hobe-locations-api/utils"
)

type picturesService struct {
	repo storage.PictureRepository
}

func NewPictureService(repo storage.PictureRepository) PicturesService {
	return &picturesService{repo: repo}
}

type PicturesService interface {
	GetURL(locationID, pictureName string) string
	Upload(file *multipart.FileHeader, subfolder string) (*models.FileInfo, error)
	validate(file *multipart.FileHeader) (bool, exceptions.PictureException)
}

func (s *picturesService) GetURL(locationID, pictureName string) string {
	return s.repo.GetURL(locationID, pictureName)
}

func (s *picturesService) Upload(file *multipart.FileHeader, subfolder string) (*models.FileInfo, error) {
	ctx := context.Background()
	_, err := s.validate(file)
	if err.IsPictureError() {
		return nil, fmt.Errorf("failed to validate file: %v", err)
	}

	result, error := s.repo.Upload(ctx, file, subfolder)
	if error != nil {
		return nil, fmt.Errorf("failed to upload file: %v", error)
	}

	return result, nil
}

func (s *picturesService) validate(file *multipart.FileHeader) (bool, exceptions.PictureException) {
	validator := utils.ImageValidator{
		MaxFileSize:   2 << 20, // 2MB
		MaxDimensions: 1024,
		AllowedTypes:  []string{"image/jpeg", "image/png", "image/jpg", "image/gif"},
	}

	return validator.ValidateBasicProperties(file)
}
