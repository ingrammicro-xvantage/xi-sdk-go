# PostQuoteToOrderV6400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traceid** | Pointer to **string** | A unique trace id to identify the issue. | [optional] 
**Type** | Pointer to **string** | Type of the error message. | [optional] 
**Message** | Pointer to **string** | A detailed error message. | [optional] 
**Fields** | Pointer to [**[]PostQuoteToOrderV6400ResponseFieldsInner**](PostQuoteToOrderV6400ResponseFieldsInner.md) |  | [optional] 

## Methods

### NewPostQuoteToOrderV6400Response

`func NewPostQuoteToOrderV6400Response() *PostQuoteToOrderV6400Response`

NewPostQuoteToOrderV6400Response instantiates a new PostQuoteToOrderV6400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostQuoteToOrderV6400ResponseWithDefaults

`func NewPostQuoteToOrderV6400ResponseWithDefaults() *PostQuoteToOrderV6400Response`

NewPostQuoteToOrderV6400ResponseWithDefaults instantiates a new PostQuoteToOrderV6400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceid

`func (o *PostQuoteToOrderV6400Response) GetTraceid() string`

GetTraceid returns the Traceid field if non-nil, zero value otherwise.

### GetTraceidOk

`func (o *PostQuoteToOrderV6400Response) GetTraceidOk() (*string, bool)`

GetTraceidOk returns a tuple with the Traceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceid

`func (o *PostQuoteToOrderV6400Response) SetTraceid(v string)`

SetTraceid sets Traceid field to given value.

### HasTraceid

`func (o *PostQuoteToOrderV6400Response) HasTraceid() bool`

HasTraceid returns a boolean if a field has been set.

### GetType

`func (o *PostQuoteToOrderV6400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostQuoteToOrderV6400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostQuoteToOrderV6400Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostQuoteToOrderV6400Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *PostQuoteToOrderV6400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostQuoteToOrderV6400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostQuoteToOrderV6400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostQuoteToOrderV6400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *PostQuoteToOrderV6400Response) GetFields() []PostQuoteToOrderV6400ResponseFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PostQuoteToOrderV6400Response) GetFieldsOk() (*[]PostQuoteToOrderV6400ResponseFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PostQuoteToOrderV6400Response) SetFields(v []PostQuoteToOrderV6400ResponseFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PostQuoteToOrderV6400Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


