package apptree

import (
	"encoding/json"
	"fmt"
)

type RecordSet struct {
	Records              []Record       `json:"records"`
	Configuration        *Configuration `json:"-"`
	MoreRecordsAvailable bool           `json:"moreRecordsAvailable"`
	Success              bool
}

const (
	RecordType_Normal     = "RECORD"
	RecordType_Attachment = "ATTACHMENT"
)

func NewRecordSet(configuration *Configuration) RecordSet {
	return RecordSet{Configuration: configuration}
}

func (rs *RecordSet) UnmarshalJSON(bytes []byte) error {
	var helper recordSetUnmarshalHelper
	err := json.Unmarshal(bytes, &helper)
	if err != nil {
		return err
	}
	parsedRecords := make([]Record, len(helper.Records))
	for index, rawRecord := range helper.Records {
		item, err := NewRecordFromJSON(rawRecord, rs.Configuration)
		if err != nil {
			return err
		}
		parsedRecords[index] = item
	}
	rs.Records = parsedRecords
	rs.MoreRecordsAvailable = helper.MoreRecordsAvailable
	rs.Success = helper.Success
	return nil
}

type recordSetUnmarshalHelper struct {
	Records              []json.RawMessage
	MoreRecordsAvailable bool `json:"moreRecordsAvailable"`
	Success              bool `json:"success"`
}

type recordUnmarshalHelper struct {
	PrimaryKey string            `json:"primaryKey"`
	RecordType string            `json:"recordType"`
	Attributes []json.RawMessage `json:"attributes"`
	CRUDStatus CRUDStatus
}

func (recHelper recordUnmarshalHelper) ToRecord(configuration *Configuration) (Record, error) {
	rec := Record{
		PrimaryKey:    recHelper.PrimaryKey,
		RecordType:    recHelper.RecordType,
		CRUDStatus:    recHelper.CRUDStatus,
		Attributes:    map[int]TypedValue{},
		Configuration: configuration,
	}
	for index, rawAttr := range recHelper.Attributes {
		err := rec.UnmarshalAttribute(index, rawAttr)
		if err != nil {
			return rec, err
		}
	}
	return rec, nil
}

type CRUDStatus string

const (
	StatusNone   CRUDStatus = "NONE"
	StatusCreate CRUDStatus = "CREATE"
	StatusUpdate CRUDStatus = "UPDATE"
	StatusDelete CRUDStatus = "DELETE"
	StatusRead   CRUDStatus = "READ"
)

type Record struct {
	PrimaryKey    string             `json:"primaryKey"`
	RecordType    string             `json:"recordType"`
	Attributes    map[int]TypedValue `json:"attributes"`
	Configuration *Configuration     `json:"-"`
	CRUDStatus    CRUDStatus
}

func NewRecordFromJSON(bytes []byte, configuration *Configuration) (Record, error) {
	recordUnmarshalHelper := recordUnmarshalHelper{}
	err := json.Unmarshal(bytes, &recordUnmarshalHelper)
	if recordUnmarshalHelper.CRUDStatus == "" {
		recordUnmarshalHelper.CRUDStatus = StatusRead
	}
	if err != nil {
		return Record{}, err
	}
	return recordUnmarshalHelper.ToRecord(configuration)
}

