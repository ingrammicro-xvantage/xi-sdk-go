# GetAccesstoken500ResponseFault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Faultstring** | Pointer to **string** |  | [optional] 
**Detail** | Pointer to [**GetAccesstoken500ResponseFaultDetail**](GetAccesstoken500ResponseFaultDetail.md) |  | [optional] 

## Methods

### NewGetAccesstoken500ResponseFault

`func NewGetAccesstoken500ResponseFault() *GetAccesstoken500ResponseFault`

NewGetAccesstoken500ResponseFault instantiates a new GetAccesstoken500ResponseFault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccesstoken500ResponseFaultWithDefaults

`func NewGetAccesstoken500ResponseFaultWithDefaults() *GetAccesstoken500ResponseFault`

NewGetAccesstoken500ResponseFaultWithDefaults instantiates a new GetAccesstoken500ResponseFault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFaultstring

`func (o *GetAccesstoken500ResponseFault) GetFaultstring() string`

GetFaultstring returns the Faultstring field if non-nil, zero value otherwise.

### GetFaultstringOk

`func (o *GetAccesstoken500ResponseFault) GetFaultstringOk() (*string, bool)`

GetFaultstringOk returns a tuple with the Faultstring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultstring

`func (o *GetAccesstoken500ResponseFault) SetFaultstring(v string)`

SetFaultstring sets Faultstring field to given value.

### HasFaultstring

`func (o *GetAccesstoken500ResponseFault) HasFaultstring() bool`

HasFaultstring returns a boolean if a field has been set.

### GetDetail

`func (o *GetAccesstoken500ResponseFault) GetDetail() GetAccesstoken500ResponseFaultDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetAccesstoken500ResponseFault) GetDetailOk() (*GetAccesstoken500ResponseFaultDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetAccesstoken500ResponseFault) SetDetail(v GetAccesstoken500ResponseFaultDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GetAccesstoken500ResponseFault) HasDetail() bool`

HasDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


