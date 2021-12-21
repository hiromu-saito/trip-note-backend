package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApp - Start our application
func StartApp() *gin.Engine {

	mapUrls()

	log.Println("Start App...")
	return router
}
