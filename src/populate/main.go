package main

import (
	"fmt"

	dbDriver "github.com/slimcdk/nin-graph/pkg/database"
	"github.com/slimcdk/nin-graph/pkg/domains/dnk"
)

func main() {
	fmt.Printf("Go token populator!\n")

	db, err := dbDriver.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to database:", version)

	dnk.SpawnPopulationWorkers(db)

}
