# \OrdersV6API

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrderdetailsV6**](OrdersV6API.md#GetOrderdetailsV6) | **Get** /resellers/v6/orders/{ordernumber} | Get Order Details v6



## GetOrderdetailsV6

> OrderDetailResponse GetOrderdetailsV6(ctx, ordernumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).IngramOrderDate(ingramOrderDate).VendorNumber(vendorNumber).SimulateStatus(simulateStatus).IsIml(isIml).RegionCode(regionCode).Execute()

Get Order Details v6



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	xi_sdk_resellers "github.com/ingrammicro-xvantage/xi-sdk-resellers-go ingrammicro-xvantage/xi-sdk-resellers-go"
)

func main() {
	ordernumber := "20-RD3QV" // string | The Ingram Micro sales order number.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems.
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany. (optional)
	ingramOrderDate := time.Now() // string | The date and time in UTC format that the order was created. (optional)
	vendorNumber := "vendorNumber_example" // string | Vendor Number. (optional)
	simulateStatus := "simulateStatus_example" // string | Order response for various order statuses. Not for use in production. (optional)
	isIml := true // bool | True/False only for IML customers. (optional)
	regionCode := "regionCode_example" // string | Region code for sandbox testing - Not for use in production. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV6API.GetOrderdetailsV6(context.Background(), ordernumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).IngramOrderDate(ingramOrderDate).VendorNumber(vendorNumber).SimulateStatus(simulateStatus).IsIml(isIml).RegionCode(regionCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV6API.GetOrderdetailsV6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrderdetailsV6`: OrderDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV6API.GetOrderdetailsV6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ordernumber** | **string** | The Ingram Micro sales order number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderdetailsV6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany. | 
 **ingramOrderDate** | **string** | The date and time in UTC format that the order was created. | 
 **vendorNumber** | **string** | Vendor Number. | 
 **simulateStatus** | **string** | Order response for various order statuses. Not for use in production. | 
 **isIml** | **bool** | True/False only for IML customers. | 
 **regionCode** | **string** | Region code for sandbox testing - Not for use in production. | 

### Return type

[**OrderDetailResponse**](OrderDetailResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

