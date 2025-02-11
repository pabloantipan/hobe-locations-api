package storage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/utils"
)

type PictureRepository struct {
	bucketName string
	client     *storage.Client
	bucket     *storage.BucketHandle
}

func NewPictureRepository(client *storage.Client) PictureRepositoryInterface {
	bucketName := "hobe-locations-api-pictures"

	return &PictureRepository{
		bucketName: bucketName,
		client:     client,
		bucket:     client.Bucket(bucketName),
	}
}

type PictureRepositoryInterface interface {
	GetURL(pictureName string) string
	Upload(ctx context.Context, file *multipart.FileHeader) (*models.FileInfo, error)
}

func (r *PictureRepository) Upload(ctx context.Context, file *multipart.FileHeader) (*models.FileInfo, error) {
	obj := r.bucket.Object(file.Filename)

	writer := obj.NewWriter(ctx)
	writer.ContentType = utils.FileHeadertContentType(file)

	// Copy the file data to the object
	filex, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}

	if _, err := io.Copy(writer, filex); err != nil {
		return nil, fmt.Errorf("failed to copy file to storage: %v", err)
	}
	defer filex.Close()

	// Close the writer
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close writer: %v", err)
	}

	// Make the file public
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, fmt.Errorf("failed to make file public: %v", err)
	}

	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get file attributes: %v", err)
	}

	return &models.FileInfo{
		Name:        file.Filename,
		Size:        attrs.Size,
		ContentType: attrs.ContentType,
		URL:         fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.bucketName, file.Filename),
		UploadedAt:  attrs.Created,
	}, nil

}

func (r *PictureRepository) GetURL(filename string) string {
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", r.bucketName, filename)
}
