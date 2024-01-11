# OrderSearchResponseOrdersInnerSubOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest to the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**SubOrderTotal** | Pointer to **float32** | The total for the suborder. | [optional] 
**SubOrderStatus** | Pointer to **string** | The status of the suborder. One of:- Shipped, Canceled, Backordered, Processing, On Hold, Delivered | [optional] 
**Links** | Pointer to [**[]OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner**](OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner.md) | Link to Order Details for the sub order(s). | [optional] 

## Methods

### NewOrderSearchResponseOrdersInnerSubOrdersInner

`func NewOrderSearchResponseOrdersInnerSubOrdersInner() *OrderSearchResponseOrdersInnerSubOrdersInner`

NewOrderSearchResponseOrdersInnerSubOrdersInner instantiates a new OrderSearchResponseOrdersInnerSubOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResponseOrdersInnerSubOrdersInnerWithDefaults

`func NewOrderSearchResponseOrdersInnerSubOrdersInnerWithDefaults() *OrderSearchResponseOrdersInnerSubOrdersInner`

NewOrderSearchResponseOrdersInnerSubOrdersInnerWithDefaults instantiates a new OrderSearchResponseOrdersInnerSubOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetSubOrderTotal

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetSubOrderTotal() float32`

GetSubOrderTotal returns the SubOrderTotal field if non-nil, zero value otherwise.

### GetSubOrderTotalOk

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetSubOrderTotalOk() (*float32, bool)`

GetSubOrderTotalOk returns a tuple with the SubOrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderTotal

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) SetSubOrderTotal(v float32)`

SetSubOrderTotal sets SubOrderTotal field to given value.

### HasSubOrderTotal

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) HasSubOrderTotal() bool`

HasSubOrderTotal returns a boolean if a field has been set.

### GetSubOrderStatus

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetSubOrderStatus() string`

GetSubOrderStatus returns the SubOrderStatus field if non-nil, zero value otherwise.

### GetSubOrderStatusOk

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetSubOrderStatusOk() (*string, bool)`

GetSubOrderStatusOk returns a tuple with the SubOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderStatus

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) SetSubOrderStatus(v string)`

SetSubOrderStatus sets SubOrderStatus field to given value.

### HasSubOrderStatus

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) HasSubOrderStatus() bool`

HasSubOrderStatus returns a boolean if a field has been set.

### GetLinks

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetLinks() []OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) GetLinksOk() (*[]OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) SetLinks(v []OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderSearchResponseOrdersInnerSubOrdersInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


