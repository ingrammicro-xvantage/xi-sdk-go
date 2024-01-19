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

// checks if the ValidateQuoteResponseVmfAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateQuoteResponseVmfAdditionalAttributesInner{}

// ValidateQuoteResponseVmfAdditionalAttributesInner struct for ValidateQuoteResponseVmfAdditionalAttributesInner
type ValidateQuoteResponseVmfAdditionalAttributesInner struct {
	// The name of the header level field.
	AttributeName *string `json:"attributeName,omitempty"`
	// The value of the header level field.
	AttributeValue *string `json:"attributeValue,omitempty"`
	// The description of the header level field.
	AttributeDescription *string `json:"attributeDescription,omitempty"`
}

// NewValidateQuoteResponseVmfAdditionalAttributesInner instantiates a new ValidateQuoteResponseVmfAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateQuoteResponseVmfAdditionalAttributesInner() *ValidateQuoteResponseVmfAdditionalAttributesInner {
	this := ValidateQuoteResponseVmfAdditionalAttributesInner{}
	return &this
}

// NewValidateQuoteResponseVmfAdditionalAttributesInnerWithDefaults instantiates a new ValidateQuoteResponseVmfAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateQuoteResponseVmfAdditionalAttributesInnerWithDefaults() *ValidateQuoteResponseVmfAdditionalAttributesInner {
	this := ValidateQuoteResponseVmfAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetAttributeDescription returns the AttributeDescription field value if set, zero value otherwise.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) GetAttributeDescription() string {
	if o == nil || IsNil(o.AttributeDescription) {
		var ret string
		return ret
	}
	return *o.AttributeDescription
}

// GetAttributeDescriptionOk returns a tuple with the AttributeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) GetAttributeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeDescription) {
		return nil, false
	}
	return o.AttributeDescription, true
}

// HasAttributeDescription returns a boolean if a field has been set.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) HasAttributeDescription() bool {
	if o != nil && !IsNil(o.AttributeDescription) {
		return true
	}

	return false
}

// SetAttributeDescription gets a reference to the given string and assigns it to the AttributeDescription field.
func (o *ValidateQuoteResponseVmfAdditionalAttributesInner) SetAttributeDescription(v string) {
	o.AttributeDescription = &v
}

func (o ValidateQuoteResponseVmfAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateQuoteResponseVmfAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !IsNil(o.AttributeDescription) {
		toSerialize["attributeDescription"] = o.AttributeDescription
	}
	return toSerialize, nil
}

type NullableValidateQuoteResponseVmfAdditionalAttributesInner struct {
	value *ValidateQuoteResponseVmfAdditionalAttributesInner
	isSet bool
}

func (v NullableValidateQuoteResponseVmfAdditionalAttributesInner) Get() *ValidateQuoteResponseVmfAdditionalAttributesInner {
	return v.value
}

func (v *NullableValidateQuoteResponseVmfAdditionalAttributesInner) Set(val *ValidateQuoteResponseVmfAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateQuoteResponseVmfAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateQuoteResponseVmfAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateQuoteResponseVmfAdditionalAttributesInner(val *ValidateQuoteResponseVmfAdditionalAttributesInner) *NullableValidateQuoteResponseVmfAdditionalAttributesInner {
	return &NullableValidateQuoteResponseVmfAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableValidateQuoteResponseVmfAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateQuoteResponseVmfAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


