/*
edge-api

Testing EdgeFirewallsFunctionAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package edgeapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_edgeapi_EdgeFirewallsFunctionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EdgeFirewallsFunctionAPIService CreateEdgeFirewallFunction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeFirewallId string

		resp, httpRes, err := apiClient.EdgeFirewallsFunctionAPI.CreateEdgeFirewallFunction(context.Background(), edgeFirewallId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeFirewallsFunctionAPIService DestroyEdgeFirewallFunction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeFirewallId string
		var id string

		resp, httpRes, err := apiClient.EdgeFirewallsFunctionAPI.DestroyEdgeFirewallFunction(context.Background(), edgeFirewallId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeFirewallsFunctionAPIService ListEdgeFirewallFunction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeFirewallId string

		resp, httpRes, err := apiClient.EdgeFirewallsFunctionAPI.ListEdgeFirewallFunction(context.Background(), edgeFirewallId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeFirewallsFunctionAPIService PartialUpdateEdgeFirewallFunction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeFirewallId string
		var id string

		resp, httpRes, err := apiClient.EdgeFirewallsFunctionAPI.PartialUpdateEdgeFirewallFunction(context.Background(), edgeFirewallId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeFirewallsFunctionAPIService RetrieveEdgeFirewallFunction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeFirewallId string
		var id string

		resp, httpRes, err := apiClient.EdgeFirewallsFunctionAPI.RetrieveEdgeFirewallFunction(context.Background(), edgeFirewallId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EdgeFirewallsFunctionAPIService UpdateEdgeFirewallFunction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var edgeFirewallId string
		var id string

		resp, httpRes, err := apiClient.EdgeFirewallsFunctionAPI.UpdateEdgeFirewallFunction(context.Background(), edgeFirewallId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
