package main

import (
	"github.com/VegimagDevs/vegimag-api/handlers"
	"github.com/VegimagDevs/vegimag-api/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

func run(ctx *cli.Context) error {
	storage := storage.New(&storage.Config{
		Path: ctx.String("storage-path"),
	})

	if err := storage.Open(); err != nil {
		return err
	}

	router := gin.Default()
	router.Use(cors.Default())

	handlers := handlers.New(&handlers.Config{
		Storage: storage,
	})

	router.POST("/users", handlers.CreateUser)
	router.POST("/sessions", handlers.CreateSession)

	return router.Run()
}

func main() {
	app := &cli.App{
		Name: "vegimag-api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "storage-path",
				Value: "data.db",
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
