# OrderCreateV7RequestLinesInnerWarrantyInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HardwareLineLink** | Pointer to **string** | Customer line number of the warranty product in this request for linkage, either hardwareLineLink or warrantyLineLink can be used in a line | [optional] 
**WarrantyLineLink** | Pointer to **string** | Customer line number of the hardware product in this request for linkage, either hardwareLineLink or warrantyLineLink can be used in a line. | [optional] 
**DirectLineLink** | Pointer to **string** | Unique value to link hardware and warranty lines. Should be used only when products are purchased from both Ingram and/or vendor but the warranty is purchased through Ingram for them. | [optional] 
**SerialInfo** | Pointer to [**[]OrderCreateV7RequestLinesInnerWarrantyInfoInnerSerialInfoInner**](OrderCreateV7RequestLinesInnerWarrantyInfoInnerSerialInfoInner.md) |  | [optional] 

## Methods

### NewOrderCreateV7RequestLinesInnerWarrantyInfoInner

`func NewOrderCreateV7RequestLinesInnerWarrantyInfoInner() *OrderCreateV7RequestLinesInnerWarrantyInfoInner`

NewOrderCreateV7RequestLinesInnerWarrantyInfoInner instantiates a new OrderCreateV7RequestLinesInnerWarrantyInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateV7RequestLinesInnerWarrantyInfoInnerWithDefaults

`func NewOrderCreateV7RequestLinesInnerWarrantyInfoInnerWithDefaults() *OrderCreateV7RequestLinesInnerWarrantyInfoInner`

NewOrderCreateV7RequestLinesInnerWarrantyInfoInnerWithDefaults instantiates a new OrderCreateV7RequestLinesInnerWarrantyInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHardwareLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetHardwareLineLink() string`

GetHardwareLineLink returns the HardwareLineLink field if non-nil, zero value otherwise.

### GetHardwareLineLinkOk

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetHardwareLineLinkOk() (*string, bool)`

GetHardwareLineLinkOk returns a tuple with the HardwareLineLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) SetHardwareLineLink(v string)`

SetHardwareLineLink sets HardwareLineLink field to given value.

### HasHardwareLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) HasHardwareLineLink() bool`

HasHardwareLineLink returns a boolean if a field has been set.

### GetWarrantyLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetWarrantyLineLink() string`

GetWarrantyLineLink returns the WarrantyLineLink field if non-nil, zero value otherwise.

### GetWarrantyLineLinkOk

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetWarrantyLineLinkOk() (*string, bool)`

GetWarrantyLineLinkOk returns a tuple with the WarrantyLineLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) SetWarrantyLineLink(v string)`

SetWarrantyLineLink sets WarrantyLineLink field to given value.

### HasWarrantyLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) HasWarrantyLineLink() bool`

HasWarrantyLineLink returns a boolean if a field has been set.

### GetDirectLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetDirectLineLink() string`

GetDirectLineLink returns the DirectLineLink field if non-nil, zero value otherwise.

### GetDirectLineLinkOk

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetDirectLineLinkOk() (*string, bool)`

GetDirectLineLinkOk returns a tuple with the DirectLineLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) SetDirectLineLink(v string)`

SetDirectLineLink sets DirectLineLink field to given value.

### HasDirectLineLink

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) HasDirectLineLink() bool`

HasDirectLineLink returns a boolean if a field has been set.

### GetSerialInfo

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetSerialInfo() []OrderCreateV7RequestLinesInnerWarrantyInfoInnerSerialInfoInner`

GetSerialInfo returns the SerialInfo field if non-nil, zero value otherwise.

### GetSerialInfoOk

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) GetSerialInfoOk() (*[]OrderCreateV7RequestLinesInnerWarrantyInfoInnerSerialInfoInner, bool)`

GetSerialInfoOk returns a tuple with the SerialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialInfo

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) SetSerialInfo(v []OrderCreateV7RequestLinesInnerWarrantyInfoInnerSerialInfoInner)`

SetSerialInfo sets SerialInfo field to given value.

### HasSerialInfo

`func (o *OrderCreateV7RequestLinesInnerWarrantyInfoInner) HasSerialInfo() bool`

HasSerialInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


