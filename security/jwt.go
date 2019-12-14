package security

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"piedpiper/database"
	"piedpiper/models"
	"piedpiper/utils/log"
	"strings"
	"time"
)

// AuthJWTUser handles the jwt part
func AuthJWTUser() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:            "user",
		SigningAlgorithm: "HS256",
		Key:              []byte("w2fwdsxxEw3w3Ko3llexyMpJEoFeLAwde"),
		Timeout:          9999 * time.Hour,
		MaxRefresh:       9999 * time.Hour,

		Authenticator: func(c *gin.Context) (i interface{}, e error) {
			var login models.LoginUser
			err := c.BindJSON(&login)
			if err != nil {
				log.Error(err)
				return nil, err
			}

			login.Email = strings.ToLower(login.Email)

			user, err := database.Login(login.Email, login.Password)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			return &models.User{
				ID:   user.ID,
				Role: user.Role,
			}, nil
		},
		// maps the claims in the JWT.
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id":   v.ID,
					"role": v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			jwtClaims := jwt.ExtractClaims(c)
			return jwtClaims["role"] == "user"
		},

		TokenHeadName: "Token",
		TimeFunc:      time.Now,
	}
	return authMiddleware
}
