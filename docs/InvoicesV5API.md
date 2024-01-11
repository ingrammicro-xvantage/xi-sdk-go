# \InvoicesV5API

All URIs are relative to *https://api.ingrammicro.com:443/sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoices**](InvoicesV5API.md#GetInvoices) | **Get** /resellers/v5/invoices/{invoiceNumber} | Get Invoice Details



## GetInvoices

> InvoiceDetails GetInvoices(ctx, invoiceNumber).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).Execute()

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
	invoiceNumber := "invoiceNumber_example" // string | Ingram Micro Invoice Number (default to "20-RCW67-11")
	customerNumber := "customerNumber_example" // string | Your unique Ingram Micro customer number (default to "20-222222")
	isoCountryCode := "isoCountryCode_example" // string | ISO 2 char country code (default to "US")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesV5API.GetInvoices(context.Background(), invoiceNumber).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesV5API.GetInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoices`: InvoiceDetails
	fmt.Fprintf(os.Stdout, "Response from `InvoicesV5API.GetInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceNumber** | **string** | Ingram Micro Invoice Number | [default to &quot;20-RCW67-11&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerNumber** | **string** | Your unique Ingram Micro customer number | [default to &quot;20-222222&quot;]
 **isoCountryCode** | **string** | ISO 2 char country code | [default to &quot;US&quot;]

### Return type

[**InvoiceDetails**](InvoiceDetails.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

