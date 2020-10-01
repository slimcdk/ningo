package database

import "os"

var mySQLHost = getEnv("MYSQL_HOST", "localhost")
var mySQLPort = getEnv("MYSQL_PORT", "3306")
var mySQLDb = getEnv("MYSQL_DB", "nin_graph")
var mySQLDbUser = getEnv("MYSQL_USER", "root")
var mySQLDbUserPass = getEnv("MYSQL_PASS", "root")

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
