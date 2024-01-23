# OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShipmentDate** | Pointer to **string** | The date the line item was shipped. | [optional] 
**ShipFromWarehouseId** | Pointer to **string** | The ID of the warehouse the product will ship from. | [optional] 
**WarehouseName** | Pointer to **string** | \&quot;\&quot; | [optional] 
**CarrierCode** | Pointer to **string** | The carrier code for the shipment containing the  line item. | [optional] 
**CarrierName** | Pointer to **string** | The name of the carrier of the shipment containing   the line item. | [optional] 
**PackageDetails** | Pointer to [**[]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner**](OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner.md) |  | [optional] 

## Methods

### NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner

`func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner`

NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerWithDefaults

`func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner`

NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShipmentDate

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipmentDate() string`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipmentDateOk() (*string, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetShipmentDate(v string)`

SetShipmentDate sets ShipmentDate field to given value.

### HasShipmentDate

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasShipmentDate() bool`

HasShipmentDate returns a boolean if a field has been set.

### GetShipFromWarehouseId

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string`

GetShipFromWarehouseId returns the ShipFromWarehouseId field if non-nil, zero value otherwise.

### GetShipFromWarehouseIdOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool)`

GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromWarehouseId

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string)`

SetShipFromWarehouseId sets ShipFromWarehouseId field to given value.

### HasShipFromWarehouseId

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool`

HasShipFromWarehouseId returns a boolean if a field has been set.

### GetWarehouseName

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetWarehouseName() string`

GetWarehouseName returns the WarehouseName field if non-nil, zero value otherwise.

### GetWarehouseNameOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetWarehouseNameOk() (*string, bool)`

GetWarehouseNameOk returns a tuple with the WarehouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseName

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetWarehouseName(v string)`

SetWarehouseName sets WarehouseName field to given value.

### HasWarehouseName

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasWarehouseName() bool`

HasWarehouseName returns a boolean if a field has been set.

### GetCarrierCode

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetPackageDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetPackageDetails() []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner`

GetPackageDetails returns the PackageDetails field if non-nil, zero value otherwise.

### GetPackageDetailsOk

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) GetPackageDetailsOk() (*[]OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner, bool)`

GetPackageDetailsOk returns a tuple with the PackageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) SetPackageDetails(v []OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner)`

SetPackageDetails sets PackageDetails field to given value.

### HasPackageDetails

`func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInner) HasPackageDetails() bool`

HasPackageDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


