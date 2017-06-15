package apptree

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

//TODO: Write ListItem json marshaling/unmarshaling testss

func TestSetters(t *testing.T) {
	config := Configuration{}
	attributes := []ConfigurationAttribute{}
	attr0 := ConfigurationAttribute{
		Name:  "TextValue",
		Type:  Type_Text,
		Index: 0,
	}
	attr1 := ConfigurationAttribute{
		Name:  "FloatValue",
		Type:  Type_Float,
		Index: 1,
	}
	attr2 := ConfigurationAttribute{
		Name:  "IntValue",
		Type:  Type_Int,
		Index: 2,
	}
	attr3 := ConfigurationAttribute{
		Name:  "ListItem",
		Type:  Type_ListItem,
		Index: 3,
	}
	attr4 := ConfigurationAttribute{
		Name:  "DateRange",
		Type:  Type_DateRange,
		Index: 4,
	}
	attr5 := ConfigurationAttribute{
		Name:  "TimeInterval",
		Type:  Type_TimeInterval,
		Index: 5,
	}
	attr6 := ConfigurationAttribute{
		Name:  "DateTimeRange",
		Type:  Type_DateTimeRange,
		Index: 6,
	}
	attr7 := ConfigurationAttribute{
		Name:  "Image",
		Type:  Type_Image,
		Index: 7,
	}
	attr8 := ConfigurationAttribute{
		Name:  "Location",
		Type:  Type_Location,
		Index: 8,
	}
	attr9 := ConfigurationAttribute{
		Name:  "Color",
		Type:  Type_Color,
		Index: 9,
	}

	attributes = append(attributes, attr0, attr1, attr2, attr3, attr4, attr5, attr6, attr7, attr8, attr9)

	config.Attributes = attributes

	record := NewItem(&config)
	record.PrimaryKey = "1234"
	record.SetValue(0, NewTextValue("Test Text"))
	record.SetValue(1, FloatValue{Value: 1.0})
	record.SetValue(2, IntValue{Value: 2})
	record.SetValue(3, ListItemValue{ListItem: NewListItem("Test")})
	record.SetValue(4, DateRangeValue{Value: NewDateRange(time.Now(), time.Now().AddDate(0, 1, 2))})
	record.SetValue(5, TimeIntervalValue{Value: 54})
	record.SetValue(6, DateTimeRangeValue{Value: NewDateTimeRange(time.Now(), time.Now().AddDate(0, 1, 2))})
	record.SetValue(7, ImageValue{Value: NewImage("http://fakeImage.com", "someUploadKey", nil)})
	record.SetValue(8, LocationValue{Value: NewLocation(14.5677, 176.245356, 14, 13.456, 1, 234.78, NewNullDateTime(time.Now(), true))})
	record.SetValue(9, ColorValue{Value: NewColor(76, 175, 80, 0)})

	if record.PrimaryKey != "1234" {
		t.Fatalf("Primary key should be 1234")
	}
	val, _ := record.GetTextValue(0)

	if val.Value != "Test Text" {
		t.Fatalf("Attribute 0 should be %s", `Test Text`)
	}
	fVal, _ := record.GetFloatValue(1)
	if fVal.Value != 1.0 {
		t.Fatal("Attribute 1 should be 1.0")
	}
	iVal, _ := record.GetIntValue(2)
	if iVal.Value != 2 {
		t.Fatal("Attribute 2 should be 2")
	}
	lVal, _ := record.GetListItemValue(3)
	if lVal.ListItem.Value != "Test" {
		t.Fatal("Attribute 3 should be 2")
	}
	drVal, _ := record.GetDateRangeValue(4)
	if drVal.Value.ToDate.Year() != 2017 {
		t.Fatal(fmt.Printf("Date attribute has incorrect date: %d", drVal.Value.ToDate.Year()))
	}
	tiVal, _ := record.GetTimeIntervalValue(5)
	if tiVal.Value != 54 {
		t.Fatal("Attribute 5 should be 54")
	}
	dtrVal, _ := record.GetDateTimeRangeValue(6)
	if dtrVal.Value.ToDate.Year() != 2017 {
		t.Fatal(fmt.Printf("Date time attribute has incorrect date: %d", dtrVal.Value.ToDate.Year()))
	}
	imageVal, _ := record.GetImageValue(7)
	if imageVal.Value.ImageURL != "http://fakeImage.com" || imageVal.Value.UploadKey != "someUploadKey" {
		t.Fatal("Attribute 7 Image values set incorrectly")
	}
	locVal, _ := record.GetLocationValue(8)
	location := locVal.Value
	if location.Timestamp.Date.Year() != 2017 {
		t.Fatal("Location timestamp set incorrectly")
	}
	if location.Elevation <= 0 || location.Latitude <= 0 || location.Bearing <= 0 || location.Accuracy <= 0 || location.Speed <= 0 {
		t.Fatal("Location values set incorrectly")
	}
	cVal, _ := record.GetColorValue(9)
	color := cVal.Value
	if color.Red != 76 || color.Green != 175 || color.Blue != 80 || color.Alpha != 0 {
		t.Fatal("Color values set incorrectly")
	}
}

