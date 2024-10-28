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

// checks if the ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner{}

// ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner struct for ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner
type ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner struct {
	// Provides the details of the product.
	SubscriptionPeriodUnit *string `json:"subscriptionPeriodUnit,omitempty"`
	// The URL endpoint for accessing the relevant data..
	SubscriptionPeriod *string `json:"subscriptionPeriod,omitempty"`
}

// NewProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner instantiates a new ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner() *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner {
	this := ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner{}
	return &this
}

// NewProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInnerWithDefaults instantiates a new ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInnerWithDefaults() *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner {
	this := ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner{}
	return &this
}

// GetSubscriptionPeriodUnit returns the SubscriptionPeriodUnit field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) GetSubscriptionPeriodUnit() string {
	if o == nil || IsNil(o.SubscriptionPeriodUnit) {
		var ret string
		return ret
	}
	return *o.SubscriptionPeriodUnit
}

// GetSubscriptionPeriodUnitOk returns a tuple with the SubscriptionPeriodUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) GetSubscriptionPeriodUnitOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionPeriodUnit) {
		return nil, false
	}
	return o.SubscriptionPeriodUnit, true
}

// HasSubscriptionPeriodUnit returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) HasSubscriptionPeriodUnit() bool {
	if o != nil && !IsNil(o.SubscriptionPeriodUnit) {
		return true
	}

	return false
}

// SetSubscriptionPeriodUnit gets a reference to the given string and assigns it to the SubscriptionPeriodUnit field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) SetSubscriptionPeriodUnit(v string) {
	o.SubscriptionPeriodUnit = &v
}

// GetSubscriptionPeriod returns the SubscriptionPeriod field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) GetSubscriptionPeriod() string {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		var ret string
		return ret
	}
	return *o.SubscriptionPeriod
}

// GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) GetSubscriptionPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionPeriod) {
		return nil, false
	}
	return o.SubscriptionPeriod, true
}

// HasSubscriptionPeriod returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) HasSubscriptionPeriod() bool {
	if o != nil && !IsNil(o.SubscriptionPeriod) {
		return true
	}

	return false
}

// SetSubscriptionPeriod gets a reference to the given string and assigns it to the SubscriptionPeriod field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) SetSubscriptionPeriod(v string) {
	o.SubscriptionPeriod = &v
}

func (o ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionPeriodUnit) {
		toSerialize["subscriptionPeriodUnit"] = o.SubscriptionPeriodUnit
	}
	if !IsNil(o.SubscriptionPeriod) {
		toSerialize["subscriptionPeriod"] = o.SubscriptionPeriod
	}
	return toSerialize, nil
}

type NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner struct {
	value *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner
	isSet bool
}

func (v NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) Get() *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner {
	return v.value
}

func (v *NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) Set(val *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner(val *ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) *NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner {
	return &NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner{value: val, isSet: true}
}

func (v NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


