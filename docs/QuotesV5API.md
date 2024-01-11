# \QuotesV5API

All URIs are relative to *https://api.ingrammicro.com:443/sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV5QuotesDetails**](QuotesV5API.md#GetV5QuotesDetails) | **Get** /resellers/v5/quote/{quoteNumber} | Get Quote Details
[**PostV5QuotesSearch**](QuotesV5API.md#PostV5QuotesSearch) | **Post** /resellers/v5/quote/search | Search Quotes



## GetV5QuotesDetails

> QuoteDetails GetV5QuotesDetails(ctx, quoteNumber).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).ThirdPartySource(thirdPartySource).Execute()

Get Quote Details



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
	quoteNumber := "quoteNumber_example" // string | Ingram Micro Quote Number (default to "QUO-25576-C8S2W7")
	customerNumber := "customerNumber_example" // string | Your Ingram Micro unique customer number (default to "20-222222")
	isoCountryCode := "isoCountryCode_example" // string |  (default to "US")
	thirdPartySource := "thirdPartySource_example" // string | Unique identifier used to identify the third party source accessing the services (optional) (default to "customer")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesV5API.GetV5QuotesDetails(context.Background(), quoteNumber).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).ThirdPartySource(thirdPartySource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesV5API.GetV5QuotesDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV5QuotesDetails`: QuoteDetails
	fmt.Fprintf(os.Stdout, "Response from `QuotesV5API.GetV5QuotesDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteNumber** | **string** | Ingram Micro Quote Number | [default to &quot;QUO-25576-C8S2W7&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetV5QuotesDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerNumber** | **string** | Your Ingram Micro unique customer number | [default to &quot;20-222222&quot;]
 **isoCountryCode** | **string** |  | [default to &quot;US&quot;]
 **thirdPartySource** | **string** | Unique identifier used to identify the third party source accessing the services | [default to &quot;customer&quot;]

### Return type

[**QuoteDetails**](QuoteDetails.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV5QuotesSearch

> QuoteListResponse PostV5QuotesSearch(ctx).QuoteListRequest(quoteListRequest).Execute()

Search Quotes



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
	quoteListRequest := *openapiclient.NewQuoteListRequest() // QuoteListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesV5API.PostV5QuotesSearch(context.Background()).QuoteListRequest(quoteListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesV5API.PostV5QuotesSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV5QuotesSearch`: QuoteListResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesV5API.PostV5QuotesSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV5QuotesSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteListRequest** | [**QuoteListRequest**](QuoteListRequest.md) |  | 

### Return type

[**QuoteListResponse**](QuoteListResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