func TestRelationship(t *testing.T) {
	var configuration Configuration
	err := json.Unmarshal([]byte(Config1), &configuration)
	if err != nil {
		t.Error(err)
	}

	rec := NewItem(&configuration)
	child, err := rec.NewChildAtIndex(19)
	if err != nil {
		t.Error(err)
	}
	child.PrimaryKey = "999"
	child.SetValue(1, TextValue{Value: "Normal"})
	testChildVal, err := rec.GetRelationshipValue(19)
	if err != nil {
		t.Error(err)
	}
	testChild := testChildVal.Items[0]
	txtVal, _ := testChild.GetTextValue(1)

	if txtVal.Value != "Normal" {
		t.Fatal("New child attribute 1 should be normal")
	}

}

func TestParseRecordSet(t *testing.T) {
	var configuration Configuration
	err := json.Unmarshal([]byte(Config1), &configuration)
	if err != nil {
		t.Error(err)
	}
	if len(configuration.Attributes) != 28 {
		t.Fatalf("Invalid # of attributes. Expected 28, got %d", len(configuration.Attributes))
	}
	recordSet := NewRecordSet(&configuration)
	err = json.Unmarshal([]byte(Config1RecordSet), &recordSet)
	if err != nil {
		t.Error(err)
	}
	if len(recordSet.Records) != 5 {
		t.Fatalf("Expected 5 records, record set size was %d", len(recordSet.Records))
	}
	record := recordSet.Records[0]
	if record.PrimaryKey != "68839407" {
		t.Fatal("Expecting primary key 68839407")
	}
	attr := record.GetValue(2)
	if attr.ValueType() != Type_Text {
		t.Fatalf("Invalid value type %s", attr.ValueType())
	}
	if attr.(TextValue).Value != "88024417" {
		t.Fatal("Expecting text value of `88024417`")
	}

	attr = record.GetValue(19)
	if attr.ValueType() != Type_Relationship {
		t.Fatal("Expected a relationship type at index 19")
	}
	if len(attr.(RelationshipValue).Items) != 2 {
		t.Fatal("Expected 2 sub items at index 19")
	}
	relationship := attr.(RelationshipValue).Items[0]
	if relationship.PrimaryKey != "231342466" {
		t.Fatal("Expecting sub item primary key of 231342466")
	}
	attr = relationship.GetValue(1)
	if attr.ValueType() != Type_Text {
		t.Fatal("Expecting text for subItem 0-1")
	}

	attr, err = record.GetDateTimeValue(1)
	compareDate, err := time.Parse("2006-01-02 15:04:05", "2017-05-10 12:31:36")
	if err != nil {
		t.Error(err)
	}
	if attr.(DateTimeValue).Value != compareDate {
		t.Fatalf("Expecting 0.3 to be a date of %v", compareDate)
	}
}

func TestMarshalUnmarshalRecord(t *testing.T) {
	var configuration Configuration
	err := json.Unmarshal([]byte(Config1), &configuration)
	if err != nil {
		t.Error(err)
	}
	recordSet := NewRecordSet(&configuration)
	err = json.Unmarshal([]byte(Config1RecordSet), &recordSet)
	record := recordSet.Records[0]
	b, err := json.Marshal(&record)
	if err != nil {
		t.Error(err)
	}
	unmarshaledRecord, err := NewRecordFromJSON(b, &configuration)
	if err != nil {
		t.Error(err)
	}
	if unmarshaledRecord.PrimaryKey != "68839407" {
		t.Fatalf("Unmarshaled record has incorrect primary key %s", unmarshaledRecord.PrimaryKey)
	}
	attr := unmarshaledRecord.GetValue(19)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 19")
	}
	if attr.ValueType() != Type_Relationship {
		t.Fatalf("Expected relationship for attribute 19")
	}
	childItem := attr.(RelationshipValue).Items[0]
	if childItem.GetValue(1).ValueType() != Type_Text {
		t.Fatal("Child Index 1 has wrong type")
	}
	if childItem.GetValue(1).(TextValue).Value != "Approved" {
		t.Fatal("Child index 1 != Approved")
	}
}
