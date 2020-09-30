package dnk

import "time"

// MappingTableRow contains sequence and gender data for a year
type MappingTableRow struct {
	Year          []int
	SequenceRange []int
	Sex           []string
}

// MappingTable First year HAS to be first and last year HAS to be last
var MappingTable []MappingTableRow = []MappingTableRow{
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

var dateWeights []int = []int{4, 3, 2, 7, 6, 5}
var sequenceWeights []int = []int{4, 3, 2, 1}

var startTime = time.Date(1858, 1, 1, 0, 0, 0, 0, time.UTC)
var endTime = time.Date(2058, 1, 1, 0, 0, 0, 0, time.UTC)
