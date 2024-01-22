/*
Reseller API

Testing QuotesV5APIService

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

func Test_xi.sdk.resellers_QuotesV5APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QuotesV5APIService GetV5QuotesDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var quoteNumber string

		resp, httpRes, err := apiClient.QuotesV5API.GetV5QuotesDetails(context.Background(), quoteNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuotesV5APIService PostV5QuotesSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QuotesV5API.PostV5QuotesSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}