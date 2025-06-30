# azionapi-v4-go-sdk

Golang SDK for Azion APIs (v4)


## Azion Go SDK Usage Example - Edge Connectors API (v4)

This example demonstrates how to use the Azion API Go SDK v4 to manage **Edge Connectors** via Azion's API. It includes authentication, creating, retrieving, updating, listing, and deleting Edge Connectors.

### 📦 Installation

First, install the Azion API Go SDK v4 for Stage:

```bash
go get github.com/aziontech/azionapi-v4-go-sdk-dev
```

---

### 🔑 Authentication

You need a valid Personal Token from Azion. Set it in the `Authorization` header.

---

### 🛠️ Example Usage

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	sdk "github.com/aziontech/azionapi-v4-go-sdk-dev/edge"
)

func main() {
	token := "YOUR_PERSONAL_TOKEN"
	baseURL := "https://api.azionapi.net"

	// Setup configuration
	conf := sdk.NewConfiguration()
	conf.HTTPClient = &http.Client{Timeout: 50 * time.Second}
	conf.AddDefaultHeader("Authorization", "token " + token)
	conf.AddDefaultHeader("Accept", "application/json")
	conf.UserAgent = "ExampleClient/1.0"
	conf.Servers = sdk.ServerConfigurations{
		{URL: baseURL},
	}

	client := sdk.NewAPIClient(conf)
	ctx := context.Background()

	// --- Create an Edge Connector ---
	createReq := sdk.EdgeConnectorPolymorphicRequest{
		// Fill with required fields, e.g.:
		// Name: "my-connector",
	}
	createResp, httpResp, err := client.EdgeConnectorsAPI.CreateEdgeConnector(ctx).
		EdgeConnectorPolymorphicRequest(createReq).
		Execute()
	if err != nil {
		fmt.Println("Error creating Edge Connector:", err)
		return
	}
	fmt.Printf("Created Edge Connector ID: %s\n", createResp.Data.GetUuid())

	// --- Retrieve an Edge Connector ---
	id := createResp.Data.GetUuid()
	getResp, httpResp, err := client.EdgeConnectorsAPI.RetrieveEdgeConnector(ctx, id).Execute()
	if err != nil {
		fmt.Println("Error retrieving Edge Connector:", err)
		return
	}
	fmt.Printf("Retrieved Edge Connector Name: %s\n", getResp.Data.GetName())

	// --- Update an Edge Connector ---
	updateReq := sdk.PatchedEdgeConnectorPolymorphicRequest{
		// Name: sdk.PtrString("updated-connector-name"),
	}
	updateResp, httpResp, err := client.EdgeConnectorsAPI.PartialUpdateEdgeConnector(ctx, id).
		PatchedEdgeConnectorPolymorphicRequest(updateReq).
		Execute()
	if err != nil {
		fmt.Println("Error updating Edge Connector:", err)
		return
	}
	fmt.Printf("Updated Edge Connector Name: %s\n", updateResp.Data.GetName())

	// --- List Edge Connectors ---
	listResp, httpResp, err := client.EdgeConnectorsAPI.ListEdgeConnectors(ctx).
		Page(1).
		PageSize(10).
		Ordering("id").
		Execute()
	if err != nil {
		fmt.Println("Error listing Edge Connectors:", err)
		return
	}
	for _, connector := range listResp.Results {
		fmt.Printf("Connector ID: %s | Name: %s\n", connector.GetUuid(), connector.GetName())
	}

	// --- Delete Edge Connector ---
	_, httpResp, err = client.EdgeConnectorsAPI.DestroyEdgeConnector(ctx, id).Execute()
	if err != nil {
		fmt.Println("Error deleting Edge Connector:", err)
		return
	}
	fmt.Println("Edge Connector deleted.")
}
```

---

### 📘 Resources

- [Azion API Documentation](https://api.azion.com)
- [Azion CLI (azion-cli)](https://github.com/aziontech/azion)
- [Production Azion API Go SDK v4](https://github.com/aziontech/azionapi-v4-go-sdk)

---

### ⚠️ Notes

- Always set the appropriate `Accept` and `Authorization` headers.
- Make sure the fields required in your specific connector implementation are set.
- The SDK uses polymorphic request/response types, so adjust based on connector type.
