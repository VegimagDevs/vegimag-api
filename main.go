package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ginRouter := gin.Default()

	ginRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	ginRouter.Run()
}
