package apptree

import (
	"encoding/json"
	"testing"
	"time"
)

//TODO: Write ListItem json marshaling/unmarshaling testss

var testDateTime, _ = time.Parse(dateTimeFormat, "2017-07-16 05:38:00")
var testDateTime2, _ = time.Parse(dateTimeFormat, "2017-07-17 06:38:00")
var testDate, _ = time.Parse(dateFormat, "2017-07-16")
var testDate2, _ = time.Parse(dateFormat, "2017-07-18")
var testLocation = NewLocation(122.111, -34.13, 10.0, 10.0, 15.0, 100, NewDateTime(testDateTime))

func sampleConfig() Configuration {
	config := Configuration{}
	attributes := []ConfigurationAttribute{}
	attributes = append(attributes, ConfigurationAttribute{Name: "TextValue", Type: Type_Text, Index: 0})
	attributes = append(attributes, ConfigurationAttribute{Name: "FloatValue", Type: Type_Float, Index: 1})
	attributes = append(attributes, ConfigurationAttribute{Name: "IntValue", Type: Type_Int, Index: 2})
	attributes = append(attributes, ConfigurationAttribute{Name: "BoolValue", Type: Type_Boolean, Index: 3})
	attributes = append(attributes, ConfigurationAttribute{Name: "ColorValue", Type: Type_Color, Index: 4})
	attributes = append(attributes, ConfigurationAttribute{Name: "ListItem", Type: Type_ListItem, Index: 5})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateItem", Type: Type_Date, Index: 6})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateTimeItem", Type: Type_DateTime, Index: 7})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateRange", Type: Type_DateRange, Index: 8})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateTimeRange", Type: Type_DateTimeRange, Index: 9})
	attributes = append(attributes, ConfigurationAttribute{Name: "TimeInterval", Type: Type_TimeInterval, Index: 10})
	attributes = append(attributes, ConfigurationAttribute{Name: "ImageValue", Type: Type_Image, Index: 11})
	attributes = append(attributes, ConfigurationAttribute{Name: "LocationValue", Type: Type_Location, Index: 12})

	//Null Checks

	attributes = append(attributes, ConfigurationAttribute{Name: "TextValue", Type: Type_Text, Index: 13})
	attributes = append(attributes, ConfigurationAttribute{Name: "FloatValue", Type: Type_Float, Index: 14})
	attributes = append(attributes, ConfigurationAttribute{Name: "IntValue", Type: Type_Int, Index: 15})
	attributes = append(attributes, ConfigurationAttribute{Name: "BoolValue", Type: Type_Boolean, Index: 16})
	attributes = append(attributes, ConfigurationAttribute{Name: "ColorValue", Type: Type_Color, Index: 17})
	attributes = append(attributes, ConfigurationAttribute{Name: "ListItem", Type: Type_ListItem, Index: 18})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateItem", Type: Type_Date, Index: 19})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateTimeItem", Type: Type_DateTime, Index: 20})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateRange", Type: Type_DateRange, Index: 21})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateTimeRange", Type: Type_DateTimeRange, Index: 22})
	attributes = append(attributes, ConfigurationAttribute{Name: "TimeInterval", Type: Type_TimeInterval, Index: 23})
	attributes = append(attributes, ConfigurationAttribute{Name: "ImageValue", Type: Type_Image, Index: 24})
	attributes = append(attributes, ConfigurationAttribute{Name: "LocationValue", Type: Type_Location, Index: 25})

	//Relationships
	childConfig := sampleChildConfig()
	attributes = append(attributes, ConfigurationAttribute{Name: "ToManyRelationship", Type: Type_Relationship, Index: 26, RelatedConfiguration: &childConfig})
	attributes = append(attributes, ConfigurationAttribute{Name: "SingleRelationship", Type: Type_SingleRelationship, Index: 27, RelatedConfiguration: &childConfig})

	config.Attributes = attributes
	return config
}

