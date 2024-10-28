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

// checks if the ProductSearchResponseSubscriptionCatalogInnerPlansInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSearchResponseSubscriptionCatalogInnerPlansInner{}

// ProductSearchResponseSubscriptionCatalogInnerPlansInner struct for ProductSearchResponseSubscriptionCatalogInnerPlansInner
type ProductSearchResponseSubscriptionCatalogInnerPlansInner struct {
	// ID of the Plan.
	PlanId *string `json:"planId,omitempty"`
	// Name of the Plan.
	PlanName *string `json:"planName,omitempty"`
	// The description of the Plan
	PlanDescription *string `json:"planDescription,omitempty"`
	SubscriptionPeriodSummary []ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner `json:"subscriptionPeriodSummary,omitempty"`
	Links []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner `json:"links,omitempty"`
}

// NewProductSearchResponseSubscriptionCatalogInnerPlansInner instantiates a new ProductSearchResponseSubscriptionCatalogInnerPlansInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSearchResponseSubscriptionCatalogInnerPlansInner() *ProductSearchResponseSubscriptionCatalogInnerPlansInner {
	this := ProductSearchResponseSubscriptionCatalogInnerPlansInner{}
	return &this
}

// NewProductSearchResponseSubscriptionCatalogInnerPlansInnerWithDefaults instantiates a new ProductSearchResponseSubscriptionCatalogInnerPlansInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSearchResponseSubscriptionCatalogInnerPlansInnerWithDefaults() *ProductSearchResponseSubscriptionCatalogInnerPlansInner {
	this := ProductSearchResponseSubscriptionCatalogInnerPlansInner{}
	return &this
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetPlanId() string {
	if o == nil || IsNil(o.PlanId) {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlanId) {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) HasPlanId() bool {
	if o != nil && !IsNil(o.PlanId) {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) SetPlanId(v string) {
	o.PlanId = &v
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) SetPlanName(v string) {
	o.PlanName = &v
}

// GetPlanDescription returns the PlanDescription field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetPlanDescription() string {
	if o == nil || IsNil(o.PlanDescription) {
		var ret string
		return ret
	}
	return *o.PlanDescription
}

// GetPlanDescriptionOk returns a tuple with the PlanDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetPlanDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PlanDescription) {
		return nil, false
	}
	return o.PlanDescription, true
}

// HasPlanDescription returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) HasPlanDescription() bool {
	if o != nil && !IsNil(o.PlanDescription) {
		return true
	}

	return false
}

// SetPlanDescription gets a reference to the given string and assigns it to the PlanDescription field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) SetPlanDescription(v string) {
	o.PlanDescription = &v
}

// GetSubscriptionPeriodSummary returns the SubscriptionPeriodSummary field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetSubscriptionPeriodSummary() []ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner {
	if o == nil || IsNil(o.SubscriptionPeriodSummary) {
		var ret []ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner
		return ret
	}
	return o.SubscriptionPeriodSummary
}

// GetSubscriptionPeriodSummaryOk returns a tuple with the SubscriptionPeriodSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetSubscriptionPeriodSummaryOk() ([]ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner, bool) {
	if o == nil || IsNil(o.SubscriptionPeriodSummary) {
		return nil, false
	}
	return o.SubscriptionPeriodSummary, true
}

// HasSubscriptionPeriodSummary returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) HasSubscriptionPeriodSummary() bool {
	if o != nil && !IsNil(o.SubscriptionPeriodSummary) {
		return true
	}

	return false
}

// SetSubscriptionPeriodSummary gets a reference to the given []ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner and assigns it to the SubscriptionPeriodSummary field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) SetSubscriptionPeriodSummary(v []ProductSearchResponseSubscriptionCatalogInnerPlansInnerSubscriptionPeriodSummaryInner) {
	o.SubscriptionPeriodSummary = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetLinks() []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) GetLinksOk() ([]ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner and assigns it to the Links field.
func (o *ProductSearchResponseSubscriptionCatalogInnerPlansInner) SetLinks(v []ProductSearchResponseSubscriptionCatalogInnerPlansInnerLinksInner) {
	o.Links = v
}

func (o ProductSearchResponseSubscriptionCatalogInnerPlansInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSearchResponseSubscriptionCatalogInnerPlansInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlanId) {
		toSerialize["planId"] = o.PlanId
	}
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.PlanDescription) {
		toSerialize["planDescription"] = o.PlanDescription
	}
	if !IsNil(o.SubscriptionPeriodSummary) {
		toSerialize["subscriptionPeriodSummary"] = o.SubscriptionPeriodSummary
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableProductSearchResponseSubscriptionCatalogInnerPlansInner struct {
	value *ProductSearchResponseSubscriptionCatalogInnerPlansInner
	isSet bool
}

func (v NullableProductSearchResponseSubscriptionCatalogInnerPlansInner) Get() *ProductSearchResponseSubscriptionCatalogInnerPlansInner {
	return v.value
}

func (v *NullableProductSearchResponseSubscriptionCatalogInnerPlansInner) Set(val *ProductSearchResponseSubscriptionCatalogInnerPlansInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSearchResponseSubscriptionCatalogInnerPlansInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSearchResponseSubscriptionCatalogInnerPlansInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSearchResponseSubscriptionCatalogInnerPlansInner(val *ProductSearchResponseSubscriptionCatalogInnerPlansInner) *NullableProductSearchResponseSubscriptionCatalogInnerPlansInner {
	return &NullableProductSearchResponseSubscriptionCatalogInnerPlansInner{value: val, isSet: true}
}

func (v NullableProductSearchResponseSubscriptionCatalogInnerPlansInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSearchResponseSubscriptionCatalogInnerPlansInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


