package dnk

import (
	"math"
	"time"
)

// IntIntoBins transforms an integer into chuncks such that the total sum is equal the number
func intIntoBins(number int, nBins uint) []int {

	var n int = int(number / int(nBins))
	var bins []int = make([]int, int(nBins))

	for i := 0; i < int(nBins); i++ {
		bins[i] = n
	}
	bins[len(bins)-1] += number % int(nBins)
	return bins
}

func weightedToken(tokenSequenceArray, weights []int) int {
	var sum int = 0
	for i := 0; i < len(tokenSequenceArray); i++ {
		sum += tokenSequenceArray[i] * weights[i]
	}
	return sum
}

func tokenSeries(sequence, sum int) int {

	// zero series
	if sum%11 == 0 {
		return 0
	}

	// if token is for female
	if sequence%2 == 0 {
		if sequence < 6 {
			return sequence / 2
		}
		return int(math.Abs(float64(4 - (sequence%6)/2.0 - 1)))
	}

	// token must belong to a male and be non zero
	return (sequence%6-3)/2 + 2
}

func rowsForYear(year int) []mappingTableRow {

	var rows []mappingTableRow

	for _, row := range mappingTable {
		if row.Year[0] <= year && year <= row.Year[1] {
			rows = append(rows, row)
		}
	}
	return rows
}


func TotalTokensAvailable() uint {
	var total uint

	for _, row := range mappingTable {
		start := time.Date(row.Year[0], 1, 1, 0, 0, 0, 0, time.UTC)
		end := time.Date(row.Year[1]+1, 1, 1, 0, 0, 0, 0, time.UTC) 
		days := uint(math.Abs(float64(end.Sub(start).Hours() / 24)))

		total += days * uint(row.SequenceRange[1]-row.SequenceRange[0]+1)
	}

	return total
}

