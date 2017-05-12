package apptree

import (
	"encoding/json"
	"testing"
	"time"
)

func TestParseRecordSet(t *testing.T) {
	var configuration Configuration
	err := json.Unmarshal([]byte(ConfigJSON), &configuration)
	if err != nil {
		t.Error(err)
	}
	if len(configuration.Attributes) != 21 {
		t.Fatalf("Invalid # of attributes. Expected 21, got %d", len(configuration.Attributes))
	}
	recordSet := NewRecordSet(configuration)
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
	if attr.ValueType() != Text {
		t.Fatalf("Invalid value type %s", attr.ValueType())
	}
	if attr.(TextValue).Value != "Normal" {
		t.Fatal("Expecting text value of `Normal`")
	}

	attr = record.GetValue(21)
	if attr.ValueType() != Relationship {
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
	if attr.ValueType() != DateTime {
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
	recordSet := NewRecordSet(configuration)
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
	if attr.ValueType() != Relationship {
		t.Fatalf("Expected relationship for attribute 21")
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
                ]
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
        }
    ],
    "serviceFilterParameters": null,
    "contextInfo": {},
    "dependentLists": null,
    "platformVersion": "5.5"
}`
