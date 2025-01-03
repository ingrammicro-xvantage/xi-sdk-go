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

// checks if the ReturnsCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnsCreateResponse{}

// ReturnsCreateResponse struct for ReturnsCreateResponse
type ReturnsCreateResponse struct {
	ReturnsClaims []ReturnsCreateResponseReturnsClaimsInner `json:"returnsClaims,omitempty"`
}

// NewReturnsCreateResponse instantiates a new ReturnsCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnsCreateResponse() *ReturnsCreateResponse {
	this := ReturnsCreateResponse{}
	return &this
}

// NewReturnsCreateResponseWithDefaults instantiates a new ReturnsCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnsCreateResponseWithDefaults() *ReturnsCreateResponse {
	this := ReturnsCreateResponse{}
	return &this
}

// GetReturnsClaims returns the ReturnsClaims field value if set, zero value otherwise.
func (o *ReturnsCreateResponse) GetReturnsClaims() []ReturnsCreateResponseReturnsClaimsInner {
	if o == nil || IsNil(o.ReturnsClaims) {
		var ret []ReturnsCreateResponseReturnsClaimsInner
		return ret
	}
	return o.ReturnsClaims
}

// GetReturnsClaimsOk returns a tuple with the ReturnsClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnsCreateResponse) GetReturnsClaimsOk() ([]ReturnsCreateResponseReturnsClaimsInner, bool) {
	if o == nil || IsNil(o.ReturnsClaims) {
		return nil, false
	}
	return o.ReturnsClaims, true
}

// HasReturnsClaims returns a boolean if a field has been set.
func (o *ReturnsCreateResponse) HasReturnsClaims() bool {
	if o != nil && !IsNil(o.ReturnsClaims) {
		return true
	}

	return false
}

// SetReturnsClaims gets a reference to the given []ReturnsCreateResponseReturnsClaimsInner and assigns it to the ReturnsClaims field.
func (o *ReturnsCreateResponse) SetReturnsClaims(v []ReturnsCreateResponseReturnsClaimsInner) {
	o.ReturnsClaims = v
}

func (o ReturnsCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnsCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReturnsClaims) {
		toSerialize["returnsClaims"] = o.ReturnsClaims
	}
	return toSerialize, nil
}

type NullableReturnsCreateResponse struct {
	value *ReturnsCreateResponse
	isSet bool
}

func (v NullableReturnsCreateResponse) Get() *ReturnsCreateResponse {
	return v.value
}

func (v *NullableReturnsCreateResponse) Set(val *ReturnsCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnsCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnsCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnsCreateResponse(val *ReturnsCreateResponse) *NullableReturnsCreateResponse {
	return &NullableReturnsCreateResponse{value: val, isSet: true}
}

func (v NullableReturnsCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnsCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


