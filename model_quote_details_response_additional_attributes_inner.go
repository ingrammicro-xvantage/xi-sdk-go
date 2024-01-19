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

// checks if the QuoteDetailsResponseAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsResponseAdditionalAttributesInner{}

// QuoteDetailsResponseAdditionalAttributesInner struct for QuoteDetailsResponseAdditionalAttributesInner
type QuoteDetailsResponseAdditionalAttributesInner struct {
	// estimateId - is the identification number for an estimate provided by Cisco for a quote.  dealId - is the identification number for the specific deal pricing related to a Cisco quote  vendorName - Name of Vendor associated with the quote.  vendorMessage - Vendor Message is associated with primary vendor in the quote.  In cases where a vendor requires a message be presented in the quote, the vendor name and message will be retreived and must be included in the quote vendor message fields.
	AttributeName *string `json:"attributeName,omitempty"`
	// The attribute field data.
	AttributeValue *string `json:"attributeValue,omitempty"`
}

// NewQuoteDetailsResponseAdditionalAttributesInner instantiates a new QuoteDetailsResponseAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsResponseAdditionalAttributesInner() *QuoteDetailsResponseAdditionalAttributesInner {
	this := QuoteDetailsResponseAdditionalAttributesInner{}
	return &this
}

// NewQuoteDetailsResponseAdditionalAttributesInnerWithDefaults instantiates a new QuoteDetailsResponseAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsResponseAdditionalAttributesInnerWithDefaults() *QuoteDetailsResponseAdditionalAttributesInner {
	this := QuoteDetailsResponseAdditionalAttributesInner{}
	return &this
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *QuoteDetailsResponseAdditionalAttributesInner) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *QuoteDetailsResponseAdditionalAttributesInner) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *QuoteDetailsResponseAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *QuoteDetailsResponseAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

func (o QuoteDetailsResponseAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsResponseAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	return toSerialize, nil
}

type NullableQuoteDetailsResponseAdditionalAttributesInner struct {
	value *QuoteDetailsResponseAdditionalAttributesInner
	isSet bool
}

func (v NullableQuoteDetailsResponseAdditionalAttributesInner) Get() *QuoteDetailsResponseAdditionalAttributesInner {
	return v.value
}

func (v *NullableQuoteDetailsResponseAdditionalAttributesInner) Set(val *QuoteDetailsResponseAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsResponseAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsResponseAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsResponseAdditionalAttributesInner(val *QuoteDetailsResponseAdditionalAttributesInner) *NullableQuoteDetailsResponseAdditionalAttributesInner {
	return &NullableQuoteDetailsResponseAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableQuoteDetailsResponseAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsResponseAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


