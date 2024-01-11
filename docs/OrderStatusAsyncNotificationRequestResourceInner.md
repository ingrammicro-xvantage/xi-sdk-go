# OrderStatusAsyncNotificationRequestResourceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The event name sent in the event request. | [optional] 
**OrderNumber** | Pointer to **string** | The Ingram Micro order number. | [optional] 
**CustomerOrderNumber** | Pointer to **string** | The reseller&#39;s unique PO/Order number. | [optional] 
**OrderEntryTimeStamp** | Pointer to **string** | The timestamp at which the order was created. | [optional] 
**Lines** | Pointer to [**[]OrderStatusAsyncNotificationRequestResourceInnerLinesInner**](OrderStatusAsyncNotificationRequestResourceInnerLinesInner.md) | The line-level details for the order. | [optional] 
**Links** | Pointer to [**[]OrderStatusAsyncNotificationRequestResourceInnerLinksInner**](OrderStatusAsyncNotificationRequestResourceInnerLinksInner.md) | Link to Order Details for the order(s). | [optional] 

## Methods

### NewOrderStatusAsyncNotificationRequestResourceInner

`func NewOrderStatusAsyncNotificationRequestResourceInner() *OrderStatusAsyncNotificationRequestResourceInner`

NewOrderStatusAsyncNotificationRequestResourceInner instantiates a new OrderStatusAsyncNotificationRequestResourceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusAsyncNotificationRequestResourceInnerWithDefaults

`func NewOrderStatusAsyncNotificationRequestResourceInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInner`

NewOrderStatusAsyncNotificationRequestResourceInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OrderStatusAsyncNotificationRequestResourceInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *OrderStatusAsyncNotificationRequestResourceInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInner) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.

### HasOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInner) HasOrderNumber() bool`

HasOrderNumber returns a boolean if a field has been set.

### GetCustomerOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetCustomerOrderNumber() string`

GetCustomerOrderNumber returns the CustomerOrderNumber field if non-nil, zero value otherwise.

### GetCustomerOrderNumberOk

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetCustomerOrderNumberOk() (*string, bool)`

GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInner) SetCustomerOrderNumber(v string)`

SetCustomerOrderNumber sets CustomerOrderNumber field to given value.

### HasCustomerOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInner) HasCustomerOrderNumber() bool`

HasCustomerOrderNumber returns a boolean if a field has been set.

### GetOrderEntryTimeStamp

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetOrderEntryTimeStamp() string`

GetOrderEntryTimeStamp returns the OrderEntryTimeStamp field if non-nil, zero value otherwise.

### GetOrderEntryTimeStampOk

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetOrderEntryTimeStampOk() (*string, bool)`

GetOrderEntryTimeStampOk returns a tuple with the OrderEntryTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderEntryTimeStamp

`func (o *OrderStatusAsyncNotificationRequestResourceInner) SetOrderEntryTimeStamp(v string)`

SetOrderEntryTimeStamp sets OrderEntryTimeStamp field to given value.

### HasOrderEntryTimeStamp

`func (o *OrderStatusAsyncNotificationRequestResourceInner) HasOrderEntryTimeStamp() bool`

HasOrderEntryTimeStamp returns a boolean if a field has been set.

### GetLines

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetLines() []OrderStatusAsyncNotificationRequestResourceInnerLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetLinesOk() (*[]OrderStatusAsyncNotificationRequestResourceInnerLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderStatusAsyncNotificationRequestResourceInner) SetLines(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderStatusAsyncNotificationRequestResourceInner) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLinks

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetLinks() []OrderStatusAsyncNotificationRequestResourceInnerLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrderStatusAsyncNotificationRequestResourceInner) GetLinksOk() (*[]OrderStatusAsyncNotificationRequestResourceInnerLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrderStatusAsyncNotificationRequestResourceInner) SetLinks(v []OrderStatusAsyncNotificationRequestResourceInnerLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OrderStatusAsyncNotificationRequestResourceInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


