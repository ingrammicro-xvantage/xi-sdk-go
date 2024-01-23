# ErrorResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique Id to identify error. | [optional] 
**Type** | Pointer to **string** | Describes the type of the error. | [optional] 
**Message** | Pointer to **string** | Describes the error message. | [optional] 
**Fields** | Pointer to [**[]ErrorResponseErrorsInnerFieldsInner**](ErrorResponseErrorsInnerFieldsInner.md) |  | [optional] 

## Methods

### NewErrorResponseErrorsInner

`func NewErrorResponseErrorsInner() *ErrorResponseErrorsInner`

NewErrorResponseErrorsInner instantiates a new ErrorResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseErrorsInnerWithDefaults

`func NewErrorResponseErrorsInnerWithDefaults() *ErrorResponseErrorsInner`

NewErrorResponseErrorsInnerWithDefaults instantiates a new ErrorResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorResponseErrorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseErrorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseErrorsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorResponseErrorsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ErrorResponseErrorsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponseErrorsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponseErrorsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorResponseErrorsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorResponseErrorsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *ErrorResponseErrorsInner) GetFields() []ErrorResponseErrorsInnerFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ErrorResponseErrorsInner) GetFieldsOk() (*[]ErrorResponseErrorsInnerFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ErrorResponseErrorsInner) SetFields(v []ErrorResponseErrorsInnerFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ErrorResponseErrorsInner) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


