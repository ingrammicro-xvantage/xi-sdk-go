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

// checks if the OrderDetailB2BLinesInnerEstimatedDatesInnerShip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerEstimatedDatesInnerShip{}

// OrderDetailB2BLinesInnerEstimatedDatesInnerShip struct for OrderDetailB2BLinesInnerEstimatedDatesInnerShip
type OrderDetailB2BLinesInnerEstimatedDatesInnerShip struct {
	// Date type. Example Single or multiple dates.
	ShipDateType *string `json:"shipDateType,omitempty"`
	ShipDateRange *OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange `json:"shipDateRange,omitempty"`
	// Source of the shipment.
	ShipSource *string `json:"shipSource,omitempty"`
	// Shipment description.
	ShipDescription *string `json:"shipDescription,omitempty"`
	// Ship date.
	ShipDate *string `json:"shipDate,omitempty"`
}

// NewOrderDetailB2BLinesInnerEstimatedDatesInnerShip instantiates a new OrderDetailB2BLinesInnerEstimatedDatesInnerShip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerEstimatedDatesInnerShip() *OrderDetailB2BLinesInnerEstimatedDatesInnerShip {
	this := OrderDetailB2BLinesInnerEstimatedDatesInnerShip{}
	return &this
}

// NewOrderDetailB2BLinesInnerEstimatedDatesInnerShipWithDefaults instantiates a new OrderDetailB2BLinesInnerEstimatedDatesInnerShip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerEstimatedDatesInnerShipWithDefaults() *OrderDetailB2BLinesInnerEstimatedDatesInnerShip {
	this := OrderDetailB2BLinesInnerEstimatedDatesInnerShip{}
	return &this
}

// GetShipDateType returns the ShipDateType field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDateType() string {
	if o == nil || IsNil(o.ShipDateType) {
		var ret string
		return ret
	}
	return *o.ShipDateType
}

// GetShipDateTypeOk returns a tuple with the ShipDateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ShipDateType) {
		return nil, false
	}
	return o.ShipDateType, true
}

// HasShipDateType returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) HasShipDateType() bool {
	if o != nil && !IsNil(o.ShipDateType) {
		return true
	}

	return false
}

// SetShipDateType gets a reference to the given string and assigns it to the ShipDateType field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) SetShipDateType(v string) {
	o.ShipDateType = &v
}

// GetShipDateRange returns the ShipDateRange field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDateRange() OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange {
	if o == nil || IsNil(o.ShipDateRange) {
		var ret OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange
		return ret
	}
	return *o.ShipDateRange
}

// GetShipDateRangeOk returns a tuple with the ShipDateRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDateRangeOk() (*OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange, bool) {
	if o == nil || IsNil(o.ShipDateRange) {
		return nil, false
	}
	return o.ShipDateRange, true
}

// HasShipDateRange returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) HasShipDateRange() bool {
	if o != nil && !IsNil(o.ShipDateRange) {
		return true
	}

	return false
}

// SetShipDateRange gets a reference to the given OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange and assigns it to the ShipDateRange field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) SetShipDateRange(v OrderDetailB2BLinesInnerEstimatedDatesInnerShipShipDateRange) {
	o.ShipDateRange = &v
}

// GetShipSource returns the ShipSource field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipSource() string {
	if o == nil || IsNil(o.ShipSource) {
		var ret string
		return ret
	}
	return *o.ShipSource
}

// GetShipSourceOk returns a tuple with the ShipSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ShipSource) {
		return nil, false
	}
	return o.ShipSource, true
}

// HasShipSource returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) HasShipSource() bool {
	if o != nil && !IsNil(o.ShipSource) {
		return true
	}

	return false
}

// SetShipSource gets a reference to the given string and assigns it to the ShipSource field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) SetShipSource(v string) {
	o.ShipSource = &v
}

// GetShipDescription returns the ShipDescription field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDescription() string {
	if o == nil || IsNil(o.ShipDescription) {
		var ret string
		return ret
	}
	return *o.ShipDescription
}

// GetShipDescriptionOk returns a tuple with the ShipDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShipDescription) {
		return nil, false
	}
	return o.ShipDescription, true
}

// HasShipDescription returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) HasShipDescription() bool {
	if o != nil && !IsNil(o.ShipDescription) {
		return true
	}

	return false
}

// SetShipDescription gets a reference to the given string and assigns it to the ShipDescription field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) SetShipDescription(v string) {
	o.ShipDescription = &v
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDate() string {
	if o == nil || IsNil(o.ShipDate) {
		var ret string
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) GetShipDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShipDate) {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) HasShipDate() bool {
	if o != nil && !IsNil(o.ShipDate) {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given string and assigns it to the ShipDate field.
func (o *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) SetShipDate(v string) {
	o.ShipDate = &v
}

func (o OrderDetailB2BLinesInnerEstimatedDatesInnerShip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerEstimatedDatesInnerShip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipDateType) {
		toSerialize["shipDateType"] = o.ShipDateType
	}
	if !IsNil(o.ShipDateRange) {
		toSerialize["shipDateRange"] = o.ShipDateRange
	}
	if !IsNil(o.ShipSource) {
		toSerialize["shipSource"] = o.ShipSource
	}
	if !IsNil(o.ShipDescription) {
		toSerialize["shipDescription"] = o.ShipDescription
	}
	if !IsNil(o.ShipDate) {
		toSerialize["shipDate"] = o.ShipDate
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip struct {
	value *OrderDetailB2BLinesInnerEstimatedDatesInnerShip
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip) Get() *OrderDetailB2BLinesInnerEstimatedDatesInnerShip {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip) Set(val *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip(val *OrderDetailB2BLinesInnerEstimatedDatesInnerShip) *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip {
	return &NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerEstimatedDatesInnerShip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


