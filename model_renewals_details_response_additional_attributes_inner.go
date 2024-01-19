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

// checks if the RenewalsDetailsResponseAdditionalAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsDetailsResponseAdditionalAttributesInner{}

// RenewalsDetailsResponseAdditionalAttributesInner struct for RenewalsDetailsResponseAdditionalAttributesInner
type RenewalsDetailsResponseAdditionalAttributesInner struct {
	// The description of the additional attribute.
	AttributeDescription *string `json:"attributeDescription,omitempty"`
	// The value of the additional attribute.
	AttributeValue *string `json:"attributeValue,omitempty"`
	// The attribute start date.
	StartDate *string `json:"startDate,omitempty"`
	// The attribute expiration date.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// Is the line item consolidated? Yes or No.
	IsConsolidated *string `json:"isConsolidated,omitempty"`
}

// NewRenewalsDetailsResponseAdditionalAttributesInner instantiates a new RenewalsDetailsResponseAdditionalAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsDetailsResponseAdditionalAttributesInner() *RenewalsDetailsResponseAdditionalAttributesInner {
	this := RenewalsDetailsResponseAdditionalAttributesInner{}
	return &this
}

// NewRenewalsDetailsResponseAdditionalAttributesInnerWithDefaults instantiates a new RenewalsDetailsResponseAdditionalAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsDetailsResponseAdditionalAttributesInnerWithDefaults() *RenewalsDetailsResponseAdditionalAttributesInner {
	this := RenewalsDetailsResponseAdditionalAttributesInner{}
	return &this
}

// GetAttributeDescription returns the AttributeDescription field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetAttributeDescription() string {
	if o == nil || IsNil(o.AttributeDescription) {
		var ret string
		return ret
	}
	return *o.AttributeDescription
}

// GetAttributeDescriptionOk returns a tuple with the AttributeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetAttributeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeDescription) {
		return nil, false
	}
	return o.AttributeDescription, true
}

// HasAttributeDescription returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) HasAttributeDescription() bool {
	if o != nil && !IsNil(o.AttributeDescription) {
		return true
	}

	return false
}

// SetAttributeDescription gets a reference to the given string and assigns it to the AttributeDescription field.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) SetAttributeDescription(v string) {
	o.AttributeDescription = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) SetStartDate(v string) {
	o.StartDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetIsConsolidated returns the IsConsolidated field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetIsConsolidated() string {
	if o == nil || IsNil(o.IsConsolidated) {
		var ret string
		return ret
	}
	return *o.IsConsolidated
}

// GetIsConsolidatedOk returns a tuple with the IsConsolidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) GetIsConsolidatedOk() (*string, bool) {
	if o == nil || IsNil(o.IsConsolidated) {
		return nil, false
	}
	return o.IsConsolidated, true
}

// HasIsConsolidated returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) HasIsConsolidated() bool {
	if o != nil && !IsNil(o.IsConsolidated) {
		return true
	}

	return false
}

// SetIsConsolidated gets a reference to the given string and assigns it to the IsConsolidated field.
func (o *RenewalsDetailsResponseAdditionalAttributesInner) SetIsConsolidated(v string) {
	o.IsConsolidated = &v
}

func (o RenewalsDetailsResponseAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsDetailsResponseAdditionalAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttributeDescription) {
		toSerialize["attributeDescription"] = o.AttributeDescription
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.IsConsolidated) {
		toSerialize["isConsolidated"] = o.IsConsolidated
	}
	return toSerialize, nil
}

type NullableRenewalsDetailsResponseAdditionalAttributesInner struct {
	value *RenewalsDetailsResponseAdditionalAttributesInner
	isSet bool
}

func (v NullableRenewalsDetailsResponseAdditionalAttributesInner) Get() *RenewalsDetailsResponseAdditionalAttributesInner {
	return v.value
}

func (v *NullableRenewalsDetailsResponseAdditionalAttributesInner) Set(val *RenewalsDetailsResponseAdditionalAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsDetailsResponseAdditionalAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsDetailsResponseAdditionalAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsDetailsResponseAdditionalAttributesInner(val *RenewalsDetailsResponseAdditionalAttributesInner) *NullableRenewalsDetailsResponseAdditionalAttributesInner {
	return &NullableRenewalsDetailsResponseAdditionalAttributesInner{value: val, isSet: true}
}

func (v NullableRenewalsDetailsResponseAdditionalAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsDetailsResponseAdditionalAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


