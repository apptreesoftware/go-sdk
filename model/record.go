package model

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

type RecordSet struct {
	Records       []Record      `json:"records"`
	Configuration Configuration `json:"-"`
}

func NewRecordSet(configuration Configuration) RecordSet {
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
		item := NewItem(&rs.Configuration)
		err := json.Unmarshal(rawRecord, &item)
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
	Configuration Configuration      `json:"-"`
}

func NewItem(configuration *Configuration) Record {
	return Record{Attributes: map[int]TypedValue{}, Configuration: *configuration}
}

func (item Record) SetValue(index int, value TypedValue) error {
	configAttribute := &ConfigurationAttribute{}
	configAttribute, err := item.Configuration.getConfigurationAttribute(index)
	if err != nil {
		return err
	}
	if configAttribute.Type != value.ValueType() {
		return SetAttributeError{givenType: value.ValueType(), expectedType: configAttribute.Type, index: index}
	}
	item.Attributes[index] = value
	return nil
}

func (item Record) GetValue(index int) TypedValue {
	return item.Attributes[index]
}

func (item *Record) UnmarshalJSON(bytes []byte) error {
	var container map[string]interface{}
	json.Unmarshal(bytes, &container)
	item.PrimaryKey = container["primaryKey"].(string)
	item.RecordType = container["recordType"].(string)
	item.Attributes = map[int]TypedValue{}
	var nilValue interface{} = nil
	rawAttributes := container["attributes"].([]interface{})
	for index, attributeData := range rawAttributes {
		configAttribute, err := item.Configuration.getConfigurationAttribute(index)
		if err != nil {
			log.Println(err)
			continue
		}
		var value TypedValue
		var parseErr error
		if attributeData == nilValue {
			continue
		} else {
			switch configAttribute.Type {
			case Relationship:
				var childItems []Record
				var rawChildren = attributeData.([]interface{})
				for _, rawChild := range rawChildren {
					childByte, parseErr := json.Marshal(rawChild)
					if parseErr != nil {
						log.Println(parseErr)
						continue
					}
					var childItem = Record{Configuration: configAttribute.RelatedConfiguration}
					parseErr = json.Unmarshal(childByte, &childItem)
					if parseErr != nil {
						log.Println(parseErr)
					}
					childItems = append(childItems, childItem)
				}
				value = RelationshipValue{Items: childItems}
			case Text:
				value = TextValue{Value: attributeData.(string)}
			case Float:
				var floatValue, parseErr = strconv.ParseFloat(attributeData.(string), 64)
				if parseErr != nil {
					log.Println(parseErr)
					continue
				}
				value = FloatValue{Value: floatValue}
			case Int:
				var intValue, parseErr = strconv.ParseInt(attributeData.(string), 10, 64)
				if parseErr != nil {
					log.Println(parseErr)
					continue
				}
				value = IntValue{Value: intValue}
			case DateTime:
				timeString := attributeData.(string)
				var date, parseErr = time.Parse(`2006-01-02 15:04:05`, timeString)
				if parseErr != nil {
					log.Println(parseErr)
					continue
				}
				value = DateTimeValue{Value: date, HasTime: true}
			case Date:
				var date, parseErr = time.Parse(`"2006-01-02"`, attributeData.(string))
				if parseErr != nil {
					log.Println(parseErr)
					continue
				}
				value = DateTimeValue{Value: date, HasTime: false}
			case ListItem:
				var listItem ListElement
				parseErr = json.Unmarshal([]byte(attributeData.(string)), &listItem)
				if parseErr != nil {
					log.Println(parseErr)
					continue
				}
				value = ListItemValue{ListItem: listItem}
			default:
				continue
			}
		}
		item.Attributes[index] = value
	}
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
	attributes := make([]TypedValue, item.Configuration.MaxAttributeIndex())
	for key, val := range item.Attributes {
		attributes[key] = val
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
