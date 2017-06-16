package apptree

import (
	"encoding/json"
	"fmt"
)

type ToManyRelationship struct {
	Items []Record
}

func (ToManyRelationship) ValueType() Type {
	return Type_Relationship
}

func (val ToManyRelationship) ToString() string {
	stringVal := fmt.Sprintf("%i items", len(val.Items))
	return stringVal
}

func (v ToManyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Items)
}

func (v ToManyRelationship) IsNull() bool {
	return false
}
