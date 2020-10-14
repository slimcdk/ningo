package storage

import (
	"os"
	"strconv"
)

var (
	mySQLHost       = getEnv("MYSQL_HOST", "localhost")
	mySQLPort       = getEnv("MYSQL_PORT", "13306")
	mySQLDb         = getEnv("MYSQL_DB", "rainbow_nin")
	mySQLDbUser     = getEnv("MYSQL_USER", "nin")
	mySQLDbUserPass = getEnv("MYSQL_PASS", "rainbow")
	dbConnections   = getEnv("MYSQL_CONNS", "100")

	// MaxConns is the maximum of concurrent database connctions
	MaxConns = 0
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func setupVars() error {

	// Setup MaxConns (parse int)
	n, err := strconv.Atoi(dbConnections)
	if err != nil {
		return err
	}
	MaxConns = n

	// Everything went well
	return nil
}
