package db

import (
	"gopkg.in/mgo.v2"
	"surf-cams/app/utils"
)

var connectionString = utils.GetDBConnectionString()

// Connect to database - returns a session
func Connect() *mgo.Session {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	return session
}
