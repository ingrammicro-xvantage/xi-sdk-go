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
)


// StockUpdateAPIService StockUpdateAPI service
type StockUpdateAPIService service

type ApiResellersV1WebhooksAvailabilityupdatePostRequest struct {
	ctx context.Context
	ApiService *StockUpdateAPIService
	targeturl *string
	xHubSignature *string
	availabilityAsyncNotificationRequest *AvailabilityAsyncNotificationRequest
}

// The webhook url where the request needs to sent.
func (r ApiResellersV1WebhooksAvailabilityupdatePostRequest) Targeturl(targeturl string) ApiResellersV1WebhooksAvailabilityupdatePostRequest {
	r.targeturl = &targeturl
	return r
}

// Ingram Micro creates a signature token by use of a secret key + Event ID. The algorithm to generate the secret ley is given at link https://developer.ingrammicro.com/reseller/article/how-use-webhook-secret-key. Use the event Id in the below sample along with your secret key to generate the key. Alternatively, to send try this out, use a random text to see how it works.
func (r ApiResellersV1WebhooksAvailabilityupdatePostRequest) XHubSignature(xHubSignature string) ApiResellersV1WebhooksAvailabilityupdatePostRequest {
	r.xHubSignature = &xHubSignature
	return r
}

func (r ApiResellersV1WebhooksAvailabilityupdatePostRequest) AvailabilityAsyncNotificationRequest(availabilityAsyncNotificationRequest AvailabilityAsyncNotificationRequest) ApiResellersV1WebhooksAvailabilityupdatePostRequest {
	r.availabilityAsyncNotificationRequest = &availabilityAsyncNotificationRequest
	return r
}

func (r ApiResellersV1WebhooksAvailabilityupdatePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResellersV1WebhooksAvailabilityupdatePostExecute(r)
}

/*
ResellersV1WebhooksAvailabilityupdatePost Stock Update

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResellersV1WebhooksAvailabilityupdatePostRequest
*/
func (a *StockUpdateAPIService) ResellersV1WebhooksAvailabilityupdatePost(ctx context.Context) ApiResellersV1WebhooksAvailabilityupdatePostRequest {
	return ApiResellersV1WebhooksAvailabilityupdatePostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StockUpdateAPIService) ResellersV1WebhooksAvailabilityupdatePostExecute(r ApiResellersV1WebhooksAvailabilityupdatePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StockUpdateAPIService.ResellersV1WebhooksAvailabilityupdatePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v1/webhooks/availabilityupdate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targeturl == nil {
		return nil, reportError("targeturl is required and must be specified")
	}
	if r.xHubSignature == nil {
		return nil, reportError("xHubSignature is required and must be specified")
	}
	if r.availabilityAsyncNotificationRequest == nil {
		return nil, reportError("availabilityAsyncNotificationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "targeturl", r.targeturl, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-hub-signature", r.xHubSignature, "simple", "")
	// body params
	localVarPostBody = r.availabilityAsyncNotificationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
