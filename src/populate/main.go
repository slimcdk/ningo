package main

import (
	"fmt"
	"log"
	"time"

	"github.com/slimcdk/rainbow-nin/pkg/domains/dnk"
	"github.com/slimcdk/rainbow-nin/pkg/storage"
)

func main() {
	fmt.Println("National Identity Number Rainbow!")

	db, err := storage.Init()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	log.Printf("Populating %d tokens for nation %s..", dnk.TotalTokens, dnk.ISO3301.Alpha3)
	err = dnk.SpawnPopulationWorkers(db, true)
	if err != nil {
		log.Fatalf("Population for %s did not finish %a", dnk.ISO3301.Alpha3, err)
		err = dnk.SpawnPopulationWorkers(db, true)
		if err != nil {
			panic(err)
		}
	} else {
		log.Printf("Done after %s after %s", dnk.ISO3301.Alpha3, time.Since(start))
	}
}
