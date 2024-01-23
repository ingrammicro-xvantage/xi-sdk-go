# \QuotesV4API

All URIs are relative to *https://api.ingrammicro.com:443/sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV4Quotedetails**](QuotesV4API.md#PostV4Quotedetails) | **Post** /quotes/v1/quotedetails | Get Quote Details
[**PostV4Quotesearch**](QuotesV4API.md#PostV4Quotesearch) | **Post** /quotes/v1/quotes | Get Quote List



## PostV4Quotedetails

> QuoteDetailsResponse PostV4Quotedetails(ctx).QuoteDetailsRequest(quoteDetailsRequest).Execute()

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
	quoteDetailsRequest := *openapiclient.NewQuoteDetailsRequest() // QuoteDetailsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuotesV4API.PostV4Quotedetails(context.Background()).QuoteDetailsRequest(quoteDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesV4API.PostV4Quotedetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Quotedetails`: QuoteDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesV4API.PostV4Quotedetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4QuotedetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteDetailsRequest** | [**QuoteDetailsRequest**](QuoteDetailsRequest.md) |  | 

### Return type

[**QuoteDetailsResponse**](QuoteDetailsResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV4Quotesearch

> QuoteListResponse PostV4Quotesearch(ctx).QuoteListRequest(quoteListRequest).Execute()

Get Quote List



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
	resp, r, err := apiClient.QuotesV4API.PostV4Quotesearch(context.Background()).QuoteListRequest(quoteListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesV4API.PostV4Quotesearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Quotesearch`: QuoteListResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesV4API.PostV4Quotesearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4QuotesearchRequest struct via the builder pattern


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

