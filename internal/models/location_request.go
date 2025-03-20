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
	UserID         string                  `form:"userId" binding:"required" example:"1"`
	UserEmail      string                  `form:"userEmail" binding:"required" example:"john@doe.com"`
	UserFirebaseID string                  `form:"userFirebaseId" binding:"required" example:"123456"`
	Name           string                  `form:"name" binding:"required" example:"John Doe"`
	Address        string                  `form:"address" binding:"required" example:"Av. Corrientes 1234"`
	Comment        string                  `form:"comment" binding:"required" example:"This is a description"`
	Latitude       float64                 `form:"latitude" binding:"required" example:"-34.603722"`
	Longitude      float64                 `form:"longitude" binding:"required" example:"-58.381592"`
	Accuracy       float64                 `form:"accuracy" binding:"required" example:"0.0001"`
	PointType      string                  `form:"pointType" binding:"required" example:"ruco"`
	MenCount       int                     `form:"menCount" binding:"required" example:"2"`
	WomenCount     int                     `form:"womenCount" binding:"required" example:"2"`
	HasMigrants    bool                    `form:"hasMigrants" binding:"required" example:"true"`
	CanSurvey      bool                    `form:"canSurvey" binding:"required" example:"true"`
	Pictures       []*multipart.FileHeader `form:"pictures" binding:"required"`
}

func validateFormData(form *multipart.Form) (*LocationRequest, error) {
	var req LocationRequest

	var errorMessages = make([]string, 0)

	if userIdSlice, ok := form.Value["userId"]; ok && len(userIdSlice) > 0 {
		req.UserID = userIdSlice[0]
	} else {
		errorMessages = append(errorMessages, "userId is required")
	}

	if userEmailSlice, ok := form.Value["userEmail"]; ok && len(userEmailSlice) > 0 {
		req.UserEmail = userEmailSlice[0]
	} else {
		errorMessages = append(errorMessages, "userEmail is required")
	}

	if userFirebaseIdSlice, ok := form.Value["userFirebaseId"]; ok && len(userFirebaseIdSlice) > 0 {
		req.UserFirebaseID = userFirebaseIdSlice[0]
	} else {
		errorMessages = append(errorMessages, "userFirebaseId is required")
	}

	if namesSlice, ok := form.Value["name"]; ok && len(namesSlice) > 0 {
		req.Name = namesSlice[0]
	} else {
		errorMessages = append(errorMessages, "name are required")
	}

	if addressSlice, ok := form.Value["address"]; ok && len(addressSlice) > 0 {
		req.Address = addressSlice[0]
	} else {
		errorMessages = append(errorMessages, "address is required")
	}

	if commentSlice, ok := form.Value["comment"]; ok && len(commentSlice) > 0 {
		req.Comment = commentSlice[0]
	} else {
		errorMessages = append(errorMessages, "last names are required")
	}

	if latitudeSlice, ok := form.Value["latitude"]; ok && len(latitudeSlice) > 0 {
		req.Latitude = utils.ParseEnvFloat64(latitudeSlice[0])
	} else {
		errorMessages = append(errorMessages, "latitud is required")
	}

	if longitudeSlice, ok := form.Value["longitude"]; ok && len(longitudeSlice) > 0 {
		req.Longitude = utils.ParseEnvFloat64(longitudeSlice[0])
	} else {
		errorMessages = append(errorMessages, "longitude is required")
	}

	if accuracySlice, ok := form.Value["accuracy"]; ok && len(accuracySlice) > 0 {
		req.Accuracy = utils.ParseEnvFloat64(accuracySlice[0])
	} else {
		errorMessages = append(errorMessages, "accuracy is required")
	}

	if pointTypeSlice, ok := form.Value["pointType"]; ok && len(pointTypeSlice) > 0 {
		req.PointType = pointTypeSlice[0]
	} else {
		errorMessages = append(errorMessages, "pointType is required")
	}

	if menCountSlice, ok := form.Value["menCount"]; ok && len(menCountSlice) > 0 {
		menCount, err := utils.ParseStringToInt(menCountSlice[0])
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("menCount: %v", err))
		}
		req.MenCount = menCount
	} else {
		errorMessages = append(errorMessages, "menCount is required")
	}

	if womenCountSlice, ok := form.Value["womenCount"]; ok && len(womenCountSlice) > 0 {
		womenCount, err := utils.ParseStringToInt(womenCountSlice[0])
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("womenCount: %v", err))
		}
		req.WomenCount = womenCount
	} else {
		errorMessages = append(errorMessages, "womenCount is required")
	}

	if hasMigrantsSlice, ok := form.Value["hasMigrants"]; ok && len(hasMigrantsSlice) > 0 {
		hasMigrants, err := utils.ParseStringToBool(hasMigrantsSlice[0])
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("hasMigrants: %v", err))
		}
		req.HasMigrants = hasMigrants
	} else {
		errorMessages = append(errorMessages, "hasMigrants is required")
	}

	if canSurveySlice, ok := form.Value["canSurvey"]; ok && len(canSurveySlice) > 0 {
		canSurvey, err := utils.ParseStringToBool(canSurveySlice[0])
		if err != nil {
			errorMessages = append(errorMessages, fmt.Sprintf("canSurvey: %v", err))
		}
		req.CanSurvey = canSurvey
	} else {
		errorMessages = append(errorMessages, "canSurvey is required")
	}

	if picturesHeaders, ok := form.File["pictures[]"]; ok {
		req.Pictures = picturesHeaders
	} else {
		req.Pictures = []*multipart.FileHeader{}
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
