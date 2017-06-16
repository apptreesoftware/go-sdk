package apptree

import (
	"time"
)

const dateFormat = "2006-01-02"

type Date struct {
	Time  time.Time
	Valid bool
}

func (Date) ValueType() Type {
	return Type_Date
}

func (d Date) IsNull() bool {
	return !d.Valid
}

func NewDate(time time.Time) Date {
	return Date{Time: time, Valid: true}
}

func NullDate() Date {
	return Date{Valid: false}
}

func (t Date) MarshalText() ([]byte, error) {
	if !t.Valid {
		return []byte(`null`), nil
	}
	return []byte(t.Time.Format(dateFormat)), nil
}

func (t *Date) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		t.Valid = false
		return nil
	}
	time, err := time.Parse(dateFormat, str)
	if err != nil {
		return err
	}
	t.Time = time
	t.Valid = true
	return nil
}

func (t *Date) SetValid(v time.Time) {
	t.Time = v
	t.Valid = true
}

// Ptr returns a pointer to this Time's value, or a nil pointer if this Time is null.
func (t Date) Ptr() *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}
