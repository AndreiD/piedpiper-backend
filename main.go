package main

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"os/signal"
	"piedpiper/configs"
	"piedpiper/database"
	"piedpiper/utils/log"
	"strconv"
	"time"
)

const version = "1.0 Merc"

var router *gin.Engine

var configuration *configs.ViperConfiguration

func main() {
	configuration = configs.NewConfiguration()
	configuration.Init()

	debug := configuration.GetBool("debug")
	log.Init(debug)
	log.Println("=======================================")
	log.Println("Starting PiedPiper " + version)
	log.Printf("Running on http://%s:%d", configuration.Get("server.host"), configuration.GetInt("server.port"))
	log.Println("=======================================")

	err := database.InitDatabase(configuration.Get("database.mongoURI"), configuration.Get("database.dbname"))
	if err != nil {
		log.Fatalf("can't connect to database %s", err.Error())
	}

	router = gin.New()
	if configuration.GetBool("debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(requestIDMiddleware())
	router.Use(corsMiddleware())
	router.Use(configurationMiddleware(configuration))

	InitializeRouter()

	server := &http.Server{
		Addr:           configuration.Get("server.host") + ":" + strconv.Itoa(configuration.GetInt("server.port")),
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 10, // 1Mb
	}
	server.SetKeepAlivesEnabled(true)

	// Serve'em
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen failed: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("initiated server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("Server Exiting. Bye!")
}

// requestIDMiddleware adds x-request-id
func requestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept,"+
				" origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// configurationMiddleware will add the configuration to the context
func configurationMiddleware(config *configs.ViperConfiguration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("configuration", config)
		c.Next()
	}
}
