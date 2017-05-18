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

func TestParseRecordSet(t *testing.T) {
	var configuration Configuration
	err := json.Unmarshal([]byte(ConfigJSON), &configuration)
	if err != nil {
		t.Error(err)
	}
	if len(configuration.Attributes) != 27 {
		t.Fatalf("Invalid # of attributes. Expected 27, got %d", len(configuration.Attributes))
	}
	recordSet := NewRecordSet(&configuration)
	err = json.Unmarshal([]byte(DataSetJSON), &recordSet)
	if err != nil {
		t.Error(err)
	}
	if len(recordSet.Records) != 1 {
		t.Fatalf("Expected 1 record, record set size was %d", len(recordSet.Records))
	}
	record := recordSet.Records[0]
	if record.PrimaryKey != "12345" {
		t.Fatal("Expecting primary key 12345")
	}
	attr := record.GetValue(1)
	if attr.ValueType() != Type_Text {
		t.Fatalf("Invalid value type %s", attr.ValueType())
	}
	if attr.(TextValue).Value != "Normal" {
		t.Fatal("Expecting text value of `Normal`")
	}

	attr = record.GetValue(21)
	if attr.ValueType() != Type_Relationship {
		t.Fatal("Expected a relationship type at index 21")
	}
	if len(attr.(RelationshipValue).Items) != 1 {
		t.Fatal("Expected 1 sub item at index 21")
	}
	relationship := attr.(RelationshipValue).Items[0]
	if relationship.PrimaryKey != "54321" {
		t.Fatal("Expecting sub item primary key of 54321")
	}
	attr = relationship.GetValue(3)
	if attr.ValueType() != Type_DateTime {
		t.Fatal("Expecting date time for subItem 0-3")
	}
	compareDate, err := time.Parse("2006-01-02 15:04:05", "2017-05-21 16:21:07")
	if err != nil {
		t.Error(err)
	}
	if attr.(DateTimeValue).Value != compareDate {
		t.Fatalf("Expecting 0.3 to be a date of %v", compareDate)
	}
}

func TestMarshalUnmarshalRecord(t *testing.T) {
	var configuration Configuration
	err := json.Unmarshal([]byte(ConfigJSON), &configuration)
	if err != nil {
		t.Error(err)
	}
	recordSet := NewRecordSet(&configuration)
	err = json.Unmarshal([]byte(DataSetJSON), &recordSet)
	record := recordSet.Records[0]
	b, err := json.Marshal(&record)
	if err != nil {
		t.Error(err)
	}
	unmarshaledRecord, err := NewRecordFromJSON(b, &configuration)
	if err != nil {
		t.Error(err)
	}
	if unmarshaledRecord.PrimaryKey != "12345" {
		t.Fatalf("Unmarshaled record has incorrect primary key %s", unmarshaledRecord.PrimaryKey)
	}
	attr := unmarshaledRecord.GetValue(21)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 21")
	}
	if attr.ValueType() != Type_Relationship {
		t.Fatalf("Expected relationship for attribute 21")
	}
	childItem := attr.(RelationshipValue).Items[0]
	if childItem.GetValue(1).ValueType() != Type_Text {
		t.Fatal("Child Index 1 has wrong type")
	}
	if childItem.GetValue(1).(TextValue).Value != "Pending" {
		t.Fatal("Child index 1 != Pending")
	}

	attr = unmarshaledRecord.GetValue(22)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 22")
	}
	if attr.ValueType() != Type_DateRange {
		t.Fatal("Expected date time range for attribute 22")
	}
	attr = unmarshaledRecord.GetValue(23)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 23")
	}
	if attr.ValueType() != Type_TimeInterval {
		t.Fatal("Expected time interval for attribute 23")
	}
	attr = unmarshaledRecord.GetValue(24)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 24")
	}
	if attr.ValueType() != Type_DateTimeRange {
		t.Fatal("Expected date time range for attribute 24")
	}
	attr = unmarshaledRecord.GetValue(25)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 25")
	}
	if attr.ValueType() != Type_Image {
		t.Fatal("Expected image for attribute 25")
	}
	attr = unmarshaledRecord.GetValue(26)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 26")
	}
	if attr.ValueType() != Type_Location {
		t.Fatal("Expected location for atttribute 26")
	}
	attr = unmarshaledRecord.GetValue(27)
	if attr == nil {
		t.Fatal("Unexpected nil attribute for 27")
	}
	if attr.ValueType() != Type_Color {
		t.Fatal("Expected color for attribute 27")
	}
	t.Logf("Unmarshaled record has %d attributes", len(unmarshaledRecord.Attributes))
}

