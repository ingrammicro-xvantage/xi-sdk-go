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

// checks if the ProductDetailResponseSubscriptionDetailsInnerGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetailResponseSubscriptionDetailsInnerGroupsInner{}

// ProductDetailResponseSubscriptionDetailsInnerGroupsInner struct for ProductDetailResponseSubscriptionDetailsInnerGroupsInner
type ProductDetailResponseSubscriptionDetailsInnerGroupsInner struct {
	// Name of the group.
	GroupName *string `json:"groupName,omitempty"`
	// Name of the subscription group.
	GroupDescription *string `json:"groupDescription,omitempty"`
}

// NewProductDetailResponseSubscriptionDetailsInnerGroupsInner instantiates a new ProductDetailResponseSubscriptionDetailsInnerGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailResponseSubscriptionDetailsInnerGroupsInner() *ProductDetailResponseSubscriptionDetailsInnerGroupsInner {
	this := ProductDetailResponseSubscriptionDetailsInnerGroupsInner{}
	return &this
}

// NewProductDetailResponseSubscriptionDetailsInnerGroupsInnerWithDefaults instantiates a new ProductDetailResponseSubscriptionDetailsInnerGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailResponseSubscriptionDetailsInnerGroupsInnerWithDefaults() *ProductDetailResponseSubscriptionDetailsInnerGroupsInner {
	this := ProductDetailResponseSubscriptionDetailsInnerGroupsInner{}
	return &this
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) SetGroupName(v string) {
	o.GroupName = &v
}

// GetGroupDescription returns the GroupDescription field value if set, zero value otherwise.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) GetGroupDescription() string {
	if o == nil || IsNil(o.GroupDescription) {
		var ret string
		return ret
	}
	return *o.GroupDescription
}

// GetGroupDescriptionOk returns a tuple with the GroupDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) GetGroupDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.GroupDescription) {
		return nil, false
	}
	return o.GroupDescription, true
}

// HasGroupDescription returns a boolean if a field has been set.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) HasGroupDescription() bool {
	if o != nil && !IsNil(o.GroupDescription) {
		return true
	}

	return false
}

// SetGroupDescription gets a reference to the given string and assigns it to the GroupDescription field.
func (o *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) SetGroupDescription(v string) {
	o.GroupDescription = &v
}

func (o ProductDetailResponseSubscriptionDetailsInnerGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDetailResponseSubscriptionDetailsInnerGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.GroupDescription) {
		toSerialize["groupDescription"] = o.GroupDescription
	}
	return toSerialize, nil
}

type NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner struct {
	value *ProductDetailResponseSubscriptionDetailsInnerGroupsInner
	isSet bool
}

func (v NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner) Get() *ProductDetailResponseSubscriptionDetailsInnerGroupsInner {
	return v.value
}

func (v *NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner) Set(val *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailResponseSubscriptionDetailsInnerGroupsInner(val *ProductDetailResponseSubscriptionDetailsInnerGroupsInner) *NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner {
	return &NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner{value: val, isSet: true}
}

func (v NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDetailResponseSubscriptionDetailsInnerGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


