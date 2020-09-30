package dnk

import "time"

// Token stores the data for a single token
type Token struct {
	Token string
	//	Nation     string
	Attributes DNKAttributes
}

// DNKAttributes are special attributes for this token
type DNKAttributes struct {
	Date         time.Time
	Sequence     string
	ControlDigit string
	TokenSeries  string
	Sex          string
	Sum          int
}
