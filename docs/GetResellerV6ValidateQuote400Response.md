# GetResellerV6ValidateQuote400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traceid** | Pointer to **string** | Unique Id to identify error. | [optional] 
**Type** | Pointer to **string** | Describes the type of the error. | [optional] 
**Message** | Pointer to **string** | A detailed error message. | [optional] 
**Fields** | Pointer to [**[]GetResellerV6ValidateQuote400ResponseFieldsInner**](GetResellerV6ValidateQuote400ResponseFieldsInner.md) |  | [optional] 

## Methods

### NewGetResellerV6ValidateQuote400Response

`func NewGetResellerV6ValidateQuote400Response() *GetResellerV6ValidateQuote400Response`

NewGetResellerV6ValidateQuote400Response instantiates a new GetResellerV6ValidateQuote400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetResellerV6ValidateQuote400ResponseWithDefaults

`func NewGetResellerV6ValidateQuote400ResponseWithDefaults() *GetResellerV6ValidateQuote400Response`

NewGetResellerV6ValidateQuote400ResponseWithDefaults instantiates a new GetResellerV6ValidateQuote400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceid

`func (o *GetResellerV6ValidateQuote400Response) GetTraceid() string`

GetTraceid returns the Traceid field if non-nil, zero value otherwise.

### GetTraceidOk

`func (o *GetResellerV6ValidateQuote400Response) GetTraceidOk() (*string, bool)`

GetTraceidOk returns a tuple with the Traceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceid

`func (o *GetResellerV6ValidateQuote400Response) SetTraceid(v string)`

SetTraceid sets Traceid field to given value.

### HasTraceid

`func (o *GetResellerV6ValidateQuote400Response) HasTraceid() bool`

HasTraceid returns a boolean if a field has been set.

### GetType

`func (o *GetResellerV6ValidateQuote400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetResellerV6ValidateQuote400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetResellerV6ValidateQuote400Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetResellerV6ValidateQuote400Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *GetResellerV6ValidateQuote400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetResellerV6ValidateQuote400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetResellerV6ValidateQuote400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetResellerV6ValidateQuote400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *GetResellerV6ValidateQuote400Response) GetFields() []GetResellerV6ValidateQuote400ResponseFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetResellerV6ValidateQuote400Response) GetFieldsOk() (*[]GetResellerV6ValidateQuote400ResponseFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetResellerV6ValidateQuote400Response) SetFields(v []GetResellerV6ValidateQuote400ResponseFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GetResellerV6ValidateQuote400Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


