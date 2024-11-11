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
	"reflect"
)


// ProductCatalogAPIService ProductCatalogAPI service
type ProductCatalogAPIService service

type ApiGetResellerV6ProductdetailRequest struct {
	ctx context.Context
	ApiService *ProductCatalogAPIService
	ingramPartNumber string
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMSenderID *string
}

// Your unique Ingram Micro customer number
func (r ApiGetResellerV6ProductdetailRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellerV6ProductdetailRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellerV6ProductdetailRequest) IMCountryCode(iMCountryCode string) ApiGetResellerV6ProductdetailRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction accross all the systems
func (r ApiGetResellerV6ProductdetailRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellerV6ProductdetailRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Sender Identification text
func (r ApiGetResellerV6ProductdetailRequest) IMSenderID(iMSenderID string) ApiGetResellerV6ProductdetailRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiGetResellerV6ProductdetailRequest) Execute() (*ProductDetailResponse, *http.Response, error) {
	return r.ApiService.GetResellerV6ProductdetailExecute(r)
}

/*
GetResellerV6Productdetail Product Details

Search all the product-related details using a unique Ingram Part Number. Currently, this API is available in the USA, India, and Netherlands.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ingramPartNumber Ingram Micro unique part number for the product
 @return ApiGetResellerV6ProductdetailRequest
*/
func (a *ProductCatalogAPIService) GetResellerV6Productdetail(ctx context.Context, ingramPartNumber string) ApiGetResellerV6ProductdetailRequest {
	return ApiGetResellerV6ProductdetailRequest{
		ApiService: a,
		ctx: ctx,
		ingramPartNumber: ingramPartNumber,
	}
}

