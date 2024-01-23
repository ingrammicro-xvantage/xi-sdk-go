# OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierCode** | Pointer to **string** | The carrier code for the shipment containing the line item. | [optional] 
**CarrierName** | Pointer to **string** | The name of the carrier of the shipment containing the line item. | [optional] 
**Quantity** | Pointer to **int32** | The quantity shipped of the line item. | [optional] 
**ShippedDate** | Pointer to **string** | The actual date when line item shipped. | [optional] 
**EstimatedDeliveryDate** | Pointer to **string** | The date the line item is expected to be delivered. | [optional] 
**DeliveredDate** | Pointer to **string** | The actual date of delivery of the line item. | [optional] 
**CarrierPickupDate** | Pointer to **string** | The actual date when carrier picked up line item. | [optional] 
**TrackingDetails** | Pointer to [**[]OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner**](OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner.md) | The tracking details for the shipment containing the line item. | [optional] 

## Methods

### NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner

`func NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner() *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner`

NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerWithDefaults

`func NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerWithDefaults() *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner`

NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerWithDefaults instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierCode

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetShippedDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetShippedDate() string`

GetShippedDate returns the ShippedDate field if non-nil, zero value otherwise.

### GetShippedDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetShippedDateOk() (*string, bool)`

GetShippedDateOk returns a tuple with the ShippedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetShippedDate(v string)`

SetShippedDate sets ShippedDate field to given value.

### HasShippedDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasShippedDate() bool`

HasShippedDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetEstimatedDeliveryDate() string`

GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetEstimatedDeliveryDateOk() (*string, bool)`

GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetEstimatedDeliveryDate(v string)`

SetEstimatedDeliveryDate sets EstimatedDeliveryDate field to given value.

### HasEstimatedDeliveryDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasEstimatedDeliveryDate() bool`

HasEstimatedDeliveryDate returns a boolean if a field has been set.

### GetDeliveredDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetDeliveredDate() string`

GetDeliveredDate returns the DeliveredDate field if non-nil, zero value otherwise.

### GetDeliveredDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetDeliveredDateOk() (*string, bool)`

GetDeliveredDateOk returns a tuple with the DeliveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetDeliveredDate(v string)`

SetDeliveredDate sets DeliveredDate field to given value.

### HasDeliveredDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasDeliveredDate() bool`

HasDeliveredDate returns a boolean if a field has been set.

### GetCarrierPickupDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierPickupDate() string`

GetCarrierPickupDate returns the CarrierPickupDate field if non-nil, zero value otherwise.

### GetCarrierPickupDateOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierPickupDateOk() (*string, bool)`

GetCarrierPickupDateOk returns a tuple with the CarrierPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierPickupDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetCarrierPickupDate(v string)`

SetCarrierPickupDate sets CarrierPickupDate field to given value.

### HasCarrierPickupDate

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasCarrierPickupDate() bool`

HasCarrierPickupDate returns a boolean if a field has been set.

### GetTrackingDetails

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetTrackingDetails() []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner`

GetTrackingDetails returns the TrackingDetails field if non-nil, zero value otherwise.

### GetTrackingDetailsOk

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetTrackingDetailsOk() (*[]OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner, bool)`

GetTrackingDetailsOk returns a tuple with the TrackingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingDetails

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetTrackingDetails(v []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner)`

SetTrackingDetails sets TrackingDetails field to given value.

### HasTrackingDetails

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasTrackingDetails() bool`

HasTrackingDetails returns a boolean if a field has been set.

### SetTrackingDetailsNil

`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetTrackingDetailsNil(b bool)`

 SetTrackingDetailsNil sets the value for TrackingDetails to be an explicit nil

### UnsetTrackingDetails
`func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) UnsetTrackingDetails()`

UnsetTrackingDetails ensures that no value is present for TrackingDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


