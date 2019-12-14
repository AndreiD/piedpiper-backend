package handlers

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"piedpiper/database"
	"piedpiper/models"
	"piedpiper/utils/log"
	"regexp"
	"strings"
)

// Register creates a new user
func Register(c *gin.Context) {
	var registerModel models.RegisterUser

	err := c.BindJSON(&registerModel)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload " + err.Error()})
		return
	}

	err = newUserValidation(registerModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registerModel.LastIP = c.Request.Header.Get("Cf-Connecting-Ip")
	registerModel.CFCookie = c.Request.Header.Get("Cookie")

	err = database.CreateUser(registerModel)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "seems there's already an user registered with your" +
				"email and/or phone"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user created ok"})
}

// GetUser ..
func GetUser(c *gin.Context) {
	userID := c.Param("id")

	if len(userID) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := database.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// GetUserByToken ...
func GetUserByToken(c *gin.Context) {
	// get the user id from jwt
	jwtClaims := jwt.ExtractClaims(c)

	id, ok := jwtClaims["id"].(string)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid request"})
		return
	}

	user, err := database.GetAuthenticatedUser(id)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

// UpdateUserDetails ...
func UpdateUserDetails(c *gin.Context) {
	jwtClaims := jwt.ExtractClaims(c)

	id, ok := jwtClaims["id"].(string)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid request"})
		return
	}

	var updatePayload models.UserUpdate

	err := c.BindJSON(&updatePayload)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload " + err.Error()})
		return
	}

	err = validateUpdateUser(updatePayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePayload.LastIP = c.Request.Header.Get("Cf-Connecting-Ip")
	updatePayload.CFCookie = c.Request.Header.Get("Cookie")

	err = database.UpdateUser(id, updatePayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}

func validateUpdateUser(payload models.UserUpdate) error {
	if len(payload.Country) < 2 {
		return fmt.Errorf("invalid country")
	}
	if len(payload.Address) < 2 {
		return fmt.Errorf("invalid address")
	}
	if len(payload.City) < 2 {
		return fmt.Errorf("invalid city")
	}
	return nil
}

// validates the form
func newUserValidation(user models.RegisterUser) error {
	if !validateEmail(user.Email) {
		return fmt.Errorf("please provide a valid email address")
	}
	if len(user.Password) < 8 {
		return fmt.Errorf("please enter a minimum of 8 characters for the password")
	}
	if len(user.FirstName) < 2 {
		return fmt.Errorf("invalid first name")
	}
	if len(user.LastName) < 2 {
		return fmt.Errorf("invalid last name")
	}
	if len(user.Country) < 2 {
		return fmt.Errorf("invalid country")
	}
	if len(user.Address) < 2 {
		return fmt.Errorf("invalid address")
	}
	if len(user.City) < 2 {
		return fmt.Errorf("invalid city")
	}
	return nil
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?" +
		"(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return Re.MatchString(email)
}
