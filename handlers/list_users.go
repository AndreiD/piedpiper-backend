package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"piedpiper/database"
	"strconv"
)

// ListUsers creates a new user
func ListUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "500")
	pagex, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limitx, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if limitx > 500 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "there's a limit on how many records you can query"})
		return
	}
	if pagex < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page should be greater than 0"})
		return
	}

	users, err := database.ListUsersDB(pagex, limitx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
