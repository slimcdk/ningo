package dnk

import (
	"time"
	"github.com/pariz/gountries"
)

// ISO3301 standard for Denmark
var iso3301, _ = gountries.New().FindCountryByName("DNK")


// mappingTable Contains data about sequence ranges for specific years. (Sequence encodes decade)
var mappingTable []mappingTableRow = []mappingTableRow{
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
var startTime = time.Date(1858, 1, 1, 0, 0, 0, 0, time.UTC) // 1858-01-01
var endTime = time.Date(2058, 1, 1, 0, 0, 0, 0, time.UTC)   // 2057-12-31


// Weights for the token components
var dateWeights []int = []int{4, 3, 2, 7, 6, 5}
var sequenceWeights []int = []int{4, 3, 2, 1}

// Total number of tokens available
var TotalTokens uint = TotalTokensAvailable() // 344061000