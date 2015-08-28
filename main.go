package main

import (
	"github.com/gin-gonic/gin"
	"surf-cams/app/surfcam"
)

func main() {
	router := gin.Default()
	router.GET("/:id", surfcam.GetCamByID)
	router.Run(":9999")
}
