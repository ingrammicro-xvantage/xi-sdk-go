# QuoteToOrderDetailsDTOLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line item number for reference in their system. | [optional] 
**IngramPartNumber** | Pointer to **string** | Unique IngramMicro part number. | [optional] 
**Quantity** | Pointer to **string** | The quantity of the line item. | [optional] 
**VmfAdditionalAttributesLines** | Pointer to [**[]QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner**](QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner.md) | The object containing the list of fields required at a line level by the vendor. | [optional] 

## Methods

### NewQuoteToOrderDetailsDTOLinesInner

`func NewQuoteToOrderDetailsDTOLinesInner() *QuoteToOrderDetailsDTOLinesInner`

NewQuoteToOrderDetailsDTOLinesInner instantiates a new QuoteToOrderDetailsDTOLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteToOrderDetailsDTOLinesInnerWithDefaults

`func NewQuoteToOrderDetailsDTOLinesInnerWithDefaults() *QuoteToOrderDetailsDTOLinesInner`

NewQuoteToOrderDetailsDTOLinesInnerWithDefaults instantiates a new QuoteToOrderDetailsDTOLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLineNumber

`func (o *QuoteToOrderDetailsDTOLinesInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *QuoteToOrderDetailsDTOLinesInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *QuoteToOrderDetailsDTOLinesInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *QuoteToOrderDetailsDTOLinesInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *QuoteToOrderDetailsDTOLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *QuoteToOrderDetailsDTOLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *QuoteToOrderDetailsDTOLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *QuoteToOrderDetailsDTOLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *QuoteToOrderDetailsDTOLinesInner) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *QuoteToOrderDetailsDTOLinesInner) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *QuoteToOrderDetailsDTOLinesInner) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *QuoteToOrderDetailsDTOLinesInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetVmfAdditionalAttributesLines

`func (o *QuoteToOrderDetailsDTOLinesInner) GetVmfAdditionalAttributesLines() []QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner`

GetVmfAdditionalAttributesLines returns the VmfAdditionalAttributesLines field if non-nil, zero value otherwise.

### GetVmfAdditionalAttributesLinesOk

`func (o *QuoteToOrderDetailsDTOLinesInner) GetVmfAdditionalAttributesLinesOk() (*[]QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner, bool)`

GetVmfAdditionalAttributesLinesOk returns a tuple with the VmfAdditionalAttributesLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmfAdditionalAttributesLines

`func (o *QuoteToOrderDetailsDTOLinesInner) SetVmfAdditionalAttributesLines(v []QuoteToOrderDetailsDTOLinesInnerVmfAdditionalAttributesLinesInner)`

SetVmfAdditionalAttributesLines sets VmfAdditionalAttributesLines field to given value.

### HasVmfAdditionalAttributesLines

`func (o *QuoteToOrderDetailsDTOLinesInner) HasVmfAdditionalAttributesLines() bool`

HasVmfAdditionalAttributesLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