var DataSetJSON = `{
    "success": true,
    "message": null,
    "showMessageAsAlert": false,
    "totalRecords": 1,
    "numberOfRecords": 1,
    "moreRecordsAvailable": false,
    "records": [
        {
            "primaryKey": "12345",
            "CRUDStatus": "READ",
            "clientKey": null,
            "recordType": "RECORD",
            "status": "NONE",
            "attributes": [
                null,
                "Normal",
                "12345",
                "Test Requisition",
                "This is a test requisition",
                null,
                "John",
                "Doe",
                "john@example.com",
                null,
                "123-456-7890",
                "123 Fake Street",
                "New York City",
                "NY",
                "12345",
                "USA",
                "United States",
                "2017-05-21 16:21:07",
                "12345",
                "Acme",
                "12345",
                [
                    {
                        "primaryKey": "54321",
                        "CRUDStatus": "READ",
                        "clientKey": null,
                        "recordType": "RECORD",
                        "status": "NONE",
                        "attributes": [
                            null,
                            "Pending",
                            "1",
                            "2017-05-21 16:21:07",
                            "Fancy AppTree T-Shirt",
                            "EA",
                            "EA",
                            "2",
                            "12-1234567-1234",
                            "20"
                        ]
                    }
                ],
                {
                	"from": "2017-05-17",
                	"to": "2017-05-28"
                },
                "54",
                {
                	"from": "2017-05-18 15:00:00",
                	"to": "2017-05-20 01:00:00"
                },
                {
                	"imageURL": "http://fakeImage.com",
                	"uploadKey": "someUploadKey"
                },
                {
                	"latitude": 13.56789,
                	"longitude": 23.98475,
                	"bearing": 0.2,
                	"speed": 87.345,
                	"accuracy": 4,
                	"elevation": 100,
                	"timestamp": "2017-05-28 01:00:00"
                },
                {
                	"r": 76,
                	"g": 175,
                	"b": 80,
                	"a": 0
                }
            ]
        }
    ]
}`

var ConfigJSON = `{
    "success": true,
    "message": null,
    "showMessageAsAlert": false,
    "async": false,
    "name": "Jaggaer Purchase Requisition Test",
    "attributes": [
        {
            "name": "Id",
            "attributeType": "Integer",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 0
        },
        {
            "name": "Type",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 1
        },
        {
            "name": "External Request Number",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 2
        },
        {
            "name": "Name",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 3
        },
        {
            "name": "Requestor First Name",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 6
        },
        {
            "name": "Requestor Last Name",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 7
        },
        {
            "name": "Requestor Email",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 8
        },
        {
            "name": "Requestor Phone",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 10
        },
        {
            "name": "Description",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 4
        },
        {
            "name": "Create Date",
            "attributeType": "DateTime",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 5
        },
        {
            "name": "Ship To Address Lines",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 11
        },
        {
            "name": "Ship To Address City",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 12
        },
        {
            "name": "Ship To Address State",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 13
        },
        {
            "name": "Ship To Address Postal Code",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 14
        },
        {
            "name": "Ship To Address ISO Country Codes",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 15
        },
        {
            "name": "Ship To Address Countries",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 16
        },
        {
            "name": "Requested Delivery Date",
            "attributeType": "DateTime",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 17
        },
        {
            "name": "Supplier Group ID",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 18
        },
        {
            "name": "Supplier Group Name",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 19
        },
        {
            "name": "Supplier Group Fulfillment Address ID",
            "attributeType": "Text",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 20
        },
        {
            "name": "Purchase Requisition lines",
            "relatedService": {
                "success": true,
                "message": null,
                "showMessageAsAlert": false,
                "async": false,
                "name": "Purchase Requisition Lines",
                "attributes": [
                    {
                        "name": "Id",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 0
                    },
                    {
                        "name": "Workflow Status",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 1
                    },
                    {
                        "name": "Line Number",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 2
                    },
                    {
                        "name": "Requested Delivery Date",
                        "attributeType": "DateTime",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 3
                    },
                    {
                        "name": "Description",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 4
                    },
                    {
                        "name": "Product Unit of Measure",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 5
                    },
                    {
                        "name": "Product Size",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 6
                    },
                    {
                        "name": "Lead Time Days",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 7
                    },
                    {
                        "name": "Barcode",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 8
                    },
                    {
                        "name": "Quantity",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 9
                    }
                ],
                "serviceFilterParameters": null,
                "contextInfo": {},
                "attributeConfigurationForIndexMap": {
                    "0": {
                        "name": "Id",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 0
                    },
                    "1": {
                        "name": "Workflow Status",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 1
                    },
                    "2": {
                        "name": "Line Number",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 2
                    },
                    "3": {
                        "name": "Requested Delivery Date",
                        "attributeType": "DateTime",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 3
                    },
                    "4": {
                        "name": "Description",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 4
                    },
                    "5": {
                        "name": "Product Unit of Measure",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 5
                    },
                    "6": {
                        "name": "Product Size",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 6
                    },
                    "7": {
                        "name": "Lead Time Days",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 7
                    },
                    "8": {
                        "name": "Barcode",
                        "attributeType": "Text",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 8
                    },
                    "9": {
                        "name": "Quantity",
                        "attributeType": "Integer",
                        "create": false,
                        "createRequired": false,
                        "update": false,
                        "updateRequired": false,
                        "search": false,
                        "searchRequired": false,
                        "attributeIndex": 9
                    }
                },
                "dependentLists": null,
                "platformVersion": "5.5"
            },
            "attributeType": "Relationship",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 21
        },
        {
            "name": "Available Dates",
            "attributeType": "DateRange",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 22
        },
        {
            "name": "Time Interval",
            "attributeType": "TimeInterval",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 23
        },
        {
            "name": "Time Range",
            "attributeType": "DateTimeRange",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 24
        },
        {
            "name": "Image",
            "attributeType": "Image",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 25
        },
        {
            "name": "Location",
            "attributeType": "Location",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 26
        },
        {
            "name": "Color",
            "attributeType": "Color",
            "create": false,
            "createRequired": false,
            "update": false,
            "updateRequired": false,
            "search": false,
            "searchRequired": false,
            "attributeIndex": 27
        }
    ],
    "serviceFilterParameters": null,
    "contextInfo": {},
    "dependentLists": null,
    "platformVersion": "5.5"
}`
