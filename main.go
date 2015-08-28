package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"surf-cams/app/surfcam"
)

var port = getPort()

func main() {
	router := gin.Default()
	router.GET("/:id", surfcam.GetCamByID)
	router.Run(":" + port)
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	return port
}