// Execute executes the request
//  @return ProductDetailResponse
func (a *ProductCatalogAPIService) GetResellerV6ProductdetailExecute(r ApiGetResellerV6ProductdetailRequest) (*ProductDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductCatalogAPIService.GetResellerV6Productdetail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/catalog/details/{ingramPartNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"ingramPartNumber"+"}", url.PathEscape(parameterValueToString(r.ingramPartNumber, "ingramPartNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.ingramPartNumber) > 6 {
		return localVarReturnValue, nil, reportError("ingramPartNumber must have less than 6 elements")
	}
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "simple", "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "simple", "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "simple", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
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
			var v ErrorResponse
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

type ApiGetResellerV6ProductdetailcmpRequest struct {
	ctx context.Context
	ApiService *ProductCatalogAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMSenderID *string
	planName *string
	planId *string
	vendorPartNumber *string
}

// Your unique Ingram Micro customer number
func (r ApiGetResellerV6ProductdetailcmpRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellerV6ProductdetailcmpRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiGetResellerV6ProductdetailcmpRequest) IMCountryCode(iMCountryCode string) ApiGetResellerV6ProductdetailcmpRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems
func (r ApiGetResellerV6ProductdetailcmpRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellerV6ProductdetailcmpRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Sender Identification text
func (r ApiGetResellerV6ProductdetailcmpRequest) IMSenderID(iMSenderID string) ApiGetResellerV6ProductdetailcmpRequest {
	r.iMSenderID = &iMSenderID
	return r
}

// Name of the subscription plan
func (r ApiGetResellerV6ProductdetailcmpRequest) PlanName(planName string) ApiGetResellerV6ProductdetailcmpRequest {
	r.planName = &planName
	return r
}

// Id of the subscription plan.   &lt;span style&#x3D;&#39;color:red&#39;&gt;To search for details of subscription products, customer must pass either vendorPartNumber, planName or planId.&lt;/span&gt;
func (r ApiGetResellerV6ProductdetailcmpRequest) PlanId(planId string) ApiGetResellerV6ProductdetailcmpRequest {
	r.planId = &planId
	return r
}

// Vendor’s part number for the product.
func (r ApiGetResellerV6ProductdetailcmpRequest) VendorPartNumber(vendorPartNumber string) ApiGetResellerV6ProductdetailcmpRequest {
	r.vendorPartNumber = &vendorPartNumber
	return r
}

func (r ApiGetResellerV6ProductdetailcmpRequest) Execute() (*ProductDetailResponse, *http.Response, error) {
	return r.ApiService.GetResellerV6ProductdetailcmpExecute(r)
}

/*
GetResellerV6Productdetailcmp Product Details

Search all the product-related details.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResellerV6ProductdetailcmpRequest
*/
func (a *ProductCatalogAPIService) GetResellerV6Productdetailcmp(ctx context.Context) ApiGetResellerV6ProductdetailcmpRequest {
	return ApiGetResellerV6ProductdetailcmpRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductDetailResponse
func (a *ProductCatalogAPIService) GetResellerV6ProductdetailcmpExecute(r ApiGetResellerV6ProductdetailcmpRequest) (*ProductDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductCatalogAPIService.GetResellerV6Productdetailcmp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/catalog/details"

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

	if r.planName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "planName", r.planName, "form", "")
	}
	if r.planId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "planId", r.planId, "form", "")
	}
	if r.vendorPartNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vendorPartNumber", r.vendorPartNumber, "form", "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "simple", "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "simple", "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "simple", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
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
			var v ErrorResponse
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

type ApiGetResellerV6ProductsearchRequest struct {
	ctx context.Context
	ApiService *ProductCatalogAPIService
	iMCustomerNumber *string
	iMCorrelationID *string
	iMCountryCode *string
	pageNumber *int32
	pageSize *int32
	iMSenderID *string
	type_ *string
	hasDiscounts *string
	vendor *[]string
	vendorPartNumber *[]string
	acceptLanguage *string
	vendorNumber *string
	keyword *[]string
	category *string
	skipAuthorisation *string
	groupName *string
	planID *string
	showGroupInfo *bool
}

// Your unique Ingram Micro customer number
func (r ApiGetResellerV6ProductsearchRequest) IMCustomerNumber(iMCustomerNumber string) ApiGetResellerV6ProductsearchRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Unique transaction number to identify each transaction accross all the systems
func (r ApiGetResellerV6ProductsearchRequest) IMCorrelationID(iMCorrelationID string) ApiGetResellerV6ProductsearchRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Two-character ISO country code.
func (r ApiGetResellerV6ProductsearchRequest) IMCountryCode(iMCountryCode string) ApiGetResellerV6ProductsearchRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Current page number. Default is 1
func (r ApiGetResellerV6ProductsearchRequest) PageNumber(pageNumber int32) ApiGetResellerV6ProductsearchRequest {
	r.pageNumber = &pageNumber
	return r
}

// Number of records required in the call - max records 100 per page
func (r ApiGetResellerV6ProductsearchRequest) PageSize(pageSize int32) ApiGetResellerV6ProductsearchRequest {
	r.pageSize = &pageSize
	return r
}

// Sender Identification text
func (r ApiGetResellerV6ProductsearchRequest) IMSenderID(iMSenderID string) ApiGetResellerV6ProductsearchRequest {
	r.iMSenderID = &iMSenderID
	return r
}

// The SKU type of product. One of Physical, Digital, or Any.
func (r ApiGetResellerV6ProductsearchRequest) Type_(type_ string) ApiGetResellerV6ProductsearchRequest {
	r.type_ = &type_
	return r
}

// Specifies if there are discounts available for the product.
func (r ApiGetResellerV6ProductsearchRequest) HasDiscounts(hasDiscounts string) ApiGetResellerV6ProductsearchRequest {
	r.hasDiscounts = &hasDiscounts
	return r
}

// The name of the vendor/manufacturer of the product.
func (r ApiGetResellerV6ProductsearchRequest) Vendor(vendor []string) ApiGetResellerV6ProductsearchRequest {
	r.vendor = &vendor
	return r
}

// The vendors part number for the product.
func (r ApiGetResellerV6ProductsearchRequest) VendorPartNumber(vendorPartNumber []string) ApiGetResellerV6ProductsearchRequest {
	r.vendorPartNumber = &vendorPartNumber
	return r
}

// Header to the API calls, the content will help us identify the response language.
func (r ApiGetResellerV6ProductsearchRequest) AcceptLanguage(acceptLanguage string) ApiGetResellerV6ProductsearchRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// Vendor number of the product
func (r ApiGetResellerV6ProductsearchRequest) VendorNumber(vendorNumber string) ApiGetResellerV6ProductsearchRequest {
	r.vendorNumber = &vendorNumber
	return r
}

// Keyword search,can be ingram part number or vendor part number or product title or vendor nameKeyword search. Can be Ingram Micro part number, vender part number, product title, or vendor name.
func (r ApiGetResellerV6ProductsearchRequest) Keyword(keyword []string) ApiGetResellerV6ProductsearchRequest {
	r.keyword = &keyword
	return r
}

// The category of the product. Example: Displays.
func (r ApiGetResellerV6ProductsearchRequest) Category(category string) ApiGetResellerV6ProductsearchRequest {
	r.category = &category
	return r
}

// This parameter is True when you want Skip the authorization, so template will work like current B2b template.
func (r ApiGetResellerV6ProductsearchRequest) SkipAuthorisation(skipAuthorisation string) ApiGetResellerV6ProductsearchRequest {
	r.skipAuthorisation = &skipAuthorisation
	return r
}

// Name of the Product Group
func (r ApiGetResellerV6ProductsearchRequest) GroupName(groupName string) ApiGetResellerV6ProductsearchRequest {
	r.groupName = &groupName
	return r
}

// ID of the plan
func (r ApiGetResellerV6ProductsearchRequest) PlanID(planID string) ApiGetResellerV6ProductsearchRequest {
	r.planID = &planID
	return r
}

// In case of value true, below Group related information will displayed without the plan info. Group Name, Group Description, Number of plans, link in the group. A link will be provided if customer want to see all the plans in that group.
func (r ApiGetResellerV6ProductsearchRequest) ShowGroupInfo(showGroupInfo bool) ApiGetResellerV6ProductsearchRequest {
	r.showGroupInfo = &showGroupInfo
	return r
}

func (r ApiGetResellerV6ProductsearchRequest) Execute() (*ProductSearchResponse, *http.Response, error) {
	return r.ApiService.GetResellerV6ProductsearchExecute(r)
}

/*
GetResellerV6Productsearch Search Products

Search the Ingram Micro product catalog by providing any of the information in the keyword(Ingram part number / vendor part number/ product description / UPC

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetResellerV6ProductsearchRequest
*/
func (a *ProductCatalogAPIService) GetResellerV6Productsearch(ctx context.Context) ApiGetResellerV6ProductsearchRequest {
	return ApiGetResellerV6ProductsearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProductSearchResponse
func (a *ProductCatalogAPIService) GetResellerV6ProductsearchExecute(r ApiGetResellerV6ProductsearchRequest) (*ProductSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProductSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductCatalogAPIService.GetResellerV6Productsearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/catalog"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.iMCustomerNumber == nil {
		return localVarReturnValue, nil, reportError("iMCustomerNumber is required and must be specified")
	}
	if strlen(*r.iMCustomerNumber) > 10 {
		return localVarReturnValue, nil, reportError("iMCustomerNumber must have less than 10 elements")
	}
	if r.iMCorrelationID == nil {
		return localVarReturnValue, nil, reportError("iMCorrelationID is required and must be specified")
	}
	if strlen(*r.iMCorrelationID) > 32 {
		return localVarReturnValue, nil, reportError("iMCorrelationID must have less than 32 elements")
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

	if r.pageNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNumber", r.pageNumber, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.hasDiscounts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hasDiscounts", r.hasDiscounts, "form", "")
	}
	if r.vendor != nil {
		t := *r.vendor
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "vendor", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "vendor", t, "form", "multi")
		}
	}
	if r.vendorPartNumber != nil {
		t := *r.vendorPartNumber
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "vendorPartNumber", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "vendorPartNumber", t, "form", "multi")
		}
	}
	if r.vendorNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vendorNumber", r.vendorNumber, "form", "")
	}
	if r.keyword != nil {
		t := *r.keyword
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "keyword", t, "form", "multi")
		}
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "form", "")
	}
	if r.skipAuthorisation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skipAuthorisation", r.skipAuthorisation, "form", "")
	}
	if r.groupName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "groupName", r.groupName, "form", "")
	}
	if r.planID != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "planID", r.planID, "form", "")
	}
	if r.showGroupInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "showGroupInfo", r.showGroupInfo, "form", "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "simple", "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "simple", "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "simple", "")
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "simple", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
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
			var v ErrorResponse
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

