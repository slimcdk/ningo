package dnk

import (
	"fmt"
	"sync"
	"time"
)

func StartPopulationWorkers() {
	fmt.Println("DNK Population Worker")
	var workers uint = 200
	var wg sync.WaitGroup

	// wg.Add(1)
	// PopulationWorker(&wg, time.Date(1995, 3, 7, 0, 0, 0, 0, time.UTC), time.Date(1995, 3, 8, 0, 0, 0, 0, time.UTC))
	// return

	var days int = int(endTime.Sub(startTime).Hours() / 24)
	var dayBins []int = intIntoBins(days, workers)

	current := startTime
	// Spawn workers
	for i := 0; i < len(dayBins); i++ {
		wg.Add(1)
		next := current.AddDate(0, 0, dayBins[i])
		go PopulationWorker(&wg, current, next)
		current = next
	}
	wg.Wait()
}

func PopulationWorker(wg *sync.WaitGroup, start time.Time, end time.Time) {
	defer wg.Done()
	count := 0

	// Keep going as long as we are not at the end date
	var current = start
	for current.Before(end) {

		// Go through the table and find the row for current year
		currentYear, currentMonth, currentDay := current.Date()
		for i := 0; i < len(MappingTable); i++ {
			if MappingTable[i].Year[0] <= currentYear && currentYear <= MappingTable[i].Year[1] {

				// Construct date part of token
				//var tokenDate string = current.Format("ddMMyy")
				var tokenDateArray = []int{
					int(currentDay / 10),
					int(currentDay % 10),
					int(currentMonth / 12),
					int(currentMonth % 12),
					int(currentYear / 10 % 10),
					int(currentYear % 10),
				}
				var tokenDateSum int = weightedToken(tokenDateArray, dateWeights)

				// Go through the sequence number range for that year
				for j := MappingTable[i].SequenceRange[0]; j <= MappingTable[i].SequenceRange[1]; j++ {

					var controlDigit int = int(j % 10)
					var tokenSequenceArray []int = []int{
						int(j / 1000),
						int(j/100) % 10,
						int(j % 100 / 10),
						controlDigit,
					}

					var tokenSequenceSum int = weightedToken(tokenSequenceArray, sequenceWeights)
					var tokenSum int = tokenDateSum + tokenSequenceSum

					tokenAttributes := DNKAttributes{
						Date:         current,
						Sequence:     fmt.Sprintf("%03d", j),
						ControlDigit: fmt.Sprintf("%d", controlDigit),
						TokenSeries:  fmt.Sprintf("%d. series", tokenSeries(j, tokenSum)),
						Sex:          MappingTable[i].Sex[j%2],
						Sum:          tokenSum,
					}
					tokenData := Token{
						Token:      fmt.Sprintf("%02d%02d%02d-%04d", currentDay, currentMonth, currentYear%100, j),
						Attributes: tokenAttributes,
					}
					count++
					tokenData = tokenData
					//fmt.Printf("%s -> %d \r", tokenData.Token, count)
					//fmt.Println(tokenData.Token)
				}
			}
		}

		// Go to next day
		current = current.AddDate(0, 0, 1)
	}

	fmt.Printf("Generated %d tokens for year range %s %s\n", count, start, end)
}
