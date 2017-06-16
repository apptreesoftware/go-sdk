package apptree

import (
	"database/sql"
	"strconv"
)

// Int is an nullable int64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type TimeInterval struct {
	sql.NullInt64
}

func (i TimeInterval) ValueType() Type {
	return Type_TimeInterval
}

func (i TimeInterval) IsNull() bool {
	return !i.Valid
}

// NewInt creates a new Int
func NewTimeInterval(i int64) TimeInterval {
	return TimeInterval{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: true,
		},
	}
}

func NullTimeInterval() TimeInterval {
	return TimeInterval{
		NullInt64: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
	}
}

// IntFromPtr creates a new Int that be null if i is nil.
func TimeIntervalFromPtr(i *int64) TimeInterval {
	if i == nil {
		return NullTimeInterval()
	}
	return NewTimeInterval(*i)
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int if the input is a blank or not an integer.
// It will return an error if the input is not an integer, blank, or "null".
func (i *TimeInterval) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		i.Valid = false
		return nil
	}
	var err error
	i.Int64, err = strconv.ParseInt(string(text), 10, 64)
	i.Valid = err == nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Int is null.
func (i TimeInterval) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte(`null`), nil
	}
	return []byte(strconv.FormatInt(i.Int64, 10)), nil
}

// SetValid changes this Int's value and also sets it to be non-null.
func (i *TimeInterval) SetValid(n int64) {
	i.Int64 = n
	i.Valid = true
}

// Ptr returns a pointer to this Int's value, or a nil pointer if this Int is null.
func (i TimeInterval) Ptr() *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

// IsZero returns true for invalid Ints, for future omitempty support (Go 1.4?)
// A non-null Int with a 0 value will not be considered zero.
func (i TimeInterval) IsZero() bool {
	return !i.Valid
}
