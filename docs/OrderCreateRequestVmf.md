# OrderCreateRequestVmf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendAuthNumber** | Pointer to **string** | Authorization number provided by vendor to Ingram&#39;s reseller. Orders will be placed on hold without this value, vendor specific mandatory field - please reach out Ingram Sales team for list of vendor for whom this is mandatory. | [optional] 

## Methods

### NewOrderCreateRequestVmf

`func NewOrderCreateRequestVmf() *OrderCreateRequestVmf`

NewOrderCreateRequestVmf instantiates a new OrderCreateRequestVmf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestVmfWithDefaults

`func NewOrderCreateRequestVmfWithDefaults() *OrderCreateRequestVmf`

NewOrderCreateRequestVmfWithDefaults instantiates a new OrderCreateRequestVmf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendAuthNumber

`func (o *OrderCreateRequestVmf) GetVendAuthNumber() string`

GetVendAuthNumber returns the VendAuthNumber field if non-nil, zero value otherwise.

### GetVendAuthNumberOk

`func (o *OrderCreateRequestVmf) GetVendAuthNumberOk() (*string, bool)`

GetVendAuthNumberOk returns a tuple with the VendAuthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendAuthNumber

`func (o *OrderCreateRequestVmf) SetVendAuthNumber(v string)`

SetVendAuthNumber sets VendAuthNumber field to given value.

### HasVendAuthNumber

`func (o *OrderCreateRequestVmf) HasVendAuthNumber() bool`

HasVendAuthNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


