# OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ordernumber** | **string** | Ingram micro sales order number | 
**Entrytimestamp** | **string** | The order creation date-time in UTC format | 
**Customerordernumber** | Pointer to **string** | PO/Order number submitted while creating the order | [optional] 
**Suborders** | Pointer to [**[]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner**](OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner.md) | An order MAY get divided into various sub orders, for example if the SKUs are being shipped from different warehouse. | [optional] 
**Links** | Pointer to [**OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks**](OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks.md) |  | [optional] 

## Methods

### NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInner

`func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInner(ordernumber string, entrytimestamp string, ) *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner`

NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInner instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerWithDefaults

`func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerWithDefaults() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner`

NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerWithDefaults instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrdernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetOrdernumber() string`

GetOrdernumber returns the Ordernumber field if non-nil, zero value otherwise.

### GetOrdernumberOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetOrdernumberOk() (*string, bool)`

GetOrdernumberOk returns a tuple with the Ordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) SetOrdernumber(v string)`

SetOrdernumber sets Ordernumber field to given value.


### GetEntrytimestamp

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetEntrytimestamp() string`

GetEntrytimestamp returns the Entrytimestamp field if non-nil, zero value otherwise.

### GetEntrytimestampOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetEntrytimestampOk() (*string, bool)`

GetEntrytimestampOk returns a tuple with the Entrytimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrytimestamp

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) SetEntrytimestamp(v string)`

SetEntrytimestamp sets Entrytimestamp field to given value.


### GetCustomerordernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetCustomerordernumber() string`

GetCustomerordernumber returns the Customerordernumber field if non-nil, zero value otherwise.

### GetCustomerordernumberOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetCustomerordernumberOk() (*string, bool)`

GetCustomerordernumberOk returns a tuple with the Customerordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerordernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) SetCustomerordernumber(v string)`

SetCustomerordernumber sets Customerordernumber field to given value.

### HasCustomerordernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) HasCustomerordernumber() bool`

HasCustomerordernumber returns a boolean if a field has been set.

### GetSuborders

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetSuborders() []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner`

GetSuborders returns the Suborders field if non-nil, zero value otherwise.

### GetSubordersOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetSubordersOk() (*[]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner, bool)`

GetSubordersOk returns a tuple with the Suborders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuborders

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) SetSuborders(v []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner)`

SetSuborders sets Suborders field to given value.

### HasSuborders

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) HasSuborders() bool`

HasSuborders returns a boolean if a field has been set.

### GetLinks

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetLinks() OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) GetLinksOk() (*OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) SetLinks(v OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


