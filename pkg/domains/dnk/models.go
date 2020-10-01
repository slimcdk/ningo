package dnk

import (
	"database/sql"
	"time"
)

// Token stores the data for a single token
type Token struct {
	Token string
	//	Nation     string
	Attributes Attributes
}

// Attributes are special attributes for this token
type Attributes struct {
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
