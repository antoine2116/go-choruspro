package choruspro

import (
	"time"
)

type Date struct {
	time.Time
}

func (t *Date) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse("\"2006-01-02\"", string(data))
	return err
}

type Datetime struct {
	time.Time
}

func (t *Datetime) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse("\"2006-01-02 15:04\"", string(data))
	return err
}
