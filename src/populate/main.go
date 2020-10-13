package main

import (
	"fmt"
	"log"
	"time"

	storage "github.com/slimcdk/rainbow-nin/pkg/domains/database"
	"github.com/slimcdk/rainbow-nin/pkg/domains/dnk"
)

func main() {
	fmt.Printf("Go token populator!\n")

	db, err := storage.Init()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	log.Printf("Populating %d tokens for nation %s..", dnk.TotalTokens, dnk.ISO3301.Alpha3)
	dnk.SpawnPopulationWorkers()
	log.Printf("Done after %s after %s", dnk.ISO3301.Alpha3, time.Since(start))

}