func sampleChildConfig() Configuration {
	config := Configuration{}
	attributes := []ConfigurationAttribute{}
	attributes = append(attributes, ConfigurationAttribute{Name: "TextValue", Type: Type_Text, Index: 0})
	attributes = append(attributes, ConfigurationAttribute{Name: "FloatValue", Type: Type_Float, Index: 1})
	attributes = append(attributes, ConfigurationAttribute{Name: "IntValue", Type: Type_Int, Index: 2})
	attributes = append(attributes, ConfigurationAttribute{Name: "BoolValue", Type: Type_Boolean, Index: 3})
	attributes = append(attributes, ConfigurationAttribute{Name: "ColorValue", Type: Type_Color, Index: 4})
	attributes = append(attributes, ConfigurationAttribute{Name: "ListItem", Type: Type_ListItem, Index: 5})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateItem", Type: Type_Date, Index: 6})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateTimeItem", Type: Type_DateTime, Index: 7})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateRange", Type: Type_DateRange, Index: 8})
	attributes = append(attributes, ConfigurationAttribute{Name: "DateTimeRange", Type: Type_DateTimeRange, Index: 9})
	attributes = append(attributes, ConfigurationAttribute{Name: "TimeInterval", Type: Type_TimeInterval, Index: 10})
	attributes = append(attributes, ConfigurationAttribute{Name: "ImageValue", Type: Type_Image, Index: 11})
	attributes = append(attributes, ConfigurationAttribute{Name: "LocationValue", Type: Type_Location, Index: 12})
	config.Attributes = attributes
	return config
}

func sampleRecord(t *testing.T) *Record {
	config := sampleConfig()
	record := NewItem(&config)
	record.PrimaryKey = "1234"
	record.SetValue(0, NewString("Test Text"))
	record.SetValue(1, FloatFrom(1.0))
	record.SetValue(2, NewInt(2))
	record.SetValue(3, NewBool(true))
	record.SetValue(4, NewColor(76, 175, 80, 10))
	record.SetValue(5, NewListItem("Test Item"))
	record.SetValue(6, NewDate(testDate))
	record.SetValue(7, NewDateTime(testDateTime))
	record.SetValue(8, NewDateRange(testDate, testDate2))
	record.SetValue(9, NewDateTimeRange(testDateTime, testDateTime2))
	record.SetValue(10, NewTimeInterval(100))
	record.SetValue(11, NewImage("http://fakeImage.com"))
	record.SetValue(12, testLocation)

	//Skip 13 - 25 for null checks

	//Relationships
	childRec, err := record.AddToManyChildAtIndex(26)
	if err != nil {
		panic(err)
	}
	childRec.PrimaryKey = "2345-1"
	childRec.SetValue(0, NewString("Test Child Text"))
	childRec.SetValue(1, FloatFrom(1.0))
	childRec.SetValue(2, NewInt(2))
	childRec.SetValue(3, NewBool(true))
	childRec.SetValue(4, NewColor(76, 175, 80, 10))
	childRec.SetValue(5, NewListItem("Test Item"))
	childRec.SetValue(6, NewDate(testDate))
	childRec.SetValue(7, NewDateTime(testDateTime))
	childRec.SetValue(8, NewDateRange(testDate, testDate2))
	childRec.SetValue(9, NewDateTimeRange(testDateTime, testDateTime2))
	childRec.SetValue(10, NewTimeInterval(100))
	childRec.SetValue(11, NewImage("http://fakeImage.com"))
	childRec.SetValue(12, testLocation)

	childRec2, err := record.AddToManyChildAtIndex(26)
	if err != nil {
		panic(err)
	}
	childRec2.PrimaryKey = "2345-2"
	childRec2.SetValue(0, NewString("Test Child Text 2"))
	childRec2.SetValue(1, FloatFrom(1.0))
	childRec2.SetValue(2, NewInt(2))
	childRec2.SetValue(3, NewBool(true))
	childRec2.SetValue(4, NewColor(76, 175, 80, 10))
	childRec2.SetValue(5, NewListItem("Test Item"))
	childRec2.SetValue(6, NewDate(testDate))
	childRec2.SetValue(7, NewDateTime(testDateTime))
	childRec2.SetValue(8, NewDateRange(testDate, testDate2))
	childRec2.SetValue(9, NewDateTimeRange(testDateTime, testDateTime2))
	childRec2.SetValue(10, NewTimeInterval(100))
	childRec2.SetValue(11, NewImage("http://fakeImage.com"))
	childRec2.SetValue(12, testLocation)

	singleChild, err := record.NewToOneRelationshipAtIndex(27)
	if err != nil {
		panic(err)
	}
	singleChild.PrimaryKey = "2345-2"
	singleChild.SetValue(0, NewString("Single Child Test"))
	singleChild.SetValue(1, FloatFrom(1.0))
	singleChild.SetValue(2, NewInt(2))
	singleChild.SetValue(3, NewBool(true))
	singleChild.SetValue(4, NewColor(76, 175, 80, 10))
	singleChild.SetValue(5, NewListItem("Test Item"))
	singleChild.SetValue(6, NewDate(testDate))
	singleChild.SetValue(7, NewDateTime(testDateTime))
	singleChild.SetValue(8, NewDateRange(testDate, testDate2))
	singleChild.SetValue(9, NewDateTimeRange(testDateTime, testDateTime2))
	singleChild.SetValue(10, NewTimeInterval(100))
	singleChild.SetValue(11, NewImage("http://fakeImage.com"))
	singleChild.SetValue(12, testLocation)

	return &record
}

