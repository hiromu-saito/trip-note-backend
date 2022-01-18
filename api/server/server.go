package server

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp - Start our application
func StartApp() *gin.Engine {

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8081",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Set-Cookie",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	mapUrls()

	log.Println("Start App...")
	return router
}
