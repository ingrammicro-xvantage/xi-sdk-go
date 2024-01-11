# FreightRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The ID references the reseller&#39;s address in Ingram Micro&#39;s system for shipping. Provided to resellers during the onboarding process. | [optional] 
**IngramPartNumber** | Pointer to **string** | The unique IngramMicro part number. | [optional] 
**Quantity** | Pointer to **string** | The requested quantity of the line item. | [optional] 
**WarehouseId** | Pointer to **string** | The ID of the warehouse the line item will ship from. | [optional] 
**CarrierCode** | Pointer to **string** | The code for the shipping carrier for the line item. | [optional] 

## Methods

### NewFreightRequestLinesInner

`func NewFreightRequestLinesInner() *FreightRequestLinesInner`

NewFreightRequestLinesInner instantiates a new FreightRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreightRequestLinesInnerWithDefaults

`func NewFreightRequestLinesInnerWithDefaults() *FreightRequestLinesInner`

NewFreightRequestLinesInnerWithDefaults instantiates a new FreightRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLineNumber

`func (o *FreightRequestLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *FreightRequestLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *FreightRequestLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *FreightRequestLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *FreightRequestLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *FreightRequestLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *FreightRequestLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *FreightRequestLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *FreightRequestLinesInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *FreightRequestLinesInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *FreightRequestLinesInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *FreightRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetWarehouseId

`func (o *FreightRequestLinesInner) GetWarehouseId() string`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *FreightRequestLinesInner) GetWarehouseIdOk() (*string, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *FreightRequestLinesInner) SetWarehouseId(v string)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *FreightRequestLinesInner) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetCarrierCode

`func (o *FreightRequestLinesInner) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *FreightRequestLinesInner) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *FreightRequestLinesInner) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *FreightRequestLinesInner) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


