package apptree

import (
	"encoding/json"
	"time"
)

type DateRange struct {
	Valid    bool
	ToDate   time.Time
	FromDate time.Time
}

func (DateRange) ValueType() Type {
	return Type_DateRange
}

func (rng DateRange) IsNull() bool {
	return !rng.Valid
}

func (rng *DateRange) UnmarshalText(bytes []byte) error {
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
	rng.Valid = true
	return nil
}

func (rng DateRange) MarshalText() ([]byte, error) {
	return json.Marshal(&struct {
		To   string `json:"to"`
		From string `json:"from"`
	}{
		To:   rng.ToDate.Format(DateFormat),
		From: rng.FromDate.Format(DateFormat),
	})
}

func NewDateRange(fromDate time.Time, toDate time.Time) DateRange {
	return DateRange{FromDate: fromDate, ToDate: toDate, Valid: true}
}

func NullDateRange() DateRange {
	return DateRange{Valid: false}
}
