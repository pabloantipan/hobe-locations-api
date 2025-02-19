package models

import (
	"fmt"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pabloantipan/hobe-locations-api/utils"
)

type LocationRequest struct {
	ID        string                  `form:"id" example:"1"`
	Name      string                  `form:"name" binding:"required" example:"John Doe"`
	Comment   string                  `form:"comment" binding:"required" example:"This is a description"`
	Latitude  float64                 `form:"latitude" binding:"required" example:"-34.603722"`
	Longitude float64                 `form:"longitude" binding:"required" example:"-58.381592"`
	Pictures  []*multipart.FileHeader `form:"pictures" binding:"required"`
	Address   string                  `form:"address" binding:"required" example:"Av. Corrientes 1234"`
}

func validateFormData(form *multipart.Form) (*LocationRequest, error) {
	var req LocationRequest

	var errorMessages = make([]string, 0)

	if namesSlice, ok := form.Value["name"]; ok && len(namesSlice) > 0 {
		req.Name = namesSlice[0]
	} else {
		errorMessages = append(errorMessages, "names are required")
	}

	if commentSlice, ok := form.Value["comment"]; ok && len(commentSlice) > 0 {
		req.Comment = commentSlice[0]
	} else {
		errorMessages = append(errorMessages, "last names are required")
	}

	if latitudeSlice, ok := form.Value["latitude"]; ok && len(latitudeSlice) > 0 {
		req.Latitude = utils.ParseEnvFloat64(latitudeSlice[0])
	} else {
		errorMessages = append(errorMessages, "nickname is required")
	}

	if longitudeSlice, ok := form.Value["longitude"]; ok && len(longitudeSlice) > 0 {
		req.Longitude = utils.ParseEnvFloat64(longitudeSlice[0])
	} else {
		errorMessages = append(errorMessages, "nickname is required")
	}

	if picturesHeaders, ok := form.File["pictures"]; ok {
		req.Pictures = picturesHeaders
	} else {
		req.Pictures = []*multipart.FileHeader{}
	}

	if addressSlice, ok := form.Value["address"]; ok && len(addressSlice) > 0 {
		req.Address = addressSlice[0]
	} else {
		errorMessages = append(errorMessages, "nickname is required")
	}

	if len(errorMessages) > 0 {
		return nil, fmt.Errorf("%s", strings.Join(errorMessages, "\n"))
	}

	return &req, nil
}

func validateFields(req *LocationRequest) (*LocationRequest, error) {
	validate := validator.New()

	validate.RegisterValidation("name", func(fl validator.FieldLevel) bool {
		name := fl.Field().String()
		if name == "" {
			return true
		}

		nameRegex := regexp.MustCompile(`^[a-zA-Z\s\-'\. ]+$`)
		return nameRegex.MatchString(name)
	})

	validate.RegisterValidation("nickname", func(fl validator.FieldLevel) bool {
		nickname := fl.Field().String()
		if nickname == "" {
			return true
		}
		nicknameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
		return nicknameRegex.MatchString(nickname)
	})

	if err := validate.Struct(req); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			errorMessages := make([]string, 0, len(validationErrors))
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Email":
					errorMessages = append(errorMessages, "Invalid email address.")
				case "Names":
					errorMessages = append(errorMessages, "Names are required and can only contain letters, spaces, hyphens, apostrophes, and periods.")
				case "LastNames":
					errorMessages = append(errorMessages, "Last names are required and can only contain letters, spaces, hyphens, apostrophes, and periods.")
				case "Nickname":
					errorMessages = append(errorMessages, "Nickname is required and can only contain letters, numbers, and underscores.")
				case "Photo":
					errorMessages = append(errorMessages, "Photo is optional.")
				default:
					errorMessages = append(errorMessages, fieldError.Error())
				}
			}
			return nil, fmt.Errorf("%v", errorMessages)
		}
		return nil, err
	}

	return req, nil
}

func ValidateLocationRequest(form *multipart.Form) (*LocationRequest, error) {

	req, err := validateFormData(form)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req, err = validateFields(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return req, nil
}
