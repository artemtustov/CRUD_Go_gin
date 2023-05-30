package app

import (
	"CRUD_Go_gin/internal/database"
	"CRUD_Go_gin/internal/services"
	"CRUD_Go_gin/internal/transport"
	logr "github.com/sirupsen/logrus"
	"log"
	"os"
)

func Run() {
	logr.SetFormatter(&logr.TextFormatter{})
	logr.SetOutput(os.Stdout)
	logr.SetLevel(logr.InfoLevel)

	services.Books = database.NewBooks()

	log.Fatal(transport.Run(transport.NewServer()))
}
