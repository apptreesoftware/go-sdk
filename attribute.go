package apptree

import (
	"encoding/json"
	"fmt"
	"time"
)

type TypedValue interface {
	ValueType() Type
	ToString() string
}

type TextValue struct {
	Value string
}

func NewTextValue(val string) TextValue {
	return TextValue{Value: val}
}

func (v TextValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, v.Value)), nil
}

func (TextValue) ValueType() Type {
	return Text
}

func (v TextValue) ToString() string {
	return v.Value
}

type FloatValue struct {
	Value float64
}

func (v FloatValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, v.ToString())), nil
}

func (FloatValue) ValueType() Type {
	return Float
}

func (v FloatValue) ToString() string {
	return fmt.Sprintf("%f", v.Value)
}

type IntValue struct {
	Value int64
}

func (IntValue) ValueType() Type {
	return Int
}

func (v IntValue) ToString() string {
	return fmt.Sprintf("%d", v.Value)
}

func (v IntValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, v.ToString())), nil
}

type DateTimeValue struct {
	Value   time.Time
	HasTime bool
}

func (v DateTimeValue) ValueType() Type {
	if v.HasTime {
		return DateTime
	}
	return Date
}

func (v DateTimeValue) ToString() string {
	if v.HasTime {
		return v.Value.Format(`2006-01-02 15:04:05`)
	} else {
		return v.Value.Format(`2006-01-02`)
	}
}

func (v DateTimeValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, v.ToString())), nil
}

type ListItemValue struct {
	ListItem ListElement
}

func (ListItemValue) ValueType() Type {
	return ListItem
}

func (v ListItemValue) ToString() string {
	return v.ListItem.Value
}

func (v ListItemValue) MarshalText() (text []byte, err error) {
	bytes, err := json.Marshal(v.ListItem)
	if err != nil {
		return nil, err
	}
	return []byte(bytes), nil
}

type RelationshipValue struct {
	Items []Record
}

func (RelationshipValue) ValueType() Type {
	return Relationship
}

func (val RelationshipValue) ToString() string {
	stringVal := fmt.Sprintf("%i items", len(val.Items))
	return stringVal
}

func (v RelationshipValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Items)
}
