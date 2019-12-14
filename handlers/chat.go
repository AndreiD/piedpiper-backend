package handlers

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"piedpiper/database"
	"piedpiper/models"
	"piedpiper/utils/log"
	"strconv"
	"time"
)

// ListMyChats lists a user's chats
func ListMyChats(c *gin.Context) {

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
	if user.Email == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "something fishy is going on..."})
		return
	}

	room, err := database.ListUsersChats(id, pagex, limitx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, room)

}

// PostMessage posts a chat
func PostMessage(c *gin.Context) {

	var chat models.Chat
	err := c.BindJSON(&chat)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload " + err.Error()})
		return
	}

	// get the user id from jwt
	jwtClaims := jwt.ExtractClaims(c)

	id, ok := jwtClaims["id"].(string)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "invalid request"})
		return
	}

	chat.FromUserID = id

	err = newChatValidation(chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//registerModel.LastIP = c.Request.Header.Get("Cf-Connecting-Ip")
	//registerModel.CFCookie = c.Request.Header.Get("Cookie")

	chat.CreatedAt = time.Now().Unix()

	err = database.CreateChat(chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(200)
}

// validates the chat
func newChatValidation(chat models.Chat) error {
	if len(chat.FromUserID) != 24 {
		return fmt.Errorf("invalid from user id")
	}
	if len(chat.ToUserID) != 24 {
		return fmt.Errorf("invalid to user id")
	}

	// TODO: verify that the from and to exist!

	if len(chat.Message) > 1000 {
		return fmt.Errorf("message length is limitted to 1000 characters")
	}
	return nil
}