type ApiPostPriceandavailabilityRequest struct {
	ctx context.Context
	ApiService *ProductCatalogAPIService
	includeAvailability *bool
	includePricing *bool
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	priceAndAvailabilityRequest *PriceAndAvailabilityRequest
	includeProductAttributes *bool
	iMSenderID *string
}

// Pass boolean value as input, if true the response will contain warehouse availability details, if false the response will not hold warehouse availability details
func (r ApiPostPriceandavailabilityRequest) IncludeAvailability(includeAvailability bool) ApiPostPriceandavailabilityRequest {
	r.includeAvailability = &includeAvailability
	return r
}

// Pass boolean value as input, if true the response will contain Pricing details of the Product, if false the response will not hold Pricing details.
func (r ApiPostPriceandavailabilityRequest) IncludePricing(includePricing bool) ApiPostPriceandavailabilityRequest {
	r.includePricing = &includePricing
	return r
}

// Your unique Ingram Micro customer number.
func (r ApiPostPriceandavailabilityRequest) IMCustomerNumber(iMCustomerNumber string) ApiPostPriceandavailabilityRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiPostPriceandavailabilityRequest) IMCountryCode(iMCountryCode string) ApiPostPriceandavailabilityRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiPostPriceandavailabilityRequest) IMCorrelationID(iMCorrelationID string) ApiPostPriceandavailabilityRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

