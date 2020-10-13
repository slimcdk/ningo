package dnk

import (
	"fmt"

	"github.com/enriquebris/goworkerpool"
)

func PopulateDatabase() {

}

// SpawnPopulationWorkers sd
func SpawnPopulationWorkers() {

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
			Db:   nil,
		})
	}

	// Kill all workers after the currently enqueued jobs get processed
	pool.LateKillAllWorkers()

	// Wait while at least one worker is alive
	pool.Wait()
}

func singleDayPopulationWorker(data interface{}) bool {
	// Check if we have the data we need
	wData, ok := data.(workerData)
	if !ok {
		return false
	}

	currentDate := wData.Date

	// Generate the tokens
	_, err := generateTokensForDay(currentDate)
	if err != nil {
		return false
	}
	return true
}
