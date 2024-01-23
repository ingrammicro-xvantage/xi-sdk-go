# \ProductCatalogV4API

All URIs are relative to *https://api.ingrammicro.com:443/sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV4Multiskupriceandstock**](ProductCatalogV4API.md#PostV4Multiskupriceandstock) | **Post** /products/v4/multiskupriceandstock | Product availability for upto 50 SKUs
[**PostV4Productsearch**](ProductCatalogV4API.md#PostV4Productsearch) | **Post** /products/v4/productsearch | Real-time Product Search



## PostV4Multiskupriceandstock

> MultiSKUPriceAndStockResponse PostV4Multiskupriceandstock(ctx).MultiSKUPriceAndStockRequest(multiSKUPriceAndStockRequest).Execute()

Product availability for upto 50 SKUs



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
	multiSKUPriceAndStockRequest := *openapiclient.NewMultiSKUPriceAndStockRequest() // MultiSKUPriceAndStockRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogV4API.PostV4Multiskupriceandstock(context.Background()).MultiSKUPriceAndStockRequest(multiSKUPriceAndStockRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogV4API.PostV4Multiskupriceandstock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Multiskupriceandstock`: MultiSKUPriceAndStockResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogV4API.PostV4Multiskupriceandstock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4MultiskupriceandstockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiSKUPriceAndStockRequest** | [**MultiSKUPriceAndStockRequest**](MultiSKUPriceAndStockRequest.md) |  | 

### Return type

[**MultiSKUPriceAndStockResponse**](MultiSKUPriceAndStockResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV4Productsearch

> ProductSearchResponse PostV4Productsearch(ctx).ProductSearchRequest(productSearchRequest).Execute()

Real-time Product Search



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
	productSearchRequest := *openapiclient.NewProductSearchRequest() // ProductSearchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogV4API.PostV4Productsearch(context.Background()).ProductSearchRequest(productSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogV4API.PostV4Productsearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV4Productsearch`: ProductSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogV4API.PostV4Productsearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV4ProductsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productSearchRequest** | [**ProductSearchRequest**](ProductSearchRequest.md) |  | 

### Return type

[**ProductSearchResponse**](ProductSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

