# OrderCreateResponseOrdersInnerMiscellaneousChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**ChargeLineReference** | Pointer to **string** | Impulse line number for the miscellaneous charge. | [optional] 
**ChargeDescription** | Pointer to **string** | Description of the miscellaneous charges for the order. | [optional] 
**ChargeAmount** | Pointer to **float32** | The total amount of miscellaneous charges for the order. | [optional] 

## Methods

### NewOrderCreateResponseOrdersInnerMiscellaneousChargesInner

`func NewOrderCreateResponseOrdersInnerMiscellaneousChargesInner() *OrderCreateResponseOrdersInnerMiscellaneousChargesInner`

NewOrderCreateResponseOrdersInnerMiscellaneousChargesInner instantiates a new OrderCreateResponseOrdersInnerMiscellaneousChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateResponseOrdersInnerMiscellaneousChargesInnerWithDefaults

`func NewOrderCreateResponseOrdersInnerMiscellaneousChargesInnerWithDefaults() *OrderCreateResponseOrdersInnerMiscellaneousChargesInner`

NewOrderCreateResponseOrdersInnerMiscellaneousChargesInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerMiscellaneousChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetChargeLineReference

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeLineReference() string`

GetChargeLineReference returns the ChargeLineReference field if non-nil, zero value otherwise.

### GetChargeLineReferenceOk

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeLineReferenceOk() (*string, bool)`

GetChargeLineReferenceOk returns a tuple with the ChargeLineReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeLineReference

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetChargeLineReference(v string)`

SetChargeLineReference sets ChargeLineReference field to given value.

### HasChargeLineReference

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasChargeLineReference() bool`

HasChargeLineReference returns a boolean if a field has been set.

### GetChargeDescription

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeDescription() string`

GetChargeDescription returns the ChargeDescription field if non-nil, zero value otherwise.

### GetChargeDescriptionOk

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeDescriptionOk() (*string, bool)`

GetChargeDescriptionOk returns a tuple with the ChargeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDescription

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetChargeDescription(v string)`

SetChargeDescription sets ChargeDescription field to given value.

### HasChargeDescription

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasChargeDescription() bool`

HasChargeDescription returns a boolean if a field has been set.

### GetChargeAmount

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeAmount() float32`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) GetChargeAmountOk() (*float32, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) SetChargeAmount(v float32)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *OrderCreateResponseOrdersInnerMiscellaneousChargesInner) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


