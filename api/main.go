package main

import (
	"log"

	"github.com/hiromu-saito/trip-note-backend/database"
	"github.com/hiromu-saito/trip-note-backend/server"
)

func main() {
	log.Println("Start App...")
	database.Connect()
	server.StartApp().Run()
}
