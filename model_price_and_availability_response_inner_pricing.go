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

// checks if the PriceAndAvailabilityResponseInnerPricing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityResponseInnerPricing{}

// PriceAndAvailabilityResponseInnerPricing struct for PriceAndAvailabilityResponseInnerPricing
type PriceAndAvailabilityResponseInnerPricing struct {
	// The 3-digit ISO currency code.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The retail price of the product.
	RetailPrice NullableFloat32 `json:"retailPrice,omitempty"`
	// Minimum Advertised Price (MAP). If required by the vendor, resellers can not sell below MAP price.
	MapPrice NullableFloat32 `json:"mapPrice,omitempty"`
	// The price customer pays after all special pricing and discounts have been applied.
	CustomerPrice NullableFloat32 `json:"customerPrice,omitempty"`
	// Boolean values specifies whether special Bid discounts are available for the product.
	SpecialBidPricingAvailable NullableBool `json:"specialBidPricingAvailable,omitempty"`
	// Boolean values specifies whether web Discounts are available for the product.
	WebDiscountsAvailable NullableBool `json:"webDiscountsAvailable,omitempty"`
}

// NewPriceAndAvailabilityResponseInnerPricing instantiates a new PriceAndAvailabilityResponseInnerPricing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityResponseInnerPricing() *PriceAndAvailabilityResponseInnerPricing {
	this := PriceAndAvailabilityResponseInnerPricing{}
	return &this
}

// NewPriceAndAvailabilityResponseInnerPricingWithDefaults instantiates a new PriceAndAvailabilityResponseInnerPricing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityResponseInnerPricingWithDefaults() *PriceAndAvailabilityResponseInnerPricing {
	this := PriceAndAvailabilityResponseInnerPricing{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PriceAndAvailabilityResponseInnerPricing) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetRetailPrice returns the RetailPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerPricing) GetRetailPrice() float32 {
	if o == nil || IsNil(o.RetailPrice.Get()) {
		var ret float32
		return ret
	}
	return *o.RetailPrice.Get()
}

// GetRetailPriceOk returns a tuple with the RetailPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerPricing) GetRetailPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetailPrice.Get(), o.RetailPrice.IsSet()
}

// HasRetailPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasRetailPrice() bool {
	if o != nil && o.RetailPrice.IsSet() {
		return true
	}

	return false
}

// SetRetailPrice gets a reference to the given NullableFloat32 and assigns it to the RetailPrice field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetRetailPrice(v float32) {
	o.RetailPrice.Set(&v)
}
// SetRetailPriceNil sets the value for RetailPrice to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) SetRetailPriceNil() {
	o.RetailPrice.Set(nil)
}

// UnsetRetailPrice ensures that no value is present for RetailPrice, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) UnsetRetailPrice() {
	o.RetailPrice.Unset()
}

// GetMapPrice returns the MapPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerPricing) GetMapPrice() float32 {
	if o == nil || IsNil(o.MapPrice.Get()) {
		var ret float32
		return ret
	}
	return *o.MapPrice.Get()
}

// GetMapPriceOk returns a tuple with the MapPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerPricing) GetMapPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MapPrice.Get(), o.MapPrice.IsSet()
}

// HasMapPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasMapPrice() bool {
	if o != nil && o.MapPrice.IsSet() {
		return true
	}

	return false
}

// SetMapPrice gets a reference to the given NullableFloat32 and assigns it to the MapPrice field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetMapPrice(v float32) {
	o.MapPrice.Set(&v)
}
// SetMapPriceNil sets the value for MapPrice to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) SetMapPriceNil() {
	o.MapPrice.Set(nil)
}

// UnsetMapPrice ensures that no value is present for MapPrice, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) UnsetMapPrice() {
	o.MapPrice.Unset()
}

// GetCustomerPrice returns the CustomerPrice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerPricing) GetCustomerPrice() float32 {
	if o == nil || IsNil(o.CustomerPrice.Get()) {
		var ret float32
		return ret
	}
	return *o.CustomerPrice.Get()
}

// GetCustomerPriceOk returns a tuple with the CustomerPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerPricing) GetCustomerPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerPrice.Get(), o.CustomerPrice.IsSet()
}

// HasCustomerPrice returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasCustomerPrice() bool {
	if o != nil && o.CustomerPrice.IsSet() {
		return true
	}

	return false
}

// SetCustomerPrice gets a reference to the given NullableFloat32 and assigns it to the CustomerPrice field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetCustomerPrice(v float32) {
	o.CustomerPrice.Set(&v)
}
// SetCustomerPriceNil sets the value for CustomerPrice to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) SetCustomerPriceNil() {
	o.CustomerPrice.Set(nil)
}

