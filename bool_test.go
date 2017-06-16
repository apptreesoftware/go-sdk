package apptree

import (
	"encoding/json"
	"testing"
)

var (
	boolJSON  = []byte(`"Y"`)
	falseJSON = []byte(`"N"`)
)

func TestBoolFrom(t *testing.T) {
	b := NewBool(true)
	assertBool(t, b, "BoolFrom()")

	zero := NewBool(false)
	if !zero.Valid {
		t.Error("BoolFrom(false)", "is invalid, but should be valid")
	}
}

func TestBoolFromPtr(t *testing.T) {
	n := true
	bptr := &n
	b := BoolFromPtr(bptr)
	assertBool(t, b, "BoolFromPtr()")

	null := BoolFromPtr(nil)
	assertNullBool(t, null, "BoolFromPtr(nil)")
}

func TestUnmarshalBool(t *testing.T) {
	var b Bool
	err := json.Unmarshal(boolJSON, &b)
	maybePanic(err)
	assertBool(t, b, "bool json")

	var null Bool
	err = json.Unmarshal(nullJSON, &null)
	maybePanic(err)
	assertNullBool(t, null, "null json")

	var badType Bool
	err = json.Unmarshal(intJSON, &badType)
	if err == nil {
		panic("err should not be nil")
	}
	assertNullBool(t, badType, "wrong type json")
}

func TestTextUnmarshalBool(t *testing.T) {
	var b Bool
	err := b.UnmarshalText([]byte("Y"))
	maybePanic(err)
	assertBool(t, b, "UnmarshalText() bool")

	var zero Bool
	err = zero.UnmarshalText([]byte("N"))
	maybePanic(err)
	assertFalseBool(t, zero, "UnmarshalText() false")

	var blank Bool
	err = blank.UnmarshalText([]byte(""))
	maybePanic(err)
	assertNullBool(t, blank, "UnmarshalText() empty bool")

	var null Bool
	err = null.UnmarshalText([]byte("null"))
	maybePanic(err)
	assertNullBool(t, null, `UnmarshalText() "null"`)

	var invalid Bool
	err = invalid.UnmarshalText([]byte(":D"))
	if err == nil {
		panic("err should not be nil")
	}
	assertNullBool(t, invalid, "invalid json")
}

func TestMarshalBool(t *testing.T) {
	b := NewBool(true)
	data, err := json.Marshal(b)
	maybePanic(err)
	assertJSONEquals(t, data, `"Y"`, "non-empty json marshal")

	zero := NewBool(false)
	data, err = json.Marshal(zero)
	maybePanic(err)
	assertJSONEquals(t, data, `"N"`, "zero json marshal")
}

func TestBoolPointer(t *testing.T) {
	b := NewBool(true)
	ptr := b.Ptr()
	if *ptr != true {
		t.Errorf("bad %s bool: %#v ≠ %v\n", "pointer", ptr, true)
	}

	null := NullBool()
	ptr = null.Ptr()
	if ptr != nil {
		t.Errorf("bad %s bool: %#v ≠ %s\n", "nil pointer", ptr, "nil")
	}
}

func TestBoolIsZero(t *testing.T) {
	b := NewBool(true)
	if b.IsZero() {
		t.Errorf("IsZero() should be false")
	}

	null := NullBool()
	if !null.IsZero() {
		t.Errorf("IsZero() should be true")
	}

	zero := NewBool(false)
	if zero.IsZero() {
		t.Errorf("IsZero() should be false")
	}
}

func TestBoolSetValid(t *testing.T) {
	change := NullBool()
	assertNullBool(t, change, "SetValid()")
	change.SetValid(true)
	assertBool(t, change, "SetValid()")
}

func TestBoolScan(t *testing.T) {
	var b Bool
	err := b.Scan(true)
	maybePanic(err)
	assertBool(t, b, "scanned bool")

	var null Bool
	err = null.Scan(nil)
	maybePanic(err)
	assertNullBool(t, null, "scanned null")
}

func assertBool(t *testing.T, b Bool, from string) {
	if b.Bool != true {
		t.Errorf("bad %s bool: %v ≠ %v\n", from, b.Bool, true)
	}
	if !b.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}

func assertFalseBool(t *testing.T, b Bool, from string) {
	if b.Bool != false {
		t.Errorf("bad %s bool: %v ≠ %v\n", from, b.Bool, false)
	}
	if !b.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}

func assertNullBool(t *testing.T, b Bool, from string) {
	if b.Valid {
		t.Error(from, "is valid, but should be invalid")
	}
}
