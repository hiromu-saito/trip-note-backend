package main

import (
	"log"

	"github.com/hiromu-saito/trip-note-backend/server"
)

func main() {
	log.Println("Start App...")

	server.StartApp().Run()
}
