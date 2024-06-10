# PostCreateorderV7400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traceid** | Pointer to **string** | A unique trace id to identify the issue. | [optional] 
**Type** | Pointer to **string** | Type of the error message. | [optional] 
**Message** | Pointer to **string** | A detailed error message. | [optional] 
**Fields** | Pointer to [**[]PostCreateorderV7400ResponseFieldsInner**](PostCreateorderV7400ResponseFieldsInner.md) |  | [optional] 

## Methods

### NewPostCreateorderV7400Response

`func NewPostCreateorderV7400Response() *PostCreateorderV7400Response`

NewPostCreateorderV7400Response instantiates a new PostCreateorderV7400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateorderV7400ResponseWithDefaults

`func NewPostCreateorderV7400ResponseWithDefaults() *PostCreateorderV7400Response`

NewPostCreateorderV7400ResponseWithDefaults instantiates a new PostCreateorderV7400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceid

`func (o *PostCreateorderV7400Response) GetTraceid() string`

GetTraceid returns the Traceid field if non-nil, zero value otherwise.

### GetTraceidOk

`func (o *PostCreateorderV7400Response) GetTraceidOk() (*string, bool)`

GetTraceidOk returns a tuple with the Traceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceid

`func (o *PostCreateorderV7400Response) SetTraceid(v string)`

SetTraceid sets Traceid field to given value.

### HasTraceid

`func (o *PostCreateorderV7400Response) HasTraceid() bool`

HasTraceid returns a boolean if a field has been set.

### GetType

`func (o *PostCreateorderV7400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCreateorderV7400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCreateorderV7400Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostCreateorderV7400Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMessage

`func (o *PostCreateorderV7400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostCreateorderV7400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostCreateorderV7400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostCreateorderV7400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetFields

`func (o *PostCreateorderV7400Response) GetFields() []PostCreateorderV7400ResponseFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PostCreateorderV7400Response) GetFieldsOk() (*[]PostCreateorderV7400ResponseFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PostCreateorderV7400Response) SetFields(v []PostCreateorderV7400ResponseFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PostCreateorderV7400Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


