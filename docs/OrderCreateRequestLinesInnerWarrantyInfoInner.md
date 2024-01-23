# OrderCreateRequestLinesInnerWarrantyInfoInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectLineLink** | Pointer to **string** | Unique value to link hardware and warranty lines. Should be used only when products are purchased from both Ingram and/or vendor but the warranty is purchased through Ingram for them. | [optional] 
**WarrantyLineLink** | Pointer to **string** | Customer line number of the hardware product in this request for linkage, either hardwareLineLink or warrantyLineLink can be used in a line. | [optional] 
**HardwareLineLink** | Pointer to **string** | Customer line number of the warranty product in this request for linkage, either hardwareLineLink or warrantyLineLink can be used in a line  | [optional] 
**SerialInfo** | Pointer to [**[]OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner**](OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner.md) | Serial information of the hardware to be associated with the warranty, applicable on post sale orders. | [optional] 

## Methods

### NewOrderCreateRequestLinesInnerWarrantyInfoInner

`func NewOrderCreateRequestLinesInnerWarrantyInfoInner() *OrderCreateRequestLinesInnerWarrantyInfoInner`

NewOrderCreateRequestLinesInnerWarrantyInfoInner instantiates a new OrderCreateRequestLinesInnerWarrantyInfoInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateRequestLinesInnerWarrantyInfoInnerWithDefaults

`func NewOrderCreateRequestLinesInnerWarrantyInfoInnerWithDefaults() *OrderCreateRequestLinesInnerWarrantyInfoInner`

NewOrderCreateRequestLinesInnerWarrantyInfoInnerWithDefaults instantiates a new OrderCreateRequestLinesInnerWarrantyInfoInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetDirectLineLink() string`

GetDirectLineLink returns the DirectLineLink field if non-nil, zero value otherwise.

### GetDirectLineLinkOk

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetDirectLineLinkOk() (*string, bool)`

GetDirectLineLinkOk returns a tuple with the DirectLineLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetDirectLineLink(v string)`

SetDirectLineLink sets DirectLineLink field to given value.

### HasDirectLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasDirectLineLink() bool`

HasDirectLineLink returns a boolean if a field has been set.

### GetWarrantyLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetWarrantyLineLink() string`

GetWarrantyLineLink returns the WarrantyLineLink field if non-nil, zero value otherwise.

### GetWarrantyLineLinkOk

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetWarrantyLineLinkOk() (*string, bool)`

GetWarrantyLineLinkOk returns a tuple with the WarrantyLineLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetWarrantyLineLink(v string)`

SetWarrantyLineLink sets WarrantyLineLink field to given value.

### HasWarrantyLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasWarrantyLineLink() bool`

HasWarrantyLineLink returns a boolean if a field has been set.

### GetHardwareLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetHardwareLineLink() string`

GetHardwareLineLink returns the HardwareLineLink field if non-nil, zero value otherwise.

### GetHardwareLineLinkOk

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetHardwareLineLinkOk() (*string, bool)`

GetHardwareLineLinkOk returns a tuple with the HardwareLineLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetHardwareLineLink(v string)`

SetHardwareLineLink sets HardwareLineLink field to given value.

### HasHardwareLineLink

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasHardwareLineLink() bool`

HasHardwareLineLink returns a boolean if a field has been set.

### GetSerialInfo

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetSerialInfo() []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner`

GetSerialInfo returns the SerialInfo field if non-nil, zero value otherwise.

### GetSerialInfoOk

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetSerialInfoOk() (*[]OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner, bool)`

GetSerialInfoOk returns a tuple with the SerialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialInfo

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetSerialInfo(v []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner)`

SetSerialInfo sets SerialInfo field to given value.

### HasSerialInfo

`func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasSerialInfo() bool`

HasSerialInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


