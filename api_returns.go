/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ReturnsAPIService ReturnsAPI service
type ReturnsAPIService service

type ApiGetResellersV6ReturnsdetailsRequest struct {
	ctx context.Context
	ApiService *ReturnsAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	caseRequestNumber string
	iMSenderID *string
}

// Your unique Ingram Micro customer number.
func (r ApiGetResellersV6ReturnsdetailsRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellersV6ReturnsdetailsRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellersV6ReturnsdetailsRequest) IMCountryCode(iMCountryCode string) ApiGetResellersV6ReturnsdetailsRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetResellersV6ReturnsdetailsRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellersV6ReturnsdetailsRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiGetResellersV6ReturnsdetailsRequest) IMSenderID(iMSenderID string) ApiGetResellersV6ReturnsdetailsRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiGetResellersV6ReturnsdetailsRequest) Execute() (*ReturnsDetailsResponse, *http.Response, error) {
	return r.ApiService.GetResellersV6ReturnsdetailsExecute(r)
}

/*
GetResellersV6Returnsdetails Returns Details

The Returns Details API will retrieve all the details related to the specific caseRequestNumber.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param caseRequestNumber A unique return request number.
 @return ApiGetResellersV6ReturnsdetailsRequest
*/
func (a *ReturnsAPIService) GetResellersV6Returnsdetails(ctx context.Context, caseRequestNumber string) ApiGetResellersV6ReturnsdetailsRequest {
	return ApiGetResellersV6ReturnsdetailsRequest{
		ApiService: a,
		ctx: ctx,
		caseRequestNumber: caseRequestNumber,
	}
}

