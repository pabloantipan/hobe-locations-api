package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseEnvFloat64(key string) float64 {
	keyVal, err := strconv.ParseFloat(key, 64)
	if err == nil {
		return keyVal
	}

	value, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return 0.0
	}
	return value
}

func ParseStringToInt(s string) (int, error) {
	// Check for empty string
	if s == "" {
		return 0, errors.New("cannot parse an empty string")
	}

	// Use strconv.Atoi to convert string to int
	i, err := strconv.Atoi(s)
	if err != nil {
		// Provide more specific error messages based on the error type
		if numError, ok := err.(*strconv.NumError); ok {
			switch numError.Err {
			case strconv.ErrRange:
				return 0, fmt.Errorf("value '%s' is out of range for int type", s)
			case strconv.ErrSyntax:
				return 0, fmt.Errorf("value '%s' is not a valid number", s)
			}
		}
		// Return the original error if it's not one we specifically handle
		return 0, fmt.Errorf("failed to parse '%s': %w", s, err)
	}

	return i, nil
}

func ParseStringToBool(s string) (bool, error) {
	// Check for empty string
	if s == "" {
		return false, errors.New("cannot parse an empty string")
	}

	// Trim spaces and convert to lowercase for more flexible parsing
	s = strings.TrimSpace(strings.ToLower(s))

	// Use strconv.ParseBool to convert string to bool
	// This accepts: 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("'%s' is not a valid boolean value", s)
	}

	return b, nil
}
