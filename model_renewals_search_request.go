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

// checks if the RenewalsSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchRequest{}

// RenewalsSearchRequest struct for RenewalsSearchRequest
type RenewalsSearchRequest struct {
	Status *RenewalsSearchRequestStatus `json:"status,omitempty"`
	DataType *RenewalsSearchRequestDataType `json:"dataType,omitempty"`
	// The name of the Vendor.
	Vendor *string `json:"vendor,omitempty"`
	// The name of the enduser. 
	EndUser *string `json:"endUser,omitempty"`
}

// NewRenewalsSearchRequest instantiates a new RenewalsSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchRequest() *RenewalsSearchRequest {
	this := RenewalsSearchRequest{}
	return &this
}

// NewRenewalsSearchRequestWithDefaults instantiates a new RenewalsSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchRequestWithDefaults() *RenewalsSearchRequest {
	this := RenewalsSearchRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RenewalsSearchRequest) GetStatus() RenewalsSearchRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret RenewalsSearchRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequest) GetStatusOk() (*RenewalsSearchRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RenewalsSearchRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given RenewalsSearchRequestStatus and assigns it to the Status field.
func (o *RenewalsSearchRequest) SetStatus(v RenewalsSearchRequestStatus) {
	o.Status = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *RenewalsSearchRequest) GetDataType() RenewalsSearchRequestDataType {
	if o == nil || IsNil(o.DataType) {
		var ret RenewalsSearchRequestDataType
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequest) GetDataTypeOk() (*RenewalsSearchRequestDataType, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *RenewalsSearchRequest) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given RenewalsSearchRequestDataType and assigns it to the DataType field.
func (o *RenewalsSearchRequest) SetDataType(v RenewalsSearchRequestDataType) {
	o.DataType = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *RenewalsSearchRequest) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequest) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *RenewalsSearchRequest) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *RenewalsSearchRequest) SetVendor(v string) {
	o.Vendor = &v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *RenewalsSearchRequest) GetEndUser() string {
	if o == nil || IsNil(o.EndUser) {
		var ret string
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchRequest) GetEndUserOk() (*string, bool) {
	if o == nil || IsNil(o.EndUser) {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *RenewalsSearchRequest) HasEndUser() bool {
	if o != nil && !IsNil(o.EndUser) {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given string and assigns it to the EndUser field.
func (o *RenewalsSearchRequest) SetEndUser(v string) {
	o.EndUser = &v
}

func (o RenewalsSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.EndUser) {
		toSerialize["endUser"] = o.EndUser
	}
	return toSerialize, nil
}

type NullableRenewalsSearchRequest struct {
	value *RenewalsSearchRequest
	isSet bool
}

func (v NullableRenewalsSearchRequest) Get() *RenewalsSearchRequest {
	return v.value
}

func (v *NullableRenewalsSearchRequest) Set(val *RenewalsSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchRequest(val *RenewalsSearchRequest) *NullableRenewalsSearchRequest {
	return &NullableRenewalsSearchRequest{value: val, isSet: true}
}

func (v NullableRenewalsSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

