/*
Reseller API Documentation

Testing DealsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package xi-sdk-resellers-go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_xi-sdk-resellers-go_DealsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DealsAPIService GetResellersV6Dealsdetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dealId string

		resp, httpRes, err := apiClient.DealsAPI.GetResellersV6Dealsdetails(context.Background(), dealId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DealsAPIService GetResellersV6Dealssearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DealsAPI.GetResellersV6Dealssearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
