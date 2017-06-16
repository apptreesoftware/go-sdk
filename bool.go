package apptree

import (
	"database/sql"
	"errors"
)

type Bool struct {
	sql.NullBool
}

func (Bool) ValueType() Type {
	return Type_Boolean
}

func (b Bool) IsNull() bool {
	return !b.Valid
}

// NewBool creates a new Bool
func NewBool(b bool) Bool {
	return Bool{
		NullBool: sql.NullBool{
			Bool:  b,
			Valid: true,
		},
	}
}

func NullBool() Bool {
	return Bool{
		NullBool: sql.NullBool{
			Bool:  false,
			Valid: false,
		},
	}
}

// BoolFromPtr creates a new Bool that will be null if f is nil.
func BoolFromPtr(b *bool) Bool {
	if b == nil {
		return NullBool()
	}
	return NewBool(*b)
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Bool if the input is a blank or not an integer.
// It will return an error if the input is not an integer, blank, or "null".
func (b *Bool) UnmarshalText(text []byte) error {
	str := string(text)
	switch str {
	case "", "null":
		b.Valid = false
		return nil
	case "Y":
		b.Bool = true
	case "N":
		b.Bool = false
	default:
		b.Valid = false
		return errors.New("invalid input:" + str)
	}
	b.Valid = true
	return nil
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Bool is null.
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	if !b.Bool {
		return []byte(`"N"`), nil
	}
	return []byte(`"Y"`), nil
}

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Bool is null.
func (b Bool) MarshalText() ([]byte, error) {
	if !b.Valid {
		return []byte{}, nil
	}
	if !b.Bool {
		return []byte("N"), nil
	}
	return []byte("Y"), nil
}

// SetValid changes this Bool's value and also sets it to be non-null.
func (b *Bool) SetValid(v bool) {
	b.Bool = v
	b.Valid = true
}

// Ptr returns a pointer to this Bool's value, or a nil pointer if this Bool is null.
func (b Bool) Ptr() *bool {
	if !b.Valid {
		return nil
	}
	return &b.Bool
}

// IsZero returns true for invalid Bools, for future omitempty support (Go 1.4?)
// A non-null Bool with a 0 value will not be considered zero.
func (b Bool) IsZero() bool {
	return !b.Valid
}
