package apptree

import (
	"encoding/json"
)

type SingleRelationship struct {
	Valid  bool
	Record Record
}

// NewBool creates a new Bool
func NewSingleRelationship(record Record, valid bool) SingleRelationship {
	return SingleRelationship{
		Valid:  valid,
		Record: record,
	}
}

func NullSingleRelationship() SingleRelationship {
	return SingleRelationship{Valid: false}
}

func (SingleRelationship) ValueType() Type {
	return Type_SingleRelationship
}

func (v SingleRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Record)
}

func (l SingleRelationship) IsNull() bool {
	return !l.Valid
}
