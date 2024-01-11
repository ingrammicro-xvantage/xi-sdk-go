# PriceAndAvailabilityRequestProductsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramPartNumber** | Pointer to **string** | Ingram Micro unique part number for the product. | [optional] 
**VendorPartNumber** | Pointer to **string** | Vendor’s part number for the product. | [optional] 
**CustomerPartNumber** | Pointer to **string** | Reseller/end-user’s part number for the product. | [optional] 
**Upc** | Pointer to **string** | The UPC code for the product. Consists of 12 numeric digits that are uniquely assigned to each trade item. | [optional] 
**QuantityRequested** | Pointer to **string** | Number of quantity of the Product. | [optional] 
**AdditionalAttributes** | Pointer to [**[]PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner**](PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner.md) |  | [optional] 

## Methods

### NewPriceAndAvailabilityRequestProductsInner

`func NewPriceAndAvailabilityRequestProductsInner() *PriceAndAvailabilityRequestProductsInner`

NewPriceAndAvailabilityRequestProductsInner instantiates a new PriceAndAvailabilityRequestProductsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceAndAvailabilityRequestProductsInnerWithDefaults

`func NewPriceAndAvailabilityRequestProductsInnerWithDefaults() *PriceAndAvailabilityRequestProductsInner`

NewPriceAndAvailabilityRequestProductsInnerWithDefaults instantiates a new PriceAndAvailabilityRequestProductsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetCustomerPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) GetCustomerPartNumber() string`

GetCustomerPartNumber returns the CustomerPartNumber field if non-nil, zero value otherwise.

### GetCustomerPartNumberOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetCustomerPartNumberOk() (*string, bool)`

GetCustomerPartNumberOk returns a tuple with the CustomerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) SetCustomerPartNumber(v string)`

SetCustomerPartNumber sets CustomerPartNumber field to given value.

### HasCustomerPartNumber

`func (o *PriceAndAvailabilityRequestProductsInner) HasCustomerPartNumber() bool`

HasCustomerPartNumber returns a boolean if a field has been set.

### GetUpc

`func (o *PriceAndAvailabilityRequestProductsInner) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *PriceAndAvailabilityRequestProductsInner) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *PriceAndAvailabilityRequestProductsInner) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### GetQuantityRequested

`func (o *PriceAndAvailabilityRequestProductsInner) GetQuantityRequested() string`

GetQuantityRequested returns the QuantityRequested field if non-nil, zero value otherwise.

### GetQuantityRequestedOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetQuantityRequestedOk() (*string, bool)`

GetQuantityRequestedOk returns a tuple with the QuantityRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityRequested

`func (o *PriceAndAvailabilityRequestProductsInner) SetQuantityRequested(v string)`

SetQuantityRequested sets QuantityRequested field to given value.

### HasQuantityRequested

`func (o *PriceAndAvailabilityRequestProductsInner) HasQuantityRequested() bool`

HasQuantityRequested returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *PriceAndAvailabilityRequestProductsInner) GetAdditionalAttributes() []PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *PriceAndAvailabilityRequestProductsInner) GetAdditionalAttributesOk() (*[]PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *PriceAndAvailabilityRequestProductsInner) SetAdditionalAttributes(v []PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *PriceAndAvailabilityRequestProductsInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


