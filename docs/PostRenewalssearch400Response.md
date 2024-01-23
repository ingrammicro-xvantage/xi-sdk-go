# PostRenewalssearch400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Traceid** | Pointer to **string** | Unique Id to identify error. | [optional] 
**Type** | Pointer to **string** | Describes the type of the error. | [optional] 
**Fields** | Pointer to [**[]GetResellerV6ValidateQuote400ResponseFieldsInner**](GetResellerV6ValidateQuote400ResponseFieldsInner.md) |  | [optional] 

## Methods

### NewPostRenewalssearch400Response

`func NewPostRenewalssearch400Response() *PostRenewalssearch400Response`

NewPostRenewalssearch400Response instantiates a new PostRenewalssearch400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRenewalssearch400ResponseWithDefaults

`func NewPostRenewalssearch400ResponseWithDefaults() *PostRenewalssearch400Response`

NewPostRenewalssearch400ResponseWithDefaults instantiates a new PostRenewalssearch400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceid

`func (o *PostRenewalssearch400Response) GetTraceid() string`

GetTraceid returns the Traceid field if non-nil, zero value otherwise.

### GetTraceidOk

`func (o *PostRenewalssearch400Response) GetTraceidOk() (*string, bool)`

GetTraceidOk returns a tuple with the Traceid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceid

`func (o *PostRenewalssearch400Response) SetTraceid(v string)`

SetTraceid sets Traceid field to given value.

### HasTraceid

`func (o *PostRenewalssearch400Response) HasTraceid() bool`

HasTraceid returns a boolean if a field has been set.

### GetType

`func (o *PostRenewalssearch400Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostRenewalssearch400Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostRenewalssearch400Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostRenewalssearch400Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFields

`func (o *PostRenewalssearch400Response) GetFields() []GetResellerV6ValidateQuote400ResponseFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *PostRenewalssearch400Response) GetFieldsOk() (*[]GetResellerV6ValidateQuote400ResponseFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *PostRenewalssearch400Response) SetFields(v []GetResellerV6ValidateQuote400ResponseFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *PostRenewalssearch400Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


