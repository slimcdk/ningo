package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/slimcdk/rainbow-nin/pkg/domains/dnk"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mySQLHost       = getEnv("MYSQL_HOST", "localhost")
	mySQLPort       = getEnv("MYSQL_PORT", "13306")
	mySQLDb         = getEnv("MYSQL_DB", "rainbow_nin")
	mySQLDbUser     = getEnv("MYSQL_USER", "nin")
	mySQLDbUserPass = getEnv("MYSQL_PASS", "rainbow")
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func main() {
	fmt.Printf("Go token populator!\n")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", mySQLDbUser, mySQLDbUserPass, mySQLHost, mySQLPort, mySQLDb)
	storage, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	start := time.Now()
	log.Printf("Populating %d tokens for nation %s..", dnk.TotalTokens, dnk.ISO3301.Alpha3)
	dnk.SpawnPopulationWorkers(storage)
	log.Printf("Done after %s after %s", dnk.ISO3301.Alpha3, time.Since(start))

}
