package app

import (
	"CRUD_Go_gin/internal/database"
	"CRUD_Go_gin/internal/transport"
	"log"
)

func Run() {
	if err := database.InitDatabase(); err != nil {
		log.Fatal(err)
		return
	}

	log.Fatal(transport.Run(transport.NewServer()))
}
