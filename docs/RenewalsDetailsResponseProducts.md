# RenewalsDetailsResponseProducts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramLineNumber** | Pointer to **string** | Unique Ingram Micro line number. | [optional] 
**ProductDescription** | Pointer to **string** | The description of the product. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the line item. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique IngramMicro part number. | [optional] 
**ManufacturerPartNumber** | Pointer to **string** | The manufacturer&#39;s part number for the line item. | [optional] 
**UnitPrice** | Pointer to **float32** | The unit price of the line item. | [optional] 
**Quantity** | Pointer to **int32** | The quantity of the line item. | [optional] 
**IsConsolidated** | Pointer to **string** | Is the line item consolidated? Yes or No. | [optional] 

## Methods

### NewRenewalsDetailsResponseProducts

`func NewRenewalsDetailsResponseProducts() *RenewalsDetailsResponseProducts`

NewRenewalsDetailsResponseProducts instantiates a new RenewalsDetailsResponseProducts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewalsDetailsResponseProductsWithDefaults

`func NewRenewalsDetailsResponseProductsWithDefaults() *RenewalsDetailsResponseProducts`

NewRenewalsDetailsResponseProductsWithDefaults instantiates a new RenewalsDetailsResponseProducts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramLineNumber

`func (o *RenewalsDetailsResponseProducts) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *RenewalsDetailsResponseProducts) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *RenewalsDetailsResponseProducts) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *RenewalsDetailsResponseProducts) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetProductDescription

`func (o *RenewalsDetailsResponseProducts) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *RenewalsDetailsResponseProducts) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *RenewalsDetailsResponseProducts) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *RenewalsDetailsResponseProducts) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *RenewalsDetailsResponseProducts) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *RenewalsDetailsResponseProducts) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *RenewalsDetailsResponseProducts) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *RenewalsDetailsResponseProducts) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *RenewalsDetailsResponseProducts) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *RenewalsDetailsResponseProducts) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *RenewalsDetailsResponseProducts) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *RenewalsDetailsResponseProducts) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetManufacturerPartNumber

`func (o *RenewalsDetailsResponseProducts) GetManufacturerPartNumber() string`

GetManufacturerPartNumber returns the ManufacturerPartNumber field if non-nil, zero value otherwise.

### GetManufacturerPartNumberOk

`func (o *RenewalsDetailsResponseProducts) GetManufacturerPartNumberOk() (*string, bool)`

GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerPartNumber

`func (o *RenewalsDetailsResponseProducts) SetManufacturerPartNumber(v string)`

SetManufacturerPartNumber sets ManufacturerPartNumber field to given value.

### HasManufacturerPartNumber

`func (o *RenewalsDetailsResponseProducts) HasManufacturerPartNumber() bool`

HasManufacturerPartNumber returns a boolean if a field has been set.

### GetUnitPrice

`func (o *RenewalsDetailsResponseProducts) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *RenewalsDetailsResponseProducts) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *RenewalsDetailsResponseProducts) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *RenewalsDetailsResponseProducts) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *RenewalsDetailsResponseProducts) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *RenewalsDetailsResponseProducts) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *RenewalsDetailsResponseProducts) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *RenewalsDetailsResponseProducts) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetIsConsolidated

`func (o *RenewalsDetailsResponseProducts) GetIsConsolidated() string`

GetIsConsolidated returns the IsConsolidated field if non-nil, zero value otherwise.

### GetIsConsolidatedOk

`func (o *RenewalsDetailsResponseProducts) GetIsConsolidatedOk() (*string, bool)`

GetIsConsolidatedOk returns a tuple with the IsConsolidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConsolidated

`func (o *RenewalsDetailsResponseProducts) SetIsConsolidated(v string)`

SetIsConsolidated sets IsConsolidated field to given value.

### HasIsConsolidated

`func (o *RenewalsDetailsResponseProducts) HasIsConsolidated() bool`

HasIsConsolidated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


