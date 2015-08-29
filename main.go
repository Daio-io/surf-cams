package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"surf-cams/app/surfcam"
)

var port = getPort()

func main() {
	router := gin.Default()
	router.GET("/id/:id", surfcam.GetCamByID)
	router.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.Run(":" + port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	return port
}
