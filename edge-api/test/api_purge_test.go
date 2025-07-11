/*
edge-api

Testing PurgeAPIService

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

func Test_edgeapi_PurgeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PurgeAPIService CreatePurgeRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var purgeType string

		resp, httpRes, err := apiClient.PurgeAPI.CreatePurgeRequest(context.Background(), purgeType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
