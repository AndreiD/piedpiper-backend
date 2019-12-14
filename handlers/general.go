package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// Ping responds with pong
func Ping(c *gin.Context) {
	c.String(200, "pong at "+fmt.Sprint(time.Now().Unix()))
}

// Logout does nothing
func Logout(c *gin.Context) {
	c.Status(200)
}
