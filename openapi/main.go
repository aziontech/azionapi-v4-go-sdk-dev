package openapi

import (
	"encoding/json"
	"fmt"
	"log"
)

// ExampleUnmarshalApplicationRuleEngineNoArgs demonstrates unmarshaling JSON into ApplicationRuleEngineNoArgs
func ExampleUnmarshalApplicationRuleEngineNoArgs() {
	// JSON string representing the behavior
	jsonData := `{"type": "deny"}`

	// Unmarshal JSON into the struct
	var behavior ApplicationRuleEngineNoArgs
	err := json.Unmarshal([]byte(jsonData), &behavior)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Printf("Behavior type: %s\n", behavior.Type)
}
