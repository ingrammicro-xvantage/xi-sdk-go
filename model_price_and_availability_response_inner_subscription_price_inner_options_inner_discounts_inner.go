/*
XI Sdk Resellers

For Resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner{}

// PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner struct for PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner
type PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner struct {
	VolumeDiscounts []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner `json:"volumeDiscounts,omitempty"`
	SpecialPricing []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerSpecialPricingInner `json:"specialPricing,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner {
	this := PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerWithDefaults() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner {
	this := PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner{}
	return &this
}

// GetVolumeDiscounts returns the VolumeDiscounts field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) GetVolumeDiscounts() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner {
	if o == nil || IsNil(o.VolumeDiscounts) {
		var ret []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner
		return ret
	}
	return o.VolumeDiscounts
}

// GetVolumeDiscountsOk returns a tuple with the VolumeDiscounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) GetVolumeDiscountsOk() ([]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner, bool) {
	if o == nil || IsNil(o.VolumeDiscounts) {
		return nil, false
	}
	return o.VolumeDiscounts, true
}

// HasVolumeDiscounts returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) HasVolumeDiscounts() bool {
	if o != nil && !IsNil(o.VolumeDiscounts) {
		return true
	}

	return false
}

// SetVolumeDiscounts gets a reference to the given []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner and assigns it to the VolumeDiscounts field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) SetVolumeDiscounts(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) {
	o.VolumeDiscounts = v
}

// GetSpecialPricing returns the SpecialPricing field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) GetSpecialPricing() []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerSpecialPricingInner {
	if o == nil || IsNil(o.SpecialPricing) {
		var ret []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerSpecialPricingInner
		return ret
	}
	return o.SpecialPricing
}

// GetSpecialPricingOk returns a tuple with the SpecialPricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) GetSpecialPricingOk() ([]PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerSpecialPricingInner, bool) {
	if o == nil || IsNil(o.SpecialPricing) {
		return nil, false
	}
	return o.SpecialPricing, true
}

// HasSpecialPricing returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) HasSpecialPricing() bool {
	if o != nil && !IsNil(o.SpecialPricing) {
		return true
	}

	return false
}

// SetSpecialPricing gets a reference to the given []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerSpecialPricingInner and assigns it to the SpecialPricing field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) SetSpecialPricing(v []PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerSpecialPricingInner) {
	o.SpecialPricing = v
}

func (o PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VolumeDiscounts) {
		toSerialize["volumeDiscounts"] = o.VolumeDiscounts
	}
	if !IsNil(o.SpecialPricing) {
		toSerialize["specialPricing"] = o.SpecialPricing
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner struct {
	value *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) Get() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) Set(val *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner(val *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner {
	return &NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


