# OrderDetailResponseMiscellaneousChargesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubOrderNumber** | Pointer to **string** | The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number. | [optional] 
**ChargeLineReference** | Pointer to **string** | Impulse line number for the miscellaneous charge. | [optional] 
**ChargeDescription** | Pointer to **string** | Description of the miscellaneous charges. | [optional] 
**ChargeAmount** | Pointer to **float64** | The amount of miscellaneous charges. | [optional] 

## Methods

### NewOrderDetailResponseMiscellaneousChargesInner

`func NewOrderDetailResponseMiscellaneousChargesInner() *OrderDetailResponseMiscellaneousChargesInner`

NewOrderDetailResponseMiscellaneousChargesInner instantiates a new OrderDetailResponseMiscellaneousChargesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDetailResponseMiscellaneousChargesInnerWithDefaults

`func NewOrderDetailResponseMiscellaneousChargesInnerWithDefaults() *OrderDetailResponseMiscellaneousChargesInner`

NewOrderDetailResponseMiscellaneousChargesInnerWithDefaults instantiates a new OrderDetailResponseMiscellaneousChargesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubOrderNumber

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetSubOrderNumber() string`

GetSubOrderNumber returns the SubOrderNumber field if non-nil, zero value otherwise.

### GetSubOrderNumberOk

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetSubOrderNumberOk() (*string, bool)`

GetSubOrderNumberOk returns a tuple with the SubOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrderNumber

`func (o *OrderDetailResponseMiscellaneousChargesInner) SetSubOrderNumber(v string)`

SetSubOrderNumber sets SubOrderNumber field to given value.

### HasSubOrderNumber

`func (o *OrderDetailResponseMiscellaneousChargesInner) HasSubOrderNumber() bool`

HasSubOrderNumber returns a boolean if a field has been set.

### GetChargeLineReference

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeLineReference() string`

GetChargeLineReference returns the ChargeLineReference field if non-nil, zero value otherwise.

### GetChargeLineReferenceOk

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeLineReferenceOk() (*string, bool)`

GetChargeLineReferenceOk returns a tuple with the ChargeLineReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeLineReference

`func (o *OrderDetailResponseMiscellaneousChargesInner) SetChargeLineReference(v string)`

SetChargeLineReference sets ChargeLineReference field to given value.

### HasChargeLineReference

`func (o *OrderDetailResponseMiscellaneousChargesInner) HasChargeLineReference() bool`

HasChargeLineReference returns a boolean if a field has been set.

### GetChargeDescription

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeDescription() string`

GetChargeDescription returns the ChargeDescription field if non-nil, zero value otherwise.

### GetChargeDescriptionOk

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeDescriptionOk() (*string, bool)`

GetChargeDescriptionOk returns a tuple with the ChargeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeDescription

`func (o *OrderDetailResponseMiscellaneousChargesInner) SetChargeDescription(v string)`

SetChargeDescription sets ChargeDescription field to given value.

### HasChargeDescription

`func (o *OrderDetailResponseMiscellaneousChargesInner) HasChargeDescription() bool`

HasChargeDescription returns a boolean if a field has been set.

### GetChargeAmount

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeAmount() float64`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeAmountOk() (*float64, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *OrderDetailResponseMiscellaneousChargesInner) SetChargeAmount(v float64)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *OrderDetailResponseMiscellaneousChargesInner) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


