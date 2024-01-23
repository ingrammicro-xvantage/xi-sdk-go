# OrderModifyResponseRejectedLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramLineNumber** | Pointer to **string** | The IngramMicro line number for the failed line item. | [optional] 
**CustomerLineNumber** | Pointer to **string** | The reseller&#39;s line number of the failed line item for reference in their system. | [optional] 
**IngramPartNumber** | Pointer to **string** | The IngramMicro part number for the failed line item. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor&#39;s part number for the failed line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity ordered of the failed line item. | [optional] 
**RejectCode** | Pointer to **string** | The rejection code for the failed line item. | [optional] 
**RejectReason** | Pointer to **string** | The rejection reason for the failed line item. | [optional] 

## Methods

### NewOrderModifyResponseRejectedLineItemsInner

`func NewOrderModifyResponseRejectedLineItemsInner() *OrderModifyResponseRejectedLineItemsInner`

NewOrderModifyResponseRejectedLineItemsInner instantiates a new OrderModifyResponseRejectedLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderModifyResponseRejectedLineItemsInnerWithDefaults

`func NewOrderModifyResponseRejectedLineItemsInnerWithDefaults() *OrderModifyResponseRejectedLineItemsInner`

NewOrderModifyResponseRejectedLineItemsInnerWithDefaults instantiates a new OrderModifyResponseRejectedLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramLineNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramLineNumber() string`

GetIngramLineNumber returns the IngramLineNumber field if non-nil, zero value otherwise.

### GetIngramLineNumberOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramLineNumberOk() (*string, bool)`

GetIngramLineNumberOk returns a tuple with the IngramLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramLineNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) SetIngramLineNumber(v string)`

SetIngramLineNumber sets IngramLineNumber field to given value.

### HasIngramLineNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) HasIngramLineNumber() bool`

HasIngramLineNumber returns a boolean if a field has been set.

### GetCustomerLineNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) GetCustomerLineNumber() string`

GetCustomerLineNumber returns the CustomerLineNumber field if non-nil, zero value otherwise.

### GetCustomerLineNumberOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetCustomerLineNumberOk() (*string, bool)`

GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLineNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) SetCustomerLineNumber(v string)`

SetCustomerLineNumber sets CustomerLineNumber field to given value.

### HasCustomerLineNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) HasCustomerLineNumber() bool`

HasCustomerLineNumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderModifyResponseRejectedLineItemsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderModifyResponseRejectedLineItemsInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderModifyResponseRejectedLineItemsInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderModifyResponseRejectedLineItemsInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetRejectCode

`func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectCode() string`

GetRejectCode returns the RejectCode field if non-nil, zero value otherwise.

### GetRejectCodeOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectCodeOk() (*string, bool)`

GetRejectCodeOk returns a tuple with the RejectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectCode

`func (o *OrderModifyResponseRejectedLineItemsInner) SetRejectCode(v string)`

SetRejectCode sets RejectCode field to given value.

### HasRejectCode

`func (o *OrderModifyResponseRejectedLineItemsInner) HasRejectCode() bool`

HasRejectCode returns a boolean if a field has been set.

### GetRejectReason

`func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *OrderModifyResponseRejectedLineItemsInner) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *OrderModifyResponseRejectedLineItemsInner) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


