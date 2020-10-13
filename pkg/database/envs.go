package database

import "os"

var (
	mySQLHost       = getEnv("MYSQL_HOST", "localhost")
	mySQLPort       = getEnv("MYSQL_PORT", "3306")
	mySQLDb         = getEnv("MYSQL_DB", "nin_graph")
	mySQLDbUser     = getEnv("MYSQL_USER", "root")
	mySQLDbUserPass = getEnv("MYSQL_PASS", "root")
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
