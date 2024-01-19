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

// checks if the MultiSKUPriceAndStockResponseServiceresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiSKUPriceAndStockResponseServiceresponse{}

// MultiSKUPriceAndStockResponseServiceresponse struct for MultiSKUPriceAndStockResponseServiceresponse
type MultiSKUPriceAndStockResponseServiceresponse struct {
	Responsepreamble *PriceAndAvailabilityResponseServiceresponseResponsepreamble `json:"responsepreamble,omitempty"`
	Priceandstockresponse *MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponse `json:"priceandstockresponse,omitempty"`
}

// NewMultiSKUPriceAndStockResponseServiceresponse instantiates a new MultiSKUPriceAndStockResponseServiceresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSKUPriceAndStockResponseServiceresponse() *MultiSKUPriceAndStockResponseServiceresponse {
	this := MultiSKUPriceAndStockResponseServiceresponse{}
	return &this
}

// NewMultiSKUPriceAndStockResponseServiceresponseWithDefaults instantiates a new MultiSKUPriceAndStockResponseServiceresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSKUPriceAndStockResponseServiceresponseWithDefaults() *MultiSKUPriceAndStockResponseServiceresponse {
	this := MultiSKUPriceAndStockResponseServiceresponse{}
	return &this
}

// GetResponsepreamble returns the Responsepreamble field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockResponseServiceresponse) GetResponsepreamble() PriceAndAvailabilityResponseServiceresponseResponsepreamble {
	if o == nil || IsNil(o.Responsepreamble) {
		var ret PriceAndAvailabilityResponseServiceresponseResponsepreamble
		return ret
	}
	return *o.Responsepreamble
}

// GetResponsepreambleOk returns a tuple with the Responsepreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockResponseServiceresponse) GetResponsepreambleOk() (*PriceAndAvailabilityResponseServiceresponseResponsepreamble, bool) {
	if o == nil || IsNil(o.Responsepreamble) {
		return nil, false
	}
	return o.Responsepreamble, true
}

// HasResponsepreamble returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockResponseServiceresponse) HasResponsepreamble() bool {
	if o != nil && !IsNil(o.Responsepreamble) {
		return true
	}

	return false
}

// SetResponsepreamble gets a reference to the given PriceAndAvailabilityResponseServiceresponseResponsepreamble and assigns it to the Responsepreamble field.
func (o *MultiSKUPriceAndStockResponseServiceresponse) SetResponsepreamble(v PriceAndAvailabilityResponseServiceresponseResponsepreamble) {
	o.Responsepreamble = &v
}

// GetPriceandstockresponse returns the Priceandstockresponse field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockResponseServiceresponse) GetPriceandstockresponse() MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponse {
	if o == nil || IsNil(o.Priceandstockresponse) {
		var ret MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponse
		return ret
	}
	return *o.Priceandstockresponse
}

// GetPriceandstockresponseOk returns a tuple with the Priceandstockresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockResponseServiceresponse) GetPriceandstockresponseOk() (*MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponse, bool) {
	if o == nil || IsNil(o.Priceandstockresponse) {
		return nil, false
	}
	return o.Priceandstockresponse, true
}

// HasPriceandstockresponse returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockResponseServiceresponse) HasPriceandstockresponse() bool {
	if o != nil && !IsNil(o.Priceandstockresponse) {
		return true
	}

	return false
}

// SetPriceandstockresponse gets a reference to the given MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponse and assigns it to the Priceandstockresponse field.
func (o *MultiSKUPriceAndStockResponseServiceresponse) SetPriceandstockresponse(v MultiSKUPriceAndStockResponseServiceresponsePriceandstockresponse) {
	o.Priceandstockresponse = &v
}

func (o MultiSKUPriceAndStockResponseServiceresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiSKUPriceAndStockResponseServiceresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsepreamble) {
		toSerialize["responsepreamble"] = o.Responsepreamble
	}
	if !IsNil(o.Priceandstockresponse) {
		toSerialize["priceandstockresponse"] = o.Priceandstockresponse
	}
	return toSerialize, nil
}

type NullableMultiSKUPriceAndStockResponseServiceresponse struct {
	value *MultiSKUPriceAndStockResponseServiceresponse
	isSet bool
}

func (v NullableMultiSKUPriceAndStockResponseServiceresponse) Get() *MultiSKUPriceAndStockResponseServiceresponse {
	return v.value
}

func (v *NullableMultiSKUPriceAndStockResponseServiceresponse) Set(val *MultiSKUPriceAndStockResponseServiceresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSKUPriceAndStockResponseServiceresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSKUPriceAndStockResponseServiceresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSKUPriceAndStockResponseServiceresponse(val *MultiSKUPriceAndStockResponseServiceresponse) *NullableMultiSKUPriceAndStockResponseServiceresponse {
	return &NullableMultiSKUPriceAndStockResponseServiceresponse{value: val, isSet: true}
}

func (v NullableMultiSKUPriceAndStockResponseServiceresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSKUPriceAndStockResponseServiceresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