func (rec *Record) UnmarshalAttribute(index int, data []byte) error {
	configAttribute := rec.Configuration.getConfigurationAttribute(index)
	var value TypedValue
	if configAttribute == nil {
		return nil
	}
	switch configAttribute.Type {
	case Type_Relationship:
		var childRecords []recordUnmarshalHelper
		var childItems []Record
		err := json.Unmarshal(data, &childRecords)
		if err != nil {
			return err
		}
		for _, rawChildRec := range childRecords {
			childRec, err := rawChildRec.ToRecord(configAttribute.RelatedConfiguration)
			if err != nil {
				return err
			}
			childItems = append(childItems, childRec)
		}
		value = ToManyRelationship{Items: childItems}
	case Type_SingleRelationship:
		childRec := recordUnmarshalHelper{}
		err := json.Unmarshal(data, &childRec)
		if err != nil {
			return err
		}
		record, err := childRec.ToRecord(configAttribute.RelatedConfiguration)
		if err != nil {
			return err
		}
		value = NewSingleRelationship(record, true)
	case Type_Text:
		var textVal String
		err := json.Unmarshal(data, &textVal)
		if err != nil {
			return err
		}
		value = textVal
	case Type_Float:
		var floatVal Float
		err := json.Unmarshal(data, &floatVal)
		if err != nil {
			return err
		}
		value = floatVal
	case Type_Int:
		var intVal Int
		err := json.Unmarshal(data, &intVal)
		if err != nil {
			return err
		}
		value = intVal
	case Type_TimeInterval:
		var timeInterval TimeInterval
		err := json.Unmarshal(data, &timeInterval)
		if err != nil {
			return err
		}
		value = timeInterval
	case Type_Boolean:
		var boolVal Bool
		err := json.Unmarshal(data, &boolVal)
		if err != nil {
			return err
		}
		value = boolVal

	case Type_Color:
		var colorVal Color
		err := json.Unmarshal(data, &colorVal)
		if err != nil {
			return err
		}
		value = colorVal
	case Type_ListItem:
		var listItem ListItem
		err := json.Unmarshal(data, &listItem)
		if err != nil {
			return err
		}
		value = listItem
	case Type_Date:
		var date Date
		err := json.Unmarshal(data, &date)
		if err != nil {
			return err
		}
		value = date
	case Type_DateTime:
		var date DateTime
		err := json.Unmarshal(data, &date)
		if err != nil {
			return err
		}
		value = date
	case Type_DateRange:
		var dateRange DateRange
		err := json.Unmarshal(data, &dateRange)
		if err != nil {
			return err
		}
		value = dateRange
	case Type_DateTimeRange:
		var dateRange DateTimeRange
		err := json.Unmarshal(data, &dateRange)
		if err != nil {
			return err
		}
		value = dateRange
	case Type_Image:
		var image Image
		err := json.Unmarshal(data, &image)
		if err != nil {
			return err
		}
		value = image
	case Type_Location:
		var location Location
		err := json.Unmarshal(data, &location)
		if err != nil {
			return err
		}
		value = location
	}
	if value != nil {
		rec.Attributes[index] = value
	}
	return nil
}

func NewItem(configuration *Configuration) Record {
	return Record{
		Attributes:    map[int]TypedValue{},
		Configuration: configuration,
		RecordType:    RecordType_Normal,
		CRUDStatus:    StatusRead,
	}
}

func (item *Record) setValue(value TypedValue, index int) error {
	configAttribute := item.Configuration.getConfigurationAttribute(index)
	if configAttribute == nil {
		return fmt.Errorf("No attribute found for index %d", index)
	}
	if configAttribute.Type != value.ValueType() {
		return SetAttributeError{givenType: value.ValueType(), expectedType: configAttribute.Type, index: index}
	}
	if value == nil {
		return nil
	}
	item.Attributes[index] = value
	return nil
}

func (item *Record) AddToManyChildAtIndex(index int) (*Record, error) {
	attr := item.Configuration.getConfigurationAttribute(index)
	if attr == nil {
		return nil, fmt.Errorf("No attribute found at index %d", index)
	}
	if attr.Type != Type_Relationship {
		return nil, newInvalidTypeAttributeError(index, attr.Type, Type_Relationship)
	}
	existingVal := item.Attributes[index]
	var childrenVal ToManyRelationship
	if existingVal == nil {
		childrenVal = ToManyRelationship{}
	} else {
		childrenVal = existingVal.(ToManyRelationship)
	}
	childRecord := NewItem(attr.RelatedConfiguration)
	newItems := append(childrenVal.Items, childRecord)
	childrenVal.Items = newItems
	item.Attributes[index] = childrenVal
	return &childRecord, nil
}

func (item *Record) NewToOneRelationshipAtIndex(index int) (*Record, error) {
	attr := item.Configuration.getConfigurationAttribute(index)
	if attr == nil {
		return nil, fmt.Errorf("No attribute found at index %d", index)
	}
	if attr.Type != Type_SingleRelationship {
		return nil, newInvalidTypeAttributeError(index, attr.Type, Type_SingleRelationship)
	}
	childRecord := NewItem(attr.RelatedConfiguration)
	childrenVal := NewSingleRelationship(childRecord, true)
	item.Attributes[index] = childrenVal
	return &childRecord, nil
}

func (item *Record) GetString(index int) (String, error) {
	val := item.getValue(index)
	if val == nil {
		return NullString(), nil
	}
	if transformedVal, ok := val.(String); ok {
		return transformedVal, nil
	} else {
		return String{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Text)
	}
}

