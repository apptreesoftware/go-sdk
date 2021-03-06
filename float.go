package apptree

import (
	"database/sql"
	"fmt"
	"strconv"
)

// Float is a nullable float64.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Float struct {
	sql.NullFloat64
}

func (v Float) IsNull() bool {
	return !v.Valid
}

func (Float) ValueType() Type {
	return Type_Float
}

// NewFloat creates a new Float
func NewFloat(f float64) Float {
	return Float{
		NullFloat64: sql.NullFloat64{
			Float64: f,
			Valid:   true,
		},
	}
}

func NullFloat() Float {
	return Float{
		NullFloat64: sql.NullFloat64{
			Float64: 0,
			Valid:   false,
		},
	}
}

// FloatFromPtr creates a new Float that be null if f is nil.
func FloatFromPtr(f *float64) Float {
	if f == nil {
		return NullFloat()
	}
	return NewFloat(*f)
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Float if the input is a blank or not an integer.
// It will return an error if the input is not an integer, blank, or "null".
func (f *Float) UnmarshalText(text []byte) error {
	str := string(text)
	if str == "" || str == "null" {
		f.Valid = false
		return nil
	}
	var err error
	f.Float64, err = strconv.ParseFloat(string(text), 64)
	f.Valid = err == nil
	return err
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Float is null.
func (f Float) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, strconv.FormatFloat(f.Float64, 'f', -1, 64))), nil
}

// SetValid changes this Float's value and also sets it to be non-null.
func (f *Float) SetValid(n float64) {
	f.Float64 = n
	f.Valid = true
}

// Ptr returns a pointer to this Float's value, or a nil pointer if this Float is null.
func (f Float) Ptr() *float64 {
	if !f.Valid {
		return nil
	}
	return &f.Float64
}

// IsZero returns true for invalid Floats, for future omitempty support (Go 1.4?)
// A non-null Float with a 0 value will not be considered zero.
func (f Float) IsZero() bool {
	return !f.Valid
}
