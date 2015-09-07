package main

import (
	"surf-cams/app/surfcam"
	"surf-cams/app/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/id/:id", surfcam.GetCamByID)
	router.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.Run(":" + utils.GetPort())
}
