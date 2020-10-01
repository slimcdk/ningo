package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Init initializes the database handle
func Init() (*sql.DB, error) {

	// Create the database handle, confirm driver is present
	var dbURI string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mySQLDbUser, mySQLDbUserPass, mySQLHost, mySQLPort, mySQLDb)
	db, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, err
	}

	// Database setup
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db, nil
}
