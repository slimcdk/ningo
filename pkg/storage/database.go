package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init initializes a database handle
func Init() (*gorm.DB, error) {
	setupVars()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", mySQLDbUser, mySQLDbUserPass, mySQLHost, mySQLPort, mySQLDb)
	storage, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := storage.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(MaxConns)
	return storage, nil
}
