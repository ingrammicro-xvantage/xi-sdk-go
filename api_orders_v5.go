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
	"time"
)


// OrdersV5APIService OrdersV5API service
type OrdersV5APIService service

type ApiDeleteOrdersOrderNumberRequest struct {
	ctx context.Context
	ApiService *OrdersV5APIService
	ordernumber string
	customerNumber *string
	isoCountryCode *string
	entryDate *string
}

// Your unique Ingram Micro customer number
func (r ApiDeleteOrdersOrderNumberRequest) CustomerNumber(customerNumber string) ApiDeleteOrdersOrderNumberRequest {
	r.customerNumber = &customerNumber
	return r
}

// 2 chars ISO country code
func (r ApiDeleteOrdersOrderNumberRequest) IsoCountryCode(isoCountryCode string) ApiDeleteOrdersOrderNumberRequest {
	r.isoCountryCode = &isoCountryCode
	return r
}

// Order entry date (yyyy-mm-dd)
func (r ApiDeleteOrdersOrderNumberRequest) EntryDate(entryDate string) ApiDeleteOrdersOrderNumberRequest {
	r.entryDate = &entryDate
	return r
}

func (r ApiDeleteOrdersOrderNumberRequest) Execute() (*OrderCancelResponse, *http.Response, error) {
	return r.ApiService.DeleteOrdersOrderNumberExecute(r)
}

/*
DeleteOrdersOrderNumber Cancel an Existing Order

This endpoint is a request to cancel a previously accepted order. Use your Ingram Micro sales order number to cancel an order.

 The <strong>orderNumber, isoCountryCode, customerNumber</strong> and <strong>entryDate</strong> parameters are required.

 This call must be submitted <strong>before</strong> the order is released to Ingram Micro’s warehouse. The order cannot be canceled once it is released to the warehouse.

 Direct ship orders cannot be canceled. Contact your Ingram Micro sales rep for assistance.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ordernumber Ingram Micro sales order number
 @return ApiDeleteOrdersOrderNumberRequest
*/
func (a *OrdersV5APIService) DeleteOrdersOrderNumber(ctx context.Context, ordernumber string) ApiDeleteOrdersOrderNumberRequest {
	return ApiDeleteOrdersOrderNumberRequest{
		ApiService: a,
		ctx: ctx,
		ordernumber: ordernumber,
	}
}

