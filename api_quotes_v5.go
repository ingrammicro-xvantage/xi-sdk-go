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
	"strings"
)


// QuotesV5APIService QuotesV5API service
type QuotesV5APIService service

type ApiGetV5QuotesDetailsRequest struct {
	ctx context.Context
	ApiService *QuotesV5APIService
	quoteNumber string
	customerNumber *string
	isoCountryCode *string
	thirdPartySource *string
}

// Your Ingram Micro unique customer number
func (r ApiGetV5QuotesDetailsRequest) CustomerNumber(customerNumber string) ApiGetV5QuotesDetailsRequest {
	r.customerNumber = &customerNumber
	return r
}

// 
func (r ApiGetV5QuotesDetailsRequest) IsoCountryCode(isoCountryCode string) ApiGetV5QuotesDetailsRequest {
	r.isoCountryCode = &isoCountryCode
	return r
}

// Unique identifier used to identify the third party source accessing the services
func (r ApiGetV5QuotesDetailsRequest) ThirdPartySource(thirdPartySource string) ApiGetV5QuotesDetailsRequest {
	r.thirdPartySource = &thirdPartySource
	return r
}

func (r ApiGetV5QuotesDetailsRequest) Execute() (*QuoteDetails, *http.Response, error) {
	return r.ApiService.GetV5QuotesDetailsExecute(r)
}

/*
GetV5QuotesDetails Get Quote Details

The quote details API provides all quote details associated with the quote number provided.

 The “<strong>quoteNumber</strong>”, “<strong>isoCountryCode</strong>” and “<strong>customerNumber</strong>” parameters are required.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param quoteNumber Ingram Micro Quote Number
 @return ApiGetV5QuotesDetailsRequest
*/
func (a *QuotesV5APIService) GetV5QuotesDetails(ctx context.Context, quoteNumber string) ApiGetV5QuotesDetailsRequest {
	return ApiGetV5QuotesDetailsRequest{
		ApiService: a,
		ctx: ctx,
		quoteNumber: quoteNumber,
	}
}

// Execute executes the request
//  @return QuoteDetails
func (a *QuotesV5APIService) GetV5QuotesDetailsExecute(r ApiGetV5QuotesDetailsRequest) (*QuoteDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesV5APIService.GetV5QuotesDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v5/quote/{quoteNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"quoteNumber"+"}", url.PathEscape(parameterValueToString(r.quoteNumber, "quoteNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerNumber == nil {
		return localVarReturnValue, nil, reportError("customerNumber is required and must be specified")
	}
	if r.isoCountryCode == nil {
		return localVarReturnValue, nil, reportError("isoCountryCode is required and must be specified")
	}
	if strlen(*r.isoCountryCode) < 2 {
		return localVarReturnValue, nil, reportError("isoCountryCode must have at least 2 elements")
	}
	if strlen(*r.isoCountryCode) > 2 {
		return localVarReturnValue, nil, reportError("isoCountryCode must have less than 2 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "customerNumber", r.customerNumber, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "isoCountryCode", r.isoCountryCode, "")
	if r.thirdPartySource != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "thirdPartySource", r.thirdPartySource, "")
	} else {
		var defaultValue string = "customer"
		r.thirdPartySource = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiPostV5QuotesSearchRequest struct {
	ctx context.Context
	ApiService *QuotesV5APIService
	quoteListRequest *QuoteListRequest
}

func (r ApiPostV5QuotesSearchRequest) QuoteListRequest(quoteListRequest QuoteListRequest) ApiPostV5QuotesSearchRequest {
	r.quoteListRequest = &quoteListRequest
	return r
}

func (r ApiPostV5QuotesSearchRequest) Execute() (*QuoteListResponse, *http.Response, error) {
	return r.ApiService.PostV5QuotesSearchExecute(r)
}

/*
PostV5QuotesSearch Search Quotes

This endpoint enables the retrieval and filtering of relevant quote list key criteria data, such as quote number, special bid numbers, end user name, status, and date ranges from the Ingram Micro system. By default, the Quotes endpoint retrieves quotes modified or created within the last 30 days.

 Observe these additional parameters:<ul><li>Only active quotes are available through this API.</li><li>Quotes older than 365 days are excluded by default.</li><li>You can use date range filters to retrieve quotes older than 30 days and up to 365 days.</li><li>Quotes that are in draft and closed states are excluded, and are not accessible through this API.</li></ul>

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostV5QuotesSearchRequest
*/
func (a *QuotesV5APIService) PostV5QuotesSearch(ctx context.Context) ApiPostV5QuotesSearchRequest {
	return ApiPostV5QuotesSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuoteListResponse
func (a *QuotesV5APIService) PostV5QuotesSearchExecute(r ApiPostV5QuotesSearchRequest) (*QuoteListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuoteListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuotesV5APIService.PostV5QuotesSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v5/quote/search"

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
	localVarPostBody = r.quoteListRequest
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
