/*
tls-api

Testing DigitalCertificatesCertificateSigningRequestsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tlsapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_tlsapi_DigitalCertificatesCertificateSigningRequestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DigitalCertificatesCertificateSigningRequestsAPIService CreateCertificateSigningRequestCSR", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DigitalCertificatesCertificateSigningRequestsAPI.CreateCertificateSigningRequestCSR(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
