# OrderSearchResponseOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngramOrderNumber** | Pointer to **string** | The Ingram Micro order number. | [optional] 
**IngramOrderDate** | Pointer to **string** | The date the order was created(UTC). | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s order number for reference in their system. | [optional] 
**VendorSalesOrderNumber** | Pointer to **string** | The vendor&#39;s order number.(only for D-Type Orders) | [optional] 
**VendorName** | Pointer to **string** | The name of the vendor. | [optional] 
**EndUserCompanyName** | Pointer to **string** | The company name of the end user/customer. | [optional] 
**OrderTotal** | Pointer to **float32** | The total of the order. | [optional] 
**OrderStatus** | Pointer to **string** | The header-level status of the order.(OPEN/CLOSED/CANCELLED) | [optional] 
**SubOrders** | Pointer to [**[]OrderSearchResponseOrdersInnerSubOrdersInner**](OrderSearchResponseOrdersInnerSubOrdersInner.md) | Individual Ingram Micro order numbers associated with a single reseller PO. | [optional] 
**Links** | Pointer to [**OrderSearchResponseOrdersInnerLinks**](OrderSearchResponseOrdersInnerLinks.md) |  | [optional] 

## Methods

### NewOrderSearchResponseOrdersInner

`func NewOrderSearchResponseOrdersInner() *OrderSearchResponseOrdersInner`

NewOrderSearchResponseOrdersInner instantiates a new OrderSearchResponseOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSearchResponseOrdersInnerWithDefaults

`func NewOrderSearchResponseOrdersInnerWithDefaults() *OrderSearchResponseOrdersInner`

NewOrderSearchResponseOrdersInnerWithDefaults instantiates a new OrderSearchResponseOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngramOrderNumber

`func (o *OrderSearchResponseOrdersInner) GetIngramOrderNumber() string`

GetIngramOrderNumber returns the IngramOrderNumber field if non-nil, zero value otherwise.

### GetIngramOrderNumberOk

`func (o *OrderSearchResponseOrdersInner) GetIngramOrderNumberOk() (*string, bool)`

GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderNumber

`func (o *OrderSearchResponseOrdersInner) SetIngramOrderNumber(v string)`

SetIngramOrderNumber sets IngramOrderNumber field to given value.

### HasIngramOrderNumber

`func (o *OrderSearchResponseOrdersInner) HasIngramOrderNumber() bool`

HasIngramOrderNumber returns a boolean if a field has been set.

### GetIngramOrderDate

`func (o *OrderSearchResponseOrdersInner) GetIngramOrderDate() string`

GetIngramOrderDate returns the IngramOrderDate field if non-nil, zero value otherwise.

### GetIngramOrderDateOk

`func (o *OrderSearchResponseOrdersInner) GetIngramOrderDateOk() (*string, bool)`

GetIngramOrderDateOk returns a tuple with the IngramOrderDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramOrderDate

`func (o *OrderSearchResponseOrdersInner) SetIngramOrderDate(v string)`

SetIngramOrderDate sets IngramOrderDate field to given value.

### HasIngramOrderDate

`func (o *OrderSearchResponseOrdersInner) HasIngramOrderDate() bool`

HasIngramOrderDate returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *OrderSearchResponseOrdersInner) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderSearchResponseOrdersInner) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderSearchResponseOrdersInner) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderSearchResponseOrdersInner) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetVendorSalesOrderNumber

`func (o *OrderSearchResponseOrdersInner) GetVendorSalesOrderNumber() string`

GetVendorSalesOrderNumber returns the VendorSalesOrderNumber field if non-nil, zero value otherwise.

### GetVendorSalesOrderNumberOk

`func (o *OrderSearchResponseOrdersInner) GetVendorSalesOrderNumberOk() (*string, bool)`

GetVendorSalesOrderNumberOk returns a tuple with the VendorSalesOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSalesOrderNumber

`func (o *OrderSearchResponseOrdersInner) SetVendorSalesOrderNumber(v string)`

SetVendorSalesOrderNumber sets VendorSalesOrderNumber field to given value.

### HasVendorSalesOrderNumber

`func (o *OrderSearchResponseOrdersInner) HasVendorSalesOrderNumber() bool`

HasVendorSalesOrderNumber returns a boolean if a field has been set.

### GetVendorName

`func (o *OrderSearchResponseOrdersInner) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *OrderSearchResponseOrdersInner) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *OrderSearchResponseOrdersInner) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *OrderSearchResponseOrdersInner) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetEndUserCompanyName

`func (o *OrderSearchResponseOrdersInner) GetEndUserCompanyName() string`

GetEndUserCompanyName returns the EndUserCompanyName field if non-nil, zero value otherwise.

### GetEndUserCompanyNameOk

`func (o *OrderSearchResponseOrdersInner) GetEndUserCompanyNameOk() (*string, bool)`

GetEndUserCompanyNameOk returns a tuple with the EndUserCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserCompanyName

`func (o *OrderSearchResponseOrdersInner) SetEndUserCompanyName(v string)`

SetEndUserCompanyName sets EndUserCompanyName field to given value.

### HasEndUserCompanyName

`func (o *OrderSearchResponseOrdersInner) HasEndUserCompanyName() bool`

HasEndUserCompanyName returns a boolean if a field has been set.

### GetOrderTotal

`func (o *OrderSearchResponseOrdersInner) GetOrderTotal() float32`

GetOrderTotal returns the OrderTotal field if non-nil, zero value otherwise.

### GetOrderTotalOk

`func (o *OrderSearchResponseOrdersInner) GetOrderTotalOk() (*float32, bool)`

GetOrderTotalOk returns a tuple with the OrderTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTotal

`func (o *OrderSearchResponseOrdersInner) SetOrderTotal(v float32)`

SetOrderTotal sets OrderTotal field to given value.

### HasOrderTotal

`func (o *OrderSearchResponseOrdersInner) HasOrderTotal() bool`

HasOrderTotal returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderSearchResponseOrdersInner) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderSearchResponseOrdersInner) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderSearchResponseOrdersInner) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderSearchResponseOrdersInner) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetSubOrders

`func (o *OrderSearchResponseOrdersInner) GetSubOrders() []OrderSearchResponseOrdersInnerSubOrdersInner`

GetSubOrders returns the SubOrders field if non-nil, zero value otherwise.

### GetSubOrdersOk

`func (o *OrderSearchResponseOrdersInner) GetSubOrdersOk() (*[]OrderSearchResponseOrdersInnerSubOrdersInner, bool)`

GetSubOrdersOk returns a tuple with the SubOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrders

`func (o *OrderSearchResponseOrdersInner) SetSubOrders(v []OrderSearchResponseOrdersInnerSubOrdersInner)`

SetSubOrders sets SubOrders field to given value.

### HasSubOrders

`func (o *OrderSearchResponseOrdersInner) HasSubOrders() bool`

HasSubOrders returns a boolean if a field has been set.

### GetLinks

`func (o *OrderSearchResponseOrdersInner) GetLinks() OrderSearchResponseOrdersInnerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderSearchResponseOrdersInner) GetLinksOk() (*OrderSearchResponseOrdersInnerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderSearchResponseOrdersInner) SetLinks(v OrderSearchResponseOrdersInnerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderSearchResponseOrdersInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


