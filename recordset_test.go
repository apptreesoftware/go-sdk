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

	if !record1.isEqual(&record2) {
		t.Errorf("Record 1 should match Record 2")
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

	if record1.isEqual(&record2) {
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

	if record1.isEqual(&record2) {
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

	if record1.isEqual(&record2) {
		t.Errorf("Record 1 should NOT match Record 2")
	}
}
