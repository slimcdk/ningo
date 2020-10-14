package dnk

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// Token stores the data for a single token
type Token struct {
	Token      string     `gorm:"primaryKey;type:char(11)"`
	Nation     string     `gorm:"not null;type:char(3)"`
	Attributes Attributes `gorm:"embedded"`
	CreatedAt  time.Time  `gorm:"autoCreateTime"`
}

// Attributes are special attributes for this token
type Attributes struct {
	Date         time.Time `gorm:"not null;type:date"`
	Sequence     string    `gorm:"not null;type:char(4)"`
	ControlDigit string    `gorm:"not null;type:char(1)"`
	TokenSeries  string    `gorm:"not null;type:varchar(255)"`
	Sex          string    `gorm:"not null;type:varchar(255)"`
	Sum          int       `gorm:"not null;type:smallint"`
}

// bagpack is a data holder for workers
type bagpack struct {
	Date time.Time
	Db   *gorm.DB
}

// mappingTableRow contains sequence and gender data for a year
type mappingTableRow struct {
	Year          []int
	SequenceRange []int
	Sex           []string
}

// TableName overrides the table name used by tokens to `dnk`
func (Token) TableName() string {
	return strings.ToLower(ISO3301.Alpha3)
}
