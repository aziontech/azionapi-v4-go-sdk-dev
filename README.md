# azionapi-v4-go-sdk

Golang SDK for Azion APIs (v4)


## Azion Go SDK Usage Example - Domains API

This example demonstrates how to use the Azion API Go SDK to manage domains via Azion's API, including authenticating, listing, creating, retrieving, updating, and deleting domains.

### 📦 Installation

First, install the Azion API Go SDK:

```bash
go get github.com/aziontech/azionapi-go-sdk
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

	sdk "github.com/aziontech/azionapi-go-sdk/domains"
)

func main() {
	token := "YOUR_PERSONAL_TOKEN"
	baseURL := "https://api.azionapi.net"

	// Create a configuration
	conf := sdk.NewConfiguration()
	conf.HTTPClient = &http.Client{Timeout: 30 * time.Second}
	conf.AddDefaultHeader("Authorization", "token "+token)
	conf.AddDefaultHeader("Accept", "application/json;version=3")
	conf.UserAgent = "ExampleClient/1.0"
	conf.Servers = sdk.ServerConfigurations{
		{URL: baseURL},
	}

	client := sdk.NewAPIClient(conf)

	ctx := context.Background()

	// --- Create a domain ---
	createReq := sdk.CreateDomainRequest{
		Name:             "example-domain",
		DomainName:       "example.com",
		EdgeApplicationId: 123456, // Replace with your Edge App ID
	}

	createResp, httpResp, err := client.DomainsAPI.CreateDomain(ctx).
		CreateDomainRequest(createReq).
		Execute()
	if err != nil {
		fmt.Println("Error creating domain:", err)
		return
	}
	fmt.Printf("Created domain ID: %d\n", createResp.Results.GetId())

	// --- Get a domain ---
	domainID := fmt.Sprint(createResp.Results.GetId())
	getResp, httpResp, err := client.DomainsAPI.GetDomain(ctx, domainID).Execute()
	if err != nil {
		fmt.Println("Error fetching domain:", err)
		return
	}
	fmt.Printf("Domain Name: %s\n", getResp.Results.GetDomainName())

	// --- Update a domain ---
	updateReq := sdk.UpdateDomainRequest{
		Name: "updated-domain-name",
	}
	_, httpResp, err = client.DomainsAPI.UpdateDomain(ctx, domainID).
		UpdateDomainRequest(updateReq).
		Execute()
	if err != nil {
		fmt.Println("Error updating domain:", err)
		return
	}
	fmt.Println("Domain updated.")

	// --- List domains ---
	listResp, httpResp, err := client.DomainsAPI.GetDomains(ctx).
		Page(1).
		PageSize(10).
		OrderBy("id").
		Sort("asc").
		Execute()
	if err != nil {
		fmt.Println("Error listing domains:", err)
		return
	}
	for _, d := range listResp.Results {
		fmt.Printf("ID: %d | Name: %s\n", d.GetId(), d.GetName())
	}

	// --- Delete domain ---
	_, err = client.DomainsAPI.DelDomain(ctx, domainID).Execute()
	if err != nil {
		fmt.Println("Error deleting domain:", err)
		return
	}
	fmt.Println("Domain deleted.")
}
```

---

### 📘 Resources

- [Azion API Documentation](https://api.azion.com/)
- [Azion CLI (azion-cli)](https://github.com/aziontech/azion)
- [Azion API Go SDK](https://github.com/aziontech/azionapi-v4-go-sdk)

---

### ⚠️ Notes

- Be careful with error handling. The Azion API SDK returns detailed HTTP responses which may include logs in the body.

