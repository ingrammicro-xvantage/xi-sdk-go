# QuoteListResponseQuoteSearchResponseResponsePreamble

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseStatus** | Pointer to **string** | Status of the Request - \&quot;Passed\&quot;, \&quot;Failed\&quot; | [optional] 
**ResponseStatusCode** | Pointer to **string** | responseStatusCode is the code returned in response to a request. The following Codes are returned: 200 400 500 | [optional] 
**ResponseMessage** | Pointer to **string** | 200 &#x3D; Action was successfully received, understood and accepted. 400 &#x3D; The request contains bad syntax or can not be fullfilled. This means there is a problem with the request. 500 &#x3D; The server failed to fulfill an apparently valid request. This is a temporary problem, the request should be resubmitted. | [optional] 

## Methods

### NewQuoteListResponseQuoteSearchResponseResponsePreamble

`func NewQuoteListResponseQuoteSearchResponseResponsePreamble() *QuoteListResponseQuoteSearchResponseResponsePreamble`

NewQuoteListResponseQuoteSearchResponseResponsePreamble instantiates a new QuoteListResponseQuoteSearchResponseResponsePreamble object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteListResponseQuoteSearchResponseResponsePreambleWithDefaults

`func NewQuoteListResponseQuoteSearchResponseResponsePreambleWithDefaults() *QuoteListResponseQuoteSearchResponseResponsePreamble`

NewQuoteListResponseQuoteSearchResponseResponsePreambleWithDefaults instantiates a new QuoteListResponseQuoteSearchResponseResponsePreamble object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseStatus

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatus() string`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatusOk() (*string, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) SetResponseStatus(v string)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatusCode() string`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatusCodeOk() (*string, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) SetResponseStatusCode(v string)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetResponseMessage

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.

### HasResponseMessage

`func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) HasResponseMessage() bool`

HasResponseMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


