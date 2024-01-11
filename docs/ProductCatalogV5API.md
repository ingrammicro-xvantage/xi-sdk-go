# \ProductCatalogV5API

All URIs are relative to *https://api.ingrammicro.com:443/sandbox*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV5CatalogProductsearch**](ProductCatalogV5API.md#GetV5CatalogProductsearch) | **Get** /resellers/v5/Catalog | Search Product Catalog
[**MultiSKUPriceAndStock**](ProductCatalogV5API.md#MultiSKUPriceAndStock) | **Post** /resellers/v5/Catalog/priceandavailability | Find availability of upto 50 SKUs



## GetV5CatalogProductsearch

> ProductSearchResponse GetV5CatalogProductsearch(ctx).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).PartNumber(partNumber).Execute()

Search Product Catalog



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
	customerNumber := "customerNumber_example" // string | Your unique Ingram Micro customer number (default to "20-222222")
	isoCountryCode := "isoCountryCode_example" // string | 2 chars country code (default to "US")
	partNumber := "partNumber_example" // string | Part Number can be ingram part number or vendor part number or customer part number or UPC (default to "1AQ821")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogV5API.GetV5CatalogProductsearch(context.Background()).CustomerNumber(customerNumber).IsoCountryCode(isoCountryCode).PartNumber(partNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogV5API.GetV5CatalogProductsearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV5CatalogProductsearch`: ProductSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogV5API.GetV5CatalogProductsearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV5CatalogProductsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerNumber** | **string** | Your unique Ingram Micro customer number | [default to &quot;20-222222&quot;]
 **isoCountryCode** | **string** | 2 chars country code | [default to &quot;US&quot;]
 **partNumber** | **string** | Part Number can be ingram part number or vendor part number or customer part number or UPC | [default to &quot;1AQ821&quot;]

### Return type

[**ProductSearchResponse**](ProductSearchResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultiSKUPriceAndStock

> PriceAndAvailabilityResponse MultiSKUPriceAndStock(ctx).PriceAndAvailabilityRequest(priceAndAvailabilityRequest).Execute()

Find availability of upto 50 SKUs



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
	priceAndAvailabilityRequest := *openapiclient.NewPriceAndAvailabilityRequest() // PriceAndAvailabilityRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogV5API.MultiSKUPriceAndStock(context.Background()).PriceAndAvailabilityRequest(priceAndAvailabilityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogV5API.MultiSKUPriceAndStock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultiSKUPriceAndStock`: PriceAndAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogV5API.MultiSKUPriceAndStock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultiSKUPriceAndStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceAndAvailabilityRequest** | [**PriceAndAvailabilityRequest**](PriceAndAvailabilityRequest.md) |  | 

### Return type

[**PriceAndAvailabilityResponse**](PriceAndAvailabilityResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

