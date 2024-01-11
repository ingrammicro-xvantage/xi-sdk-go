/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
)

// checks if the OrderDetailB2BLinesInnerServiceContractInfoContractInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerServiceContractInfoContractInfo{}

// OrderDetailB2BLinesInnerServiceContractInfoContractInfo struct for OrderDetailB2BLinesInnerServiceContractInfoContractInfo
type OrderDetailB2BLinesInnerServiceContractInfoContractInfo struct {
	// The description of the contract.
	ContractDescription *string `json:"contractDescription,omitempty"`
	// Contract number.
	ContractNumber *string `json:"contractNumber,omitempty"`
	// The status of the contract.
	ContractStatus *string `json:"contractStatus,omitempty"`
	// Start date of the contract.
	ContractStartDate *string `json:"contractStartDate,omitempty"`
	// End date of the contract.
	ContractEndDate *string `json:"contractEndDate,omitempty"`
	// The duration of the contract.
	ContractDuration *string `json:"contractDuration,omitempty"`
}

// NewOrderDetailB2BLinesInnerServiceContractInfoContractInfo instantiates a new OrderDetailB2BLinesInnerServiceContractInfoContractInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerServiceContractInfoContractInfo() *OrderDetailB2BLinesInnerServiceContractInfoContractInfo {
	this := OrderDetailB2BLinesInnerServiceContractInfoContractInfo{}
	return &this
}

// NewOrderDetailB2BLinesInnerServiceContractInfoContractInfoWithDefaults instantiates a new OrderDetailB2BLinesInnerServiceContractInfoContractInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerServiceContractInfoContractInfoWithDefaults() *OrderDetailB2BLinesInnerServiceContractInfoContractInfo {
	this := OrderDetailB2BLinesInnerServiceContractInfoContractInfo{}
	return &this
}

// GetContractDescription returns the ContractDescription field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractDescription() string {
	if o == nil || IsNil(o.ContractDescription) {
		var ret string
		return ret
	}
	return *o.ContractDescription
}

// GetContractDescriptionOk returns a tuple with the ContractDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ContractDescription) {
		return nil, false
	}
	return o.ContractDescription, true
}

// HasContractDescription returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) HasContractDescription() bool {
	if o != nil && !IsNil(o.ContractDescription) {
		return true
	}

	return false
}

// SetContractDescription gets a reference to the given string and assigns it to the ContractDescription field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) SetContractDescription(v string) {
	o.ContractDescription = &v
}

// GetContractNumber returns the ContractNumber field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractNumber() string {
	if o == nil || IsNil(o.ContractNumber) {
		var ret string
		return ret
	}
	return *o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ContractNumber) {
		return nil, false
	}
	return o.ContractNumber, true
}

// HasContractNumber returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) HasContractNumber() bool {
	if o != nil && !IsNil(o.ContractNumber) {
		return true
	}

	return false
}

// SetContractNumber gets a reference to the given string and assigns it to the ContractNumber field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) SetContractNumber(v string) {
	o.ContractNumber = &v
}

// GetContractStatus returns the ContractStatus field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractStatus() string {
	if o == nil || IsNil(o.ContractStatus) {
		var ret string
		return ret
	}
	return *o.ContractStatus
}

// GetContractStatusOk returns a tuple with the ContractStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ContractStatus) {
		return nil, false
	}
	return o.ContractStatus, true
}

// HasContractStatus returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) HasContractStatus() bool {
	if o != nil && !IsNil(o.ContractStatus) {
		return true
	}

	return false
}

// SetContractStatus gets a reference to the given string and assigns it to the ContractStatus field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) SetContractStatus(v string) {
	o.ContractStatus = &v
}

// GetContractStartDate returns the ContractStartDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractStartDate() string {
	if o == nil || IsNil(o.ContractStartDate) {
		var ret string
		return ret
	}
	return *o.ContractStartDate
}

// GetContractStartDateOk returns a tuple with the ContractStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.ContractStartDate) {
		return nil, false
	}
	return o.ContractStartDate, true
}

// HasContractStartDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) HasContractStartDate() bool {
	if o != nil && !IsNil(o.ContractStartDate) {
		return true
	}

	return false
}

// SetContractStartDate gets a reference to the given string and assigns it to the ContractStartDate field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) SetContractStartDate(v string) {
	o.ContractStartDate = &v
}

// GetContractEndDate returns the ContractEndDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractEndDate() string {
	if o == nil || IsNil(o.ContractEndDate) {
		var ret string
		return ret
	}
	return *o.ContractEndDate
}

// GetContractEndDateOk returns a tuple with the ContractEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.ContractEndDate) {
		return nil, false
	}
	return o.ContractEndDate, true
}

// HasContractEndDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) HasContractEndDate() bool {
	if o != nil && !IsNil(o.ContractEndDate) {
		return true
	}

	return false
}

// SetContractEndDate gets a reference to the given string and assigns it to the ContractEndDate field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) SetContractEndDate(v string) {
	o.ContractEndDate = &v
}

// GetContractDuration returns the ContractDuration field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractDuration() string {
	if o == nil || IsNil(o.ContractDuration) {
		var ret string
		return ret
	}
	return *o.ContractDuration
}

// GetContractDurationOk returns a tuple with the ContractDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) GetContractDurationOk() (*string, bool) {
	if o == nil || IsNil(o.ContractDuration) {
		return nil, false
	}
	return o.ContractDuration, true
}

// HasContractDuration returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) HasContractDuration() bool {
	if o != nil && !IsNil(o.ContractDuration) {
		return true
	}

	return false
}

// SetContractDuration gets a reference to the given string and assigns it to the ContractDuration field.
func (o *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) SetContractDuration(v string) {
	o.ContractDuration = &v
}

func (o OrderDetailB2BLinesInnerServiceContractInfoContractInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerServiceContractInfoContractInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractDescription) {
		toSerialize["contractDescription"] = o.ContractDescription
	}
	if !IsNil(o.ContractNumber) {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	if !IsNil(o.ContractStatus) {
		toSerialize["contractStatus"] = o.ContractStatus
	}
	if !IsNil(o.ContractStartDate) {
		toSerialize["contractStartDate"] = o.ContractStartDate
	}
	if !IsNil(o.ContractEndDate) {
		toSerialize["contractEndDate"] = o.ContractEndDate
	}
	if !IsNil(o.ContractDuration) {
		toSerialize["contractDuration"] = o.ContractDuration
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo struct {
	value *OrderDetailB2BLinesInnerServiceContractInfoContractInfo
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo) Get() *OrderDetailB2BLinesInnerServiceContractInfoContractInfo {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo) Set(val *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo(val *OrderDetailB2BLinesInnerServiceContractInfoContractInfo) *NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo {
	return &NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerServiceContractInfoContractInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


