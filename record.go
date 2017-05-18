package apptree

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type RecordSet struct {
	Records       []Record       `json:"records"`
	Configuration *Configuration `json:"-"`
}

type AttributeErrorType string

const (
	NullValue   AttributeErrorType = "NullValue"
	InvalidType AttributeErrorType = "InvalidType"
)

type AttributeError struct {
	Type    AttributeErrorType
	Message string
	Index   int
}

func newNullAttributeError(index int) AttributeError {
	return AttributeError{
		Index: index,
		Type:  NullValue,
	}
}

func newInvalidTypeAttributeError(index int, valType Type, expectedType Type) AttributeError {
	return AttributeError{
		Index:   index,
		Message: fmt.Sprintf("Requested type is %s but the attribute is a %s", expectedType, valType),
	}
}

func (e AttributeError) Error() string {
	switch e.Type {
	case NullValue:
		return fmt.Sprintf("Value at index %d is nil", e.Index)
	case InvalidType:
		return e.Message
	}
	return ""
}

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
	return nil
}

type recordSetUnmarshalHelper struct {
	Records []json.RawMessage
}

type Record struct {
	PrimaryKey    string             `json:"primaryKey"`
	RecordType    string             `json:"recordType"`
	Attributes    map[int]TypedValue `json:"attributes"`
	Configuration *Configuration     `json:"-"`
}

func NewRecordFromJSON(bytes []byte, configuration *Configuration) (Record, error) {
	record := Record{Attributes: map[int]TypedValue{}, Configuration: configuration}
	err := json.Unmarshal(bytes, &record)
	return record, err
}

func NewItem(configuration *Configuration) Record {
	return Record{Attributes: map[int]TypedValue{}, Configuration: configuration}
}

func (item *Record) SetValue(index int, value TypedValue) error {
	configAttribute := &ConfigurationAttribute{}
	configAttribute, err := item.Configuration.getConfigurationAttribute(index)
	if err != nil {
		return err
	}
	if value == nil {
		delete(item.Attributes, index)
		return nil
	}
	if configAttribute.Type != value.ValueType() {
		return SetAttributeError{givenType: value.ValueType(), expectedType: configAttribute.Type, index: index}
	}
	item.Attributes[index] = value
	return nil
}

func (item *Record) GetTextValue(index int) (TextValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return TextValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(TextValue); ok {
		return transformedVal, nil
	} else {
		return TextValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Text)
	}
}

func (item *Record) GetIntValue(index int) (IntValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return IntValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(IntValue); ok {
		return transformedVal, nil
	} else {
		return IntValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Int)
	}
}

func (item *Record) GetFloatValue(index int) (FloatValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return FloatValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(FloatValue); ok {
		return transformedVal, nil
	} else {
		return FloatValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Float)
	}
}

func (item *Record) GetTimeIntervalValue(index int) (TimeIntervalValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return TimeIntervalValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(TimeIntervalValue); ok {
		return transformedVal, nil
	} else {
		return TimeIntervalValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_TimeInterval)
	}
}

func (item *Record) GetBoolValue(index int) (BooleanValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return BooleanValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(BooleanValue); ok {
		return transformedVal, nil
	} else {
		return BooleanValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Boolean)
	}
}

func (item *Record) GetListItemValue(index int) (ListItemValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return ListItemValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(ListItemValue); ok {
		return transformedVal, nil
	} else {
		return ListItemValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_ListItem)
	}
}

func (item *Record) GetRelationshipValue(index int) (RelationshipValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return RelationshipValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(RelationshipValue); ok {
		return transformedVal, nil
	} else {
		return RelationshipValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Relationship)
	}
}

func (item *Record) GetDateValue(index int) (DateTimeValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return DateTimeValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(DateTimeValue); ok && !transformedVal.HasTime {
		return transformedVal, nil
	} else {
		return DateTimeValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Date)
	}
}

func (item *Record) GetDateTimeValue(index int) (DateTimeValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return DateTimeValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(DateTimeValue); ok && transformedVal.HasTime {
		return transformedVal, nil
	} else {
		return DateTimeValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_DateTime)
	}
}

func (item *Record) GetDateRangeValue(index int) (DateRangeValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return DateRangeValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(DateRangeValue); ok {
		return transformedVal, nil
	} else {
		return DateRangeValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_DateRange)
	}
}

func (item *Record) GetDateTimeRangeValue(index int) (DateTimeRangeValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return DateTimeRangeValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(DateTimeRangeValue); ok {
		return transformedVal, nil
	} else {
		return DateTimeRangeValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_DateTimeRange)
	}
}

func (item *Record) GetLocationValue(index int) (LocationValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return LocationValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(LocationValue); ok {
		return transformedVal, nil
	} else {
		return LocationValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Location)
	}
}

func (item *Record) GetColorValue(index int) (ColorValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return ColorValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(ColorValue); ok {
		return transformedVal, nil
	} else {
		return ColorValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Color)
	}
}

func (item *Record) GetSingleRelationshipValue(index int) (SingleRelationshipValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return SingleRelationshipValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(SingleRelationshipValue); ok {
		return transformedVal, nil
	} else {
		return SingleRelationshipValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_SingleRelationship)
	}
}

func (item *Record) GetImageValue(index int) (ImageValue, error) {
	val := item.GetValue(index)
	if val == nil {
		return ImageValue{}, newNullAttributeError(index)
	}
	if transformedVal, ok := val.(ImageValue); ok {
		return transformedVal, nil
	} else {
		return ImageValue{}, newInvalidTypeAttributeError(index, val.ValueType(), Type_Image)
	}
}

func (item *Record) GetValue(index int) TypedValue {
	return item.Attributes[index]
}

