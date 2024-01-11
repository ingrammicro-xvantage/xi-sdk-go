/*
Reseller API Documentation

Testing ProductCatalogV5APIService

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

func Test_xi-sdk-resellers-go_ProductCatalogV5APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProductCatalogV5APIService GetV5CatalogProductsearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductCatalogV5API.GetV5CatalogProductsearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProductCatalogV5APIService MultiSKUPriceAndStock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProductCatalogV5API.MultiSKUPriceAndStock(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
