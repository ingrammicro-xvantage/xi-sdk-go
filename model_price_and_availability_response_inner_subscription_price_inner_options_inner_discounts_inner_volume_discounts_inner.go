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

// checks if the PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner{}

// PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner struct for PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner
type PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner struct {
	// The 3-digit ISO currency code.
	CurrencyCode NullableString `json:"currencyCode,omitempty"`
	// Quantity of the line item.
	Quantity NullableInt32 `json:"quantity,omitempty"`
	// Manufacturer Suggested Retail Price.
	Msrp NullableFloat32 `json:"msrp,omitempty"`
	// The unit price of the line item.
	UnitPrice NullableFloat32 `json:"unitPrice,omitempty"`
	// Reseller’s margin percentage
	Margin NullableFloat32 `json:"margin,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner {
	this := PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInnerWithDefaults instantiates a new PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInnerWithDefaults() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner {
	this := PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode.Get()) {
		var ret string
		return ret
	}
	return *o.CurrencyCode.Get()
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrencyCode.Get(), o.CurrencyCode.IsSet()
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode.IsSet() {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given NullableString and assigns it to the CurrencyCode field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetCurrencyCode(v string) {
	o.CurrencyCode.Set(&v)
}
// SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetCurrencyCodeNil() {
	o.CurrencyCode.Set(nil)
}

// UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) UnsetCurrencyCode() {
	o.CurrencyCode.Unset()
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetMsrp returns the Msrp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMsrp() float32 {
	if o == nil || IsNil(o.Msrp.Get()) {
		var ret float32
		return ret
	}
	return *o.Msrp.Get()
}

// GetMsrpOk returns a tuple with the Msrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMsrpOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Msrp.Get(), o.Msrp.IsSet()
}

// HasMsrp returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasMsrp() bool {
	if o != nil && o.Msrp.IsSet() {
		return true
	}

	return false
}

// SetMsrp gets a reference to the given NullableFloat32 and assigns it to the Msrp field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetMsrp(v float32) {
	o.Msrp.Set(&v)
}
// SetMsrpNil sets the value for Msrp to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetMsrpNil() {
	o.Msrp.Set(nil)
}

// UnsetMsrp ensures that no value is present for Msrp, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) UnsetMsrp() {
	o.Msrp.Unset()
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetUnitPrice() float32 {
	if o == nil || IsNil(o.UnitPrice.Get()) {
		var ret float32
		return ret
	}
	return *o.UnitPrice.Get()
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetUnitPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnitPrice.Get(), o.UnitPrice.IsSet()
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasUnitPrice() bool {
	if o != nil && o.UnitPrice.IsSet() {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given NullableFloat32 and assigns it to the UnitPrice field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetUnitPrice(v float32) {
	o.UnitPrice.Set(&v)
}
// SetUnitPriceNil sets the value for UnitPrice to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetUnitPriceNil() {
	o.UnitPrice.Set(nil)
}

// UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) UnsetUnitPrice() {
	o.UnitPrice.Unset()
}

// GetMargin returns the Margin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMargin() float32 {
	if o == nil || IsNil(o.Margin.Get()) {
		var ret float32
		return ret
	}
	return *o.Margin.Get()
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMarginOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Margin.Get(), o.Margin.IsSet()
}

// HasMargin returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasMargin() bool {
	if o != nil && o.Margin.IsSet() {
		return true
	}

	return false
}

// SetMargin gets a reference to the given NullableFloat32 and assigns it to the Margin field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetMargin(v float32) {
	o.Margin.Set(&v)
}
// SetMarginNil sets the value for Margin to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetMarginNil() {
	o.Margin.Set(nil)
}

// UnsetMargin ensures that no value is present for Margin, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) UnsetMargin() {
	o.Margin.Unset()
}

func (o PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrencyCode.IsSet() {
		toSerialize["currencyCode"] = o.CurrencyCode.Get()
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
	if o.Msrp.IsSet() {
		toSerialize["msrp"] = o.Msrp.Get()
	}
	if o.UnitPrice.IsSet() {
		toSerialize["unitPrice"] = o.UnitPrice.Get()
	}
	if o.Margin.IsSet() {
		toSerialize["margin"] = o.Margin.Get()
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner struct {
	value *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) Get() *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) Set(val *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner(val *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner {
	return &NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


