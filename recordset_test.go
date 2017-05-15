package apptree

import (
	"encoding/json"
	"github.com/apptreesoftware/go-sdk/example"
	"testing"
)

func TestMarshal(t *testing.T) {
	model := example.Issue{Id: "1234"}
	configuration, err := GetConfiguration(model)
	b, err := json.Marshal(&configuration)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func TestSetters(t *testing.T) {
	config := Configuration{}
	attributes := []ConfigurationAttribute{}
	attr0 := ConfigurationAttribute{
		Name:  "TextValue",
		Type:  Text,
		Index: 0,
	}
	attr1 := ConfigurationAttribute{
		Name:  "FloatValue",
		Type:  Float,
		Index: 1,
	}
	attr2 := ConfigurationAttribute{
		Name:  "IntValue",
		Type:  Int,
		Index: 2,
	}
	attr3 := ConfigurationAttribute{
		Name:  "ListItem",
		Type:  ListItem,
		Index: 3,
	}

	attributes = append(attributes, attr0, attr1, attr2, attr3)

	config.Attributes = attributes

	record := NewItem(&config)
	record.PrimaryKey = "1234"
	record.SetValue(0, TextValue{Value: "Test Text"})
	record.SetValue(1, FloatValue{Value: 1.0})
	record.SetValue(2, IntValue{Value: 2})
	record.SetValue(3, ListItemValue{ListItem: NewListItem("Test")})

	if record.PrimaryKey != "1234" {
		t.Fatalf("Primary key should be 1234")
	}
	val := record.GetValue(0)
	if val.(TextValue).Value != "Test Text" {
		t.Fatalf("Attribute 0 should be %s", `Test Text`)
	}
	val = record.GetValue(1)
	if val.(FloatValue).Value != 1.0 {
		t.Fatal("Attribute 1 should be 1.0")
	}
	val = record.GetValue(2)
	if val.(IntValue).Value != 2 {
		t.Fatal("Attribute 2 should be 2")
	}
	val = record.GetValue(3)
	if val.(ListItemValue).ListItem.Value != "Test" {
		t.Fatal("Attribute 3 should be 2")
	}
}
