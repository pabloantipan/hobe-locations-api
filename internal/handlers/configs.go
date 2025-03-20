package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pabloantipan/hobe-locations-api/internal/constants"
)

type configsHandler struct{}

func NewConfigsHandler() ConfigsHandler {
	return &configsHandler{}
}

type ConfigsHandler interface {
	GetPointTypes(c *gin.Context)
	GetLocationOrderKeys(c *gin.Context)
}

func (h *configsHandler) GetPointTypes(c *gin.Context) {
	c.JSON(200, gin.H{"types": constants.LOCATION_POINT_TYPES})
}

func (h *configsHandler) GetLocationOrderKeys(c *gin.Context) {
	c.JSON(200, gin.H{"options": constants.LOCATION_ORDER_KEY_OPTIONS})
}