// Execute executes the request
//  @return OrderCancelResponse
func (a *OrdersV5APIService) DeleteOrdersOrderNumberExecute(r ApiDeleteOrdersOrderNumberRequest) (*OrderCancelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderCancelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV5APIService.DeleteOrdersOrderNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v5/Orders/{ordernumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"ordernumber"+"}", url.PathEscape(parameterValueToString(r.ordernumber, "ordernumber")), -1)

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
	if r.entryDate == nil {
		return localVarReturnValue, nil, reportError("entryDate is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "customerNumber", r.customerNumber, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "isoCountryCode", r.isoCountryCode, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "entryDate", r.entryDate, "")
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

type ApiGetOrdersSearchRequest struct {
	ctx context.Context
	ApiService *OrdersV5APIService
	customerNumber *string
	isocountrycode *string
	ordernumber *string
	customerordernumber *string
	orderstatus *string
	startcreatetimestamp *time.Time
	endcreatetimestamp *time.Time
	pagesize *int32
	pagenumber *int32
}

// Your unique Ingram Micro customer number
func (r ApiGetOrdersSearchRequest) CustomerNumber(customerNumber string) ApiGetOrdersSearchRequest {
	r.customerNumber = &customerNumber
	return r
}

// 2 char iso country code
func (r ApiGetOrdersSearchRequest) Isocountrycode(isocountrycode string) ApiGetOrdersSearchRequest {
	r.isocountrycode = &isocountrycode
	return r
}

// Ingram sales order number
func (r ApiGetOrdersSearchRequest) Ordernumber(ordernumber string) ApiGetOrdersSearchRequest {
	r.ordernumber = &ordernumber
	return r
}

// Search using your PO/Order number
func (r ApiGetOrdersSearchRequest) Customerordernumber(customerordernumber string) ApiGetOrdersSearchRequest {
	r.customerordernumber = &customerordernumber
	return r
}

// Ingram Micro order status
func (r ApiGetOrdersSearchRequest) Orderstatus(orderstatus string) ApiGetOrdersSearchRequest {
	r.orderstatus = &orderstatus
	return r
}

// Search start date/time in UTC format
func (r ApiGetOrdersSearchRequest) Startcreatetimestamp(startcreatetimestamp time.Time) ApiGetOrdersSearchRequest {
	r.startcreatetimestamp = &startcreatetimestamp
	return r
}

// Search end date/time in UTC format
func (r ApiGetOrdersSearchRequest) Endcreatetimestamp(endcreatetimestamp time.Time) ApiGetOrdersSearchRequest {
	r.endcreatetimestamp = &endcreatetimestamp
	return r
}

// Number of records required in the call
func (r ApiGetOrdersSearchRequest) Pagesize(pagesize int32) ApiGetOrdersSearchRequest {
	r.pagesize = &pagesize
	return r
}

// the page number reference
func (r ApiGetOrdersSearchRequest) Pagenumber(pagenumber int32) ApiGetOrdersSearchRequest {
	r.pagenumber = &pagenumber
	return r
}

func (r ApiGetOrdersSearchRequest) Execute() (*OrderSearchResponse, *http.Response, error) {
	return r.ApiService.GetOrdersSearchExecute(r)
}

/*
GetOrdersSearch Search your Orders

Search your Ingram Micro orders. This endpoint searches by multiple order parameters and supports pagination of results. Search using one or more of the parameters below:

 <ul><li>ordernumber — Ingram Micro sales order number</li><li>customerordernumber — The PO or order number provided by you when creating an order</li><li>orderstatus — user order status codes for the search, default is set to "any"</li><li>startcreatetimestamp and endcreatetimestamp — Order create date range</li></ul>

 For pagination, please use these parameters:
 <ul><li>pagesize — default 25, max 100</li><li>pagenumber — default 1</li></ul>

 Order Status Values:
 <ul><li>P – PENDING</li><li>R – RELEASED</li><li>4 – SHIPPED</li><li>I – INVOICED</li><li>V – VOIDED</li></ul>

 The search endpoint also returns HATEOAS links for order details and invoice details, if applicable.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOrdersSearchRequest
*/
func (a *OrdersV5APIService) GetOrdersSearch(ctx context.Context) ApiGetOrdersSearchRequest {
	return ApiGetOrdersSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OrderSearchResponse
func (a *OrdersV5APIService) GetOrdersSearchExecute(r ApiGetOrdersSearchRequest) (*OrderSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV5APIService.GetOrdersSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v5/Orders/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerNumber == nil {
		return localVarReturnValue, nil, reportError("customerNumber is required and must be specified")
	}
	if r.isocountrycode == nil {
		return localVarReturnValue, nil, reportError("isocountrycode is required and must be specified")
	}
	if strlen(*r.isocountrycode) < 2 {
		return localVarReturnValue, nil, reportError("isocountrycode must have at least 2 elements")
	}
	if strlen(*r.isocountrycode) > 2 {
		return localVarReturnValue, nil, reportError("isocountrycode must have less than 2 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "customerNumber", r.customerNumber, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "isocountrycode", r.isocountrycode, "")
	if r.ordernumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordernumber", r.ordernumber, "")
	}
	if r.customerordernumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerordernumber", r.customerordernumber, "")
	}
	if r.orderstatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderstatus", r.orderstatus, "")
	} else {
		var defaultValue string = "any"
		r.orderstatus = &defaultValue
	}
	if r.startcreatetimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startcreatetimestamp", r.startcreatetimestamp, "")
	}
	if r.endcreatetimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endcreatetimestamp", r.endcreatetimestamp, "")
	}
	if r.pagesize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagesize", r.pagesize, "")
	}
	if r.pagenumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pagenumber", r.pagenumber, "")
	} else {
		var defaultValue int32 = 1
		r.pagenumber = &defaultValue
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

type ApiGetV5OrdersDetailsRequest struct {
	ctx context.Context
	ApiService *OrdersV5APIService
	ordernumber string
	customernumber *string
	isocountrycode *string
	customerordernumber *string
	startcreatetimestamp *string
	endcreatetimestamp *string
	simulate *string
}

// Your unique Ingram Micro customer number
func (r ApiGetV5OrdersDetailsRequest) Customernumber(customernumber string) ApiGetV5OrdersDetailsRequest {
	r.customernumber = &customernumber
	return r
}

// 2 chars ISO country code
func (r ApiGetV5OrdersDetailsRequest) Isocountrycode(isocountrycode string) ApiGetV5OrdersDetailsRequest {
	r.isocountrycode = &isocountrycode
	return r
}

// Your PO/Order Number provide at the time of order creation
func (r ApiGetV5OrdersDetailsRequest) Customerordernumber(customerordernumber string) ApiGetV5OrdersDetailsRequest {
	r.customerordernumber = &customerordernumber
	return r
}

// Filter start date - format YYYY-MM-DD
func (r ApiGetV5OrdersDetailsRequest) Startcreatetimestamp(startcreatetimestamp string) ApiGetV5OrdersDetailsRequest {
	r.startcreatetimestamp = &startcreatetimestamp
	return r
}

// Filter end date - format YYYY-MM-DD
func (r ApiGetV5OrdersDetailsRequest) Endcreatetimestamp(endcreatetimestamp string) ApiGetV5OrdersDetailsRequest {
	r.endcreatetimestamp = &endcreatetimestamp
	return r
}

// Order response for various order statuses. Not for use in production.
func (r ApiGetV5OrdersDetailsRequest) Simulate(simulate string) ApiGetV5OrdersDetailsRequest {
	r.simulate = &simulate
	return r
}

func (r ApiGetV5OrdersDetailsRequest) Execute() (*OrderDetailResponse, *http.Response, error) {
	return r.ApiService.GetV5OrdersDetailsExecute(r)
}

/*
GetV5OrdersDetails Get Order Details

Use your Ingram Micro sales order number to search for existing orders or retrieve existing order details.

 <b>The sales order number, customer number and isoCountryCode are required parameters.</b>

 The sales order number is returned in the Order Create POST response. Ingram Micro recommends that you save this number for future uses.

 The IM sales order number can also be retrieved by searching for your existing order using the Order Search GET endpoint. You will need the customer PO number or order number that was provided at the time of order creation.

 In a case when the IM sales order number is repeated, you can refine the result by providing customer order number for additional filtering or using the date range to filter orders by creation date.

 Use the "simulate" query parameter to test the GET order response for various order statuses. This parameter is only available in the sandbox to help with development and testing of the GET order endpoint.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ordernumber Ingram Micro sales order number
 @return ApiGetV5OrdersDetailsRequest
*/
func (a *OrdersV5APIService) GetV5OrdersDetails(ctx context.Context, ordernumber string) ApiGetV5OrdersDetailsRequest {
	return ApiGetV5OrdersDetailsRequest{
		ApiService: a,
		ctx: ctx,
		ordernumber: ordernumber,
	}
}

// Execute executes the request
//  @return OrderDetailResponse
func (a *OrdersV5APIService) GetV5OrdersDetailsExecute(r ApiGetV5OrdersDetailsRequest) (*OrderDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV5APIService.GetV5OrdersDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v5/Orders/{ordernumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"ordernumber"+"}", url.PathEscape(parameterValueToString(r.ordernumber, "ordernumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customernumber == nil {
		return localVarReturnValue, nil, reportError("customernumber is required and must be specified")
	}
	if r.isocountrycode == nil {
		return localVarReturnValue, nil, reportError("isocountrycode is required and must be specified")
	}
	if strlen(*r.isocountrycode) < 2 {
		return localVarReturnValue, nil, reportError("isocountrycode must have at least 2 elements")
	}
	if strlen(*r.isocountrycode) > 2 {
		return localVarReturnValue, nil, reportError("isocountrycode must have less than 2 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "customernumber", r.customernumber, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "isocountrycode", r.isocountrycode, "")
	if r.customerordernumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customerordernumber", r.customerordernumber, "")
	}
	if r.startcreatetimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startcreatetimestamp", r.startcreatetimestamp, "")
	}
	if r.endcreatetimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endcreatetimestamp", r.endcreatetimestamp, "")
	}
	if r.simulate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "simulate", r.simulate, "")
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

type ApiPostV5OrdersCreateRequest struct {
	ctx context.Context
	ApiService *OrdersV5APIService
	orderCreateRequest *OrderCreateRequest
}

func (r ApiPostV5OrdersCreateRequest) OrderCreateRequest(orderCreateRequest OrderCreateRequest) ApiPostV5OrdersCreateRequest {
	r.orderCreateRequest = &orderCreateRequest
	return r
}

func (r ApiPostV5OrdersCreateRequest) Execute() (*OrderCreateResponse, *http.Response, error) {
	return r.ApiService.PostV5OrdersCreateExecute(r)
}

/*
PostV5OrdersCreate Create a New Order

Instantly create and place orders. The POST API supports stocked SKUs as well as licensing and warranties SKUs.

 Every order to be created with this API must complete these validations to be placed and processed:<ul><li>SKU, shipping address, product authorization and stock allocations must clear validation.</li><li>Ingram Micro Sales validates pricing, stock or other processing parameters. Ingram Micro sales may place an order a hold if revision is necessary.</li><li>Credit validation confirms available credit prior to processing an order. If an order does not clear credit validation, the Ingram Micro sales rep or accounts receivable manager will contact you for next steps.</li><li>Warehouse validation selects the location closest to the destination zip code. If the stock is not available in any of the warehouses, Ingram Micro places a backorder in the warehouse closest to the destination zip code.</li></ul>

 Ingram Micro recommends that you provide the <strong>ingrampartnumber</strong> for each SKU contained in each order.

 When using <strong>vendorpartnumber</strong> to place an order, please use the product search endpoint to find the <strong>ingrampartnumber</strong> for a specific <strong>vendorpartnumber</strong>, and then supply the <strong>ingrampartnumber</strong> to place an order.

 <strong>NOTE:</strong> You must have net terms to use the <strong>Ingram Micro Order Create API</strong>. Ingram Micro offers trade credit when using our APIs, and repayment is based on net terms. For example, if your net terms agreement is net-30, you will have 30 days to make a full payment. Ingram Micro does not allow credit card transactions for API ordering.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostV5OrdersCreateRequest
*/
func (a *OrdersV5APIService) PostV5OrdersCreate(ctx context.Context) ApiPostV5OrdersCreateRequest {
	return ApiPostV5OrdersCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OrderCreateResponse
func (a *OrdersV5APIService) PostV5OrdersCreateExecute(r ApiPostV5OrdersCreateRequest) (*OrderCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrderCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersV5APIService.PostV5OrdersCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v5/Orders"

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
	localVarPostBody = r.orderCreateRequest
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
