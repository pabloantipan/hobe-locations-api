package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type FirebaseIdentities struct {
	Email []string `json:"email"`
}

type FirebaseData struct {
	Identities     FirebaseIdentities `json:"identities"`
	SignInProvider string             `json:"sign_in_provider"`
}

type UserData struct {
	AuthTime      float64      `json:"auth_time"`
	Email         string       `json:"email"`
	EmailVerified bool         `json:"email_verified"`
	Firebase      FirebaseData `json:"firebase"`
	UserID        string       `json:"user_id"`
}

func ParseClaims(c *gin.Context) (*map[string]interface{}, error) {
	claims, exist := c.Request.Header["Claims"]
	if !exist {
		return nil, fmt.Errorf("claims not found")
	}

	// fmt.Println(claims)

	if len(claims) != 1 {
		return nil, fmt.Errorf("multiple claims found")
	}

	var userData map[string]interface{}
	err := json.Unmarshal([]byte(claims[0]), &userData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling complex JSON: %v", err)
	}

	return &userData, nil
}

func ParseClaimsAsUserData(c *gin.Context) (*UserData, error) {
	userData, err := ParseClaims(c)
	if err != nil {
		return nil, err
	}

	userDataJson, err := json.Marshal(userData)
	if err != nil {
		return nil, fmt.Errorf("error marshaling claims to JSON: %v", err)
	}

	var user UserData
	err = json.Unmarshal(userDataJson, &user)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON to UserData: %v", err)
	}

	return &user, nil
}
