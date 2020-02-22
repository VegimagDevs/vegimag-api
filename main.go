package main

import (
	"github.com/VegimagDevs/vegimag-api/handlers"
	"github.com/VegimagDevs/vegimag-api/storage"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	storage := storage.New(&storage.Config{
		Path: "data.db",
	})

	if err := storage.Open(); err != nil {
		log.WithError(err).Fatal("Error opening the database")
	}

	router := gin.Default()

	handlers := handlers.New(&handlers.Config{
		Storage: storage,
	})

	router.POST("/users", handlers.CreateUser)
	router.POST("/sessions", handlers.CreateSession)

	router.Run()
}
