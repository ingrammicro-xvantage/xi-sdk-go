# OrderStatusAsyncNotificationRequestResourceInnerLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineNumber** | Pointer to **string** | The Ingram Micro line number for the product | [optional] 
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**LineStatus** | Pointer to **string** | The status for the line item in the order. One of: Backordered, Open, Shipped | [optional] 
**IngramPartNumber** | Pointer to **string** | The Ingram Micro part number for the line item. | [optional] 
**VendorPartNumber** | Pointer to **string** | The vendor part number for the line item. | [optional] 
**RequestedQuantity** | Pointer to **string** | The quantity of the line item requested. | [optional] 
**ShippedQuantity** | Pointer to **string** | The quantity of the line item that has been shipped. | [optional] 
**BackorderedQuantity** | Pointer to **string** | The quantity of the line item that is backordered. | [optional] 
**ShipmentDetails** | Pointer to [**[]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner**](OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner.md) |  | [optional] 
**SerialNumberDetails** | Pointer to [**[]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner**](OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner.md) |  | [optional] 

## Methods

### NewOrderStatusAsyncNotificationRequestResourceInnerLinesInner

`func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInner() *OrderStatusAsyncNotificationRequestResourceInnerLinesInner`

NewOrderStatusAsyncNotificationRequestResourceInnerLinesInner instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerWithDefaults

`func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInnerLinesInner`

NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineNumber() string`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineNumberOk() (*string, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetLineNumber(v string)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### GetSubOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetLineStatus

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineStatus() string`

GetLineStatus returns the LineStatus field if non-nil, zero value otherwise.

### GetLineStatusOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetLineStatusOk() (*string, bool)`

GetLineStatusOk returns a tuple with the LineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineStatus

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetLineStatus(v string)`

SetLineStatus sets LineStatus field to given value.

### HasLineStatus

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasLineStatus() bool`

HasLineStatus returns a boolean if a field has been set.

### GetIngramPartNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetIngramPartNumber() string`

GetIngramPartNumber returns the IngramPartNumber field if non-nil, zero value otherwise.

### GetIngramPartNumberOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetIngramPartNumberOk() (*string, bool)`

GetIngramPartNumberOk returns a tuple with the IngramPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngramPartNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetIngramPartNumber(v string)`

SetIngramPartNumber sets IngramPartNumber field to given value.

### HasIngramPartNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasIngramPartNumber() bool`

HasIngramPartNumber returns a boolean if a field has been set.

### GetVendorPartNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetVendorPartNumber() string`

GetVendorPartNumber returns the VendorPartNumber field if non-nil, zero value otherwise.

### GetVendorPartNumberOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetVendorPartNumberOk() (*string, bool)`

GetVendorPartNumberOk returns a tuple with the VendorPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorPartNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetVendorPartNumber(v string)`

SetVendorPartNumber sets VendorPartNumber field to given value.

### HasVendorPartNumber

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasVendorPartNumber() bool`

HasVendorPartNumber returns a boolean if a field has been set.

### GetRequestedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetRequestedQuantity() string`

GetRequestedQuantity returns the RequestedQuantity field if non-nil, zero value otherwise.

### GetRequestedQuantityOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetRequestedQuantityOk() (*string, bool)`

GetRequestedQuantityOk returns a tuple with the RequestedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetRequestedQuantity(v string)`

SetRequestedQuantity sets RequestedQuantity field to given value.

### HasRequestedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasRequestedQuantity() bool`

HasRequestedQuantity returns a boolean if a field has been set.

### GetShippedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShippedQuantity() string`

GetShippedQuantity returns the ShippedQuantity field if non-nil, zero value otherwise.

### GetShippedQuantityOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShippedQuantityOk() (*string, bool)`

GetShippedQuantityOk returns a tuple with the ShippedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetShippedQuantity(v string)`

SetShippedQuantity sets ShippedQuantity field to given value.

### HasShippedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasShippedQuantity() bool`

HasShippedQuantity returns a boolean if a field has been set.

### GetBackorderedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetBackorderedQuantity() string`

GetBackorderedQuantity returns the BackorderedQuantity field if non-nil, zero value otherwise.

### GetBackorderedQuantityOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetBackorderedQuantityOk() (*string, bool)`

GetBackorderedQuantityOk returns a tuple with the BackorderedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetBackorderedQuantity(v string)`

SetBackorderedQuantity sets BackorderedQuantity field to given value.

### HasBackorderedQuantity

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasBackorderedQuantity() bool`

HasBackorderedQuantity returns a boolean if a field has been set.

### GetShipmentDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShipmentDetails() []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner`

GetShipmentDetails returns the ShipmentDetails field if non-nil, zero value otherwise.

### GetShipmentDetailsOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetShipmentDetailsOk() (*[]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner, bool)`

GetShipmentDetailsOk returns a tuple with the ShipmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetShipmentDetails(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner)`

SetShipmentDetails sets ShipmentDetails field to given value.

### HasShipmentDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasShipmentDetails() bool`

HasShipmentDetails returns a boolean if a field has been set.

### GetSerialNumberDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSerialNumberDetails() []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner`

GetSerialNumberDetails returns the SerialNumberDetails field if non-nil, zero value otherwise.

### GetSerialNumberDetailsOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) GetSerialNumberDetailsOk() (*[]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner, bool)`

GetSerialNumberDetailsOk returns a tuple with the SerialNumberDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) SetSerialNumberDetails(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner)`

SetSerialNumberDetails sets SerialNumberDetails field to given value.

### HasSerialNumberDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInner) HasSerialNumberDetails() bool`

HasSerialNumberDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


