package dnk

import (
	"fmt"
	"time"
)

// GenerateTokensForDay -> Computes tokens for a specific date
func GenerateTokensForDay(currentDate time.Time) ([]Token, error) {

	var tokens []Token

	// Extract date components
	currentYear, currentMonth, currentDay := currentDate.Date()

	// Construct date part of token and compute weighted sum of date
	tokenDateArray := []int{
		int(currentDay / 10),
		int(currentDay % 10),
		int(currentMonth / 12),
		int(currentMonth % 12),
		int(currentYear / 10 % 10),
		int(currentYear % 10),
	}
	tokenDateSum := weightedToken(tokenDateArray, dateWeights)

	// Go through every matching row for that year
	rows := rowsForYear(currentYear)
	for _, row := range rows {

		// Go through sequence range for that year
		for seq := row.SequenceRange[0]; seq <= row.SequenceRange[1]; seq++ {

			// Compute control digit and weighted sum of sequence number
			controlDigit := int(seq % 10)
			tokenSequenceArray := []int{
				int(seq / 1000),
				int(seq/100) % 10,
				int(seq % 100 / 10),
				controlDigit,
			}
			tokenSequenceSum := weightedToken(tokenSequenceArray, sequenceWeights)
			tokenSum := tokenDateSum + tokenSequenceSum

			// Create token entity
			tokenData := Token{
				Token: fmt.Sprintf("%02d%02d%02d-%04d", currentDay, currentMonth, currentYear%100, seq),
				Attributes: Attributes{
					Date:         currentDate,
					Sequence:     fmt.Sprintf("%03d", seq),
					ControlDigit: fmt.Sprintf("%d", controlDigit),
					TokenSeries:  fmt.Sprintf("%d. series", tokenSeries(seq, tokenSum)),
					Sex:          row.Sex[seq%2],
					Sum:          tokenSum,
				},
			}

			// Store token in array
			tokens = append(tokens, tokenData)
		}
	}

	// Return the amount of tokens created and nil error
	return tokens, nil
}