func (item *Record) SetString(index int, v String) error {
	return item.setValue(v, index)
}

func (item *Record) GetInt(index int) (Int, error) {
	val := item.getValue(index)
	if val == nil {
		return NullInt(), nil
	}
	if transformedVal, ok := val.(Int); ok {
		return transformedVal, nil
	} else {
		return Int{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Int)
	}
}

func (item *Record) SetInt(index int, v Int) error {
	return item.setValue(v, index)
}

func (item *Record) GetFloat(index int) (Float, error) {
	val := item.getValue(index)
	if val == nil {
		return NullFloat(), nil
	}
	if transformedVal, ok := val.(Float); ok {
		return transformedVal, nil
	} else {
		return Float{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Float)
	}
}

func (item *Record) SetFloat(index int, v Float) error {
	return item.setValue(v, index)
}

func (item *Record) GetTimeInterval(index int) (TimeInterval, error) {
	val := item.getValue(index)
	if val == nil {
		return NullTimeInterval(), nil
	}
	if transformedVal, ok := val.(TimeInterval); ok {
		return transformedVal, nil
	} else {
		return TimeInterval{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_TimeInterval)
	}
}

func (item *Record) SetTimeInterval(index int, v TimeInterval) error {
	return item.setValue(v, index)
}

func (item *Record) GetBool(index int) (Bool, error) {
	val := item.getValue(index)
	if val == nil {
		return NullBool(), nil
	}
	if transformedVal, ok := val.(Bool); ok {
		return transformedVal, nil
	} else {
		return Bool{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Boolean)
	}
}

func (item *Record) SetBool(index int, v Bool) error {
	return item.setValue(v, index)
}

func (item *Record) GetListItem(index int) (ListItem, error) {
	val := item.getValue(index)
	if val == nil {
		return NullListItem(), nil
	}
	if transformedVal, ok := val.(ListItem); ok {
		return transformedVal, nil
	} else {
		return ListItem{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_ListItem)
	}
}

func (item *Record) SetListItem(index int, v ListItem) error {
	return item.setValue(v, index)
}

func (item *Record) GetRelationship(index int) (ToManyRelationship, error) {
	val := item.getValue(index)
	if val == nil {
		return ToManyRelationship{}, nil
	}
	if transformedVal, ok := val.(ToManyRelationship); ok {
		return transformedVal, nil
	} else {
		return ToManyRelationship{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Relationship)
	}
}

func (item *Record) GetDate(index int) (Date, error) {
	val := item.getValue(index)
	if val == nil {
		return NullDate(), nil
	}
	if transformedVal, ok := val.(Date); ok {
		return transformedVal, nil
	} else {
		return Date{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Date)
	}
}

func (item *Record) SetDate(index int, v Date) error {
	return item.setValue(v, index)
}

func (item *Record) GetDateTime(index int) (DateTime, error) {
	val := item.getValue(index)
	if val == nil {
		return NullDateTime(), nil
	}
	if transformedVal, ok := val.(DateTime); ok {
		return transformedVal, nil
	} else {
		return DateTime{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_DateTime)
	}
}

func (item *Record) SetDateTime(index int, v DateTime) error {
	return item.setValue(v, index)
}

func (item *Record) GetDateRange(index int) (DateRange, error) {
	val := item.getValue(index)
	if val == nil {
		return NullDateRange(), nil
	}
	if transformedVal, ok := val.(DateRange); ok {
		return transformedVal, nil
	} else {
		return DateRange{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_DateRange)
	}
}

func (item *Record) SetDateRange(index int, v DateRange) error {
	return item.setValue(v, index)
}

func (item *Record) GetDateTimeRange(index int) (DateTimeRange, error) {
	val := item.getValue(index)
	if val == nil {
		return NullDateTimeRange(), nil
	}
	if transformedVal, ok := val.(DateTimeRange); ok {
		return transformedVal, nil
	} else {
		return DateTimeRange{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_DateTimeRange)
	}
}

func (item *Record) SetDateTimeRange(index int, v DateTimeRange) error {
	return item.setValue(v, index)
}

func (item *Record) GetLocation(index int) (Location, error) {
	val := item.getValue(index)
	if val == nil {
		return NullLocation(), nil
	}
	if transformedVal, ok := val.(Location); ok {
		return transformedVal, nil
	} else {
		return Location{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Location)
	}
}

func (item *Record) SetLocation(index int, v Location) error {
	return item.setValue(v, index)
}

