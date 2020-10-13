package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init initializes a connection pool to the database
func Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mySQLDbUser, mySQLDbUserPass, mySQLHost, mySQLPort, mySQLDb)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
