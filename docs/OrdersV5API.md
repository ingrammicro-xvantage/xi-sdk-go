# \OrdersV5API

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrdersOrderNumber**](OrdersV5API.md#DeleteOrdersOrderNumber) | **Delete** /resellers/v5/Orders/{ordernumber} | Cancel an Existing Order
[**GetOrdersSearch**](OrdersV5API.md#GetOrdersSearch) | **Get** /resellers/v5/Orders/search | Search your Orders
[**GetV5OrdersDetails**](OrdersV5API.md#GetV5OrdersDetails) | **Get** /resellers/v5/Orders/{ordernumber} | Get Order Details
[**PostV5OrdersCreate**](OrdersV5API.md#PostV5OrdersCreate) | **Post** /resellers/v5/Orders | Create a New Order



## DeleteOrdersOrderNumber

> OrderCancelResponse DeleteOrdersOrderNumber(ctx, ordernumber).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).EntryDate(entryDate).Execute()

Cancel an Existing Order



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
	ordernumber := "20-RD128" // string | Ingram Micro sales order number
	customerNumber := "customerNumber_example" // string | Your unique Ingram Micro customer number
	isoCountryCode := "isoCountryCode_example" // string | 2 chars ISO country code
	entryDate := "entryDate_example" // string | Order entry date (yyyy-mm-dd) (default to "2020-04-03")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV5API.DeleteOrdersOrderNumber(context.Background(), ordernumber).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).EntryDate(entryDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV5API.DeleteOrdersOrderNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrdersOrderNumber`: OrderCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV5API.DeleteOrdersOrderNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ordernumber** | **string** | Ingram Micro sales order number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrdersOrderNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerNumber** | **string** | Your unique Ingram Micro customer number | 
 **isoCountryCode** | **string** | 2 chars ISO country code | 
 **entryDate** | **string** | Order entry date (yyyy-mm-dd) | [default to &quot;2020-04-03&quot;]

### Return type

[**OrderCancelResponse**](OrderCancelResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersSearch

> OrderSearchResponse GetOrdersSearch(ctx).CustomerNumber(customerNumber).Isocountrycode(isocountrycode).Ordernumber(ordernumber).Customerordernumber(customerordernumber).Orderstatus(orderstatus).Startcreatetimestamp(startcreatetimestamp).Endcreatetimestamp(endcreatetimestamp).Pagesize(pagesize).Pagenumber(pagenumber).Execute()

Search your Orders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	customerNumber := "20-222222" // string | Your unique Ingram Micro customer number
	isocountrycode := "US" // string | 2 char iso country code
	ordernumber := "ordernumber_example" // string | Ingram sales order number (optional)
	customerordernumber := "ZENPO1" // string | Search using your PO/Order number (optional)
	orderstatus := "orderstatus_example" // string | Ingram Micro order status (optional) (default to "any")
	startcreatetimestamp := time.Now() // time.Time | Search start date/time in UTC format (optional)
	endcreatetimestamp := time.Now() // time.Time | Search end date/time in UTC format (optional)
	pagesize := int32(56) // int32 | Number of records required in the call (optional)
	pagenumber := int32(56) // int32 | the page number reference (optional) (default to 1)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV5API.GetOrdersSearch(context.Background()).CustomerNumber(customerNumber).Isocountrycode(isocountrycode).Ordernumber(ordernumber).Customerordernumber(customerordernumber).Orderstatus(orderstatus).Startcreatetimestamp(startcreatetimestamp).Endcreatetimestamp(endcreatetimestamp).Pagesize(pagesize).Pagenumber(pagenumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV5API.GetOrdersSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersSearch`: OrderSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV5API.GetOrdersSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerNumber** | **string** | Your unique Ingram Micro customer number | 
 **isocountrycode** | **string** | 2 char iso country code | 
 **ordernumber** | **string** | Ingram sales order number | 
 **customerordernumber** | **string** | Search using your PO/Order number | 
 **orderstatus** | **string** | Ingram Micro order status | [default to &quot;any&quot;]
 **startcreatetimestamp** | **time.Time** | Search start date/time in UTC format | 
 **endcreatetimestamp** | **time.Time** | Search end date/time in UTC format | 
 **pagesize** | **int32** | Number of records required in the call | 
 **pagenumber** | **int32** | the page number reference | [default to 1]

### Return type

[**OrderSearchResponse**](OrderSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV5OrdersDetails

> OrderDetailResponse GetV5OrdersDetails(ctx, ordernumber).Customernumber(customernumber).Isocountrycode(isocountrycode).Customerordernumber(customerordernumber).Startcreatetimestamp(startcreatetimestamp).Endcreatetimestamp(endcreatetimestamp).Simulate(simulate).Execute()

Get Order Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	ordernumber := "20-RD128" // string | Ingram Micro sales order number
	customernumber := "customernumber_example" // string | Your unique Ingram Micro customer number (default to "20-222222")
	isocountrycode := "isocountrycode_example" // string | 2 chars ISO country code (default to "US")
	customerordernumber := "customerordernumber_example" // string | Your PO/Order Number provide at the time of order creation (optional)
	startcreatetimestamp := time.Now() // string | Filter start date - format YYYY-MM-DD (optional)
	endcreatetimestamp := "2020-04-20" // string | Filter end date - format YYYY-MM-DD (optional)
	simulate := "simulate_example" // string | Order response for various order statuses. Not for use in production. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV5API.GetV5OrdersDetails(context.Background(), ordernumber).Customernumber(customernumber).Isocountrycode(isocountrycode).Customerordernumber(customerordernumber).Startcreatetimestamp(startcreatetimestamp).Endcreatetimestamp(endcreatetimestamp).Simulate(simulate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV5API.GetV5OrdersDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV5OrdersDetails`: OrderDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV5API.GetV5OrdersDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ordernumber** | **string** | Ingram Micro sales order number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV5OrdersDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customernumber** | **string** | Your unique Ingram Micro customer number | [default to &quot;20-222222&quot;]
 **isocountrycode** | **string** | 2 chars ISO country code | [default to &quot;US&quot;]
 **customerordernumber** | **string** | Your PO/Order Number provide at the time of order creation | 
 **startcreatetimestamp** | **string** | Filter start date - format YYYY-MM-DD | 
 **endcreatetimestamp** | **string** | Filter end date - format YYYY-MM-DD | 
 **simulate** | **string** | Order response for various order statuses. Not for use in production. | 

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


## PostV5OrdersCreate

> OrderCreateResponse PostV5OrdersCreate(ctx).OrderCreateRequest(orderCreateRequest).Execute()

Create a New Order



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
	orderCreateRequest := *openapiclient.NewOrderCreateRequest() // OrderCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV5API.PostV5OrdersCreate(context.Background()).OrderCreateRequest(orderCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV5API.PostV5OrdersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV5OrdersCreate`: OrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV5API.PostV5OrdersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV5OrdersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderCreateRequest** | [**OrderCreateRequest**](OrderCreateRequest.md) |  | 

### Return type

[**OrderCreateResponse**](OrderCreateResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

