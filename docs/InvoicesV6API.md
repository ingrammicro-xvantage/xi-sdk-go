# \InvoicesV6API

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoicedetailsV6**](InvoicesV6API.md#GetInvoicedetailsV6) | **Get** /resellers/v6/invoices/{invoicenumber} | Get Invoice Details v6



## GetInvoicedetailsV6

> InvoiceDetailResponse GetInvoicedetailsV6(ctx, invoicenumber).Version(version).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMApplicationID(iMApplicationID).CustomerType(customerType).IncludeSerialNumbers(includeSerialNumbers).Execute()

Get Invoice Details v6



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	xi_sdk_resellers "https://github.com/dprakash2101/GoSDK" ingrammicro-xvantage/xi-sdk-resellers-go ingrammicro-xvantage/xi-sdk-resellers-go
)

func main() {
	invoicenumber := "335238411" // string | The Ingram Micro invoice number.
	version := "20-222222" // string | Version of codebase.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	iMApplicationID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany.
	customerType := "invoice" // string | it should be invoice or order (optional)
	includeSerialNumbers := false // bool | if serial in the response send as true or else false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesV6API.GetInvoicedetailsV6(context.Background(), invoicenumber).Version(version).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMApplicationID(iMApplicationID).CustomerType(customerType).IncludeSerialNumbers(includeSerialNumbers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesV6API.GetInvoicedetailsV6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoicedetailsV6`: InvoiceDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `InvoicesV6API.GetInvoicedetailsV6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoicenumber** | **string** | The Ingram Micro invoice number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicedetailsV6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | Version of codebase. | 
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **iMApplicationID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany. | 
 **customerType** | **string** | it should be invoice or order | 
 **includeSerialNumbers** | **bool** | if serial in the response send as true or else false | 

### Return type

[**InvoiceDetailResponse**](InvoiceDetailResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

