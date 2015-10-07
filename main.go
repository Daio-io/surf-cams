package main

import (
	"github.com/gin-gonic/gin"
	"surf-cams/app/surfcam"
	"surf-cams/app/utils"
)

func main() {
	router := gin.Default()
	router.GET("/id/:id", surfcam.GetCamByID)
	router.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.Run(":" + utils.GetPort())
}
