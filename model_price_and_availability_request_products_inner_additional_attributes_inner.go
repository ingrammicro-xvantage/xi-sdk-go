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

// checks if the PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner{}

// PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner struct for PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner
type PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner struct {
	// governmentprogramcode: Special Discount details will be provided based on the governmentprogramcode if available. shiptostatebrazil: Attribute Specific to Brazil. shipfrombranchnumber: If provided, displays only the availability of the specified branch number.
	AttributeName NullableString `json:"attributeName,omitempty"`
	// key value pair -key value.
	AttributeValue NullableString `json:"attributeValue,omitempty"`
}

// NewPriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner instantiates a new PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner() *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner {
	this := PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner{}
	return &this
}

// NewPriceAndAvailabilityRequestProductsInnerAdditionalAttributesInnerWithDefaults instantiates a new PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceAndAvailabilityRequestProductsInnerAdditionalAttributesInnerWithDefaults() *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner {
	this := PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName.Get()) {
		var ret string
		return ret
	}
	return *o.AttributeName.Get()
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeName.Get(), o.AttributeName.IsSet()
}

// HasAttributeName returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && o.AttributeName.IsSet() {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given NullableString and assigns it to the AttributeName field.
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName.Set(&v)
}
// SetAttributeNameNil sets the value for AttributeName to be an explicit nil
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) SetAttributeNameNil() {
	o.AttributeName.Set(nil)
}

// UnsetAttributeName ensures that no value is present for AttributeName, not even an explicit nil
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) UnsetAttributeName() {
	o.AttributeName.Unset()
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue.Get()) {
		var ret string
		return ret
	}
	return *o.AttributeValue.Get()
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeValue.Get(), o.AttributeValue.IsSet()
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && o.AttributeValue.IsSet() {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given NullableString and assigns it to the AttributeValue field.
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue.Set(&v)
}
// SetAttributeValueNil sets the value for AttributeValue to be an explicit nil
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) SetAttributeValueNil() {
	o.AttributeValue.Set(nil)
}

// UnsetAttributeValue ensures that no value is present for AttributeValue, not even an explicit nil
func (o *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) UnsetAttributeValue() {
	o.AttributeValue.Unset()
}

func (o PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AttributeName.IsSet() {
		toSerialize["attributeName"] = o.AttributeName.Get()
	}
	if o.AttributeValue.IsSet() {
		toSerialize["attributeValue"] = o.AttributeValue.Get()
	}
	return toSerialize, nil
}

type NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner struct {
	value *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner
	isSet bool
}

func (v NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) Get() *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner {
	return v.value
}

func (v *NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) Set(val *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner(val *PriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) *NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner {
	return &NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceAndAvailabilityRequestProductsInnerAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


