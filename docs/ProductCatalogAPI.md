# \ProductCatalogAPI

All URIs are relative to *https://api.ingrammicro.com:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResellerV6Productdetail**](ProductCatalogAPI.md#GetResellerV6Productdetail) | **Get** /resellers/v6/catalog/details/{ingramPartNumber} | Product Details
[**GetResellerV6Productsearch**](ProductCatalogAPI.md#GetResellerV6Productsearch) | **Get** /resellers/v6/catalog | Search Products
[**PostPriceandavailability**](ProductCatalogAPI.md#PostPriceandavailability) | **Post** /resellers/v6/catalog/priceandavailability | Price and Availability



## GetResellerV6Productdetail

> ProductDetailResponse GetResellerV6Productdetail(ctx, ingramPartNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).Execute()

Product Details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	xi_sdk_resellers "https://github.com/dprakash2101/GoSDK" ingrammicro-xvantage/xi-sdk-resellers-go
)

func main() {
	ingramPartNumber := "6YE881" // string | Ingram Micro unique part number for the product
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems
	iMSenderID := "MyCompany" // string | Sender Identification text (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.GetResellerV6Productdetail(context.Background(), ingramPartNumber).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.GetResellerV6Productdetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellerV6Productdetail`: ProductDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.GetResellerV6Productdetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ingramPartNumber** | **string** | Ingram Micro unique part number for the product | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResellerV6ProductdetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems | 
 **iMSenderID** | **string** | Sender Identification text | 

### Return type

[**ProductDetailResponse**](ProductDetailResponse.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResellerV6Productsearch

> ProductSearchResponse GetResellerV6Productsearch(ctx).IMCustomerNumber(iMCustomerNumber).IMCorrelationID(iMCorrelationID).IMCountryCode(iMCountryCode).PageNumber(pageNumber).PageSize(pageSize).IMSenderID(iMSenderID).Type_(type_).HasDiscounts(hasDiscounts).Vendor(vendor).VendorPartNumber(vendorPartNumber).AcceptLanguage(acceptLanguage).VendorNumber(vendorNumber).Keyword(keyword).Category(category).SkipAuthorisation(skipAuthorisation).Execute()

Search Products



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	xi_sdk_resellers "https://github.com/dprakash2101/GoSDK" ingrammicro-xvantage/xi-sdk-resellers-go
)

func main() {
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction accross all the systems
	iMCountryCode := "US" // string | Two-character ISO country code.
	pageNumber := int32(56) // int32 | Current page number. Default is 1 (optional)
	pageSize := int32(56) // int32 | Number of records required in the call - max records 100 per page (optional)
	iMSenderID := "MyCompany" // string | Sender Identification text (optional)
	type_ := "type__example" // string | The SKU type of product. One of Physical, Digital, or Any. (optional)
	hasDiscounts := "true" // string | Specifies if there are discounts available for the product. (optional)
	vendor := []string{"Inner_example"} // []string | The name of the vendor/manufacturer of the product. (optional)
	vendorPartNumber := []string{"Inner_example"} // []string | The vendors part number for the product. (optional)
	acceptLanguage := "acceptLanguage_example" // string | Header to the API calls, the content will help us identify the response language. (optional) (default to "en")
	vendorNumber := "vendorNumber_example" // string | Vendor number of the product (optional)
	keyword := []string{"Inner_example"} // []string | Keyword search,can be ingram part number or vendor part number or product title or vendor nameKeyword search. Can be Ingram Micro part number, vender part number, product title, or vendor name. (optional)
	category := "Accessories" // string | The category of the product. Example: Displays. (optional)
	skipAuthorisation := "true" // string | This parameter is True when you want Skip the authorization, so template will work like current B2b template. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.GetResellerV6Productsearch(context.Background()).IMCustomerNumber(iMCustomerNumber).IMCorrelationID(iMCorrelationID).IMCountryCode(iMCountryCode).PageNumber(pageNumber).PageSize(pageSize).IMSenderID(iMSenderID).Type_(type_).HasDiscounts(hasDiscounts).Vendor(vendor).VendorPartNumber(vendorPartNumber).AcceptLanguage(acceptLanguage).VendorNumber(vendorNumber).Keyword(keyword).Category(category).SkipAuthorisation(skipAuthorisation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.GetResellerV6Productsearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResellerV6Productsearch`: ProductSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.GetResellerV6Productsearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResellerV6ProductsearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction accross all the systems | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **pageNumber** | **int32** | Current page number. Default is 1 | 
 **pageSize** | **int32** | Number of records required in the call - max records 100 per page | 
 **iMSenderID** | **string** | Sender Identification text | 
 **type_** | **string** | The SKU type of product. One of Physical, Digital, or Any. | 
 **hasDiscounts** | **string** | Specifies if there are discounts available for the product. | 
 **vendor** | **[]string** | The name of the vendor/manufacturer of the product. | 
 **vendorPartNumber** | **[]string** | The vendors part number for the product. | 
 **acceptLanguage** | **string** | Header to the API calls, the content will help us identify the response language. | [default to &quot;en&quot;]
 **vendorNumber** | **string** | Vendor number of the product | 
 **keyword** | **[]string** | Keyword search,can be ingram part number or vendor part number or product title or vendor nameKeyword search. Can be Ingram Micro part number, vender part number, product title, or vendor name. | 
 **category** | **string** | The category of the product. Example: Displays. | 
 **skipAuthorisation** | **string** | This parameter is True when you want Skip the authorization, so template will work like current B2b template. | 

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


## PostPriceandavailability

> []PriceAndAvailabilityResponseInner PostPriceandavailability(ctx).IncludeAvailability(includeAvailability).IncludePricing(includePricing).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).PriceAndAvailabilityRequest(priceAndAvailabilityRequest).IncludeProductAttributes(includeProductAttributes).IMSenderID(iMSenderID).Execute()

Price and Availability



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	xi_sdk_resellers "https://github.com/dprakash2101/GoSDK" ingrammicro-xvantage/xi-sdk-resellers-go
)

func main() {
	includeAvailability := true // bool | Pass boolean value as input, if true the response will contain warehouse availability details, if false the response will not hold warehouse availability details
	includePricing := true // bool | Pass boolean value as input, if true the response will contain Pricing details of the Product, if false the response will not hold Pricing details.
	iMCustomerNumber := "20-222222" // string | Your unique Ingram Micro customer number.
	iMCountryCode := "US" // string | Two-character ISO country code.
	iMCorrelationID := "fbac82ba-cf0a-4bcf-fc03-0c5084" // string | Unique transaction number to identify each transaction across all the systems.
	priceAndAvailabilityRequest := *openapiclient.NewPriceAndAvailabilityRequest() // PriceAndAvailabilityRequest | 
	includeProductAttributes := true // bool | Pass boolean value as input, if true the response will contain detailed attributes related to the Product, if false or not sent the response will contain very few Product details. (optional)
	iMSenderID := "MyCompany" // string | Unique value used to identify the sender of the transaction. Example: MyCompany (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductCatalogAPI.PostPriceandavailability(context.Background()).IncludeAvailability(includeAvailability).IncludePricing(includePricing).IMCustomerNumber(iMCustomerNumber).IMCountryCode(iMCountryCode).IMCorrelationID(iMCorrelationID).PriceAndAvailabilityRequest(priceAndAvailabilityRequest).IncludeProductAttributes(includeProductAttributes).IMSenderID(iMSenderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductCatalogAPI.PostPriceandavailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPriceandavailability`: []PriceAndAvailabilityResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProductCatalogAPI.PostPriceandavailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPriceandavailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeAvailability** | **bool** | Pass boolean value as input, if true the response will contain warehouse availability details, if false the response will not hold warehouse availability details | 
 **includePricing** | **bool** | Pass boolean value as input, if true the response will contain Pricing details of the Product, if false the response will not hold Pricing details. | 
 **iMCustomerNumber** | **string** | Your unique Ingram Micro customer number. | 
 **iMCountryCode** | **string** | Two-character ISO country code. | 
 **iMCorrelationID** | **string** | Unique transaction number to identify each transaction across all the systems. | 
 **priceAndAvailabilityRequest** | [**PriceAndAvailabilityRequest**](PriceAndAvailabilityRequest.md) |  | 
 **includeProductAttributes** | **bool** | Pass boolean value as input, if true the response will contain detailed attributes related to the Product, if false or not sent the response will contain very few Product details. | 
 **iMSenderID** | **string** | Unique value used to identify the sender of the transaction. Example: MyCompany | 

### Return type

[**[]PriceAndAvailabilityResponseInner**](PriceAndAvailabilityResponseInner.md)

### Authorization

[application](../README.md#application)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

