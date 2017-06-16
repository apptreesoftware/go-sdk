package apptree

import (
	"time"
)

const dateTimeFormat = "2006-01-02 15:04:05"

type DateTime struct {
	Time  time.Time
	Valid bool
}

func (t DateTime) IsNull() bool {
	return !t.Valid
}

func (DateTime) ValueType() Type {
	return Type_DateTime
}

func NewDateTime(time time.Time) DateTime {
	return DateTime{Time: time, Valid: true}
}

func NullDateTime() DateTime {
	return DateTime{Valid: false}
}

func (t DateTime) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte(`null`), nil
	}
	return []byte(t.Time.Format(dateTimeFormat)), nil
}

func (t *DateTime) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}
	time, err := time.Parse(dateTimeFormat, str)
	if err != nil {
		return err
	}
	t.Time = time
	t.Valid = true
	return nil
}

func (t *DateTime) SetValid(v time.Time) {
	t.Time = v
	t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t DateTime) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
