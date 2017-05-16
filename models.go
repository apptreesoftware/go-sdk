package apptree

import (
	"encoding/json"
	"mime/multipart"
	"time"
)

const DateTimeFormat = `2006-01-02 15:04:05`
const DateFormat = `2006-01-02`

type DateTimeRange struct {
	ToDate   time.Time `json:"-"`
	FromDate time.Time `json:"-"`
}

func (rng *DateTimeRange) UnmarshalJSON(bytes []byte) error {
	var values map[string]string
	json.Unmarshal(bytes, &values)
	date, err := time.Parse(DateTimeFormat, values["from"])
	if err == nil {
		rng.FromDate = date
	}
	date, err = time.Parse(DateTimeFormat, values["to"])
	if err == nil {
		rng.ToDate = date
	}
	return nil
}

func (rng *DateTimeRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		To   string `json:"to"`
		From string `json:"from"`
	}{
		To:   rng.ToDate.Format(DateTimeFormat),
		From: rng.FromDate.Format(DateTimeFormat),
	})
}

func NewDateTimeRange(fromDate time.Time, toDate time.Time) DateTimeRange {
	return DateTimeRange{FromDate: fromDate, ToDate: toDate}
}

type DateRange struct {
	ToDate   time.Time
	FromDate time.Time
}

func (rng *DateRange) UnmarshalJSON(bytes []byte) error {
	var values map[string]string
	json.Unmarshal(bytes, &values)
	date, err := time.Parse(DateFormat, values["from"])
	if err == nil {
		rng.FromDate = date
	}
	date, err = time.Parse(DateFormat, values["to"])
	if err == nil {
		rng.ToDate = date
	}
	return nil
}

func (rng *DateRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		To   string `json:"to"`
		From string `json:"from"`
	}{
		To:   rng.ToDate.Format(DateFormat),
		From: rng.FromDate.Format(DateFormat),
	})
}

func NewDateRange(fromDate time.Time, toDate time.Time) DateRange {
	return DateRange{FromDate: fromDate, ToDate: toDate}
}

type Image struct {
	ImageURL  string         `json:"imageURL"`
	UploadKey string         `json:"uploadKey"`
	FilePart  multipart.File `json:"-"`
}

type Location struct {
	Latitude  float32
	Longitude float32
	Bearing   float32
	Speed     float32
	Accuracy  float32
	Elevation float32
	Timestamp time.Time
}

type Color struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
	Alpha int `json:"a"`
}