// Execute executes the request
//  @return ReturnsDetailsResponse
func (a *ReturnsAPIService) GetResellersV6ReturnsdetailsExecute(r ApiGetResellersV6ReturnsdetailsRequest) (*ReturnsDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnsDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsAPIService.GetResellersV6Returnsdetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/returns/{caseRequestNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"caseRequestNumber"+"}", url.PathEscape(parameterValueToString(r.caseRequestNumber, "caseRequestNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMCustomerNumber == nil {
		return localVarReturnValue, nil, reportError("iMCustomerNumber is required and must be specified")
	}
	if strlen(*r.iMCustomerNumber) > 10 {
		return localVarReturnValue, nil, reportError("iMCustomerNumber must have less than 10 elements")
	}
	if r.iMCountryCode == nil {
		return localVarReturnValue, nil, reportError("iMCountryCode is required and must be specified")
	}
	if strlen(*r.iMCountryCode) < 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have at least 2 elements")
	}
	if strlen(*r.iMCountryCode) > 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have less than 2 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v PostRenewalssearch400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v PostCreateorderV7500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetResellersV6ReturnssearchRequest struct {
	ctx context.Context
	ApiService *ReturnsAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMSenderID *string
	caseRequestNumber *string
	invoiceNumber *string
	returnClaimId *string
	referenceNumber *string
	ingramPartNumber *string
	vendorPartNumber *string
	returnStatusIn *string
	claimStatusIn *string
	createdOnBt *string
	modifiedOnBt *string
	returnReasonIn *string
	page *string
	size *string
	sort *string
	sortingColumnName *string
}

// Your unique Ingram Micro customer number.
func (r ApiGetResellersV6ReturnssearchRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellersV6ReturnssearchRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellersV6ReturnssearchRequest) IMCountryCode(iMCountryCode string) ApiGetResellersV6ReturnssearchRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetResellersV6ReturnssearchRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellersV6ReturnssearchRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiGetResellersV6ReturnssearchRequest) IMSenderID(iMSenderID string) ApiGetResellersV6ReturnssearchRequest {
	r.iMSenderID = &iMSenderID
	return r
}

// A unique return request number.
func (r ApiGetResellersV6ReturnssearchRequest) CaseRequestNumber(caseRequestNumber string) ApiGetResellersV6ReturnssearchRequest {
	r.caseRequestNumber = &caseRequestNumber
	return r
}

// The Invoice number for the order.
func (r ApiGetResellersV6ReturnssearchRequest) InvoiceNumber(invoiceNumber string) ApiGetResellersV6ReturnssearchRequest {
	r.invoiceNumber = &invoiceNumber
	return r
}

// A unique return claim Id.
func (r ApiGetResellersV6ReturnssearchRequest) ReturnClaimId(returnClaimId string) ApiGetResellersV6ReturnssearchRequest {
	r.returnClaimId = &returnClaimId
	return r
}

// The reference number for the return.
func (r ApiGetResellersV6ReturnssearchRequest) ReferenceNumber(referenceNumber string) ApiGetResellersV6ReturnssearchRequest {
	r.referenceNumber = &referenceNumber
	return r
}

// Unique IngramMicro part number.
func (r ApiGetResellersV6ReturnssearchRequest) IngramPartNumber(ingramPartNumber string) ApiGetResellersV6ReturnssearchRequest {
	r.ingramPartNumber = &ingramPartNumber
	return r
}

// The vendor&#39;s part number.
func (r ApiGetResellersV6ReturnssearchRequest) VendorPartNumber(vendorPartNumber string) ApiGetResellersV6ReturnssearchRequest {
	r.vendorPartNumber = &vendorPartNumber
	return r
}

// Comma-separated values of pre-defined status. Open, Approved, Partially Approved, Denied, Voided.
func (r ApiGetResellersV6ReturnssearchRequest) ReturnStatusIn(returnStatusIn string) ApiGetResellersV6ReturnssearchRequest {
	r.returnStatusIn = &returnStatusIn
	return r
}

// Comma-separated values of pre-defined status. Open, Approved, Partially Approved, Denied, Voided.
func (r ApiGetResellersV6ReturnssearchRequest) ClaimStatusIn(claimStatusIn string) ApiGetResellersV6ReturnssearchRequest {
	r.claimStatusIn = &claimStatusIn
	return r
}

// The date on which the return request was created. 
func (r ApiGetResellersV6ReturnssearchRequest) CreatedOnBt(createdOnBt string) ApiGetResellersV6ReturnssearchRequest {
	r.createdOnBt = &createdOnBt
	return r
}

// The date on which the return request was last updated.
func (r ApiGetResellersV6ReturnssearchRequest) ModifiedOnBt(modifiedOnBt string) ApiGetResellersV6ReturnssearchRequest {
	r.modifiedOnBt = &modifiedOnBt
	return r
}

// Comma separated Pre-defined value. test, (EW) Express Warehousing, (AR) Account Receivables, (BB) Buy Back, (BE) Stock Balance Exception, (BO) Bill Only, (CE) Credit Dept Use - Credit Exception, (CF) Configuration Fee, (CS ) Customer Service Discretion, (CS1) Customer Service Discretion CS Error, (DE) Defective Exception, (DF) Defective Items, (DI) Direct Credit, (DM) Damaged from Carrier, (DN) Damaged Without Product, (DT) Direct/ Special Order, (DT1) Direct Ship billed, not shipped., (FO) Freight Out, (FX) No-Scan, (IN) Incomplete, (LS) Lost Shipment, (MN) Minimum Order Fee Credit, (OS) Over Shipment, (PR) Pricing Error, (RF) Refusal Credit, (RI) Re-Invoice, (RP) Return For Repair, (RT) Return Not Credited, (RTD) RCN, (SB) Stock Balance, (SD) Sales Discretion, (SH) Incorrect Shipping And Handling, (SS) Short Shipment, (SSK) Short Ship kit, (SW) Sales Writeoff, (TE) Opened Return, (TR) Training Refund, (TX) Tax Credit, (WS) Wrong Sales Sealed, (WW) Wrong Warehouse, (FS) Warehouse Failed Serial# Capture, Latin America Vebdor Credits, Select Source, ITAD - Trade-in Credit, Withholding Tax
func (r ApiGetResellersV6ReturnssearchRequest) ReturnReasonIn(returnReasonIn string) ApiGetResellersV6ReturnssearchRequest {
	r.returnReasonIn = &returnReasonIn
	return r
}

// Number of page.
func (r ApiGetResellersV6ReturnssearchRequest) Page(page string) ApiGetResellersV6ReturnssearchRequest {
	r.page = &page
	return r
}

// The submitted pagesize, default is 25
func (r ApiGetResellersV6ReturnssearchRequest) Size(size string) ApiGetResellersV6ReturnssearchRequest {
	r.size = &size
	return r
}

// Refers to the column selected to apply the sorting criteria.
func (r ApiGetResellersV6ReturnssearchRequest) Sort(sort string) ApiGetResellersV6ReturnssearchRequest {
	r.sort = &sort
	return r
}

// The column name which will be sorted on.
func (r ApiGetResellersV6ReturnssearchRequest) SortingColumnName(sortingColumnName string) ApiGetResellersV6ReturnssearchRequest {
	r.sortingColumnName = &sortingColumnName
	return r
}

func (r ApiGetResellersV6ReturnssearchRequest) Execute() (*ReturnsSearchResponse, *http.Response, error) {
	return r.ApiService.GetResellersV6ReturnssearchExecute(r)
}

/*
GetResellersV6Returnssearch Returns Search

The Returns Search API, by default, will retrieve all the returns that are associated with the customer’s account. The customer will be able to search returns using the query parameters. The Returns Search response will return the following information:  returnClaimId, caseRequestNumber, createdOn, referenceNumber, and returnReason.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResellersV6ReturnssearchRequest
*/
func (a *ReturnsAPIService) GetResellersV6Returnssearch(ctx context.Context) ApiGetResellersV6ReturnssearchRequest {
	return ApiGetResellersV6ReturnssearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnsSearchResponse
func (a *ReturnsAPIService) GetResellersV6ReturnssearchExecute(r ApiGetResellersV6ReturnssearchRequest) (*ReturnsSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnsSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsAPIService.GetResellersV6Returnssearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/returns/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMCustomerNumber == nil {
		return localVarReturnValue, nil, reportError("iMCustomerNumber is required and must be specified")
	}
	if strlen(*r.iMCustomerNumber) > 10 {
		return localVarReturnValue, nil, reportError("iMCustomerNumber must have less than 10 elements")
	}
	if r.iMCountryCode == nil {
		return localVarReturnValue, nil, reportError("iMCountryCode is required and must be specified")
	}
	if strlen(*r.iMCountryCode) < 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have at least 2 elements")
	}
	if strlen(*r.iMCountryCode) > 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have less than 2 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
	}

	if r.caseRequestNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "caseRequestNumber", r.caseRequestNumber, "")
	}
	if r.invoiceNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "invoiceNumber", r.invoiceNumber, "")
	}
	if r.returnClaimId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "returnClaimId", r.returnClaimId, "")
	}
	if r.referenceNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "referenceNumber", r.referenceNumber, "")
	}
	if r.ingramPartNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ingramPartNumber", r.ingramPartNumber, "")
	}
	if r.vendorPartNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vendorPartNumber", r.vendorPartNumber, "")
	}
	if r.returnStatusIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "returnStatus-in", r.returnStatusIn, "")
	}
	if r.claimStatusIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "claimStatus-in", r.claimStatusIn, "")
	}
	if r.createdOnBt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdOn-bt", r.createdOnBt, "")
	}
	if r.modifiedOnBt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "modifiedOn-bt", r.modifiedOnBt, "")
	}
	if r.returnReasonIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "returnReason-in", r.returnReasonIn, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.size != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "size", r.size, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	}
	if r.sortingColumnName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortingColumnName", r.sortingColumnName, "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v PostRenewalssearch400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v PostCreateorderV7500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPostReturnscreateRequest struct {
	ctx context.Context
	ApiService *ReturnsAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMSenderID *string
	returnsCreateRequest *ReturnsCreateRequest
}

