package surfcam

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

var connectionString = "mongodb://localhost/"

// GetCamByID handler
func GetCamByID(c *gin.Context) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	spotid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	col := session.DB("surfcam").C("webcams")
	result := SurfCam{}
	err = col.Find(bson.M{"spotid": spotid}).One(&result)
	if err != nil {
		c.JSON(200, gin.H{"status": "not found"})
	} else {
		c.JSON(200, &result)
	}

}
