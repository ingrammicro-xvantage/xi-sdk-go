# OrderCreateResponseOrdersInnerLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**IngramLineNumber** | Pointer to **string** | The Ingram Micro line number for the product. | [optional] 
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line number for reference in their system. | [optional] 
**LineStatus** | Pointer to **string** | The status for the line item in the order. One of: Backordered, Open | [optional] 
**IngramPartNumber** | Pointer to **string** | The Ingram Micro part number for the line item. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor part number for the line item. | [optional] 
**UnitPrice** | Pointer to **float32** | The unit price for the line item. | [optional] 
**ExtendedUnitPrice** | Pointer to **float32** | The extended list price (unit price X quantity) for the line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity of the line item ordered. | [optional] 
**QuantityConfirmed** | Pointer to **int32** | The quantity of the line item that has been confirmed. | [optional] 
**QuantityBackOrdered** | Pointer to **int32** | The quantity of the line item that is backordered. | [optional] 
**SpecialBidNumber** | Pointer to **string** | The bid number for the line item provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers. | [optional] 
**Notes** | Pointer to **string** | Line-level notes. | [optional] 
**ShipmentDetails** | Pointer to [**[]OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner**](OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner.md) | The shipment details for the line item. | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner**](OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner.md) | SAP requested and country-specific line level details. | [optional] 

## Methods

### NewOrderCreateResponseOrdersInnerLinesInner

`func NewOrderCreateResponseOrdersInnerLinesInner() *OrderCreateResponseOrdersInnerLinesInner`

NewOrderCreateResponseOrdersInnerLinesInner instantiates a new OrderCreateResponseOrdersInnerLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseOrdersInnerLinesInnerWithDefaults

`func NewOrderCreateResponseOrdersInnerLinesInnerWithDefaults() *OrderCreateResponseOrdersInnerLinesInner`

NewOrderCreateResponseOrdersInnerLinesInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetIngramLineNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetLineStatus

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetLineStatus() string`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetLineStatusOk() (*string, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetLineStatus(v string)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetUnitPrice

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetExtendedUnitPrice

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetExtendedUnitPrice() float32`

GetExtendedUnitPrice returns the ExtendedUnitPrice field if non-nil, zero value otherwise.

### GetExtendedUnitPriceOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetExtendedUnitPriceOk() (*float32, bool)`

GetExtendedUnitPriceOk returns a tuple with the ExtendedUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedUnitPrice

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetExtendedUnitPrice(v float32)`

SetExtendedUnitPrice sets ExtendedUnitPrice field to given value.

### HasExtendedUnitPrice

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasExtendedUnitPrice() bool`

HasExtendedUnitPrice returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetQuantityConfirmed

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetQuantityConfirmed() int32`

GetQuantityConfirmed returns the QuantityConfirmed field if non-nil, zero value otherwise.

### GetQuantityConfirmedOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetQuantityConfirmedOk() (*int32, bool)`

GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConfirmed

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetQuantityConfirmed(v int32)`

SetQuantityConfirmed sets QuantityConfirmed field to given value.

### HasQuantityConfirmed

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasQuantityConfirmed() bool`

HasQuantityConfirmed returns a boolean if a field has been set.

### GetQuantityBackOrdered

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetQuantityBackOrdered() int32`

GetQuantityBackOrdered returns the QuantityBackOrdered field if non-nil, zero value otherwise.

### GetQuantityBackOrderedOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetQuantityBackOrderedOk() (*int32, bool)`

GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackOrdered

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetQuantityBackOrdered(v int32)`

SetQuantityBackOrdered sets QuantityBackOrdered field to given value.

### HasQuantityBackOrdered

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasQuantityBackOrdered() bool`

HasQuantityBackOrdered returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetNotes

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetShipmentDetails() []OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetShipmentDetailsOk() (*[]OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetShipmentDetails(v []OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetAdditionalAttributes() []OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderCreateResponseOrdersInnerLinesInner) GetAdditionalAttributesOk() (*[]OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderCreateResponseOrdersInnerLinesInner) SetAdditionalAttributes(v []OrderCreateResponseOrdersInnerLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderCreateResponseOrdersInnerLinesInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


