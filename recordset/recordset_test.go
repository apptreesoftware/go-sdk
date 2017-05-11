package recordset

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
