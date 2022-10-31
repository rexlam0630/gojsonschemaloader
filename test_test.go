package gojsonschemaloader

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSchemaPasse(t *testing.T) {
	schemaLoader := NewReferenceLoader("file:///Users/rex.lam/Documents/Projects/Testing/cf-testing/OrderPatchBuyIntentionRequest.json")
	sc, err := NewSchemaLoader().Compile(schemaLoader)
	if err != nil {
		t.Error(err)
	}

	exp := sc.GetExportSchema()
	fmt.Println(exp)

	j, err := json.MarshalIndent(exp, "", "  ")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(j))
}
