package utils

import "os"

// GetPort - Get the port number set
// by the environment PORT
// or return default
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	return port
}

// GetDBConnectionString - Get the connection string
// set by environment DB_CONNECTION
// or return default
func GetDBConnectionString() string {
	connectionString := os.Getenv("DB_CONNECTION")
	if connectionString == "" {
		connectionString = "mongodb://localhost/surfcam"
	}
	return connectionString
}
