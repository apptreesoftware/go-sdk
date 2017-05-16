package apptree

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetConfiguration(v interface{}) (*Configuration, error) {
	config := Configuration{}
	var attributes []ConfigurationAttribute
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("apptree")
		attr := ConfigurationAttribute{}
		attr.Name = field.Name
		attr.Index = i
		attr.Type = inferDataTypeFromField(field)
		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		err := enhanceAttributeFromTag(&attr, tag)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)
	}
	config.Attributes = attributes
	return &config, nil
}

func inferDataTypeFromField(field reflect.StructField) Type {
	switch field.Type.Name() {
	case "string":
		return Type_Text
	case "Time":
		return Type_DateTime
	}
	return Type_Text
}

func enhanceAttributeFromTag(attribute *ConfigurationAttribute, tag string) error {
	infoArray := strings.Split(tag, ";")
	for _, info := range infoArray {
		components := strings.Split(info, "=")
		if len(components) != 2 {
			continue
		}
		key := components[0]
		value := components[1]
		switch strings.ToLower(key) {
		case "index":
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			attribute.Index = intValue
		case "type":
			var t Type
			switch value {
			case "ListItem":
				t = Type_ListItem
			case "Text":
				t = Type_Text
			case "DateTime":
				t = Type_DateTime
			case "Date":
				t = Type_Date
			case "Relationship":
				t = Type_Relationship
			case "Int":
				t = Type_Int
			case "Float":
				t = Type_Float
			case "TimeInterval":
				t = Type_TimeInterval
			case "Boolean":
				t = Type_Boolean
			case "DateRange":
				t = Type_DateRange
			case "DateTimeRange":
				t = Type_DateTimeRange
			case "Image":
				t = Type_Image
			case "Location":
				t = Type_Location
			case "SingleRelationship":
				t = Type_SingleRelationship
			case "Color":
				t = Type_Color
			default:
				t = Type_Text
			}
			attribute.Type = t
		case "name":
			attribute.Name = value
		}
	}
	return nil
}
