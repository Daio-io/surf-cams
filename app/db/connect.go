package db

import (
	"gopkg.in/mgo.v2"
	"os"
)

var connectionString = getConnectionString()

// Connect to database - returns a session
func Connect() *mgo.Session {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	return session
}

func getConnectionString() string {
	connectionString := os.Getenv("DB_CONNECTION")
	if connectionString == "" {
		connectionString = "mongodb://localhost/surfcam"
	}
	return connectionString
}
