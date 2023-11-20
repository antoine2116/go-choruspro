package choruspro

import (
	"time"
)

// Date is a custom type to handle date format (YYYY-MM-DD)
type Date struct {
	time.Time
}

// UnmarshalJSON is a custom unmarshaler for Date
func (t *Date) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse("\"2006-01-02\"", string(data))
	return err
}

// Datetime is a custom type to handle datetime format (YYYY-MM-DD HH:MM)
type Datetime struct {
	time.Time
}

// UnmarshalJSON is a custom unmarshaler for Datetime
func (t *Datetime) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse("\"2006-01-02 15:04\"", string(data))
	return err
}
