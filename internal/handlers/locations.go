package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pabloantipan/hobe-locations-api/internal/bussines"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
)

type LocationsHandler struct {
	business bussines.LocationBusinessInterface
}

func NewLocationsHandler(b bussines.LocationBusinessInterface) LocationsHandlerInterface {
	return &LocationsHandler{business: b}
}

type LocationsHandlerInterface interface {
	Add(c *gin.Context)
}

func (h *LocationsHandler) Add(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't get form data"})
		return
	}

	latitude, err := strconv.ParseFloat(form.Value["Latitude"][0], 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude must be a number"})
		return
	}

	longitude, err := strconv.ParseFloat(form.Value["Longitude"][0], 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Longitude must be a number"})
		return
	}

	var request = models.LocationRequest{
		Address:   form.Value["Address"][0],
		Comment:   form.Value["Comment"][0],
		Latitude:  latitude,
		Longitude: longitude,
		Name:      form.Value["Name"][0],
		Pictures:  form.File["pictures[]"],
	}

	location, err := h.business.Add(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, location)
}
