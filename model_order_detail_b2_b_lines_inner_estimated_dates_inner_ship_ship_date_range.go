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

// checks if the OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange{}

// OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange struct for OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange
type OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange struct {
	// Start Date.
	StartDate *string `json:"startDate,omitempty"`
	// End Date.
	EndDate *string `json:"endDate,omitempty"`
}

// NewOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange instantiates a new OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange() *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange {
	this := OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange{}
	return &this
}

// NewOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRangeWithDefaults instantiates a new OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRangeWithDefaults() *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange {
	this := OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) SetEndDate(v string) {
	o.EndDate = &v
}

func (o OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange struct {
	value *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) Get() *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) Set(val *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange(val *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange {
	return &NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


