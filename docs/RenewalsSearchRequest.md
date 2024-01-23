# RenewalsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**RenewalsSearchRequestStatus**](RenewalsSearchRequestStatus.md) |  | [optional] 
**DataType** | Pointer to [**RenewalsSearchRequestDataType**](RenewalsSearchRequestDataType.md) |  | [optional] 
**Vendor** | Pointer to **string** | The name of the Vendor. | [optional] 
**EndUser** | Pointer to **string** | The name of the enduser.  | [optional] 

## Methods

### NewRenewalsSearchRequest

`func NewRenewalsSearchRequest() *RenewalsSearchRequest`

NewRenewalsSearchRequest instantiates a new RenewalsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewalsSearchRequestWithDefaults

`func NewRenewalsSearchRequestWithDefaults() *RenewalsSearchRequest`

NewRenewalsSearchRequestWithDefaults instantiates a new RenewalsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RenewalsSearchRequest) GetStatus() RenewalsSearchRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RenewalsSearchRequest) GetStatusOk() (*RenewalsSearchRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RenewalsSearchRequest) SetStatus(v RenewalsSearchRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RenewalsSearchRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDataType

`func (o *RenewalsSearchRequest) GetDataType() RenewalsSearchRequestDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *RenewalsSearchRequest) GetDataTypeOk() (*RenewalsSearchRequestDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *RenewalsSearchRequest) SetDataType(v RenewalsSearchRequestDataType)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *RenewalsSearchRequest) HasDataType() bool`

HasDataType returns a boolean if a field has been set.

### GetVendor

`func (o *RenewalsSearchRequest) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *RenewalsSearchRequest) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *RenewalsSearchRequest) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *RenewalsSearchRequest) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetEndUser

`func (o *RenewalsSearchRequest) GetEndUser() string`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *RenewalsSearchRequest) GetEndUserOk() (*string, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *RenewalsSearchRequest) SetEndUser(v string)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *RenewalsSearchRequest) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


