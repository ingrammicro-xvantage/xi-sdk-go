# OrderModifyResponseLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**IngramLineNumber** | Pointer to **string** | The IngramMicro line number. | [optional] 
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line number for reference in their system. | [optional] 
**IngramPartNumber** | Pointer to **string** | The unique IngramMicro part number for the line item. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity ordered of the line item. | [optional] 
**QuantityConfirmed** | Pointer to **int32** | The quantity confirmed of the line item. | [optional] 
**QuantityBackOrdered** | Pointer to **int32** | The quantity backordered of the line item. | [optional] 
**ShipmentDetails** | Pointer to [**OrderModifyResponseLinesInnerShipmentDetails**](OrderModifyResponseLinesInnerShipmentDetails.md) |  | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderModifyResponseLinesInnerAdditionalAttributesInner**](OrderModifyResponseLinesInnerAdditionalAttributesInner.md) | SAP requested and country-specific line level details. | [optional] 
**Notes** | Pointer to **string** | Line-level notes for the order. | [optional] 

## Methods

### NewOrderModifyResponseLinesInner

`func NewOrderModifyResponseLinesInner() *OrderModifyResponseLinesInner`

NewOrderModifyResponseLinesInner instantiates a new OrderModifyResponseLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderModifyResponseLinesInnerWithDefaults

`func NewOrderModifyResponseLinesInnerWithDefaults() *OrderModifyResponseLinesInner`

NewOrderModifyResponseLinesInnerWithDefaults instantiates a new OrderModifyResponseLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderModifyResponseLinesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderModifyResponseLinesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderModifyResponseLinesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderModifyResponseLinesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetIngramLineNumber

`func (o *OrderModifyResponseLinesInner) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *OrderModifyResponseLinesInner) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *OrderModifyResponseLinesInner) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *OrderModifyResponseLinesInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *OrderModifyResponseLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderModifyResponseLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderModifyResponseLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderModifyResponseLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderModifyResponseLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderModifyResponseLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderModifyResponseLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderModifyResponseLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderModifyResponseLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderModifyResponseLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderModifyResponseLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderModifyResponseLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderModifyResponseLinesInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderModifyResponseLinesInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderModifyResponseLinesInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderModifyResponseLinesInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetQuantityConfirmed

`func (o *OrderModifyResponseLinesInner) GetQuantityConfirmed() int32`

GetQuantityConfirmed returns the QuantityConfirmed field if non-nil, zero value otherwise.

### GetQuantityConfirmedOk

`func (o *OrderModifyResponseLinesInner) GetQuantityConfirmedOk() (*int32, bool)`

GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConfirmed

`func (o *OrderModifyResponseLinesInner) SetQuantityConfirmed(v int32)`

SetQuantityConfirmed sets QuantityConfirmed field to given value.

### HasQuantityConfirmed

`func (o *OrderModifyResponseLinesInner) HasQuantityConfirmed() bool`

HasQuantityConfirmed returns a boolean if a field has been set.

### GetQuantityBackOrdered

`func (o *OrderModifyResponseLinesInner) GetQuantityBackOrdered() int32`

GetQuantityBackOrdered returns the QuantityBackOrdered field if non-nil, zero value otherwise.

### GetQuantityBackOrderedOk

`func (o *OrderModifyResponseLinesInner) GetQuantityBackOrderedOk() (*int32, bool)`

GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBackOrdered

`func (o *OrderModifyResponseLinesInner) SetQuantityBackOrdered(v int32)`

SetQuantityBackOrdered sets QuantityBackOrdered field to given value.

### HasQuantityBackOrdered

`func (o *OrderModifyResponseLinesInner) HasQuantityBackOrdered() bool`

HasQuantityBackOrdered returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderModifyResponseLinesInner) GetShipmentDetails() OrderModifyResponseLinesInnerShipmentDetails`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderModifyResponseLinesInner) GetShipmentDetailsOk() (*OrderModifyResponseLinesInnerShipmentDetails, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderModifyResponseLinesInner) SetShipmentDetails(v OrderModifyResponseLinesInnerShipmentDetails)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderModifyResponseLinesInner) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderModifyResponseLinesInner) GetAdditionalAttributes() []OrderModifyResponseLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderModifyResponseLinesInner) GetAdditionalAttributesOk() (*[]OrderModifyResponseLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderModifyResponseLinesInner) SetAdditionalAttributes(v []OrderModifyResponseLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderModifyResponseLinesInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### GetNotes

`func (o *OrderModifyResponseLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderModifyResponseLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderModifyResponseLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderModifyResponseLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


