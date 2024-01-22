/*
Reseller API

Testing InvoicesV4APIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package xi.sdk.resellers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_xi.sdk.resellers_InvoicesV4APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InvoicesV4APIService PostV4Invoicedetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InvoicesV4API.PostV4Invoicedetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}