package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"
	"time"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/utils"
)

type PictureRepository struct {
	bucketName string
	client     *storage.Client
	bucket     *storage.BucketHandle
}

func NewPictureRepository(client *storage.Client) PictureRepositoryInterface {
	bucketName := "hobe-location-picrtures"

	return &PictureRepository{
		bucketName: bucketName,
		client:     client,
		bucket:     client.Bucket(bucketName),
	}
}

type PictureRepositoryInterface interface {
	GetURL(subfolder, pictureName string) string
	Upload(ctx context.Context, file *multipart.FileHeader, subfolder string) (*models.FileInfo, error)
}

func generateFileName(filename string) string {
	return fmt.Sprintf("%d-%d%s", uuid.New().ID(), time.Now().UnixNano(), filepath.Ext(filename))
}

func (r *PictureRepository) Upload(ctx context.Context, file *multipart.FileHeader, subfolder string) (*models.FileInfo, error) {
	newFileName := generateFileName(file.Filename)
	objectPath := file.Filename
	if subfolder != "" {
		objectPath = path.Join(subfolder, newFileName)
	}

	fileInfo := &models.FileInfo{
		Name:        newFileName,
		URL:         fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.bucketName, objectPath),
		ContentType: utils.FileHeadertContentType(file),
	}

	resultChan := make(chan *models.FileInfo, 1)
	errorChan := make(chan error, 1)

	go func() {
		obj := r.bucket.Object(objectPath)
		writer := obj.NewWriter(ctx)
		writer.ContentType = fileInfo.ContentType

		defer writer.Close()

		filex, err := file.Open()
		if err != nil {
			errorChan <- fmt.Errorf("failed to open file: %v", err)
			return
		}
		defer filex.Close()

		if _, err := io.Copy(writer, filex); err != nil {
			errorChan <- fmt.Errorf("failed to copy file to storage: %v", err)
			return
		}

		if err := writer.Close(); err != nil {
			errorChan <- fmt.Errorf("failed to close writer: %v", err)
			return
		}

		attrs, err := obj.Attrs(ctx)
		if err != nil {
			errorChan <- fmt.Errorf("failed to get file attributes: %v", err)
			return
		}

		fileInfo.Size = attrs.Size
		fileInfo.UploadedAt = attrs.Created
		fileInfo.Path = objectPath

		resultChan <- fileInfo
	}()

	select {
	case fileInfo := <-resultChan:
		return fileInfo, nil
	case err := <-errorChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("upload process timed out")
	}
}

func (r *PictureRepository) GetURL(subfolder, filename string) string {
	objectPath := filename
	if subfolder != "" {
		objectPath = path.Join(subfolder, filename)
	}

	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.bucketName, objectPath)
}