func (item *Record) UnmarshalJSON(bytes []byte) error {
	var container map[string]interface{}
	err := json.Unmarshal(bytes, &container)
	if err != nil {
		return err
	}
	err = item.unmarshalMap(container)
	if err != nil {
		return err
	}
	return nil
}

func NewRecordFromMap(container map[string]interface{}, configuration *Configuration) Record {
	record := Record{Attributes: map[int]TypedValue{}, Configuration: configuration}
	record.unmarshalMap(container)
	return record
}

func (item *Record) unmarshalMap(container map[string]interface{}) error {
	item.PrimaryKey = container["primaryKey"].(string)
	item.RecordType = container["recordType"].(string)
	attributes := make(map[int]TypedValue, item.Configuration.MaxAttributeIndex())
	var nilValue interface{} = nil
	rawAttributes := container["attributes"].([]interface{})
	for index, attributeData := range rawAttributes {
		configAttribute, err := item.Configuration.getConfigurationAttribute(index)
		if err != nil {
			log.Println(err)
			continue
		}
		//log.Printf("Parsing value: %v into %s (idx: %d) of type %s", attributeData, configAttribute.Name, configAttribute.Index, configAttribute.Type)
		var value TypedValue
		var parseErr error
		if attributeData == nilValue {
			continue
		} else {
			switch configAttribute.Type {
			case Type_Relationship:
				var childItems []Record
				var rawChildren = attributeData.([]interface{})
				for _, rawChild := range rawChildren {
					childByte, parseErr := json.Marshal(rawChild)
					if parseErr != nil {
						return parseErr
					}
					childItem, parseErr := NewRecordFromJSON(childByte, configAttribute.RelatedConfiguration)
					if parseErr != nil {
						return parseErr
					}
					childItems = append(childItems, childItem)
				}
				value = RelationshipValue{Items: childItems}
			case Type_Text:
				value = TextValue{Value: attributeData.(string)}
			case Type_Float:
				var floatValue, parseErr = strconv.ParseFloat(attributeData.(string), 64)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = FloatValue{Value: floatValue}
			case Type_Int:
				var intValue, parseErr = strconv.ParseInt(attributeData.(string), 10, 64)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = IntValue{Value: intValue}
			case Type_DateTime:
				timeString := attributeData.(string)
				var date, parseErr = time.Parse(`2006-01-02 15:04:05`, timeString)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = DateTimeValue{Value: date, HasTime: true}
			case Type_Date:
				var date, parseErr = time.Parse(`"2006-01-02"`, attributeData.(string))
				if parseErr != nil {
					return parseErr
					continue
				}
				value = DateTimeValue{Value: date, HasTime: false}
			case Type_ListItem:
				var listItem ListElement
				parseErr = json.Unmarshal([]byte(attributeData.(string)), &listItem)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = ListItemValue{ListItem: listItem}
			case Type_TimeInterval:
				var longValue, parseErr = strconv.ParseInt(attributeData.(string), 10, 64)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = TimeIntervalValue{Value: longValue}
			case Type_Boolean:
				stringVal := attributeData.(string)
				value = BooleanValue{Value: stringVal == "Y"}
			case Type_DateTimeRange:
				var dateTimeRange DateTimeRange
				var rawChild = attributeData.(map[string]interface{})
				childByte, parseErr := json.Marshal(rawChild)
				parseErr = json.Unmarshal(childByte, &dateTimeRange)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = DateTimeRangeValue{Value: dateTimeRange}
			case Type_DateRange:
				var dateRange DateRange
				var rawChild = attributeData.(map[string]interface{})
				childByte, parseErr := json.Marshal(rawChild)
				parseErr = json.Unmarshal(childByte, &dateRange)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = DateRangeValue{Value: dateRange}
			case Type_Location:
				var location Location
				var raw = attributeData.(map[string]interface{})
				childByte, parseErr := json.Marshal(raw)
				parseErr = json.Unmarshal(childByte, &location)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = LocationValue{Value: location}
			case Type_SingleRelationship:
				var dataSetItem Record
				rawItem := attributeData.(map[string]interface{})
				childByte, parseErr := json.Marshal(rawItem)
				parseErr = json.Unmarshal(childByte, &dataSetItem)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = SingleRelationshipValue{Value: dataSetItem}
			case Type_Color:
				var color Color
				var raw = attributeData.(map[string]interface{})
				childByte, parseErr := json.Marshal(raw)
				parseErr = json.Unmarshal(childByte, &color)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = ColorValue{Value: color}
			case Type_Image:
				var image Image
				var raw = attributeData.(map[string]interface{})
				childByte, parseErr := json.Marshal(raw)
				parseErr = json.Unmarshal(childByte, &image)
				if parseErr != nil {
					return parseErr
					continue
				}
				value = ImageValue{Value: image}
			default:
				continue
			}
		}
		attributes[configAttribute.Index] = value
	}
	item.Attributes = attributes
	return nil
}

func (item Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		PrimaryKey string       `json:"primaryKey"`
		RecordType string       `json:"recordType"`
		Attributes []TypedValue `json:"attributes"`
	}{
		PrimaryKey: item.PrimaryKey,
		RecordType: item.RecordType,
		Attributes: item.attributeList(),
	})
}

func (item Record) attributeList() []TypedValue {
	attributes := make([]TypedValue, item.Configuration.MaxAttributeIndex()+1)
	for key, val := range item.Attributes {
		if len(attributes) > key {
			attributes[key] = val
		}
	}
	return attributes
}

type SetAttributeError struct {
	index        int
	expectedType Type
	givenType    Type
}

func (e SetAttributeError) Error() string {
	return fmt.Sprintf("Attempting to set attribute of type %s at index %d which is of type %s", e.givenType, e.index, e.expectedType)
}
