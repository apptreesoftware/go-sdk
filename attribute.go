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
	return Type_Text
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
	return Type_Float
}

func (v FloatValue) ToString() string {
	return fmt.Sprintf("%f", v.Value)
}

type IntValue struct {
	Value int64
}

func (IntValue) ValueType() Type {
	return Type_Int
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
		return Type_DateTime
	}
	return Type_Date
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
	return Type_ListItem
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
	return Type_Relationship
}

func (val RelationshipValue) ToString() string {
	stringVal := fmt.Sprintf("%i items", len(val.Items))
	return stringVal
}

func (v RelationshipValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Items)
}

type TimeIntervalValue struct {
	Value int64
}

func (TimeIntervalValue) ValueType() Type {
	return Type_TimeInterval
}

func (v TimeIntervalValue) ToString() string {
	return fmt.Sprintf("%d", v.Value)
}

func (v TimeIntervalValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, v.ToString())), nil
}

type BooleanValue struct {
	Value bool
}

func (BooleanValue) ValueType() Type {
	return Type_Boolean
}

func (v BooleanValue) ToString() string {
	if v.Value {
		return "Y"
	} else {
		return "N"
	}
}

func (v BooleanValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, v.ToString())), nil
}

type DateRangeValue struct {
	Value DateRange
}

func (DateRangeValue) ValueType() Type {
	return Type_DateRange
}

func (v DateRangeValue) ToString() string {
	toString := v.Value.ToDate.Format(DateFormat)
	fromString := v.Value.FromDate.Format(DateFormat)
	return fmt.Sprintf(`"%s - %s"`, fromString, toString)
}

func (v DateRangeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

type DateTimeRangeValue struct {
	Value DateTimeRange
}

func (DateTimeRangeValue) ValueType() Type {
	return Type_DateTimeRange
}

func (v DateTimeRangeValue) ToString() string {
	toString := v.Value.ToDate.Format(DateTimeFormat)
	fromString := v.Value.FromDate.Format(DateTimeFormat)
	return fmt.Sprintf(`"%s - %s"`, fromString, toString)
}

func (v DateTimeRangeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

type ImageValue struct {
	Value Image
}

func (ImageValue) ValueType() Type {
	return Type_Image
}

func (v ImageValue) ToString() string {
	return fmt.Sprintf(`"%s"`, v.Value.ImageURL)
}

func (v ImageValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

type LocationValue struct {
	Value Location
}

func (LocationValue) ValueType() Type {
	return Type_Location
}

func (v LocationValue) ToString() string {
	return fmt.Sprintf(`"(%d, %d)"`, v.Value.Latitude, v.Value.Longitude)
}

func (v LocationValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value);
}

type SingleRelationshipValue struct {
	Value Record
}

func (SingleRelationshipValue) ValueType() Type {
	return Type_SingleRelationship
}

func (v SingleRelationshipValue) ToString() string {
	return fmt.Sprintf("Primary Key: %s", v.Value.PrimaryKey)
}

func (v SingleRelationshipValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}

type ColorValue struct {
	Value Color
}

func (ColorValue) ValueType() Type {
	return Type_Color
}

func (v ColorValue) ToString() string {
	return fmt.Sprintf(`"(%d, %d, %d)"`, v.Value.Red, v.Value.Green, v.Value.Blue)
}

func (v ColorValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Value)
}
