/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// OrderStatusAPIService OrderStatusAPI service
type OrderStatusAPIService service

type ApiResellersV1WebhooksOrderstatuseventPostRequest struct {
	ctx context.Context
	ApiService *OrderStatusAPIService
	targeturl *string
	xHubSignature *string
	orderStatusAsyncNotificationRequest *OrderStatusAsyncNotificationRequest
}

// The webhook url where the request needs to sent.
func (r ApiResellersV1WebhooksOrderstatuseventPostRequest) Targeturl(targeturl string) ApiResellersV1WebhooksOrderstatuseventPostRequest {
	r.targeturl = &targeturl
	return r
}

// Ingram Micro creates a signature token by use of a secret key + Event ID. The algorithm to generate the secret ley is given at link https://developer.ingrammicro.com/reseller/article/how-use-webhook-secret-key. Use the event Id in the below sample along with your secret key to generate the key. Alternatively, to send try this out, use a random text to see how it works.
func (r ApiResellersV1WebhooksOrderstatuseventPostRequest) XHubSignature(xHubSignature string) ApiResellersV1WebhooksOrderstatuseventPostRequest {
	r.xHubSignature = &xHubSignature
	return r
}

func (r ApiResellersV1WebhooksOrderstatuseventPostRequest) OrderStatusAsyncNotificationRequest(orderStatusAsyncNotificationRequest OrderStatusAsyncNotificationRequest) ApiResellersV1WebhooksOrderstatuseventPostRequest {
	r.orderStatusAsyncNotificationRequest = &orderStatusAsyncNotificationRequest
	return r
}

func (r ApiResellersV1WebhooksOrderstatuseventPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResellersV1WebhooksOrderstatuseventPostExecute(r)
}

/*
ResellersV1WebhooksOrderstatuseventPost Order Status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResellersV1WebhooksOrderstatuseventPostRequest
*/
func (a *OrderStatusAPIService) ResellersV1WebhooksOrderstatuseventPost(ctx context.Context) ApiResellersV1WebhooksOrderstatuseventPostRequest {
	return ApiResellersV1WebhooksOrderstatuseventPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OrderStatusAPIService) ResellersV1WebhooksOrderstatuseventPostExecute(r ApiResellersV1WebhooksOrderstatuseventPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrderStatusAPIService.ResellersV1WebhooksOrderstatuseventPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resellers/v1/webhooks/orderstatusevent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.targeturl == nil {
		return nil, reportError("targeturl is required and must be specified")
	}
	if r.xHubSignature == nil {
		return nil, reportError("xHubSignature is required and must be specified")
	}
	if r.orderStatusAsyncNotificationRequest == nil {
		return nil, reportError("orderStatusAsyncNotificationRequest is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "targeturl", r.targeturl, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "x-hub-signature", r.xHubSignature, "")
	// body params
	localVarPostBody = r.orderStatusAsyncNotificationRequest
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
