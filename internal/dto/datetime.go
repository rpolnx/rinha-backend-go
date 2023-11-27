package dto

import (
	"encoding/json"
	"strings"
	"time"
)

type Datetime time.Time

// Implement Marshaler and Unmarshaler interface
func (j *Datetime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = Datetime(t)
	return nil
}

func (j Datetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j Datetime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
