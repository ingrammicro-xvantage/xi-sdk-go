# \QuoteToOrderAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostQuoteToOrderV6**](QuoteToOrderAPI.md#PostQuoteToOrderV6) | **Post** /resellers/v6/q2o/orders | Quote To Order



## PostQuoteToOrderV6

> QuoteToOrderResponse PostQuoteToOrderV6(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).QuoteToOrderDetailsDTO(quoteToOrderDetailsDTO).IMSenderID(iMSenderID).Execute()

Quote To Order



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	xi_sdk_resellers "github.com/ingrammicro-xvantage/xi-sdk-resellers-go ingrammicro-xvantage/xi-sdk-resellers-go"
)

func main() {
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	quoteToOrderDetailsDTO := *openapiclient.NewQuoteToOrderDetailsDTO() // QuoteToOrderDetailsDTO | 
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuoteToOrderAPI.PostQuoteToOrderV6(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).QuoteToOrderDetailsDTO(quoteToOrderDetailsDTO).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuoteToOrderAPI.PostQuoteToOrderV6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostQuoteToOrderV6`: QuoteToOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `QuoteToOrderAPI.PostQuoteToOrderV6`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostQuoteToOrderV6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **quoteToOrderDetailsDTO** | [**QuoteToOrderDetailsDTO**](QuoteToOrderDetailsDTO.md) |  | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. | 

### Return type

[**QuoteToOrderResponse**](QuoteToOrderResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

