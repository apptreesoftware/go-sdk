package apptree

import (
	"encoding/json"
)

type ToManyRelationship struct {
	Items []Record
}

func (ToManyRelationship) ValueType() Type {
	return Type_Relationship
}

func (v ToManyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Items)
}

func (v ToManyRelationship) IsNull() bool {
	return false
}
