package dnk

import (
	"fmt"
	"time"

	"github.com/pariz/gountries"
)

var (

	// mappingTable Contains data about sequence ranges for specific years. (Sequence encodes decade)
	mappingTable []mappingTableRow = []mappingTableRow{
		{Year: []int{1858, 1899}, SequenceRange: []int{5000, 5999}, Sex: []string{"female", "male"}},
		{Year: []int{1858, 1899}, SequenceRange: []int{6000, 6999}, Sex: []string{"female", "male"}},
		{Year: []int{1858, 1899}, SequenceRange: []int{7000, 7999}, Sex: []string{"female", "male"}},
		{Year: []int{1858, 1899}, SequenceRange: []int{8000, 8999}, Sex: []string{"female", "male"}},
		{Year: []int{1900, 1999}, SequenceRange: []int{1 - 1, 999}, Sex: []string{"female", "male"}},
		{Year: []int{1900, 1999}, SequenceRange: []int{1000, 1999}, Sex: []string{"female", "male"}},
		{Year: []int{1900, 1999}, SequenceRange: []int{2000, 2999}, Sex: []string{"female", "male"}},
		{Year: []int{1900, 1999}, SequenceRange: []int{3000, 3999}, Sex: []string{"female", "male"}},
		{Year: []int{1937, 1999}, SequenceRange: []int{9000, 9999}, Sex: []string{"female", "male"}},
		{Year: []int{1937, 1999}, SequenceRange: []int{4000, 4999}, Sex: []string{"female", "male"}},
		{Year: []int{2000, 2036}, SequenceRange: []int{4000, 4999}, Sex: []string{"female", "male"}},
		{Year: []int{2000, 2036}, SequenceRange: []int{9000, 9999}, Sex: []string{"female", "male"}},
		{Year: []int{2000, 2057}, SequenceRange: []int{5000, 5999}, Sex: []string{"female", "male"}},
		{Year: []int{2000, 2057}, SequenceRange: []int{7000, 7999}, Sex: []string{"female", "male"}},
		{Year: []int{2000, 2057}, SequenceRange: []int{8000, 8999}, Sex: []string{"female", "male"}},
	}

	// Start and end for dates
	startTime = time.Date(1858, 1, 1, 0, 0, 0, 0, time.UTC) // 1858-01-01 (inclusive)
	endTime   = time.Date(2058, 1, 1, 0, 0, 0, 0, time.UTC) // 2057-12-31 (inclusive)

	// Weights for the token components [4,3,2,7,6,5,4,3,2,1]
	dateWeights     []int = []int{4, 3, 2, 7, 6, 5}
	sequenceWeights []int = []int{4, 3, 2, 1}

	// ISO3301 standard for Denmark
	ISO3301, _ = gountries.New().FindCountryByAlpha("dnk")

	// TotalTokens holds the total number of tokens available
	TotalTokens uint = TotalTokensAvailable() // 344061000

	// Database stuff
	collectionName = fmt.Sprintf("%s_%s_%s", ISO3301.Alpha3, ISO3301.Alpha2, ISO3301.CCN3)
)
