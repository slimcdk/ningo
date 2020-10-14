package dnk

import (
	"fmt"
	"log"
	"time"

	"github.com/enriquebris/goworkerpool"
	db "github.com/slimcdk/rainbow-nin/pkg/storage"
	"gorm.io/gorm"
)

// SpawnPopulationWorkers sd
func SpawnPopulationWorkers(storage *gorm.DB, resumePrevious bool) error {

	// Prepare database table
	err := storage.AutoMigrate(&Token{})
	if err != nil {
		return err
	}

	// Count the number of total days from startTime to endTime
	totalNumberOfDays := uint(endTime.Sub(startTime).Hours() / 24)
	maximumConcurrentProcesses := uint(db.MaxConns)

	// Setup worker pool
	pool, err := goworkerpool.NewPoolWithOptions(goworkerpool.PoolOptions{
		TotalInitialWorkers:          maximumConcurrentProcesses,
		MaxWorkers:                   maximumConcurrentProcesses,
		MaxOperationsInQueue:         totalNumberOfDays,
		WaitUntilInitialWorkersAreUp: false,
	})
	if err != nil {
		return err
	}

	// Continue from previous progress if told to
	startTime := continueProgress(storage, resumePrevious)

	// Enqueue jobs
	pool.SetWorkerFunc(singleDayPopulationWorker)
	for i := 0; i < int(totalNumberOfDays); i++ {
		pool.AddTask(bagpack{
			Date: startTime.AddDate(0, 0, i),
			Db:   storage,
		})
	}

	// Kill all workers after the currently enqueued jobs get processed
	err = pool.LateKillAllWorkers()
	if err != nil {
		return err
	}

	// Wait while at least one worker is alive
	err = pool.Wait()
	if err != nil {
		return err
	}

	// Verify that data store contain all tokens (or at least same amount)
	err = verifyStorage(storage)
	if err != nil {
		return err
	}

	// Everything went well
	return nil
}

func singleDayPopulationWorker(data interface{}) bool {
	// Check if we have the data we need
	bp, ok := data.(bagpack)
	if !ok {
		return false
	}

	currentDate := bp.Date
	storage := bp.Db

	// Generate the tokens
	tokens, err := generateTokensForDay(currentDate)
	if err != nil {
		return false
	}

	// Insert tokens into database
	result := storage.Create(&tokens)
	if result.Error != nil {
		log.Printf("Got error for date %s: %a \n", currentDate.Format("01-02-2006"), result.Error)
		return false
	}
	log.Printf("%s: Generated %d tokens and wrote %d \n", currentDate.Format("01-02-2006"), len(tokens), result.RowsAffected)

	// Tell worker manager that we are done
	return true
}

// Returns the queried progress or starts from the very beginning
func continueProgress(storage *gorm.DB, resumePrevious bool) time.Time {

	if !resumePrevious {
		return startTime
	}

	var result time.Time
	storage.Model(&Token{}).Select("MAX(date)").Find(&result)

	if result.After(startTime) {
		return result
	}

	return continueProgress(storage, false)
}

func verifyStorage(storage *gorm.DB) error {

	log.Print("Verifying storage.. ")

	var count int64
	storage.Model(&Token{}).Count(&count)

	if uint(count) != TotalTokens {
		return fmt.Errorf("Data is incomplete. Infact %d tokens are missing", TotalTokens-uint(count))
	}
	return nil
}
