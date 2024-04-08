# \AsyncOrderCreateAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAsyncOrderCreateV7**](AsyncOrderCreateAPI.md#PostAsyncOrderCreateV7) | **Post** /resellers/v7/orders | Async Order Create



## PostAsyncOrderCreateV7

> AsyncOrderCreateResponse PostAsyncOrderCreateV7(ctx).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).AsyncOrderCreateDTO(asyncOrderCreateDTO).IMSenderID(iMSenderID).Execute()

Async Order Create



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
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	asyncOrderCreateDTO := *openapiclient.NewAsyncOrderCreateDTO() // AsyncOrderCreateDTO | 
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AsyncOrderCreateAPI.PostAsyncOrderCreateV7(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).AsyncOrderCreateDTO(asyncOrderCreateDTO).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AsyncOrderCreateAPI.PostAsyncOrderCreateV7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAsyncOrderCreateV7`: AsyncOrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `AsyncOrderCreateAPI.PostAsyncOrderCreateV7`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAsyncOrderCreateV7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **asyncOrderCreateDTO** | [**AsyncOrderCreateDTO**](AsyncOrderCreateDTO.md) |  | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. | 

### Return type

[**AsyncOrderCreateResponse**](AsyncOrderCreateResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

