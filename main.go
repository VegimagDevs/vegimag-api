package main

import (
	_ "github.com/VegimagDevs/vegimag-api/docs"
	"github.com/VegimagDevs/vegimag-api/handlers"
	"github.com/VegimagDevs/vegimag-api/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
		Storage:              storage,
		MailgunDomain:        "mg.vegimag.org",
		MailgunPrivateAPIKey: "b485a650c588d077e0b2543568a30766-115fe3a6-78d481d5",
		MailgunSender:        "no-reply@vegimag.org",
	})

	router.POST("/users", handlers.CreateUser)
	router.GET("/users/validation-token", handlers.GetValidationToken)
	router.POST("/users/validate", handlers.ValidateUser)
	router.POST("/sessions", handlers.CreateSession)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router.Run()
}

// @title Vegimag API
// @version 1.0
// @description The API server of the Vegimag project.

// @license.name LGPL3
// @license.url http://www.gnu.org/licenses/lgpl-3.0.en.html

// @host https://api.vegimag.org
// @BasePath /

func main() {
	app := &cli.App{
		Name: "vegimag-api",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "storage-path",
				Value:   "data.db",
				EnvVars: []string{"VEGIMAG_API_STORAGE_PATH"},
			},
			&cli.StringFlag{
				Name:    "base-url",
				Value:   "https://api.vegimag.org/",
				EnvVars: []string{"VEGIMAG_API_BASE_URL"},
			},
		},
		Action: run,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
