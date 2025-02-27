package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pabloantipan/hobe-locations-api/internal/bussines"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/utils"
)

type LocationsHandler struct {
	business bussines.LocationBusinessInterface
}

func NewLocationsHandler(b bussines.LocationBusinessInterface) LocationsHandlerInterface {
	return &LocationsHandler{business: b}
}

type LocationsHandlerInterface interface {
	Add(c *gin.Context)
	GetThemByEmail(c *gin.Context)
}

func (h *LocationsHandler) Add(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get form data", "details": err.Error()})
		return
	}

	newLocation, err := models.ValidateLocationRequest(form)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "couldn't parse form data",
			"details": strings.Split(err.Error(), "\n"),
		})
		return
	}

	location, err := h.business.Add(*newLocation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, location)
}

func (h *LocationsHandler) GetThemByEmail(c *gin.Context) {
	claims, err := utils.ParseClaimsAsUserData(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	locations, err := h.business.GetThemByEmail(claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"locations": locations})

}
