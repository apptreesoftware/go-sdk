package apptree

import (
	"encoding/json"
	"testing"
)

var (
	floatJSON = []byte(`"1.2345"`)
)

func TestFloatFrom(t *testing.T) {
	f := NewFloat(1.2345)
	assertFloat(t, f, "FloatFrom()")

	zero := NewFloat(0)
	if !zero.Valid {
		t.Error("FloatFrom(0)", "is invalid, but should be valid")
	}
}

func TestFloatFromPtr(t *testing.T) {
	n := float64(1.2345)
	iptr := &n
	f := FloatFromPtr(iptr)
	assertFloat(t, f, "FloatFromPtr()")

	null := FloatFromPtr(nil)
	assertNullFloat(t, null, "FloatFromPtr(nil)")
}

func TestUnmarshalFloat(t *testing.T) {
	var f Float
	err := json.Unmarshal(floatJSON, &f)
	maybePanic(err)
	assertFloat(t, f, "float json")

	var null Float
	err = json.Unmarshal(nullJSON, &null)
	maybePanic(err)
	assertNullFloat(t, null, "null json")

	var badType Float
	err = json.Unmarshal(boolJSON, &badType)
	if err == nil {
		panic("err should not be nil")
	}
	assertNullFloat(t, badType, "wrong type json")
}

func TestFloatPointer(t *testing.T) {
	f := NewFloat(1.2345)
	ptr := f.Ptr()
	if *ptr != 1.2345 {
		t.Errorf("bad %s float: %#v ≠ %v\n", "pointer", ptr, 1.2345)
	}

	null := NullFloat()
	ptr = null.Ptr()
	if ptr != nil {
		t.Errorf("bad %s float: %#v ≠ %s\n", "nil pointer", ptr, "nil")
	}
}

func TestFloatIsZero(t *testing.T) {
	f := NewFloat(1.2345)
	if f.IsZero() {
		t.Errorf("IsZero() should be false")
	}

	null := NullFloat()
	if !null.IsZero() {
		t.Errorf("IsZero() should be true")
	}

	zero := NewFloat(0)
	if zero.IsZero() {
		t.Errorf("IsZero() should be false")
	}
}

func TestFloatSetValid(t *testing.T) {
	change := NullFloat()
	assertNullFloat(t, change, "SetValid()")
	change.SetValid(1.2345)
	assertFloat(t, change, "SetValid()")
}

func TestFloatScan(t *testing.T) {
	var f Float
	err := f.Scan(1.2345)
	maybePanic(err)
	assertFloat(t, f, "scanned float")

	var null Float
	err = null.Scan(nil)
	maybePanic(err)
	assertNullFloat(t, null, "scanned null")
}

func assertFloat(t *testing.T, f Float, from string) {
	if f.Float64 != 1.2345 {
		t.Errorf("bad %s float: %f ≠ %f\n", from, f.Float64, 1.2345)
	}
	if !f.Valid {
		t.Error(from, "is invalid, but should be valid")
	}
}

func assertNullFloat(t *testing.T, f Float, from string) {
	if f.Valid {
		t.Error(from, "is valid, but should be invalid")
	}
}
