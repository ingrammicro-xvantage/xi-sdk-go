# AsyncOrderCreateDTOLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique IngramMicro part number. | [optional] 
**Quantity** | Pointer to **string** | The quantity of the line item. | [optional] 
**UnitPrice** | Pointer to **string** | Unit Price of Item | [optional] 
**SpecialBidNumber** | Pointer to **string** |  | [optional] 
**EndUserPrice** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**EndUserInfo** | Pointer to [**[]AsyncOrderCreateDTOLinesInnerEndUserInfoInner**](AsyncOrderCreateDTOLinesInnerEndUserInfoInner.md) | The contact information for the end user/customer provided by the reseller. Used to determine pricing and discounts. | [optional] 

## Methods

### NewAsyncOrderCreateDTOLinesInner

`func NewAsyncOrderCreateDTOLinesInner() *AsyncOrderCreateDTOLinesInner`

NewAsyncOrderCreateDTOLinesInner instantiates a new AsyncOrderCreateDTOLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncOrderCreateDTOLinesInnerWithDefaults

`func NewAsyncOrderCreateDTOLinesInnerWithDefaults() *AsyncOrderCreateDTOLinesInner`

NewAsyncOrderCreateDTOLinesInnerWithDefaults instantiates a new AsyncOrderCreateDTOLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLineNumber

`func (o *AsyncOrderCreateDTOLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *AsyncOrderCreateDTOLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *AsyncOrderCreateDTOLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *AsyncOrderCreateDTOLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *AsyncOrderCreateDTOLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *AsyncOrderCreateDTOLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *AsyncOrderCreateDTOLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *AsyncOrderCreateDTOLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *AsyncOrderCreateDTOLinesInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AsyncOrderCreateDTOLinesInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AsyncOrderCreateDTOLinesInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AsyncOrderCreateDTOLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetUnitPrice

`func (o *AsyncOrderCreateDTOLinesInner) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *AsyncOrderCreateDTOLinesInner) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *AsyncOrderCreateDTOLinesInner) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *AsyncOrderCreateDTOLinesInner) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### GetSpecialBidNumber

`func (o *AsyncOrderCreateDTOLinesInner) GetSpecialBidNumber() string`

GetSpecialBidNumber returns the SpecialBidNumber field if non-nil, zero value otherwise.

### GetSpecialBidNumberOk

`func (o *AsyncOrderCreateDTOLinesInner) GetSpecialBidNumberOk() (*string, bool)`

GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialBidNumber

`func (o *AsyncOrderCreateDTOLinesInner) SetSpecialBidNumber(v string)`

SetSpecialBidNumber sets SpecialBidNumber field to given value.

### HasSpecialBidNumber

`func (o *AsyncOrderCreateDTOLinesInner) HasSpecialBidNumber() bool`

HasSpecialBidNumber returns a boolean if a field has been set.

### GetEndUserPrice

`func (o *AsyncOrderCreateDTOLinesInner) GetEndUserPrice() string`

GetEndUserPrice returns the EndUserPrice field if non-nil, zero value otherwise.

### GetEndUserPriceOk

`func (o *AsyncOrderCreateDTOLinesInner) GetEndUserPriceOk() (*string, bool)`

GetEndUserPriceOk returns a tuple with the EndUserPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserPrice

`func (o *AsyncOrderCreateDTOLinesInner) SetEndUserPrice(v string)`

SetEndUserPrice sets EndUserPrice field to given value.

### HasEndUserPrice

`func (o *AsyncOrderCreateDTOLinesInner) HasEndUserPrice() bool`

HasEndUserPrice returns a boolean if a field has been set.

### GetNotes

`func (o *AsyncOrderCreateDTOLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AsyncOrderCreateDTOLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AsyncOrderCreateDTOLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AsyncOrderCreateDTOLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetEndUserInfo

`func (o *AsyncOrderCreateDTOLinesInner) GetEndUserInfo() []AsyncOrderCreateDTOLinesInnerEndUserInfoInner`

GetEndUserInfo returns the EndUserInfo field if non-nil, zero value otherwise.

### GetEndUserInfoOk

`func (o *AsyncOrderCreateDTOLinesInner) GetEndUserInfoOk() (*[]AsyncOrderCreateDTOLinesInnerEndUserInfoInner, bool)`

GetEndUserInfoOk returns a tuple with the EndUserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserInfo

`func (o *AsyncOrderCreateDTOLinesInner) SetEndUserInfo(v []AsyncOrderCreateDTOLinesInnerEndUserInfoInner)`

SetEndUserInfo sets EndUserInfo field to given value.

### HasEndUserInfo

`func (o *AsyncOrderCreateDTOLinesInner) HasEndUserInfo() bool`

HasEndUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


