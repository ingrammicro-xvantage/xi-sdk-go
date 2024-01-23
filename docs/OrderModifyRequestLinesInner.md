# OrderModifyRequestLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramPartNumber** | Pointer to **string** | The unique IngramMicro part number. | [optional] 
**IngramLineNumber** | Pointer to **string** | The IngramMicro line number. | [optional] 
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line number for reference in their system. | [optional] 
**AddUpdateDeleteLine** | Pointer to **string** | The line number that was added, updated, or deleted. | [optional] 
**Quantity** | Pointer to **int32** | The quantity of the line item. | [optional] 
**Notes** | Pointer to **string** | The line-level notes. | [optional] 

## Methods

### NewOrderModifyRequestLinesInner

`func NewOrderModifyRequestLinesInner() *OrderModifyRequestLinesInner`

NewOrderModifyRequestLinesInner instantiates a new OrderModifyRequestLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderModifyRequestLinesInnerWithDefaults

`func NewOrderModifyRequestLinesInnerWithDefaults() *OrderModifyRequestLinesInner`

NewOrderModifyRequestLinesInnerWithDefaults instantiates a new OrderModifyRequestLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramPartNumber

`func (o *OrderModifyRequestLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderModifyRequestLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderModifyRequestLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderModifyRequestLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetIngramLineNumber

`func (o *OrderModifyRequestLinesInner) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *OrderModifyRequestLinesInner) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *OrderModifyRequestLinesInner) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *OrderModifyRequestLinesInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *OrderModifyRequestLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderModifyRequestLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderModifyRequestLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderModifyRequestLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetAddUpdateDeleteLine

`func (o *OrderModifyRequestLinesInner) GetAddUpdateDeleteLine() string`

GetAddUpdateDeleteLine returns the AddUpdateDeleteLine field if non-nil, zero value otherwise.

### GetAddUpdateDeleteLineOk

`func (o *OrderModifyRequestLinesInner) GetAddUpdateDeleteLineOk() (*string, bool)`

GetAddUpdateDeleteLineOk returns a tuple with the AddUpdateDeleteLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUpdateDeleteLine

`func (o *OrderModifyRequestLinesInner) SetAddUpdateDeleteLine(v string)`

SetAddUpdateDeleteLine sets AddUpdateDeleteLine field to given value.

### HasAddUpdateDeleteLine

`func (o *OrderModifyRequestLinesInner) HasAddUpdateDeleteLine() bool`

HasAddUpdateDeleteLine returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderModifyRequestLinesInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderModifyRequestLinesInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderModifyRequestLinesInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderModifyRequestLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetNotes

`func (o *OrderModifyRequestLinesInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderModifyRequestLinesInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderModifyRequestLinesInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderModifyRequestLinesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


