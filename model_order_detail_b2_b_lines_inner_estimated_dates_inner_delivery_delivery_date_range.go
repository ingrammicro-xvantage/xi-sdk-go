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

// checks if the OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange{}

// OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange Delivery date range.
type OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange struct {
	// Start Date.
	StartDate *string `json:"startDate,omitempty"`
	// End Date.
	EndDate *string `json:"endDate,omitempty"`
}

// NewOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange instantiates a new OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange() *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange {
	this := OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange{}
	return &this
}

// NewOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRangeWithDefaults instantiates a new OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRangeWithDefaults() *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange {
	this := OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) SetEndDate(v string) {
	o.EndDate = &v
}

func (o OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange struct {
	value *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) Get() *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) Set(val *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange(val *OrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange {
	return &NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerDeliveryDeliveryDateRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


