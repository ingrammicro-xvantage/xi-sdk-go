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


// DealsAPIService DealsAPI service
type DealsAPIService service

type ApiGetResellersV6DealsdetailsRequest struct {
	ctx context.Context
	ApiService *DealsAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	dealId string
	iMSenderID *string
}

// Your unique Ingram Micro customer number.
func (r ApiGetResellersV6DealsdetailsRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellersV6DealsdetailsRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellersV6DealsdetailsRequest) IMCountryCode(iMCountryCode string) ApiGetResellersV6DealsdetailsRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetResellersV6DealsdetailsRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellersV6DealsdetailsRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiGetResellersV6DealsdetailsRequest) IMSenderID(iMSenderID string) ApiGetResellersV6DealsdetailsRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiGetResellersV6DealsdetailsRequest) Execute() (*DealsDetailsResponse, *http.Response, error) {
	return r.ApiService.GetResellersV6DealsdetailsExecute(r)
}

/*
GetResellersV6Dealsdetails Deals Details

The Deals Details API will retrieve all the details related to the specific deal id.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dealId Unique deal ID.
 @return ApiGetResellersV6DealsdetailsRequest
*/
func (a *DealsAPIService) GetResellersV6Dealsdetails(ctx context.Context, dealId string) ApiGetResellersV6DealsdetailsRequest {
	return ApiGetResellersV6DealsdetailsRequest{
		ApiService: a,
		ctx: ctx,
		dealId: dealId,
	}
}

// Execute executes the request
//  @return DealsDetailsResponse
func (a *DealsAPIService) GetResellersV6DealsdetailsExecute(r ApiGetResellersV6DealsdetailsRequest) (*DealsDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DealsDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DealsAPIService.GetResellersV6Dealsdetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/deals/{dealId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dealId"+"}", url.PathEscape(parameterValueToString(r.dealId, "dealId")), -1)

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
			var v GetResellerV6ValidateQuote500Response
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

type ApiGetResellersV6DealssearchRequest struct {
	ctx context.Context
	ApiService *DealsAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMSenderID *string
	endUser *string
	vendor *string
	dealId *string
}

// Your unique Ingram Micro customer number.
func (r ApiGetResellersV6DealssearchRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellersV6DealssearchRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellersV6DealssearchRequest) IMCountryCode(iMCountryCode string) ApiGetResellersV6DealssearchRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiGetResellersV6DealssearchRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellersV6DealssearchRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiGetResellersV6DealssearchRequest) IMSenderID(iMSenderID string) ApiGetResellersV6DealssearchRequest {
	r.iMSenderID = &iMSenderID
	return r
}

// The end user/customer&#39;s name.
func (r ApiGetResellersV6DealssearchRequest) EndUser(endUser string) ApiGetResellersV6DealssearchRequest {
	r.endUser = &endUser
	return r
}

// The vendor&#39;s name.
func (r ApiGetResellersV6DealssearchRequest) Vendor(vendor string) ApiGetResellersV6DealssearchRequest {
	r.vendor = &vendor
	return r
}

// Deal/Special bid number.
func (r ApiGetResellersV6DealssearchRequest) DealId(dealId string) ApiGetResellersV6DealssearchRequest {
	r.dealId = &dealId
	return r
}

func (r ApiGetResellersV6DealssearchRequest) Execute() (*DealsSearchResponse, *http.Response, error) {
	return r.ApiService.GetResellersV6DealssearchExecute(r)
}

/*
GetResellersV6Dealssearch Deals Search

The Deals Search API, by default, will retrieve all the deals that are associated with the customer’s account. The customer will be able to search deals using the End-user name, Vendor name, or DealID. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResellersV6DealssearchRequest
*/
func (a *DealsAPIService) GetResellersV6Dealssearch(ctx context.Context) ApiGetResellersV6DealssearchRequest {
	return ApiGetResellersV6DealssearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DealsSearchResponse
func (a *DealsAPIService) GetResellersV6DealssearchExecute(r ApiGetResellersV6DealssearchRequest) (*DealsSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DealsSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DealsAPIService.GetResellersV6Dealssearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/deals/search"

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

	if r.endUser != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endUser", r.endUser, "")
	}
	if r.vendor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vendor", r.vendor, "")
	}
	if r.dealId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dealId", r.dealId, "")
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
			var v GetResellerV6ValidateQuote500Response
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