func TestMarshalUnmarshalConfig(t *testing.T) {
	config := sampleConfig()
	b, e := json.Marshal(&config)
	if e != nil {
		t.Error(e)
	}
	marshal1 := string(b)
	var config2 Configuration
	e = json.Unmarshal(b, &config2)
	if e != nil {
		t.Error(e)
	}
	b2, e := json.Marshal(&config2)
	marshal2 := string(b2)

	if marshal1 != marshal2 {
		t.Error("Configs do not match")
	}
}

func TestMarshalUnmarshalRecord(t *testing.T) {
	record := sampleRecord(t)
	b, err := json.Marshal(record)
	if err != nil {
		t.Error(err)
	}
	config := sampleConfig()
	t.Logf("Marshaled record %s", string(b))
	record2, err := NewRecordFromJSON(b, &config)
	if err != nil {
		t.Error(err)
	}
	checkRecordValues(record2, t)

	b2, err := json.Marshal(&record2)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Marshaled record %s", string(b2))
	record3, err := NewRecordFromJSON(b2, &config)
	if err != nil {
		t.Error(err)
	}
	checkRecordValues(record3, t)

	json1 := string(b)
	json2 := string(b2)

	if json1 != json2 {
		t.Error("JSON output does not match after unmarshaling and re-marshaling")
	}
}

func checkRecordValues(record Record, t *testing.T) {
	if record.PrimaryKey != "1234" {
		t.Fatalf("Primary key should be 1234")
	}
	val, _ := record.GetText(0)

	if val.String != "Test Text" {
		t.Fatalf("Attribute 0 should be %s", `Test Text`)
	}
	fVal, _ := record.GetFloat(1)
	if fVal.Float64 != 1.0 {
		t.Fatal("Attribute 1 should be 1.0")
	}
	iVal, _ := record.GetInt(2)
	if iVal.Int64 != 2 {
		t.Fatal("Attribute 2 should be 2")
	}
	bVal, _ := record.GetBool(3)
	if bVal.Bool != true {
		t.Errorf("Bool values set incorrectly")
	}
	color, _ := record.GetColor(4)
	if color.Red != 76 || color.Green != 175 || color.Blue != 80 || color.Alpha != 10 {
		t.Fatal("Color values set incorrectly")
	}
	lVal, _ := record.GetListItem(5)
	if lVal.Value != "Test Item" {
		t.Fatal("Attribute 5 should be Test Item")
	}
	date, _ := record.GetDate(6)
	if date.Time != testDate {
		t.Fatal("Date does not match")
	}
	dateTime, _ := record.GetDateTime(7)
	if dateTime.Time != testDateTime {
		t.Fatal("DateTime does not match")
	}
	dateRange, _ := record.GetDateRange(8)
	if dateRange.FromDate != testDate || dateRange.ToDate != testDate2 {
		t.Fatalf("Date range does not match %v - %v != %v - %v", dateRange.FromDate, dateRange.ToDate, testDate, testDate2)
	}

	dateTimeRange, _ := record.GetDateTimeRange(9)
	if dateTimeRange.FromDate != testDateTime || dateTimeRange.ToDate != testDateTime2 {
		t.Fatalf("Date range does not match %v - %v != %v - %v", dateTimeRange.FromDate, dateTimeRange.ToDate, testDateTime, testDateTime2)
	}

	timeInterval, _ := record.GetTimeInterval(10)
	if !timeInterval.Valid || timeInterval.Int64 != 100 {
		t.Fatal("Time intervals do not match")
	}

	image, _ := record.GetImage(11)
	if image.ImageURL != "http://fakeImage.com" {
		t.Fatal("Attribute 7 Image values set incorrectly")
	}

	loc, _ := record.GetLocation(12)
	if loc.Latitude != testLocation.Latitude ||
		loc.Longitude != testLocation.Longitude ||
		loc.Accuracy != testLocation.Accuracy ||
		loc.Bearing != testLocation.Bearing ||
		loc.Elevation != testLocation.Elevation ||
		loc.Speed != testLocation.Speed ||
		loc.Timestamp != testLocation.Timestamp {
		t.Errorf("Locations do not match: %v \n %v", loc, testLocation)
	}

	nullString, _ := record.GetText(13)
	if nullString.Valid {
		t.Fail()
	}
}
