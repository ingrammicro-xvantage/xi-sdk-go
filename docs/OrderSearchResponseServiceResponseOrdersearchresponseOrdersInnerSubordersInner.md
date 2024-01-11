# OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subordernumber** | Pointer to **string** | A sub order number | [optional] 
**Statuscode** | Pointer to **string** | Order status code | [optional] 
**Status** | Pointer to **string** | Details of the order statuscode - i.e. statuscode &#x3D; 4 then status &#x3D; SHIPPED | [optional] 
**Holdreasoncode** | Pointer to **string** | Will be returned in case of order on hold | [optional] 
**Holdreason** | Pointer to **string** | Reason for order hold - will be returned if the order is on hold | [optional] 
**Links** | Pointer to [**[]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner**](OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner.md) | HATEOAS links for the details and invoices of the sub-orders if available | [optional] 

## Methods

### NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner

`func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner`

NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerWithDefaults

`func NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerWithDefaults() *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner`

NewOrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerWithDefaults instantiates a new OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubordernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetSubordernumber() string`

GetSubordernumber returns the Subordernumber field if non-nil, zero value otherwise.

### GetSubordernumberOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetSubordernumberOk() (*string, bool)`

GetSubordernumberOk returns a tuple with the Subordernumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubordernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetSubordernumber(v string)`

SetSubordernumber sets Subordernumber field to given value.

### HasSubordernumber

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasSubordernumber() bool`

HasSubordernumber returns a boolean if a field has been set.

### GetStatuscode

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatuscode() string`

GetStatuscode returns the Statuscode field if non-nil, zero value otherwise.

### GetStatuscodeOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatuscodeOk() (*string, bool)`

GetStatuscodeOk returns a tuple with the Statuscode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuscode

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetStatuscode(v string)`

SetStatuscode sets Statuscode field to given value.

### HasStatuscode

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasStatuscode() bool`

HasStatuscode returns a boolean if a field has been set.

### GetStatus

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHoldreasoncode

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreasoncode() string`

GetHoldreasoncode returns the Holdreasoncode field if non-nil, zero value otherwise.

### GetHoldreasoncodeOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreasoncodeOk() (*string, bool)`

GetHoldreasoncodeOk returns a tuple with the Holdreasoncode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldreasoncode

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetHoldreasoncode(v string)`

SetHoldreasoncode sets Holdreasoncode field to given value.

### HasHoldreasoncode

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasHoldreasoncode() bool`

HasHoldreasoncode returns a boolean if a field has been set.

### GetHoldreason

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreason() string`

GetHoldreason returns the Holdreason field if non-nil, zero value otherwise.

### GetHoldreasonOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetHoldreasonOk() (*string, bool)`

GetHoldreasonOk returns a tuple with the Holdreason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldreason

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetHoldreason(v string)`

SetHoldreason sets Holdreason field to given value.

### HasHoldreason

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasHoldreason() bool`

HasHoldreason returns a boolean if a field has been set.

### GetLinks

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetLinks() []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) GetLinksOk() (*[]OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) SetLinks(v []OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderSearchResponseServiceResponseOrdersearchresponseOrdersInnerSubordersInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


