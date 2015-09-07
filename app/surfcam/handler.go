package surfcam

import (
	"strconv"
	"surf-cams/app/db"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// GetCamByID handler
func GetCamByID(c *gin.Context) {

	session := db.Connect()
	defer session.Close()

	spotid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	col := session.DB("").C("webcams")
	result := SurfCam{}
	err = col.Find(bson.M{"spotid": spotid}).One(&result)
	if err != nil {
		c.JSON(200, gin.H{"status": "not found"})
	} else {
		c.JSON(200, gin.H{"status": "success", "response": &result})
	}

}
