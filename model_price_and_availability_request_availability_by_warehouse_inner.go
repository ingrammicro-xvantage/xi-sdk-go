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

// checks if the PriceAndAvailabilityRequestAvailabilityByWarehouseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityRequestAvailabilityByWarehouseInner{}

// PriceAndAvailabilityRequestAvailabilityByWarehouseInner struct for PriceAndAvailabilityRequestAvailabilityByWarehouseInner
type PriceAndAvailabilityRequestAvailabilityByWarehouseInner struct {
	// Plant/warehouse Id of a particular location in order to get just the inventory of that location.
	AvailabilityByWarehouseId *string `json:"availabilityByWarehouseId,omitempty"`
	// Pass boolean value as input, if true the response will contain warehouse location details, if false the response will not hold warehouse location details. By default value is true.
	AvailabilityForAllLocation *bool `json:"availabilityForAllLocation,omitempty"`
}

// NewPriceAndAvailabilityRequestAvailabilityByWarehouseInner instantiates a new PriceAndAvailabilityRequestAvailabilityByWarehouseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityRequestAvailabilityByWarehouseInner() *PriceAndAvailabilityRequestAvailabilityByWarehouseInner {
	this := PriceAndAvailabilityRequestAvailabilityByWarehouseInner{}
	return &this
}

// NewPriceAndAvailabilityRequestAvailabilityByWarehouseInnerWithDefaults instantiates a new PriceAndAvailabilityRequestAvailabilityByWarehouseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityRequestAvailabilityByWarehouseInnerWithDefaults() *PriceAndAvailabilityRequestAvailabilityByWarehouseInner {
	this := PriceAndAvailabilityRequestAvailabilityByWarehouseInner{}
	return &this
}

// GetAvailabilityByWarehouseId returns the AvailabilityByWarehouseId field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityByWarehouseId() string {
	if o == nil || IsNil(o.AvailabilityByWarehouseId) {
		var ret string
		return ret
	}
	return *o.AvailabilityByWarehouseId
}

// GetAvailabilityByWarehouseIdOk returns a tuple with the AvailabilityByWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityByWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityByWarehouseId) {
		return nil, false
	}
	return o.AvailabilityByWarehouseId, true
}

// HasAvailabilityByWarehouseId returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) HasAvailabilityByWarehouseId() bool {
	if o != nil && !IsNil(o.AvailabilityByWarehouseId) {
		return true
	}

	return false
}

// SetAvailabilityByWarehouseId gets a reference to the given string and assigns it to the AvailabilityByWarehouseId field.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) SetAvailabilityByWarehouseId(v string) {
	o.AvailabilityByWarehouseId = &v
}

// GetAvailabilityForAllLocation returns the AvailabilityForAllLocation field value if set, zero value otherwise.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityForAllLocation() bool {
	if o == nil || IsNil(o.AvailabilityForAllLocation) {
		var ret bool
		return ret
	}
	return *o.AvailabilityForAllLocation
}

// GetAvailabilityForAllLocationOk returns a tuple with the AvailabilityForAllLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) GetAvailabilityForAllLocationOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailabilityForAllLocation) {
		return nil, false
	}
	return o.AvailabilityForAllLocation, true
}

// HasAvailabilityForAllLocation returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) HasAvailabilityForAllLocation() bool {
	if o != nil && !IsNil(o.AvailabilityForAllLocation) {
		return true
	}

	return false
}

// SetAvailabilityForAllLocation gets a reference to the given bool and assigns it to the AvailabilityForAllLocation field.
func (o *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) SetAvailabilityForAllLocation(v bool) {
	o.AvailabilityForAllLocation = &v
}

func (o PriceAndAvailabilityRequestAvailabilityByWarehouseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityRequestAvailabilityByWarehouseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AvailabilityByWarehouseId) {
		toSerialize["availabilityByWarehouseId"] = o.AvailabilityByWarehouseId
	}
	if !IsNil(o.AvailabilityForAllLocation) {
		toSerialize["availabilityForAllLocation"] = o.AvailabilityForAllLocation
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner struct {
	value *PriceAndAvailabilityRequestAvailabilityByWarehouseInner
	isSet bool
}

func (v NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner) Get() *PriceAndAvailabilityRequestAvailabilityByWarehouseInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner) Set(val *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner(val *PriceAndAvailabilityRequestAvailabilityByWarehouseInner) *NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner {
	return &NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityRequestAvailabilityByWarehouseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


