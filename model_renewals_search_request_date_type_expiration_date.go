/*
XI Sdk Resellers

For Ingram Micro Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the RenewalsSearchRequestDateTypeExpirationDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchRequestDateTypeExpirationDate{}

// RenewalsSearchRequestDateTypeExpirationDate struct for RenewalsSearchRequestDateTypeExpirationDate
type RenewalsSearchRequestDateTypeExpirationDate struct {
	// Custom start date for expiration date.
	CustomStartDate *string `json:"customStartDate,omitempty"`
	// Custom end date for expiration date.
	CustomEndDate *string `json:"customEndDate,omitempty"`
}

// NewRenewalsSearchRequestDateTypeExpirationDate instantiates a new RenewalsSearchRequestDateTypeExpirationDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchRequestDateTypeExpirationDate() *RenewalsSearchRequestDateTypeExpirationDate {
	this := RenewalsSearchRequestDateTypeExpirationDate{}
	return &this
}

// NewRenewalsSearchRequestDateTypeExpirationDateWithDefaults instantiates a new RenewalsSearchRequestDateTypeExpirationDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchRequestDateTypeExpirationDateWithDefaults() *RenewalsSearchRequestDateTypeExpirationDate {
	this := RenewalsSearchRequestDateTypeExpirationDate{}
	return &this
}

// GetCustomStartDate returns the CustomStartDate field value if set, zero value otherwise.
func (o *RenewalsSearchRequestDateTypeExpirationDate) GetCustomStartDate() string {
	if o == nil || IsNil(o.CustomStartDate) {
		var ret string
		return ret
	}
	return *o.CustomStartDate
}

// GetCustomStartDateOk returns a tuple with the CustomStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestDateTypeExpirationDate) GetCustomStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomStartDate) {
		return nil, false
	}
	return o.CustomStartDate, true
}

// HasCustomStartDate returns a boolean if a field has been set.
func (o *RenewalsSearchRequestDateTypeExpirationDate) HasCustomStartDate() bool {
	if o != nil && !IsNil(o.CustomStartDate) {
		return true
	}

	return false
}

// SetCustomStartDate gets a reference to the given string and assigns it to the CustomStartDate field.
func (o *RenewalsSearchRequestDateTypeExpirationDate) SetCustomStartDate(v string) {
	o.CustomStartDate = &v
}

// GetCustomEndDate returns the CustomEndDate field value if set, zero value otherwise.
func (o *RenewalsSearchRequestDateTypeExpirationDate) GetCustomEndDate() string {
	if o == nil || IsNil(o.CustomEndDate) {
		var ret string
		return ret
	}
	return *o.CustomEndDate
}

// GetCustomEndDateOk returns a tuple with the CustomEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestDateTypeExpirationDate) GetCustomEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEndDate) {
		return nil, false
	}
	return o.CustomEndDate, true
}

// HasCustomEndDate returns a boolean if a field has been set.
func (o *RenewalsSearchRequestDateTypeExpirationDate) HasCustomEndDate() bool {
	if o != nil && !IsNil(o.CustomEndDate) {
		return true
	}

	return false
}

// SetCustomEndDate gets a reference to the given string and assigns it to the CustomEndDate field.
func (o *RenewalsSearchRequestDateTypeExpirationDate) SetCustomEndDate(v string) {
	o.CustomEndDate = &v
}

func (o RenewalsSearchRequestDateTypeExpirationDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchRequestDateTypeExpirationDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomStartDate) {
		toSerialize["customStartDate"] = o.CustomStartDate
	}
	if !IsNil(o.CustomEndDate) {
		toSerialize["customEndDate"] = o.CustomEndDate
	}
	return toSerialize, nil
}

type NullableRenewalsSearchRequestDateTypeExpirationDate struct {
	value *RenewalsSearchRequestDateTypeExpirationDate
	isSet bool
}

func (v NullableRenewalsSearchRequestDateTypeExpirationDate) Get() *RenewalsSearchRequestDateTypeExpirationDate {
	return v.value
}

func (v *NullableRenewalsSearchRequestDateTypeExpirationDate) Set(val *RenewalsSearchRequestDateTypeExpirationDate) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchRequestDateTypeExpirationDate) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchRequestDateTypeExpirationDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchRequestDateTypeExpirationDate(val *RenewalsSearchRequestDateTypeExpirationDate) *NullableRenewalsSearchRequestDateTypeExpirationDate {
	return &NullableRenewalsSearchRequestDateTypeExpirationDate{value: val, isSet: true}
}

func (v NullableRenewalsSearchRequestDateTypeExpirationDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchRequestDateTypeExpirationDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


