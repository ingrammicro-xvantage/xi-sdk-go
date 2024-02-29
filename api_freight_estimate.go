/*
XI Sdk Resellers

For Ingram Micro Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

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
)


// FreightEstimateAPIService FreightEstimateAPI service
type FreightEstimateAPIService service

type ApiPostFreightestimateRequest struct {
	ctx context.Context
	ApiService *FreightEstimateAPIService
	iMCustomerNumber *string
	iMCountryCode *string
	iMCorrelationID *string
	iMCustomerContact *string
	iMSenderID *string
	freightRequest *FreightRequest
}

// Your unique Ingram Micro customer number.
func (r ApiPostFreightestimateRequest) IMCustomerNumber(iMCustomerNumber string) ApiPostFreightestimateRequest {
	r.iMCustomerNumber = &iMCustomerNumber
	return r
}

// Two-character ISO country code.
func (r ApiPostFreightestimateRequest) IMCountryCode(iMCountryCode string) ApiPostFreightestimateRequest {
	r.iMCountryCode = &iMCountryCode
	return r
}

// Unique transaction number to identify each transaction across all the systems.
func (r ApiPostFreightestimateRequest) IMCorrelationID(iMCorrelationID string) ApiPostFreightestimateRequest {
	r.iMCorrelationID = &iMCorrelationID
	return r
}

// Logged in Users email address contact.
func (r ApiPostFreightestimateRequest) IMCustomerContact(iMCustomerContact string) ApiPostFreightestimateRequest {
	r.iMCustomerContact = &iMCustomerContact
	return r
}

// Unique value used to identify the sender of the transaction.
func (r ApiPostFreightestimateRequest) IMSenderID(iMSenderID string) ApiPostFreightestimateRequest {
	r.iMSenderID = &iMSenderID
	return r
}

func (r ApiPostFreightestimateRequest) FreightRequest(freightRequest FreightRequest) ApiPostFreightestimateRequest {
	r.freightRequest = &freightRequest
	return r
}

func (r ApiPostFreightestimateRequest) Execute() (*FreightResponse, *http.Response, error) {
	return r.ApiService.PostFreightestimateExecute(r)
}

/*
PostFreightestimate Freight Estimate

The freight estimator endpoint will allow customers to understand the freight cost for an order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostFreightestimateRequest
*/
func (a *FreightEstimateAPIService) PostFreightestimate(ctx context.Context) ApiPostFreightestimateRequest {
	return ApiPostFreightestimateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FreightResponse
func (a *FreightEstimateAPIService) PostFreightestimateExecute(r ApiPostFreightestimateRequest) (*FreightResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FreightResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreightEstimateAPIService.PostFreightestimate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v6/freightestimate"

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
	if r.iMCustomerContact == nil {
		return localVarReturnValue, nil, reportError("iMCustomerContact is required and must be specified")
	}
	if strlen(*r.iMCustomerContact) > 64 {
		return localVarReturnValue, nil, reportError("iMCustomerContact must have less than 64 elements")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-CustomerContact", r.iMCustomerContact, "")
	if r.iMSenderID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "IM-SenderID", r.iMSenderID, "")
	}
	// body params
	localVarPostBody = r.freightRequest
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
