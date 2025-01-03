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

// checks if the RenewalsSearchRequestDateTypeInvoiceDate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchRequestDateTypeInvoiceDate{}

// RenewalsSearchRequestDateTypeInvoiceDate struct for RenewalsSearchRequestDateTypeInvoiceDate
type RenewalsSearchRequestDateTypeInvoiceDate struct {
	// Custom start date for invoice date.
	CustomStartDate *string `json:"customStartDate,omitempty"`
	// Custom end date for invoice date.
	CustomEndDate *string `json:"customEndDate,omitempty"`
}

// NewRenewalsSearchRequestDateTypeInvoiceDate instantiates a new RenewalsSearchRequestDateTypeInvoiceDate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchRequestDateTypeInvoiceDate() *RenewalsSearchRequestDateTypeInvoiceDate {
	this := RenewalsSearchRequestDateTypeInvoiceDate{}
	return &this
}

// NewRenewalsSearchRequestDateTypeInvoiceDateWithDefaults instantiates a new RenewalsSearchRequestDateTypeInvoiceDate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchRequestDateTypeInvoiceDateWithDefaults() *RenewalsSearchRequestDateTypeInvoiceDate {
	this := RenewalsSearchRequestDateTypeInvoiceDate{}
	return &this
}

// GetCustomStartDate returns the CustomStartDate field value if set, zero value otherwise.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) GetCustomStartDate() string {
	if o == nil || IsNil(o.CustomStartDate) {
		var ret string
		return ret
	}
	return *o.CustomStartDate
}

// GetCustomStartDateOk returns a tuple with the CustomStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) GetCustomStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomStartDate) {
		return nil, false
	}
	return o.CustomStartDate, true
}

// HasCustomStartDate returns a boolean if a field has been set.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) HasCustomStartDate() bool {
	if o != nil && !IsNil(o.CustomStartDate) {
		return true
	}

	return false
}

// SetCustomStartDate gets a reference to the given string and assigns it to the CustomStartDate field.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) SetCustomStartDate(v string) {
	o.CustomStartDate = &v
}

// GetCustomEndDate returns the CustomEndDate field value if set, zero value otherwise.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) GetCustomEndDate() string {
	if o == nil || IsNil(o.CustomEndDate) {
		var ret string
		return ret
	}
	return *o.CustomEndDate
}

// GetCustomEndDateOk returns a tuple with the CustomEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) GetCustomEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.CustomEndDate) {
		return nil, false
	}
	return o.CustomEndDate, true
}

// HasCustomEndDate returns a boolean if a field has been set.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) HasCustomEndDate() bool {
	if o != nil && !IsNil(o.CustomEndDate) {
		return true
	}

	return false
}

// SetCustomEndDate gets a reference to the given string and assigns it to the CustomEndDate field.
func (o *RenewalsSearchRequestDateTypeInvoiceDate) SetCustomEndDate(v string) {
	o.CustomEndDate = &v
}

func (o RenewalsSearchRequestDateTypeInvoiceDate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchRequestDateTypeInvoiceDate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomStartDate) {
		toSerialize["customStartDate"] = o.CustomStartDate
	}
	if !IsNil(o.CustomEndDate) {
		toSerialize["customEndDate"] = o.CustomEndDate
	}
	return toSerialize, nil
}

type NullableRenewalsSearchRequestDateTypeInvoiceDate struct {
	value *RenewalsSearchRequestDateTypeInvoiceDate
	isSet bool
}

func (v NullableRenewalsSearchRequestDateTypeInvoiceDate) Get() *RenewalsSearchRequestDateTypeInvoiceDate {
	return v.value
}

func (v *NullableRenewalsSearchRequestDateTypeInvoiceDate) Set(val *RenewalsSearchRequestDateTypeInvoiceDate) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchRequestDateTypeInvoiceDate) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchRequestDateTypeInvoiceDate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchRequestDateTypeInvoiceDate(val *RenewalsSearchRequestDateTypeInvoiceDate) *NullableRenewalsSearchRequestDateTypeInvoiceDate {
	return &NullableRenewalsSearchRequestDateTypeInvoiceDate{value: val, isSet: true}
}

func (v NullableRenewalsSearchRequestDateTypeInvoiceDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchRequestDateTypeInvoiceDate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


