package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/services"
)

type PictureHandler struct {
	service services.PictureServiceInterface
}

func NewPictureHandler(s services.PictureServiceInterface) PictureHandlerInterface {
	return &PictureHandler{service: s}
}

type PictureHandlerInterface interface {
	Upload(c *gin.Context)
}

func (h *PictureHandler) Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files[]"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no files provided"})
		return
	}

	var results = make([]models.FileInfo, 0)
	for _, fileHeader := range files {
		result, _ := h.service.Upload(fileHeader)
		results = append(results, *result)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	return
		// }

	}

	c.JSON(http.StatusCreated, results)
}
