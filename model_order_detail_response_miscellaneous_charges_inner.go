/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
)

// checks if the OrderDetailResponseMiscellaneousChargesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseMiscellaneousChargesInner{}

// OrderDetailResponseMiscellaneousChargesInner struct for OrderDetailResponseMiscellaneousChargesInner
type OrderDetailResponseMiscellaneousChargesInner struct {
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// Impulse line number for the miscellaneous charge.
	ChargeLineReference *string `json:"chargeLineReference,omitempty"`
	// Description of the miscellaneous charges.
	ChargeDescription *string `json:"chargeDescription,omitempty"`
	// The amount of miscellaneous charges.
	ChargeAmount *float64 `json:"chargeAmount,omitempty"`
}

// NewOrderDetailResponseMiscellaneousChargesInner instantiates a new OrderDetailResponseMiscellaneousChargesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseMiscellaneousChargesInner() *OrderDetailResponseMiscellaneousChargesInner {
	this := OrderDetailResponseMiscellaneousChargesInner{}
	return &this
}

// NewOrderDetailResponseMiscellaneousChargesInnerWithDefaults instantiates a new OrderDetailResponseMiscellaneousChargesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseMiscellaneousChargesInnerWithDefaults() *OrderDetailResponseMiscellaneousChargesInner {
	this := OrderDetailResponseMiscellaneousChargesInner{}
	return &this
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderDetailResponseMiscellaneousChargesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetChargeLineReference returns the ChargeLineReference field value if set, zero value otherwise.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeLineReference() string {
	if o == nil || IsNil(o.ChargeLineReference) {
		var ret string
		return ret
	}
	return *o.ChargeLineReference
}

// GetChargeLineReferenceOk returns a tuple with the ChargeLineReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeLineReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeLineReference) {
		return nil, false
	}
	return o.ChargeLineReference, true
}

// HasChargeLineReference returns a boolean if a field has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) HasChargeLineReference() bool {
	if o != nil && !IsNil(o.ChargeLineReference) {
		return true
	}

	return false
}

// SetChargeLineReference gets a reference to the given string and assigns it to the ChargeLineReference field.
func (o *OrderDetailResponseMiscellaneousChargesInner) SetChargeLineReference(v string) {
	o.ChargeLineReference = &v
}

// GetChargeDescription returns the ChargeDescription field value if set, zero value otherwise.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeDescription() string {
	if o == nil || IsNil(o.ChargeDescription) {
		var ret string
		return ret
	}
	return *o.ChargeDescription
}

// GetChargeDescriptionOk returns a tuple with the ChargeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeDescription) {
		return nil, false
	}
	return o.ChargeDescription, true
}

// HasChargeDescription returns a boolean if a field has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) HasChargeDescription() bool {
	if o != nil && !IsNil(o.ChargeDescription) {
		return true
	}

	return false
}

// SetChargeDescription gets a reference to the given string and assigns it to the ChargeDescription field.
func (o *OrderDetailResponseMiscellaneousChargesInner) SetChargeDescription(v string) {
	o.ChargeDescription = &v
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeAmount() float64 {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret float64
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) GetChargeAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *OrderDetailResponseMiscellaneousChargesInner) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given float64 and assigns it to the ChargeAmount field.
func (o *OrderDetailResponseMiscellaneousChargesInner) SetChargeAmount(v float64) {
	o.ChargeAmount = &v
}

func (o OrderDetailResponseMiscellaneousChargesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseMiscellaneousChargesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubOrderNumber) {
		toSerialize["subOrderNumber"] = o.SubOrderNumber
	}
	if !IsNil(o.ChargeLineReference) {
		toSerialize["chargeLineReference"] = o.ChargeLineReference
	}
	if !IsNil(o.ChargeDescription) {
		toSerialize["chargeDescription"] = o.ChargeDescription
	}
	if !IsNil(o.ChargeAmount) {
		toSerialize["chargeAmount"] = o.ChargeAmount
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseMiscellaneousChargesInner struct {
	value *OrderDetailResponseMiscellaneousChargesInner
	isSet bool
}

func (v NullableOrderDetailResponseMiscellaneousChargesInner) Get() *OrderDetailResponseMiscellaneousChargesInner {
	return v.value
}

func (v *NullableOrderDetailResponseMiscellaneousChargesInner) Set(val *OrderDetailResponseMiscellaneousChargesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseMiscellaneousChargesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseMiscellaneousChargesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseMiscellaneousChargesInner(val *OrderDetailResponseMiscellaneousChargesInner) *NullableOrderDetailResponseMiscellaneousChargesInner {
	return &NullableOrderDetailResponseMiscellaneousChargesInner{value: val, isSet: true}
}

func (v NullableOrderDetailResponseMiscellaneousChargesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseMiscellaneousChargesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

