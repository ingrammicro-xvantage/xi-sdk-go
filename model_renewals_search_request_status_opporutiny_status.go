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

// checks if the RenewalsSearchRequestStatusOpporutinyStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchRequestStatusOpporutinyStatus{}

// RenewalsSearchRequestStatusOpporutinyStatus struct for RenewalsSearchRequestStatusOpporutinyStatus
type RenewalsSearchRequestStatusOpporutinyStatus struct {
	// The value of opportunity status, it can be either Open or Closed.
	Value *string `json:"value,omitempty"`
	// The sub-status of Opportunity status. Possible sub-status values for Open opportunity status are Ready to order, Quote pending, Notification sent, Future, and Quote requested. Possible sub-status values for Closed opportunity status are All, Ordered, Quote closed-contract sales desk, Expired, Transition to new/upsell, Lost to competitior, Consolidated, Transitioned to cloud, Renewal went direct, EOL, End user out of business, Void, Other, and Cancelled.
	SubStatus *string `json:"subStatus,omitempty"`
}

// NewRenewalsSearchRequestStatusOpporutinyStatus instantiates a new RenewalsSearchRequestStatusOpporutinyStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchRequestStatusOpporutinyStatus() *RenewalsSearchRequestStatusOpporutinyStatus {
	this := RenewalsSearchRequestStatusOpporutinyStatus{}
	return &this
}

// NewRenewalsSearchRequestStatusOpporutinyStatusWithDefaults instantiates a new RenewalsSearchRequestStatusOpporutinyStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchRequestStatusOpporutinyStatusWithDefaults() *RenewalsSearchRequestStatusOpporutinyStatus {
	this := RenewalsSearchRequestStatusOpporutinyStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) SetValue(v string) {
	o.Value = &v
}

// GetSubStatus returns the SubStatus field value if set, zero value otherwise.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetSubStatus() string {
	if o == nil || IsNil(o.SubStatus) {
		var ret string
		return ret
	}
	return *o.SubStatus
}

// GetSubStatusOk returns a tuple with the SubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) GetSubStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SubStatus) {
		return nil, false
	}
	return o.SubStatus, true
}

// HasSubStatus returns a boolean if a field has been set.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) HasSubStatus() bool {
	if o != nil && !IsNil(o.SubStatus) {
		return true
	}

	return false
}

// SetSubStatus gets a reference to the given string and assigns it to the SubStatus field.
func (o *RenewalsSearchRequestStatusOpporutinyStatus) SetSubStatus(v string) {
	o.SubStatus = &v
}

func (o RenewalsSearchRequestStatusOpporutinyStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchRequestStatusOpporutinyStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.SubStatus) {
		toSerialize["subStatus"] = o.SubStatus
	}
	return toSerialize, nil
}

type NullableRenewalsSearchRequestStatusOpporutinyStatus struct {
	value *RenewalsSearchRequestStatusOpporutinyStatus
	isSet bool
}

func (v NullableRenewalsSearchRequestStatusOpporutinyStatus) Get() *RenewalsSearchRequestStatusOpporutinyStatus {
	return v.value
}

func (v *NullableRenewalsSearchRequestStatusOpporutinyStatus) Set(val *RenewalsSearchRequestStatusOpporutinyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchRequestStatusOpporutinyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchRequestStatusOpporutinyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchRequestStatusOpporutinyStatus(val *RenewalsSearchRequestStatusOpporutinyStatus) *NullableRenewalsSearchRequestStatusOpporutinyStatus {
	return &NullableRenewalsSearchRequestStatusOpporutinyStatus{value: val, isSet: true}
}

func (v NullableRenewalsSearchRequestStatusOpporutinyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchRequestStatusOpporutinyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


