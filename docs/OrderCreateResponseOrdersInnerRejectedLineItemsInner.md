# OrderCreateResponseOrdersInnerRejectedLineItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerLinenumber** | Pointer to **string** | The reseller&#39;s line item number of the rejected item for their reference. Number | [optional] 
**IngramPartNumber** | Pointer to **string** | The Ingram Micro part number for the rejected line item. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor part number for the rejected line item. | [optional] 
**QuantityOrdered** | Pointer to **int32** | The quantity ordered of the rejected line item. | [optional] 
**RejectCode** | Pointer to **string** | The rejection code for the rejected line item. Ex: &#39;EN&#39;  | [optional] 
**RejectReason** | Pointer to **string** | The rejection reason for the rejected line item. Ex: &#39;SKU-NOTFOUND    DF41281&#39;  | [optional] 

## Methods

### NewOrderCreateResponseOrdersInnerRejectedLineItemsInner

`func NewOrderCreateResponseOrdersInnerRejectedLineItemsInner() *OrderCreateResponseOrdersInnerRejectedLineItemsInner`

NewOrderCreateResponseOrdersInnerRejectedLineItemsInner instantiates a new OrderCreateResponseOrdersInnerRejectedLineItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseOrdersInnerRejectedLineItemsInnerWithDefaults

`func NewOrderCreateResponseOrdersInnerRejectedLineItemsInnerWithDefaults() *OrderCreateResponseOrdersInnerRejectedLineItemsInner`

NewOrderCreateResponseOrdersInnerRejectedLineItemsInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerRejectedLineItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerLinenumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetCustomerLinenumber() string`

GetCustomerLinenumber returns the CustomerLinenumber field if non-nil, zero value otherwise.

### GetCustomerLinenumberOk

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetCustomerLinenumberOk() (*string, bool)`

GetCustomerLinenumberOk returns a tuple with the CustomerLinenumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerLinenumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) SetCustomerLinenumber(v string)`

SetCustomerLinenumber sets CustomerLinenumber field to given value.

### HasCustomerLinenumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) HasCustomerLinenumber() bool`

HasCustomerLinenumber returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetQuantityOrdered

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetQuantityOrdered() int32`

GetQuantityOrdered returns the QuantityOrdered field if non-nil, zero value otherwise.

### GetQuantityOrderedOk

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetQuantityOrderedOk() (*int32, bool)`

GetQuantityOrderedOk returns a tuple with the QuantityOrdered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOrdered

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) SetQuantityOrdered(v int32)`

SetQuantityOrdered sets QuantityOrdered field to given value.

### HasQuantityOrdered

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) HasQuantityOrdered() bool`

HasQuantityOrdered returns a boolean if a field has been set.

### GetRejectCode

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetRejectCode() string`

GetRejectCode returns the RejectCode field if non-nil, zero value otherwise.

### GetRejectCodeOk

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetRejectCodeOk() (*string, bool)`

GetRejectCodeOk returns a tuple with the RejectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectCode

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) SetRejectCode(v string)`

SetRejectCode sets RejectCode field to given value.

### HasRejectCode

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) HasRejectCode() bool`

HasRejectCode returns a boolean if a field has been set.

### GetRejectReason

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *OrderCreateResponseOrdersInnerRejectedLineItemsInner) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


