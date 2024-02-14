# \InvoicesV4API

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV4Invoicedetails**](InvoicesV4API.md#PostV4Invoicedetails) | **Post** /invoices/v4/invoicedetails | Get Invoice Details



## PostV4Invoicedetails

> InvoiceDetailResponse PostV4Invoicedetails(ctx).InvoiceDetailRequest(invoiceDetailRequest).Execute()

Get Invoice Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	invoiceDetailRequest := *openapiclient.NewInvoiceDetailRequest() // InvoiceDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesV4API.PostV4Invoicedetails(context.Background()).InvoiceDetailRequest(invoiceDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesV4API.PostV4Invoicedetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Invoicedetails`: InvoiceDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesV4API.PostV4Invoicedetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4InvoicedetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceDetailRequest** | [**InvoiceDetailRequest**](InvoiceDetailRequest.md) |  | 

### Return type

[**InvoiceDetailResponse**](InvoiceDetailResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

