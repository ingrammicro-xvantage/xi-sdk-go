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


// InvoicesV4APIService InvoicesV4API service
type InvoicesV4APIService service

type ApiPostV4InvoicedetailsRequest struct {
	ctx context.Context
	ApiService *InvoicesV4APIService
	invoiceDetailRequest *InvoiceDetailRequest
}

func (r ApiPostV4InvoicedetailsRequest) InvoiceDetailRequest(invoiceDetailRequest InvoiceDetailRequest) ApiPostV4InvoicedetailsRequest {
	r.invoiceDetailRequest = &invoiceDetailRequest
	return r
}

func (r ApiPostV4InvoicedetailsRequest) Execute() (*InvoiceDetailResponse, *http.Response, error) {
	return r.ApiService.PostV4InvoicedetailsExecute(r)
}

/*
PostV4Invoicedetails Get Invoice Details

A real-time request that allows the customer to query Ingram Micro for Invoice information for a specific open or shipped order (in the past 9 months). Orders are searched using Ingram Micro Sales Order Number.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostV4InvoicedetailsRequest
*/
func (a *InvoicesV4APIService) PostV4Invoicedetails(ctx context.Context) ApiPostV4InvoicedetailsRequest {
	return ApiPostV4InvoicedetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InvoiceDetailResponse
func (a *InvoicesV4APIService) PostV4InvoicedetailsExecute(r ApiPostV4InvoicedetailsRequest) (*InvoiceDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InvoiceDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesV4APIService.PostV4Invoicedetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/invoices/v4/invoicedetails"

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
	localVarPostBody = r.invoiceDetailRequest
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
