# OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierCode** | Pointer to **string** | The code for the shipping carrier for the line item. | [optional] 
**CarrierName** | Pointer to **string** | The name of the shipping carrier for the line item. | [optional] 
**ShipFromWarehouseId** | Pointer to **string** | The ID of the warehouse the line item will ship from. | [optional] 
**ShipFromLocation** | Pointer to **string** | Location from which order is shipped. | [optional] 
**FreightAccountNumber** | Pointer to **string** | The reseller&#39;s shipping account number with carrier. Used to bill the shipping carrier directly from the reseller&#39;s account with the carrier. | [optional] 
**SignatureRequired** | Pointer to **string** | Specifies whether a signature is required for delivery. Default is False. | [optional] 
**ShippingInstructions** | Pointer to **string** | The shipping instructions for the order. | [optional] 

## Methods

### NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner

`func NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner() *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner`

NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner instantiates a new OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInnerWithDefaults

`func NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInnerWithDefaults() *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner`

NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierCode

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierCode() string`

GetCarrierCode returns the CarrierCode field if non-nil, zero value otherwise.

### GetCarrierCodeOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierCodeOk() (*string, bool)`

GetCarrierCodeOk returns a tuple with the CarrierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierCode

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetCarrierCode(v string)`

SetCarrierCode sets CarrierCode field to given value.

### HasCarrierCode

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasCarrierCode() bool`

HasCarrierCode returns a boolean if a field has been set.

### GetCarrierName

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetShipFromWarehouseId

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string`

GetShipFromWarehouseId returns the ShipFromWarehouseId field if non-nil, zero value otherwise.

### GetShipFromWarehouseIdOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool)`

GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromWarehouseId

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string)`

SetShipFromWarehouseId sets ShipFromWarehouseId field to given value.

### HasShipFromWarehouseId

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool`

HasShipFromWarehouseId returns a boolean if a field has been set.

### GetShipFromLocation

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromLocation() string`

GetShipFromLocation returns the ShipFromLocation field if non-nil, zero value otherwise.

### GetShipFromLocationOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromLocationOk() (*string, bool)`

GetShipFromLocationOk returns a tuple with the ShipFromLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipFromLocation

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetShipFromLocation(v string)`

SetShipFromLocation sets ShipFromLocation field to given value.

### HasShipFromLocation

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasShipFromLocation() bool`

HasShipFromLocation returns a boolean if a field has been set.

### GetFreightAccountNumber

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetFreightAccountNumber() string`

GetFreightAccountNumber returns the FreightAccountNumber field if non-nil, zero value otherwise.

### GetFreightAccountNumberOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetFreightAccountNumberOk() (*string, bool)`

GetFreightAccountNumberOk returns a tuple with the FreightAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightAccountNumber

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetFreightAccountNumber(v string)`

SetFreightAccountNumber sets FreightAccountNumber field to given value.

### HasFreightAccountNumber

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasFreightAccountNumber() bool`

HasFreightAccountNumber returns a boolean if a field has been set.

### GetSignatureRequired

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetSignatureRequired() string`

GetSignatureRequired returns the SignatureRequired field if non-nil, zero value otherwise.

### GetSignatureRequiredOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetSignatureRequiredOk() (*string, bool)`

GetSignatureRequiredOk returns a tuple with the SignatureRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureRequired

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetSignatureRequired(v string)`

SetSignatureRequired sets SignatureRequired field to given value.

### HasSignatureRequired

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasSignatureRequired() bool`

HasSignatureRequired returns a boolean if a field has been set.

### GetShippingInstructions

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShippingInstructions() string`

GetShippingInstructions returns the ShippingInstructions field if non-nil, zero value otherwise.

### GetShippingInstructionsOk

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShippingInstructionsOk() (*string, bool)`

GetShippingInstructionsOk returns a tuple with the ShippingInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInstructions

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetShippingInstructions(v string)`

SetShippingInstructions sets ShippingInstructions field to given value.

### HasShippingInstructions

`func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasShippingInstructions() bool`

HasShippingInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


