# OrderStatusAsyncNotificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | Field for identifying whether it is a reseller or vendor event. For eg, resellers/orders | [optional] 
**Event** | Pointer to **string** | The event sent in the request. For eg, im::create. | [optional] 
**EventTimeStamp** | Pointer to **string** | The timestamp at which the event was sent. | [optional] 
**EventId** | Pointer to **string** | A unique id used as identifier for the sepcific event and used for generating the x-hub signature. | [optional] 
**Resource** | Pointer to [**[]OrderStatusAsyncNotificationRequestResourceInner**](OrderStatusAsyncNotificationRequestResourceInner.md) |  | [optional] 

## Methods

### NewOrderStatusAsyncNotificationRequest

`func NewOrderStatusAsyncNotificationRequest() *OrderStatusAsyncNotificationRequest`

NewOrderStatusAsyncNotificationRequest instantiates a new OrderStatusAsyncNotificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusAsyncNotificationRequestWithDefaults

`func NewOrderStatusAsyncNotificationRequestWithDefaults() *OrderStatusAsyncNotificationRequest`

NewOrderStatusAsyncNotificationRequestWithDefaults instantiates a new OrderStatusAsyncNotificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *OrderStatusAsyncNotificationRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *OrderStatusAsyncNotificationRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *OrderStatusAsyncNotificationRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *OrderStatusAsyncNotificationRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetEvent

`func (o *OrderStatusAsyncNotificationRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OrderStatusAsyncNotificationRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OrderStatusAsyncNotificationRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OrderStatusAsyncNotificationRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventTimeStamp

`func (o *OrderStatusAsyncNotificationRequest) GetEventTimeStamp() string`

GetEventTimeStamp returns the EventTimeStamp field if non-nil, zero value otherwise.

### GetEventTimeStampOk

`func (o *OrderStatusAsyncNotificationRequest) GetEventTimeStampOk() (*string, bool)`

GetEventTimeStampOk returns a tuple with the EventTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeStamp

`func (o *OrderStatusAsyncNotificationRequest) SetEventTimeStamp(v string)`

SetEventTimeStamp sets EventTimeStamp field to given value.

### HasEventTimeStamp

`func (o *OrderStatusAsyncNotificationRequest) HasEventTimeStamp() bool`

HasEventTimeStamp returns a boolean if a field has been set.

### GetEventId

`func (o *OrderStatusAsyncNotificationRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *OrderStatusAsyncNotificationRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *OrderStatusAsyncNotificationRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *OrderStatusAsyncNotificationRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetResource

`func (o *OrderStatusAsyncNotificationRequest) GetResource() []OrderStatusAsyncNotificationRequestResourceInner`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OrderStatusAsyncNotificationRequest) GetResourceOk() (*[]OrderStatusAsyncNotificationRequestResourceInner, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OrderStatusAsyncNotificationRequest) SetResource(v []OrderStatusAsyncNotificationRequestResourceInner)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OrderStatusAsyncNotificationRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


