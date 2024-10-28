/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

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
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Quantity of the line item.
	Quantity *int32 `json:"quantity,omitempty"`
	// Manufacturer Suggested Retail Price.
	Msrp *float32 `json:"msrp,omitempty"`
	// The unit price of the line item.
	UnitPrice *float32 `json:"unitPrice,omitempty"`
	// Reseller’s margin percentage
	Margin *float32 `json:"margin,omitempty"`
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

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetMsrp returns the Msrp field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMsrp() float32 {
	if o == nil || IsNil(o.Msrp) {
		var ret float32
		return ret
	}
	return *o.Msrp
}

// GetMsrpOk returns a tuple with the Msrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMsrpOk() (*float32, bool) {
	if o == nil || IsNil(o.Msrp) {
		return nil, false
	}
	return o.Msrp, true
}

// HasMsrp returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasMsrp() bool {
	if o != nil && !IsNil(o.Msrp) {
		return true
	}

	return false
}

// SetMsrp gets a reference to the given float32 and assigns it to the Msrp field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetMsrp(v float32) {
	o.Msrp = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetUnitPrice() float32 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMargin() float32 {
	if o == nil || IsNil(o.Margin) {
		var ret float32
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) GetMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.Margin) {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) HasMargin() bool {
	if o != nil && !IsNil(o.Margin) {
		return true
	}

	return false
}

// SetMargin gets a reference to the given float32 and assigns it to the Margin field.
func (o *PriceAndAvailabilityResponseInnerSubscriptionPriceInnerOptionsInnerDiscountsInnerVolumeDiscountsInner) SetMargin(v float32) {
	o.Margin = &v
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
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Msrp) {
		toSerialize["msrp"] = o.Msrp
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if !IsNil(o.Margin) {
		toSerialize["margin"] = o.Margin
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


