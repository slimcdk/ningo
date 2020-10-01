package dnk

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/enriquebris/goworkerpool"
)

// SpawnPopulationWorkers sd
func SpawnPopulationWorkers(db *sql.DB) {
	start := time.Now()
	log.Printf("Spawning DNK workers..\r")

	// Count the number of total days from startTime to endTime
	var maxOperationsInQueue uint = uint(endTime.Sub(startTime).Hours() / 24)

	pool, err := goworkerpool.NewPoolWithOptions(goworkerpool.PoolOptions{
		TotalInitialWorkers:          uint(maxOperationsInQueue / 10),
		MaxWorkers:                   uint(maxOperationsInQueue),
		MaxOperationsInQueue:         maxOperationsInQueue,
		WaitUntilInitialWorkersAreUp: false,
		LogVerbose:                   false,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	pool.SetWorkerFunc(singleDayPopulationWorker)

	// Enqueue jobs
	for i := 0; i < int(maxOperationsInQueue); i++ {
		pool.AddTask(workerData{
			Date: startTime.AddDate(0, 0, i),
			Db:   db,
		})
	}

	// Kill all workers after the currently enqueued jobs get processed
	pool.LateKillAllWorkers()

	// Wait while at least one worker is alive
	pool.Wait()
	log.Printf("Spawning DNK workers.. Done after %s\r\n", time.Since(start))
}

func singleDayPopulationWorker(data interface{}) bool {
	// Check if we have the data we need
	wData, ok := data.(workerData)
	if !ok {
		fmt.Printf("No data\n")
		return false
	}

	currentDate := wData.Date

	// Generate the tokens
	_, err := GenerateTokensForDay(currentDate)
	if err != nil {
		log.Println("Error occured", err)
		return false
	}

	//log.Printf("Got data %s. Generated %d tokens!\n", currentDate, len(tokens))

	return true
}
