# OrderModifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **string** | Shipment-level notes. | [optional] 
**ShipToInfo** | Pointer to [**OrderModifyRequestShipToInfo**](OrderModifyRequestShipToInfo.md) |  | [optional] 
**Lines** | Pointer to [**[]OrderModifyRequestLinesInner**](OrderModifyRequestLinesInner.md) | The order line items. | [optional] 
**AdditionalAttributes** | Pointer to [**[]OrderModifyRequestAdditionalAttributesInner**](OrderModifyRequestAdditionalAttributesInner.md) | Header-level additional attributes. | [optional] 

## Methods

### NewOrderModifyRequest

`func NewOrderModifyRequest() *OrderModifyRequest`

NewOrderModifyRequest instantiates a new OrderModifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderModifyRequestWithDefaults

`func NewOrderModifyRequestWithDefaults() *OrderModifyRequest`

NewOrderModifyRequestWithDefaults instantiates a new OrderModifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *OrderModifyRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OrderModifyRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OrderModifyRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OrderModifyRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetShipToInfo

`func (o *OrderModifyRequest) GetShipToInfo() OrderModifyRequestShipToInfo`

GetShipToInfo returns the ShipToInfo field if non-nil, zero value otherwise.

### GetShipToInfoOk

`func (o *OrderModifyRequest) GetShipToInfoOk() (*OrderModifyRequestShipToInfo, bool)`

GetShipToInfoOk returns a tuple with the ShipToInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToInfo

`func (o *OrderModifyRequest) SetShipToInfo(v OrderModifyRequestShipToInfo)`

SetShipToInfo sets ShipToInfo field to given value.

### HasShipToInfo

`func (o *OrderModifyRequest) HasShipToInfo() bool`

HasShipToInfo returns a boolean if a field has been set.

### GetLines

`func (o *OrderModifyRequest) GetLines() []OrderModifyRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *OrderModifyRequest) GetLinesOk() (*[]OrderModifyRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *OrderModifyRequest) SetLines(v []OrderModifyRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *OrderModifyRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetAdditionalAttributes

`func (o *OrderModifyRequest) GetAdditionalAttributes() []OrderModifyRequestAdditionalAttributesInner`

GetAdditionalAttributes returns the AdditionalAttributes field if non-nil, zero value otherwise.

### GetAdditionalAttributesOk

`func (o *OrderModifyRequest) GetAdditionalAttributesOk() (*[]OrderModifyRequestAdditionalAttributesInner, bool)`

GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAttributes

`func (o *OrderModifyRequest) SetAdditionalAttributes(v []OrderModifyRequestAdditionalAttributesInner)`

SetAdditionalAttributes sets AdditionalAttributes field to given value.

### HasAdditionalAttributes

`func (o *OrderModifyRequest) HasAdditionalAttributes() bool`

HasAdditionalAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