// UnsetCustomerPrice ensures that no value is present for CustomerPrice, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) UnsetCustomerPrice() {
	o.CustomerPrice.Unset()
}

// GetSpecialBidPricingAvailable returns the SpecialBidPricingAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerPricing) GetSpecialBidPricingAvailable() bool {
	if o == nil || IsNil(o.SpecialBidPricingAvailable.Get()) {
		var ret bool
		return ret
	}
	return *o.SpecialBidPricingAvailable.Get()
}

// GetSpecialBidPricingAvailableOk returns a tuple with the SpecialBidPricingAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerPricing) GetSpecialBidPricingAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecialBidPricingAvailable.Get(), o.SpecialBidPricingAvailable.IsSet()
}

// HasSpecialBidPricingAvailable returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasSpecialBidPricingAvailable() bool {
	if o != nil && o.SpecialBidPricingAvailable.IsSet() {
		return true
	}

	return false
}

// SetSpecialBidPricingAvailable gets a reference to the given NullableBool and assigns it to the SpecialBidPricingAvailable field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetSpecialBidPricingAvailable(v bool) {
	o.SpecialBidPricingAvailable.Set(&v)
}
// SetSpecialBidPricingAvailableNil sets the value for SpecialBidPricingAvailable to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) SetSpecialBidPricingAvailableNil() {
	o.SpecialBidPricingAvailable.Set(nil)
}

// UnsetSpecialBidPricingAvailable ensures that no value is present for SpecialBidPricingAvailable, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) UnsetSpecialBidPricingAvailable() {
	o.SpecialBidPricingAvailable.Unset()
}

// GetWebDiscountsAvailable returns the WebDiscountsAvailable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityResponseInnerPricing) GetWebDiscountsAvailable() bool {
	if o == nil || IsNil(o.WebDiscountsAvailable.Get()) {
		var ret bool
		return ret
	}
	return *o.WebDiscountsAvailable.Get()
}

// GetWebDiscountsAvailableOk returns a tuple with the WebDiscountsAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityResponseInnerPricing) GetWebDiscountsAvailableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebDiscountsAvailable.Get(), o.WebDiscountsAvailable.IsSet()
}

// HasWebDiscountsAvailable returns a boolean if a field has been set.
func (o *PriceAndAvailabilityResponseInnerPricing) HasWebDiscountsAvailable() bool {
	if o != nil && o.WebDiscountsAvailable.IsSet() {
		return true
	}

	return false
}

// SetWebDiscountsAvailable gets a reference to the given NullableBool and assigns it to the WebDiscountsAvailable field.
func (o *PriceAndAvailabilityResponseInnerPricing) SetWebDiscountsAvailable(v bool) {
	o.WebDiscountsAvailable.Set(&v)
}
// SetWebDiscountsAvailableNil sets the value for WebDiscountsAvailable to be an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) SetWebDiscountsAvailableNil() {
	o.WebDiscountsAvailable.Set(nil)
}

// UnsetWebDiscountsAvailable ensures that no value is present for WebDiscountsAvailable, not even an explicit nil
func (o *PriceAndAvailabilityResponseInnerPricing) UnsetWebDiscountsAvailable() {
	o.WebDiscountsAvailable.Unset()
}

func (o PriceAndAvailabilityResponseInnerPricing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityResponseInnerPricing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if o.RetailPrice.IsSet() {
		toSerialize["retailPrice"] = o.RetailPrice.Get()
	}
	if o.MapPrice.IsSet() {
		toSerialize["mapPrice"] = o.MapPrice.Get()
	}
	if o.CustomerPrice.IsSet() {
		toSerialize["customerPrice"] = o.CustomerPrice.Get()
	}
	if o.SpecialBidPricingAvailable.IsSet() {
		toSerialize["specialBidPricingAvailable"] = o.SpecialBidPricingAvailable.Get()
	}
	if o.WebDiscountsAvailable.IsSet() {
		toSerialize["webDiscountsAvailable"] = o.WebDiscountsAvailable.Get()
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityResponseInnerPricing struct {
	value *PriceAndAvailabilityResponseInnerPricing
	isSet bool
}

func (v NullablePriceAndAvailabilityResponseInnerPricing) Get() *PriceAndAvailabilityResponseInnerPricing {
	return v.value
}

func (v *NullablePriceAndAvailabilityResponseInnerPricing) Set(val *PriceAndAvailabilityResponseInnerPricing) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityResponseInnerPricing) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityResponseInnerPricing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityResponseInnerPricing(val *PriceAndAvailabilityResponseInnerPricing) *NullablePriceAndAvailabilityResponseInnerPricing {
	return &NullablePriceAndAvailabilityResponseInnerPricing{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityResponseInnerPricing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityResponseInnerPricing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


