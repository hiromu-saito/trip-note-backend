package server

import "github.com/hirom-saito/trip-note-backend/controllers/test"

func mapUrls() {
	router.GET("/test", test.Test)
}
