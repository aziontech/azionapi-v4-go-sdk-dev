package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aziontech/azionapi-v4-go-sdk-dev/openapi"
)

func main() {
	// JSON string representing the behavior
	jsonData := `{"type": "deny"}`

	// Unmarshal JSON into the struct
	var behavior openapi.ApplicationRuleEngineNoArgs
	err := json.Unmarshal([]byte(jsonData), &behavior)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Printf("Behavior type: %s\n", behavior.Type)
}
