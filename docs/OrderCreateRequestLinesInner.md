# OrderCreateRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. The customer line number needs to be a unique numeric value between 1 and 884. In the event we receive duplicate values or alphanumeric values in the customer line number, we will re-sequence the customer line number. To prevent re-sequencing, please use a unique numeric value between 1 and 884 in the customer line number. | [optional] 
**IngramPartNumber** | Pointer to **string** | The unique IngramMicro part number. | [optional] 
**Quantity** | Pointer to **int32** | The requested quantity of the line item. | [optional] 
**SpecialBidNumber** | Pointer to **string** | The line-level bid number provided to the reseller by the vendor for special pricing and discounts. Used to track the bid number in the case of split orders or where different line items have different bid numbers. Line-level bid number take precedence over header-level bid numbers. | [optional] 
**Notes** | Pointer to **string** | Line-level notes. | [optional] 
**UnitPrice** | Pointer to **float32** | The reseller-requested unit price for the line item. The unit price is not guaranteed. | [optional] 
**EndUserPrice** | Pointer to **float32** | The end user price. | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderCreateRequestLinesInnerAdditionalAttributesInner**](OrderCreateRequestLinesInnerAdditionalAttributesInner.md) |  | [optional] 
**WarrantyInfo** | Pointer to [**[]OrderCreateRequestLinesInnerWarrantyInfoInner**](OrderCreateRequestLinesInnerWarrantyInfoInner.md) | Warranty details for the line. This is required in case of warranty orders. | [optional] 
**EndUserInfo** | Pointer to [**[]OrderCreateRequestLinesInnerEndUserInfoInner**](OrderCreateRequestLinesInnerEndUserInfoInner.md) |  | [optional] 

## Methods

### NewOrderCreateRequestLinesInner

`func NewOrderCreateRequestLinesInner() *OrderCreateRequestLinesInner`

NewOrderCreateRequestLinesInner instantiates a new OrderCreateRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestLinesInnerWithDefaults

`func NewOrderCreateRequestLinesInnerWithDefaults() *OrderCreateRequestLinesInner`

NewOrderCreateRequestLinesInnerWithDefaults instantiates a new OrderCreateRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLineNumber

`func (o *OrderCreateRequestLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderCreateRequestLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderCreateRequestLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderCreateRequestLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderCreateRequestLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderCreateRequestLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderCreateRequestLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderCreateRequestLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderCreateRequestLinesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderCreateRequestLinesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderCreateRequestLinesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderCreateRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *OrderCreateRequestLinesInner) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *OrderCreateRequestLinesInner) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *OrderCreateRequestLinesInner) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *OrderCreateRequestLinesInner) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetNotes

`func (o *OrderCreateRequestLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderCreateRequestLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderCreateRequestLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderCreateRequestLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUnitPrice

`func (o *OrderCreateRequestLinesInner) GetUnitPrice() float32`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *OrderCreateRequestLinesInner) GetUnitPriceOk() (*float32, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *OrderCreateRequestLinesInner) SetUnitPrice(v float32)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *OrderCreateRequestLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetEndUserPrice

`func (o *OrderCreateRequestLinesInner) GetEndUserPrice() float32`

GetEndUserPrice returns the EndUserPrice field if non-nil, zero value otherwise.

### GetEndUserPriceOk

`func (o *OrderCreateRequestLinesInner) GetEndUserPriceOk() (*float32, bool)`

GetEndUserPriceOk returns a tuple with the EndUserPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserPrice

`func (o *OrderCreateRequestLinesInner) SetEndUserPrice(v float32)`

SetEndUserPrice sets EndUserPrice field to given value.

### HasEndUserPrice

`func (o *OrderCreateRequestLinesInner) HasEndUserPrice() bool`

HasEndUserPrice returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderCreateRequestLinesInner) GetAdditionalAttributes() []OrderCreateRequestLinesInnerAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderCreateRequestLinesInner) GetAdditionalAttributesOk() (*[]OrderCreateRequestLinesInnerAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderCreateRequestLinesInner) SetAdditionalAttributes(v []OrderCreateRequestLinesInnerAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderCreateRequestLinesInner) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.

### GetWarrantyInfo

`func (o *OrderCreateRequestLinesInner) GetWarrantyInfo() []OrderCreateRequestLinesInnerWarrantyInfoInner`

GetWarrantyInfo returns the WarrantyInfo field if non-nil, zero value otherwise.

### GetWarrantyInfoOk

`func (o *OrderCreateRequestLinesInner) GetWarrantyInfoOk() (*[]OrderCreateRequestLinesInnerWarrantyInfoInner, bool)`

GetWarrantyInfoOk returns a tuple with the WarrantyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyInfo

`func (o *OrderCreateRequestLinesInner) SetWarrantyInfo(v []OrderCreateRequestLinesInnerWarrantyInfoInner)`

SetWarrantyInfo sets WarrantyInfo field to given value.

### HasWarrantyInfo

`func (o *OrderCreateRequestLinesInner) HasWarrantyInfo() bool`

HasWarrantyInfo returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *OrderCreateRequestLinesInner) GetEndUserInfo() []OrderCreateRequestLinesInnerEndUserInfoInner`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *OrderCreateRequestLinesInner) GetEndUserInfoOk() (*[]OrderCreateRequestLinesInnerEndUserInfoInner, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *OrderCreateRequestLinesInner) SetEndUserInfo(v []OrderCreateRequestLinesInnerEndUserInfoInner)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *OrderCreateRequestLinesInner) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


