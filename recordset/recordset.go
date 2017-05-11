package recordset

import (
	"fmt"
	"github.com/apptreesoftware/go-sdk/model"
	"reflect"
	"strconv"
	"strings"
)

func GetConfiguration(v interface{}) (*model.Configuration, error) {
	config := model.Configuration{}
	var attributes []model.ConfigurationAttribute
	t := reflect.TypeOf(v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("apptree")

		fmt.Printf("%d. %v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
		attr, err := getConfigurationAttributeFromTag(tag)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, attr)
	}
	config.Attributes = attributes
	return &config, nil
}

func getConfigurationAttributeFromTag(tag string) (model.ConfigurationAttribute, error) {
	attribute := model.ConfigurationAttribute{}
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
				return attribute, err
			}
			attribute.Index = intValue
		case "type":
			var t model.Type
			switch value {
			case "ListItem":
				t = model.ListItem
			case "Text":
				t = model.Text
			case "DateTime":
				t = model.DateTime
			case "Date":
				t = model.Date
			case "Relationship":
				t = model.Relationship
			case "Int":
				t = model.Int
			case "Float":
				t = model.Float
			default:
				t = model.Text
			}
			attribute.Type = t
		case "name":
			attribute.Name = value
		}
	}
	return attribute, nil
}
