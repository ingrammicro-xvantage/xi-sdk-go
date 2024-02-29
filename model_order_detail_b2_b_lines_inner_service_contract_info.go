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

// checks if the OrderDetailB2BLinesInnerServiceContractInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerServiceContractInfo{}

// OrderDetailB2BLinesInnerServiceContractInfo struct for OrderDetailB2BLinesInnerServiceContractInfo
type OrderDetailB2BLinesInnerServiceContractInfo struct {
	ContractInfo *OrderDetailB2BLinesInnerServiceContractInfoContractInfo `json:"contractInfo,omitempty"`
	Subscriptions *OrderDetailB2BLinesInnerServiceContractInfoSubscriptions `json:"subscriptions,omitempty"`
	LicenseInfo *OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo `json:"licenseInfo,omitempty"`
}

// NewOrderDetailB2BLinesInnerServiceContractInfo instantiates a new OrderDetailB2BLinesInnerServiceContractInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerServiceContractInfo() *OrderDetailB2BLinesInnerServiceContractInfo {
	this := OrderDetailB2BLinesInnerServiceContractInfo{}
	return &this
}

// NewOrderDetailB2BLinesInnerServiceContractInfoWithDefaults instantiates a new OrderDetailB2BLinesInnerServiceContractInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerServiceContractInfoWithDefaults() *OrderDetailB2BLinesInnerServiceContractInfo {
	this := OrderDetailB2BLinesInnerServiceContractInfo{}
	return &this
}

// GetContractInfo returns the ContractInfo field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) GetContractInfo() OrderDetailB2BLinesInnerServiceContractInfoContractInfo {
	if o == nil || IsNil(o.ContractInfo) {
		var ret OrderDetailB2BLinesInnerServiceContractInfoContractInfo
		return ret
	}
	return *o.ContractInfo
}

// GetContractInfoOk returns a tuple with the ContractInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) GetContractInfoOk() (*OrderDetailB2BLinesInnerServiceContractInfoContractInfo, bool) {
	if o == nil || IsNil(o.ContractInfo) {
		return nil, false
	}
	return o.ContractInfo, true
}

// HasContractInfo returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) HasContractInfo() bool {
	if o != nil && !IsNil(o.ContractInfo) {
		return true
	}

	return false
}

// SetContractInfo gets a reference to the given OrderDetailB2BLinesInnerServiceContractInfoContractInfo and assigns it to the ContractInfo field.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) SetContractInfo(v OrderDetailB2BLinesInnerServiceContractInfoContractInfo) {
	o.ContractInfo = &v
}

// GetSubscriptions returns the Subscriptions field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) GetSubscriptions() OrderDetailB2BLinesInnerServiceContractInfoSubscriptions {
	if o == nil || IsNil(o.Subscriptions) {
		var ret OrderDetailB2BLinesInnerServiceContractInfoSubscriptions
		return ret
	}
	return *o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) GetSubscriptionsOk() (*OrderDetailB2BLinesInnerServiceContractInfoSubscriptions, bool) {
	if o == nil || IsNil(o.Subscriptions) {
		return nil, false
	}
	return o.Subscriptions, true
}

// HasSubscriptions returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) HasSubscriptions() bool {
	if o != nil && !IsNil(o.Subscriptions) {
		return true
	}

	return false
}

// SetSubscriptions gets a reference to the given OrderDetailB2BLinesInnerServiceContractInfoSubscriptions and assigns it to the Subscriptions field.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) SetSubscriptions(v OrderDetailB2BLinesInnerServiceContractInfoSubscriptions) {
	o.Subscriptions = &v
}

// GetLicenseInfo returns the LicenseInfo field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) GetLicenseInfo() OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo {
	if o == nil || IsNil(o.LicenseInfo) {
		var ret OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo
		return ret
	}
	return *o.LicenseInfo
}

// GetLicenseInfoOk returns a tuple with the LicenseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) GetLicenseInfoOk() (*OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo, bool) {
	if o == nil || IsNil(o.LicenseInfo) {
		return nil, false
	}
	return o.LicenseInfo, true
}

// HasLicenseInfo returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) HasLicenseInfo() bool {
	if o != nil && !IsNil(o.LicenseInfo) {
		return true
	}

	return false
}

// SetLicenseInfo gets a reference to the given OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo and assigns it to the LicenseInfo field.
func (o *OrderDetailB2BLinesInnerServiceContractInfo) SetLicenseInfo(v OrderDetailB2BLinesInnerServiceContractInfoLicenseInfo) {
	o.LicenseInfo = &v
}

func (o OrderDetailB2BLinesInnerServiceContractInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerServiceContractInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractInfo) {
		toSerialize["contractInfo"] = o.ContractInfo
	}
	if !IsNil(o.Subscriptions) {
		toSerialize["subscriptions"] = o.Subscriptions
	}
	if !IsNil(o.LicenseInfo) {
		toSerialize["licenseInfo"] = o.LicenseInfo
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerServiceContractInfo struct {
	value *OrderDetailB2BLinesInnerServiceContractInfo
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfo) Get() *OrderDetailB2BLinesInnerServiceContractInfo {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfo) Set(val *OrderDetailB2BLinesInnerServiceContractInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerServiceContractInfo(val *OrderDetailB2BLinesInnerServiceContractInfo) *NullableOrderDetailB2BLinesInnerServiceContractInfo {
	return &NullableOrderDetailB2BLinesInnerServiceContractInfo{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


