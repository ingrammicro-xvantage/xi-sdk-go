# OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingNumber** | Pointer to **string** | The tracking number for the shipment containing the line item. | [optional] 
**TrackingUrl** | Pointer to **string** | The tracking URL for the shipment containing the line item. | [optional] 
**PackageWeight** | Pointer to **string** | The weight of the package for the line item. | [optional] 
**CartonNumber** | Pointer to **string** | The shipment carton number that contains the line item. | [optional] 
**QuantityInBox** | Pointer to **string** | The quantity of line items in the box. | [optional] 
**SerialNumbers** | Pointer to [**[]OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerSerialNumbersInner**](OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerSerialNumbersInner.md) |  | [optional] 

## Methods

### NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner

`func NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner() *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner`

NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerWithDefaults

`func NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerWithDefaults() *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner`

NewOrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerWithDefaults instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### GetPackageWeight

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetPackageWeight() string`

GetPackageWeight returns the PackageWeight field if non-nil, zero value otherwise.

### GetPackageWeightOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetPackageWeightOk() (*string, bool)`

GetPackageWeightOk returns a tuple with the PackageWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageWeight

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) SetPackageWeight(v string)`

SetPackageWeight sets PackageWeight field to given value.

### HasPackageWeight

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) HasPackageWeight() bool`

HasPackageWeight returns a boolean if a field has been set.

### GetCartonNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetCartonNumber() string`

GetCartonNumber returns the CartonNumber field if non-nil, zero value otherwise.

### GetCartonNumberOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetCartonNumberOk() (*string, bool)`

GetCartonNumberOk returns a tuple with the CartonNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartonNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) SetCartonNumber(v string)`

SetCartonNumber sets CartonNumber field to given value.

### HasCartonNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) HasCartonNumber() bool`

HasCartonNumber returns a boolean if a field has been set.

### GetQuantityInBox

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetQuantityInBox() string`

GetQuantityInBox returns the QuantityInBox field if non-nil, zero value otherwise.

### GetQuantityInBoxOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetQuantityInBoxOk() (*string, bool)`

GetQuantityInBoxOk returns a tuple with the QuantityInBox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityInBox

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) SetQuantityInBox(v string)`

SetQuantityInBox sets QuantityInBox field to given value.

### HasQuantityInBox

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) HasQuantityInBox() bool`

HasQuantityInBox returns a boolean if a field has been set.

### GetSerialNumbers

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetSerialNumbers() []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerSerialNumbersInner`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) GetSerialNumbersOk() (*[]OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerSerialNumbersInner, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) SetSerialNumbers(v []OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInnerSerialNumbersInner)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetailsTrackingDetailsInner) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


