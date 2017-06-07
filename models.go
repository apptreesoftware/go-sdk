package apptree

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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

func NewImage(imageURL string, uploadKey string, filePart multipart.File) Image {
	return Image{ImageURL: imageURL, UploadKey: uploadKey, FilePart: filePart}
}

type Location struct {
	Latitude  float32      `json:"latitude"`
	Longitude float32      `json:"longitude"`
	Bearing   float32      `json:"bearing"`
	Speed     float32      `json:"speed"`
	Accuracy  float32      `json:"accuracy"`
	Elevation float32      `json:"elevation"`
	Timestamp NullDateTime `json:"timestamp"`
}

func NewLocation(latitude, longitude, bearing, speed, accuracy, elevation float32, timestamp NullDateTime) Location {
	return Location{
		Latitude:  latitude,
		Longitude: longitude,
		Bearing:   bearing,
		Speed:     speed,
		Accuracy:  accuracy,
		Elevation: elevation,
		Timestamp: timestamp,
	}
}

type Color struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
	Alpha int `json:"a"`
}

func NewColor(red, green, blue, alpha int) Color {
	return Color{
		Red:   red,
		Green: green,
		Blue:  blue,
		Alpha: alpha,
	}
}

type NullDate struct {
	Date  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func NewNullDate(date time.Time, valid bool) NullDate {
	return NullDate{Date: date, Valid: valid}
}

type NullDateTime struct {
	Date  time.Time
	Valid bool // Valid is true if Time is not NULL
}

func NewNullDateTime(date time.Time, valid bool) NullDateTime {
	return NullDateTime{Date: date, Valid: valid}
}

// Scan implements the Scanner interface.
func (nt *NullDate) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}
	nt.Date, nt.Valid = value.(time.Time), true
	return nil
}

func (t *NullDate) MarshalJSON() ([]byte, error) {
	if t.Valid {
		stamp := fmt.Sprintf("\"%s\"", t.Date.Format("2006-01-02"))
		return []byte(stamp), nil
	}
	return json.Marshal(nil)
}

func (a *NullDate) UnmarshalJSON(b []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		a.Valid = false
		return nil
	}
	a.Valid = true
	a.Date = parsedTime
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullDate) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Date, nil
}

func (nt *NullDateTime) Scan(value interface{}) error {
	if value == nil {
		nt.Valid = false
		return nil
	}
	nt.Date, nt.Valid = value.(time.Time), true
	return nil
}

func (t *NullDateTime) MarshalJSON() ([]byte, error) {
	if t.Valid {
		stamp := fmt.Sprintf("\"%s\"", t.Date.Format("2006-01-02 15:04:05"))
		return []byte(stamp), nil
	}
	return json.Marshal(nil)
}

func (a *NullDateTime) UnmarshalJSON(b []byte) error {
	parsedTime, err := time.Parse(`"2006-01-02 15:04:05"`, string(b))
	if err != nil {
		a.Valid = false
		return nil
	}
	a.Valid = true
	a.Date = parsedTime
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullDateTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Date, nil
}
