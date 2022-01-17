package server

import (
	"github.com/hiromu-saito/trip-note-backend/controllers/auth"
	"github.com/hiromu-saito/trip-note-backend/controllers/memory"
	"github.com/hiromu-saito/trip-note-backend/controllers/test"
)

func mapUrls() {
	router.GET("/test", test.Test)

	//auth
	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.GET("/logout", auth.Logout)

	// memory
	router.GET("/memory", memory.GetMemories)
	router.PUT("/memory/:id", memory.UpdateMemories)
	router.DELETE("/memory/:id", memory.DeleteMemories)
}
