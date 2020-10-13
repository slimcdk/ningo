package dnk

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Token stores the data for a single token
type Token struct {
	gorm.Model
	Token string `gorm:"primaryKey"`
	//	Nation     string
	Attributes Attributes
}

// Attributes are special attributes for this token
type Attributes struct {
	gorm.Model
	Date         time.Time
	Sequence     string
	ControlDigit string
	TokenSeries  string
	Sex          string
	Sum          int
}

type workerData struct {
	Date time.Time
	Db   *sql.DB
}

// mappingTableRow contains sequence and gender data for a year
type mappingTableRow struct {
	Year          []int
	SequenceRange []int
	Sex           []string
}
