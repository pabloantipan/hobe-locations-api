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

type PictureService struct {
	repo storage.PictureRepositoryInterface
}

func NewPictureService(repo storage.PictureRepositoryInterface) PictureServiceInterface {
	return &PictureService{repo: repo}
}

type PictureServiceInterface interface {
	GetURL(pictureName string) string
	Upload(file *multipart.FileHeader) (*models.FileInfo, error)
	validate(file *multipart.FileHeader) (bool, exceptions.PictureException)
}

func (s *PictureService) GetURL(pictureName string) string {
	return s.repo.GetURL(pictureName)
}

func (s *PictureService) Upload(file *multipart.FileHeader) (*models.FileInfo, error) {
	ctx := context.Background()
	_, err := s.validate(file)
	if err.IsPictureError() {
		return nil, fmt.Errorf("failed to validate file: %v", err)
	}

	result, error := s.repo.Upload(ctx, file)
	if error != nil {
		return nil, fmt.Errorf("failed to upload file: %v", error)
	}

	return result, nil
}

func (s *PictureService) validate(file *multipart.FileHeader) (bool, exceptions.PictureException) {
	validator := utils.ImageValidator{
		MaxFileSize:   2 << 20, // 2MB
		MaxDimensions: 1024,
		AllowedTypes:  []string{"image/jpeg", "image/png"},
	}

	return validator.ValidateBasicProperties(file)
}
