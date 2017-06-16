package apptree

import (
	"encoding/json"
	"testing"
)

var (
	colorJSONStr   = []byte(`"{\"r\":10,\"g\":10,\"b\":10,\"a\":10}"`)
	colorJSON      = []byte(`{"r":10,"g":10,"b":10,"a":10}`)
	nullColorJSON  = []byte(`null`)
	emptyColorJSON = []byte(`""`)
)

func TestMarshalColor(t *testing.T) {
	color := NewColor(10, 10, 10, 10)

	b, err := json.Marshal(&color)
	if err != nil {
		t.Error(err)
	}
	parseColorStr := string(b)
	staticStr := string(colorJSONStr)
	t.Log(parseColorStr)

	if parseColorStr != staticStr {
		t.Fatalf("Got %s expected %s", string(b), string(colorJSONStr))
	}
}

func TestUnmarshalColor(t *testing.T) {
	var color Color
	err := json.Unmarshal(colorJSONStr, &color)
	if err != nil {
		t.Error(err)
	}
	if color.Valid && color.Red != 10 && color.Green != 10 && color.Blue != 10 && color.Alpha != 10 {
		t.Fail()
	}
	var color2 Color
	err = json.Unmarshal(nullColorJSON, &color2)
	if err != nil {
		t.Error(err)
	}
	if color2.Valid {
		t.Fail()
	}

	var color3 Color
	err = json.Unmarshal(emptyColorJSON, &color3)
	if err != nil {
		t.Error(err)
	}
	if color3.Valid {
		t.Fail()
	}
}

func TestMarshalUnmarshalColorValid(t *testing.T) {
	color := NewColor(10, 10, 10, 10)

	b, err := json.Marshal(&color)
	if err != nil {
		t.Error(err)
	}

	var color2 Color
	err = json.Unmarshal(b, &color2)
	if err != nil {
		t.Error(err)
	}
	if color2 != color {
		t.Logf("%v - %v", color2, color)
		t.Fail()
	}
}
