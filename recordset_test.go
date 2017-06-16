package apptree

import (
	"encoding/json"
	"testing"

	"github.com/apptreesoftware/go-sdk/example"
)

func TestMarshalConfiguration(t *testing.T) {
	model := example.Issue{Id: "1234"}
	configuration, err := GetConfiguration(model)
	b, err := json.Marshal(&configuration)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}

func TestMarshalRecord(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}

	record1, err := NewRecordFromJSON([]byte(Config1DataSetItemTxtChange), &config)
	if err != nil {
		t.Error(err)
	}

	b, err := json.Marshal(&record1)
	if err != nil {
		t.Error(err)
	}

	record2, err := NewRecordFromJSON(b, &config)
	if !record1.IsEqual(&record2) {
		t.Error("Records not equal")
	}

	if record1.CRUDStatus != record2.CRUDStatus {
		t.Errorf("CRUD Status not equal. %s != %s", record1.CRUDStatus, record2.CRUDStatus)
	}

}

func TestEqualComparison(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}

	record1, err := NewRecordFromJSON([]byte(Config1DataSetItem), &config)
	if err != nil {
		t.Error(err)
	}
	record2, err := NewRecordFromJSON([]byte(Config1DataSetItem), &config)

	if !record1.IsEqual(&record2) {
		t.Errorf("Record 1 should match Record 2")
	}
}

func TestFamisMarshaling(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(FAMISWoConfig), &config)
	if err != nil {
		t.Error(err)
	}
	record1, err := NewRecordFromJSON([]byte(FAMISDataSetItem), &config)
	if err != nil {
		t.Error(err)
	}
	if record1.PrimaryKey != "237" {
		t.Errorf("No primary key")
	}

	b, err := json.Marshal(&record1)
	if err != nil {
		t.Error(err)
	}
	record2, err := NewRecordFromJSON(b, &config)
	if err != nil {
		t.Error(err)
	}
	if record2.PrimaryKey != "237" {
		t.Errorf("No primary key")
	}

}

func TestNotEqualTextComparison(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}

	record1, err := NewRecordFromJSON([]byte(Config1DataSetItem), &config)
	if err != nil {
		t.Error(err)
	}
	record2, err := NewRecordFromJSON([]byte(Config1DataSetItemTxtChange), &config)

	if record1.IsEqual(&record2) {
		t.Errorf("Record 1 should NOT match Record 2")
	}
}

func TestNotEqualDateComparison(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}

	record1, err := NewRecordFromJSON([]byte(Config1DataSetItem), &config)
	if err != nil {
		t.Error(err)
	}
	record2, err := NewRecordFromJSON([]byte(Config1DataSetItemDteChange), &config)

	if record1.IsEqual(&record2) {
		t.Errorf("Record 1 should NOT match Record 2")
	}
}

func TestNotEqualRelationshipComparison(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}

	record1, err := NewRecordFromJSON([]byte(Config1DataSetItem), &config)
	if err != nil {
		t.Error(err)
	}
	record2, err := NewRecordFromJSON([]byte(Config1DataSetItemRelChange), &config)

	if record1.IsEqual(&record2) {
		t.Errorf("Record 1 should NOT match Record 2")
	}
}

func TestParseCRUDStatus(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}
	record1, err := NewRecordFromJSON([]byte(Config1DataSetItemTxtChange), &config)
	if err != nil {
		t.Error(err)
	}
	if record1.CRUDStatus != StatusUpdate {
		t.Errorf("Expecting status of UPDATE, got %s", record1.CRUDStatus)
	}

	record2, err := NewRecordFromJSON([]byte(Config1DataSetItem), &config)
	if err != nil {
		t.Error(err)
	}
	if record2.CRUDStatus != StatusRead {
		t.Errorf("Expecting status of UPDATE, got %s", record1.CRUDStatus)
	}
}

func TestDefaultCRUDStatus(t *testing.T) {
	var config Configuration
	err := json.Unmarshal([]byte(Config1), &config)
	if err != nil {
		t.Error(err)
	}
	record1, err := NewRecordFromJSON([]byte(Config1DataSetItemNoCRUD), &config)
	if err != nil {
		t.Error(err)
	}
	if record1.CRUDStatus != StatusRead {
		t.Errorf("Expecting status of READ, got %s", record1.CRUDStatus)
	}
}
