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
		return Text
	case "Time":
		return DateTime
	}
	return Text
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
				t = ListItem
			case "Text":
				t = Text
			case "DateTime":
				t = DateTime
			case "Date":
				t = Date
			case "Relationship":
				t = Relationship
			case "Int":
				t = Int
			case "Float":
				t = Float
			default:
				t = Text
			}
			attribute.Type = t
		case "name":
			attribute.Name = value
		}
	}
	return nil
}
