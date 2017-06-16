package apptree

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Int is an nullable int64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int struct {
	sql.NullInt64
}

func (i Int) IsNull() bool {
	return !i.Valid
}

func (i Int) ValueType() Type {
	return Type_Int
}

// NewInt creates a new Int
func NewInt(i int64) Int {
	return Int{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: true,
		},
	}
}

func NullInt() Int {
	return Int{
		NullInt64: sql.NullInt64{
			Int64: 0,
			Valid: false,
		},
	}
}

// IntFromPtr creates a new Int that be null if i is nil.
func IntFromPtr(i *int64) Int {
	if i == nil {
		return NullInt()
	}
	return NewInt(*i)
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int if the input is a blank or not an integer.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int) UnmarshalText(text []byte) error {
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

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int is null.
func (i Int) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, strconv.FormatInt(i.Int64, 10))), nil
}

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Int is null.

// SetValid changes this Int's value and also sets it to be non-null.
func (i *Int) SetValid(n int64) {
	i.Int64 = n
	i.Valid = true
}

// Ptr returns a pointer to this Int's value, or a nil pointer if this Int is null.
func (i Int) Ptr() *int64 {
	if !i.Valid {
		return nil
	}
	return &i.Int64
}

// IsZero returns true for invalid Ints, for future omitempty support (Go 1.4?)
// A non-null Int with a 0 value will not be considered zero.
func (i Int) IsZero() bool {
	return !i.Valid
}
