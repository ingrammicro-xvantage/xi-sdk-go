# MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Ingrampartnumber** | Pointer to **string** | Ingram Micro system specific SKU number for the product for which the price is requested at Ingram Micro | [optional] 
**Vendorpartnumber** | Pointer to **string** | Vendor Part Number for the product for which the price is requested at Ingram Micro | [optional] 
**UPC** | Pointer to **string** | Universal Product code for the product for which the price is requested at Ingram Micro | [optional] 
**Customerpartnumber** | Pointer to **string** | Unique identification number of customer. For this option the Ingram Micro Sales rep must set up a cross reference table.  | [optional] 
**Warehouseidlist** | Pointer to **string** | Unique identity for Ingram Micro warehouses against which stock details are returned. | [optional] 

## Methods

### NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem

`func NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem`

NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem instantiates a new MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItemWithDefaults

`func NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItemWithDefaults() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem`

NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItemWithDefaults instantiates a new MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIngrampartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetIngrampartnumber() string`

GetIngrampartnumber returns the Ingrampartnumber field if non-nil, zero value otherwise.

### GetIngrampartnumberOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetIngrampartnumberOk() (*string, bool)`

GetIngrampartnumberOk returns a tuple with the Ingrampartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngrampartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) SetIngrampartnumber(v string)`

SetIngrampartnumber sets Ingrampartnumber field to given value.

### HasIngrampartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) HasIngrampartnumber() bool`

HasIngrampartnumber returns a boolean if a field has been set.

### GetVendorpartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetVendorpartnumber() string`

GetVendorpartnumber returns the Vendorpartnumber field if non-nil, zero value otherwise.

### GetVendorpartnumberOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetVendorpartnumberOk() (*string, bool)`

GetVendorpartnumberOk returns a tuple with the Vendorpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorpartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) SetVendorpartnumber(v string)`

SetVendorpartnumber sets Vendorpartnumber field to given value.

### HasVendorpartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) HasVendorpartnumber() bool`

HasVendorpartnumber returns a boolean if a field has been set.

### GetUPC

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetUPC() string`

GetUPC returns the UPC field if non-nil, zero value otherwise.

### GetUPCOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetUPCOk() (*string, bool)`

GetUPCOk returns a tuple with the UPC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPC

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) SetUPC(v string)`

SetUPC sets UPC field to given value.

### HasUPC

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) HasUPC() bool`

HasUPC returns a boolean if a field has been set.

### GetCustomerpartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetCustomerpartnumber() string`

GetCustomerpartnumber returns the Customerpartnumber field if non-nil, zero value otherwise.

### GetCustomerpartnumberOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetCustomerpartnumberOk() (*string, bool)`

GetCustomerpartnumberOk returns a tuple with the Customerpartnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerpartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) SetCustomerpartnumber(v string)`

SetCustomerpartnumber sets Customerpartnumber field to given value.

### HasCustomerpartnumber

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) HasCustomerpartnumber() bool`

HasCustomerpartnumber returns a boolean if a field has been set.

### GetWarehouseidlist

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetWarehouseidlist() string`

GetWarehouseidlist returns the Warehouseidlist field if non-nil, zero value otherwise.

### GetWarehouseidlistOk

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) GetWarehouseidlistOk() (*string, bool)`

GetWarehouseidlistOk returns a tuple with the Warehouseidlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseidlist

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) SetWarehouseidlist(v string)`

SetWarehouseidlist sets Warehouseidlist field to given value.

### HasWarehouseidlist

`func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) HasWarehouseidlist() bool`

HasWarehouseidlist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


