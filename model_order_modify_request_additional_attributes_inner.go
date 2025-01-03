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

// checks if the OrderModifyRequestAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyRequestAdditionalAttributesInner{}

// OrderModifyRequestAdditionalAttributesInner struct for OrderModifyRequestAdditionalAttributesInner
type OrderModifyRequestAdditionalAttributesInner struct {
	// Example values are 'entryMethod', 'enableCommentsAsLines', 'regionCode'
	AttributeName *string `json:"attributeName,omitempty"`
	// Attribute Value
	AttributeValue *string `json:"attributeValue,omitempty"`
}

// NewOrderModifyRequestAdditionalAttributesInner instantiates a new OrderModifyRequestAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyRequestAdditionalAttributesInner() *OrderModifyRequestAdditionalAttributesInner {
	this := OrderModifyRequestAdditionalAttributesInner{}
	return &this
}

// NewOrderModifyRequestAdditionalAttributesInnerWithDefaults instantiates a new OrderModifyRequestAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyRequestAdditionalAttributesInnerWithDefaults() *OrderModifyRequestAdditionalAttributesInner {
	this := OrderModifyRequestAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *OrderModifyRequestAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *OrderModifyRequestAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *OrderModifyRequestAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *OrderModifyRequestAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *OrderModifyRequestAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *OrderModifyRequestAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

func (o OrderModifyRequestAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyRequestAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return toSerialize, nil
}

type NullableOrderModifyRequestAdditionalAttributesInner struct {
	value *OrderModifyRequestAdditionalAttributesInner
	isSet bool
}

func (v NullableOrderModifyRequestAdditionalAttributesInner) Get() *OrderModifyRequestAdditionalAttributesInner {
	return v.value
}

func (v *NullableOrderModifyRequestAdditionalAttributesInner) Set(val *OrderModifyRequestAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyRequestAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyRequestAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyRequestAdditionalAttributesInner(val *OrderModifyRequestAdditionalAttributesInner) *NullableOrderModifyRequestAdditionalAttributesInner {
	return &NullableOrderModifyRequestAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableOrderModifyRequestAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyRequestAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


