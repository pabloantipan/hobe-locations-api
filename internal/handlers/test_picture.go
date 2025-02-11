package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Upload(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
