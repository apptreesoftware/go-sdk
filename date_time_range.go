package apptree

import (
	"encoding/json"
	"time"
)

type DateTimeRange struct {
	Valid    bool
	ToDate   time.Time `json:"-"`
	FromDate time.Time `json:"-"`
}

func (DateTimeRange) ValueType() Type {
	return Type_DateTimeRange
}

func (rng DateTimeRange) IsNull() bool {
	return !rng.Valid
}

func (rng *DateTimeRange) UnmarshalText(bytes []byte) error {
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
	rng.Valid = true
	return nil
}

func (rng DateTimeRange) MarshalText() ([]byte, error) {
	if !rng.Valid {
		return []byte(`null`), nil
	}
	return json.Marshal(&struct {
		To   string `json:"to"`
		From string `json:"from"`
	}{
		To:   rng.ToDate.Format(DateTimeFormat),
		From: rng.FromDate.Format(DateTimeFormat),
	})
}

func NewDateTimeRange(fromDate time.Time, toDate time.Time) DateTimeRange {
	return DateTimeRange{FromDate: fromDate, ToDate: toDate, Valid: true}
}

func NullDateTimeRange() DateTimeRange {
	return DateTimeRange{Valid: false}
}
