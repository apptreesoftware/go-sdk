package apptree

import (
	"encoding/json"
	"fmt"
	"time"
)

type Location struct {
	Valid     bool
	Latitude  float32  `json:"latitude"`
	Longitude float32  `json:"longitude"`
	Bearing   float32  `json:"bearing"`
	Speed     float32  `json:"speed"`
	Accuracy  float32  `json:"accuracy"`
	Elevation float32  `json:"elevation"`
	Timestamp DateTime `json:"timestamp"`
}

func (Location) ValueType() Type {
	return Type_Location
}

func (l Location) IsNull() bool {
	return !l.Valid
}

func NewLocation(latitude, longitude, bearing, speed, accuracy, elevation float32, timestamp DateTime) Location {
	return Location{
		Valid:     true,
		Latitude:  latitude,
		Longitude: longitude,
		Bearing:   bearing,
		Speed:     speed,
		Accuracy:  accuracy,
		Elevation: elevation,
		Timestamp: timestamp,
	}
}

func NullLocation() Location {
	return Location{Valid: false}
}

func (l Location) MarshalText() ([]byte, error) {
	if !l.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(&struct {
		Latitude  float32  `json:"latitude"`
		Longitude float32  `json:"longitude"`
		Bearing   float32  `json:"bearing"`
		Speed     float32  `json:"speed"`
		Accuracy  float32  `json:"accuracy"`
		Elevation float32  `json:"elevation"`
		Timestamp DateTime `json:"timestamp"`
	}{
		Latitude:  l.Latitude,
		Longitude: l.Longitude,
		Bearing:   l.Bearing,
		Speed:     l.Speed,
		Accuracy:  l.Accuracy,
		Elevation: l.Elevation,
		Timestamp: l.Timestamp,
	})
}

type iLocation Location

func (l *Location) UnmarshalText(b []byte) error {
	str := string(b)
	if len(str) == 0 {
		l.Valid = false
		return nil
	}
	var uItem iLocation
	err := json.Unmarshal(b, &uItem)
	if err != nil {
		return err
	}

	l.Latitude = uItem.Latitude
	l.Longitude = uItem.Longitude
	l.Accuracy = uItem.Accuracy
	l.Timestamp = uItem.Timestamp
	l.Bearing = uItem.Bearing
	l.Elevation = uItem.Elevation
	l.Speed = uItem.Speed
	l.Valid = true
	return nil
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(dateTimeFormat))
	return []byte(stamp), nil
}
