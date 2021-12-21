package server

import "github.com/hiromu-saito/trip-note-backend/controllers/test"

func mapUrls() {
	router.GET("/test", test.Test)
}
