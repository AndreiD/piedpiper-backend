package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"piedpiper/caching"
	"piedpiper/handlers"
	"piedpiper/security"
	"time"
)

// InitializeRouter initialises all the routes in the app
func InitializeRouter() {
	// ----- PUBLIC APIs -----
	inMemoryStore := caching.NewInMemoryStore(1 * time.Hour)

	api := router.Group("/api")

	// info message
	api.GET("/ping", handlers.Ping)

	api.GET("/u/:id", caching.CachePage(inMemoryStore, time.Duration(3)*time.Second, handlers.GetUser))

	// get one user
	api.GET("/users", caching.CachePage(inMemoryStore, time.Duration(3)*time.Second, handlers.ListUsers))

	// register
	api.POST("/register", handlers.Register)

	// public auth
	api.POST("/login", security.AuthJWTUser().LoginHandler)
	api.GET("/logout", handlers.Logout)

	// user
	user := api.Group("/user")
	user.Use(security.AuthJWTUser().MiddlewareFunc())
	{
		// get user by it's jwt
		user.GET("/me", handlers.GetUserByToken)

		// updates the user profile
		user.PUT("/me", handlers.UpdateUserDetails)

		// chat
		user.GET("/chats", handlers.ListMyChats)
		user.POST("/chat", handlers.PostMessage)
	}

	// In case no route is found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "api endpoint not found"})
	})
}