func (item *Record) GetColor(index int) (Color, error) {
	val := item.getValue(index)
	if val == nil {
		return NullColor(), nil
	}
	if transformedVal, ok := val.(Color); ok {
		return transformedVal, nil
	} else {
		return Color{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Color)
	}
}

func (item *Record) SetColor(index int, v Color) error {
	return item.setValue(v, index)
}

func (item *Record) GetSingleRelationship(index int) (SingleRelationship, error) {
	val := item.getValue(index)
	if val == nil {
		return NewSingleRelationship(Record{}, false), nil
	}
	if transformedVal, ok := val.(SingleRelationship); ok {
		return transformedVal, nil
	} else {
		return SingleRelationship{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_SingleRelationship)
	}
}

func (item *Record) GetImage(index int) (Image, error) {
	val := item.getValue(index)
	if val == nil {
		return NullImage(), nil
	}
	if transformedVal, ok := val.(Image); ok {
		return transformedVal, nil
	} else {
		return Image{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Image)
	}
}

func (item *Record) SetImage(index int, v Image) error {
	return item.setValue(v, index)
}

func (item *Record) getValue(index int) TypedValue {
	if val, ok := item.Attributes[index]; ok {
		return val
	}
	return nullVal
}

//IsEqual does a deep comparison of a record to another record.
func (item *Record) IsEqual(otherRecord *Record) bool {
	if item.PrimaryKey != otherRecord.PrimaryKey {
		return false
	}
	for i := 0; i < 80; i++ {
		otherVal := otherRecord.getValue(i)
		val := item.getValue(i)
		if otherVal.IsNull() && val.IsNull() {
			continue
		}
		if otherVal.IsNull() && !val.IsNull() {
			return false
		}
		if !otherVal.IsNull() && val.IsNull() {
			return false
		}
		if otherVal.ValueType() != val.ValueType() {
			return false
		}
		switch val.ValueType() {
		case Type_SingleRelationship:
			valRelationship := val.(SingleRelationship).Record
			otherRelationship := otherVal.(SingleRelationship).Record
			if !valRelationship.IsEqual(&otherRelationship) {
				return false
			}
		case Type_Relationship:
			valRelationship := val.(ToManyRelationship)
			otherRelationship := otherVal.(ToManyRelationship)
			if len(valRelationship.Items) != len(otherRelationship.Items) {
				return false
			}
			for relIndex := 0; relIndex < len(valRelationship.Items); relIndex++ {
				valRelation := valRelationship.Items[relIndex]
				otherRelation := otherRelationship.Items[relIndex]
				if !valRelation.IsEqual(&otherRelation) {
					return false
				}
			}
		default:
			if val != otherVal {
				return false
			}
		}
	}
	return true
}

func (item *Record) UnmarshalJSON(bytes []byte) error {
	panic("UnmarshalJSON not not supported. Use NewRecordFromJSON instead")
}

func (item Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		PrimaryKey string       `json:"primaryKey"`
		RecordType string       `json:"recordType"`
		CRUDStatus CRUDStatus   `json:"CRUDStatus"`
		Attributes []TypedValue `json:"attributes"`
	}{
		PrimaryKey: item.PrimaryKey,
		RecordType: item.RecordType,
		CRUDStatus: item.CRUDStatus,
		Attributes: item.jsonAttributeList(),
	})
}

var nullVal = NullValue{}

func (item Record) jsonAttributeList() []TypedValue {
	attributes := make([]TypedValue, item.Configuration.MaxAttributeIndex()+1)
	length := len(attributes)
	for key, val := range item.Attributes {
		if val.IsNull() && length > key {
			attributes[key] = nullVal
		} else if len(attributes) > key {
			attributes[key] = val
		}
	}
	return attributes
}

type NullValue struct {
}

func (NullValue) ValueType() Type {
	return Type_Text
}

func (NullValue) IsNull() bool {
	return true
}

func (NullValue) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

type SetAttributeError struct {
	index        int
	expectedType Type
	givenType    Type
}

func (e SetAttributeError) Error() string {
	return fmt.Sprintf("Attempting to set attribute of type %s at index %d which is of type %s", e.givenType, e.index, e.expectedType)
}

func newInvalidTypeAttributeError(index int, real, expected Type) error {
	return fmt.Errorf("Requesting invalid type for index %d. Requested %s, real value is %s", index, expected, real)
}
