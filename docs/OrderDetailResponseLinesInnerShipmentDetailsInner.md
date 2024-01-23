# OrderDetailResponseLinesInnerShipmentDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **int32** | The quantity shipped of the line item. | [optional] 
**EstimatedShipDate** | Pointer to **string** | The estimated ship date for the line item. | [optional] 
**ShippedDate** | Pointer to **string** | The date the line item was shipped. | [optional] 
**EstimatedDeliveryDate** | Pointer to **string** | The date the line item is expected to be delivered. | [optional] 
**DeliveredDate** | Pointer to **string** | The actual date of delivery of the line item. | [optional] 
**ShipFromWarehouseId** | Pointer to **string** | The ID of the warehouse the product will ship from. | [optional] 
**ShipFromLocation** | Pointer to **string** | The city and state the line item ships from. | [optional] 
**InvoiceNumber** | Pointer to **string** | The Ingram Micro invoice number for the line item. | [optional] 
**InvoiceDate** | Pointer to **string** | The date the IngramMicro invoice was created for the line item. | [optional] 
**CarrierDetails** | Pointer to [**OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails**](OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails.md) |  | [optional] 

## Methods

### NewOrderDetailResponseLinesInnerShipmentDetailsInner

`func NewOrderDetailResponseLinesInnerShipmentDetailsInner() *OrderDetailResponseLinesInnerShipmentDetailsInner`

NewOrderDetailResponseLinesInnerShipmentDetailsInner instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseLinesInnerShipmentDetailsInnerWithDefaults

`func NewOrderDetailResponseLinesInnerShipmentDetailsInnerWithDefaults() *OrderDetailResponseLinesInnerShipmentDetailsInner`

NewOrderDetailResponseLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetEstimatedShipDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedShipDate() string`

GetEstimatedShipDate returns the EstimatedShipDate field if non-nil, zero value otherwise.

### GetEstimatedShipDateOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedShipDateOk() (*string, bool)`

GetEstimatedShipDateOk returns a tuple with the EstimatedShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedShipDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetEstimatedShipDate(v string)`

SetEstimatedShipDate sets EstimatedShipDate field to given value.

### HasEstimatedShipDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasEstimatedShipDate() bool`

HasEstimatedShipDate returns a boolean if a field has been set.

### GetShippedDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShippedDate() string`

GetShippedDate returns the ShippedDate field if non-nil, zero value otherwise.

### GetShippedDateOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShippedDateOk() (*string, bool)`

GetShippedDateOk returns a tuple with the ShippedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetShippedDate(v string)`

SetShippedDate sets ShippedDate field to given value.

### HasShippedDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasShippedDate() bool`

HasShippedDate returns a boolean if a field has been set.

### GetEstimatedDeliveryDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDate() string`

GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field if non-nil, zero value otherwise.

### GetEstimatedDeliveryDateOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDateOk() (*string, bool)`

GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDeliveryDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetEstimatedDeliveryDate(v string)`

SetEstimatedDeliveryDate sets EstimatedDeliveryDate field to given value.

### HasEstimatedDeliveryDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasEstimatedDeliveryDate() bool`

HasEstimatedDeliveryDate returns a boolean if a field has been set.

### GetDeliveredDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetDeliveredDate() string`

GetDeliveredDate returns the DeliveredDate field if non-nil, zero value otherwise.

### GetDeliveredDateOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetDeliveredDateOk() (*string, bool)`

GetDeliveredDateOk returns a tuple with the DeliveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetDeliveredDate(v string)`

SetDeliveredDate sets DeliveredDate field to given value.

### HasDeliveredDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasDeliveredDate() bool`

HasDeliveredDate returns a boolean if a field has been set.

### GetShipFromWarehouseId

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string`

GetShipFromWarehouseId returns the ShipFromWarehouseId field if non-nil, zero value otherwise.

### GetShipFromWarehouseIdOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool)`

GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromWarehouseId

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string)`

SetShipFromWarehouseId sets ShipFromWarehouseId field to given value.

### HasShipFromWarehouseId

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool`

HasShipFromWarehouseId returns a boolean if a field has been set.

### GetShipFromLocation

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromLocation() string`

GetShipFromLocation returns the ShipFromLocation field if non-nil, zero value otherwise.

### GetShipFromLocationOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromLocationOk() (*string, bool)`

GetShipFromLocationOk returns a tuple with the ShipFromLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromLocation

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetShipFromLocation(v string)`

SetShipFromLocation sets ShipFromLocation field to given value.

### HasShipFromLocation

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasShipFromLocation() bool`

HasShipFromLocation returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetCarrierDetails

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetCarrierDetails() OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails`

GetCarrierDetails returns the CarrierDetails field if non-nil, zero value otherwise.

### GetCarrierDetailsOk

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetCarrierDetailsOk() (*OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails, bool)`

GetCarrierDetailsOk returns a tuple with the CarrierDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierDetails

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetCarrierDetails(v OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails)`

SetCarrierDetails sets CarrierDetails field to given value.

### HasCarrierDetails

`func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasCarrierDetails() bool`

HasCarrierDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