func (r ApiPostPriceandavailabilityRequest) PriceAndAvailabilityRequest(priceAndAvailabilityRequest PriceAndAvailabilityRequest) ApiPostPriceandavailabilityRequest {
	r.priceAndAvailabilityRequest = &priceAndAvailabilityRequest
	return r
}

// Pass boolean value as input, if true the response will contain detailed attributes related to the Product, if false or not sent the response will contain very few Product details.
func (r ApiPostPriceandavailabilityRequest) IncludeProductAttributes(includeProductAttributes bool) ApiPostPriceandavailabilityRequest {
	r.includeProductAttributes = &includeProductAttributes
	return r
}

// Unique value used to identify the sender of the transaction. Example: MyCompany
func (r ApiPostPriceandavailabilityRequest) IMSenderID(iMSenderID string) ApiPostPriceandavailabilityRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiPostPriceandavailabilityRequest) Execute() ([]PriceAndAvailabilityResponseInner, *http.Response, error) {
	return r.ApiService.PostPriceandavailabilityExecute(r)
}

/*
PostPriceandavailability Price and Availability

The PriceAndAvailability API, will retrieve Pricing, Availability, discounts, Inventory Location, Reserve Inventory for the products upto 50 products. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostPriceandavailabilityRequest
*/
func (a *ProductCatalogAPIService) PostPriceandavailability(ctx context.Context) ApiPostPriceandavailabilityRequest {
	return ApiPostPriceandavailabilityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PriceAndAvailabilityResponseInner
func (a *ProductCatalogAPIService) PostPriceandavailabilityExecute(r ApiPostPriceandavailabilityRequest) ([]PriceAndAvailabilityResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PriceAndAvailabilityResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProductCatalogAPIService.PostPriceandavailability")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/catalog/priceandavailability"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.includeAvailability == nil {
		return localVarReturnValue, nil, reportError("includeAvailability is required and must be specified")
	}
	if r.includePricing == nil {
		return localVarReturnValue, nil, reportError("includePricing is required and must be specified")
	}
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
	if r.priceAndAvailabilityRequest == nil {
		return localVarReturnValue, nil, reportError("priceAndAvailabilityRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "includeAvailability", r.includeAvailability, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "includePricing", r.includePricing, "form", "")
	if r.includeProductAttributes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeProductAttributes", r.includeProductAttributes, "form", "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerNumber", r.iMCustomerNumber, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CountryCode", r.iMCountryCode, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CorrelationID", r.iMCorrelationID, "simple", "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "simple", "")
	}
	// body params
	localVarPostBody = r.priceAndAvailabilityRequest
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
			var v ErrorResponse
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
			var v ErrorResponse
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
