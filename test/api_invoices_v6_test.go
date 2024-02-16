/*
XI Sdk Resellers

Testing InvoicesV6APIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package xi_sdk_resellers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_xi_sdk_resellers_InvoicesV6APIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InvoicesV6APIService GetInvoicedetailsV6", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var invoicenumber string

		resp, httpRes, err := apiClient.InvoicesV6API.GetInvoicedetailsV6(context.Background(), invoicenumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
