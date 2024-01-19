/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ProductCatalogV4APIService ProductCatalogV4API service
type ProductCatalogV4APIService service

type ApiPostV4MultiskupriceandstockRequest struct {
	ctx context.Context
	ApiService *ProductCatalogV4APIService
	multiSKUPriceAndStockRequest *MultiSKUPriceAndStockRequest
}

func (r ApiPostV4MultiskupriceandstockRequest) MultiSKUPriceAndStockRequest(multiSKUPriceAndStockRequest MultiSKUPriceAndStockRequest) ApiPostV4MultiskupriceandstockRequest {
	r.multiSKUPriceAndStockRequest = &multiSKUPriceAndStockRequest
	return r
}

func (r ApiPostV4MultiskupriceandstockRequest) Execute() (*MultiSKUPriceAndStockResponse, *http.Response, error) {
	return r.ApiService.PostV4MultiskupriceandstockExecute(r)
}

/*
PostV4Multiskupriceandstock Product availability for upto 50 SKUs

Find price and availability of up to 50 SKUs in a single request. As you increase the number of items in the request response time will be extended. This transaction must not be used as a continuous cyclical call to populate availability and pricing for your full catalog. Customers that perform this activity will lose access to price and availability.

Ingram can provide a Price catalog file and an Inventory file in flat file format, which can be obtained through FTP download. Please contact 1800-616-4665 or Electronic.Services@ingrammicro.com for more information on these files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostV4MultiskupriceandstockRequest
*/
func (a *ProductCatalogV4APIService) PostV4Multiskupriceandstock(ctx context.Context) ApiPostV4MultiskupriceandstockRequest {
	return ApiPostV4MultiskupriceandstockRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MultiSKUPriceAndStockResponse
func (a *ProductCatalogV4APIService) PostV4MultiskupriceandstockExecute(r ApiPostV4MultiskupriceandstockRequest) (*MultiSKUPriceAndStockResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MultiSKUPriceAndStockResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductCatalogV4APIService.PostV4Multiskupriceandstock")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/v4/multiskupriceandstock"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.multiSKUPriceAndStockRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostV4ProductsearchRequest struct {
	ctx context.Context
	ApiService *ProductCatalogV4APIService
	productSearchRequest *ProductSearchRequest
}

func (r ApiPostV4ProductsearchRequest) ProductSearchRequest(productSearchRequest ProductSearchRequest) ApiPostV4ProductsearchRequest {
	r.productSearchRequest = &productSearchRequest
	return r
}

func (r ApiPostV4ProductsearchRequest) Execute() (*ProductSearchResponse, *http.Response, error) {
	return r.ApiService.PostV4ProductsearchExecute(r)
}

/*
PostV4Productsearch Real-time Product Search

A real time search that provides the Ingram Micro part number using the manufacturer part number.  This API is helpful to eliminate any errors when a manufactuer has the same part number and Ingram Micro has had to create multiple sku numbers 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostV4ProductsearchRequest
*/
func (a *ProductCatalogV4APIService) PostV4Productsearch(ctx context.Context) ApiPostV4ProductsearchRequest {
	return ApiPostV4ProductsearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductSearchResponse
func (a *ProductCatalogV4APIService) PostV4ProductsearchExecute(r ApiPostV4ProductsearchRequest) (*ProductSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductCatalogV4APIService.PostV4Productsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/products/v4/productsearch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.productSearchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
