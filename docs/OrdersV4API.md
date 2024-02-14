# \OrdersV4API

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV4Ordercreate**](OrdersV4API.md#PostV4Ordercreate) | **Post** /orders/v4/ordercreate | Create a new Order
[**PostV4Orderdelete**](OrdersV4API.md#PostV4Orderdelete) | **Post** /orders/v4/orderdelete | Delete an Order
[**PostV4Orderdetails**](OrdersV4API.md#PostV4Orderdetails) | **Post** /orders/v4/orderdetails | Get Order Details
[**PostV4Ordermodify**](OrdersV4API.md#PostV4Ordermodify) | **Post** /orders/v4/ordermodify | Modify an Existing Order
[**PostV4Ordersearch**](OrdersV4API.md#PostV4Ordersearch) | **Post** /orders/v4/orderlookup | Order Search



## PostV4Ordercreate

> OrderCreateResponse PostV4Ordercreate(ctx).OrderCreateRequest(orderCreateRequest).Execute()

Create a new Order



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
	resp, r, err := apiClient.OrdersV4API.PostV4Ordercreate(context.Background()).OrderCreateRequest(orderCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV4API.PostV4Ordercreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Ordercreate`: OrderCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV4API.PostV4Ordercreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4OrdercreateRequest struct via the builder pattern


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


## PostV4Orderdelete

> OrderDeleteResponse PostV4Orderdelete(ctx).OrderDeleteRequest(orderDeleteRequest).Execute()

Delete an Order



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
	orderDeleteRequest := *openapiclient.NewOrderDeleteRequest() // OrderDeleteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV4API.PostV4Orderdelete(context.Background()).OrderDeleteRequest(orderDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV4API.PostV4Orderdelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Orderdelete`: OrderDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV4API.PostV4Orderdelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4OrderdeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderDeleteRequest** | [**OrderDeleteRequest**](OrderDeleteRequest.md) |  | 

### Return type

[**OrderDeleteResponse**](OrderDeleteResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV4Orderdetails

> OrderDetailResponse PostV4Orderdetails(ctx).OrderDetailRequest(orderDetailRequest).Execute()

Get Order Details



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
	orderDetailRequest := *openapiclient.NewOrderDetailRequest() // OrderDetailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV4API.PostV4Orderdetails(context.Background()).OrderDetailRequest(orderDetailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV4API.PostV4Orderdetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Orderdetails`: OrderDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV4API.PostV4Orderdetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4OrderdetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderDetailRequest** | [**OrderDetailRequest**](OrderDetailRequest.md) |  | 

### Return type

[**OrderDetailResponse**](OrderDetailResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV4Ordermodify

> OrderModifyResponse PostV4Ordermodify(ctx).OrderModifyRequest(orderModifyRequest).Execute()

Modify an Existing Order



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
	orderModifyRequest := *openapiclient.NewOrderModifyRequest() // OrderModifyRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV4API.PostV4Ordermodify(context.Background()).OrderModifyRequest(orderModifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV4API.PostV4Ordermodify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Ordermodify`: OrderModifyResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV4API.PostV4Ordermodify`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4OrdermodifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderModifyRequest** | [**OrderModifyRequest**](OrderModifyRequest.md) |  | 

### Return type

[**OrderModifyResponse**](OrderModifyResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV4Ordersearch

> OrderSearchResponse PostV4Ordersearch(ctx).OrderSearchRequest(orderSearchRequest).Execute()

Order Search



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
	orderSearchRequest := *openapiclient.NewOrderSearchRequest() // OrderSearchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersV4API.PostV4Ordersearch(context.Background()).OrderSearchRequest(orderSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersV4API.PostV4Ordersearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Ordersearch`: OrderSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `OrdersV4API.PostV4Ordersearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4OrdersearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderSearchRequest** | [**OrderSearchRequest**](OrderSearchRequest.md) |  | 

### Return type

[**OrderSearchResponse**](OrderSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