// Your unique Ingram Micro customer number.
func (r ApiPostReturnscreateRequest) IMCustomerNumber(iMCustomerNumber string) ApiPostReturnscreateRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiPostReturnscreateRequest) IMCountryCode(iMCountryCode string) ApiPostReturnscreateRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiPostReturnscreateRequest) IMCorrelationID(iMCorrelationID string) ApiPostReturnscreateRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiPostReturnscreateRequest) IMSenderID(iMSenderID string) ApiPostReturnscreateRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiPostReturnscreateRequest) ReturnsCreateRequest(returnsCreateRequest ReturnsCreateRequest) ApiPostReturnscreateRequest {
	r.returnsCreateRequest = &returnsCreateRequest
	return r
}

func (r ApiPostReturnscreateRequest) Execute() (*ReturnsCreateResponse, *http.Response, error) {
	return r.ApiService.PostReturnscreateExecute(r)
}

/*
PostReturnscreate Returns Create

Return create endpoint will allow customers to create a return request. In order to create a request, the customer must provide either ingramPartNumber or vendorPartNumber along with other required fields listed below. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostReturnscreateRequest
*/
func (a *ReturnsAPIService) PostReturnscreate(ctx context.Context) ApiPostReturnscreateRequest {
	return ApiPostReturnscreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReturnsCreateResponse
func (a *ReturnsAPIService) PostReturnscreateExecute(r ApiPostReturnscreateRequest) (*ReturnsCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReturnsCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsAPIService.PostReturnscreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/returns/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMCustomerNumber == nil {
		return localVarReturnValue, nil, reportError("iMCustomerNumber is required and must be specified")
	}
	if strlen(*r.iMCustomerNumber) > 10 {
		return localVarReturnValue, nil, reportError("iMCustomerNumber must have less than 10 elements")
	}
	if r.iMCountryCode == nil {
		return localVarReturnValue, nil, reportError("iMCountryCode is required and must be specified")
	}
	if strlen(*r.iMCountryCode) < 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have at least 2 elements")
	}
	if strlen(*r.iMCountryCode) > 2 {
		return localVarReturnValue, nil, reportError("iMCountryCode must have less than 2 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
	}
	// body params
	localVarPostBody = r.returnsCreateRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v PostRenewalssearch400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v PostCreateorderV7500Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
