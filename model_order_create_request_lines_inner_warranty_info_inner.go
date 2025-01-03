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

// checks if the OrderCreateRequestLinesInnerWarrantyInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestLinesInnerWarrantyInfoInner{}

// OrderCreateRequestLinesInnerWarrantyInfoInner struct for OrderCreateRequestLinesInnerWarrantyInfoInner
type OrderCreateRequestLinesInnerWarrantyInfoInner struct {
	// Unique value to link hardware and warranty lines. Should be used only when products are purchased from both Ingram and/or vendor but the warranty is purchased through Ingram for them.
	DirectLineLink *string `json:"directLineLink,omitempty"`
	// Customer line number of the hardware product in this request for linkage, either hardwareLineLink or warrantyLineLink can be used in a line.
	WarrantyLineLink *string `json:"warrantyLineLink,omitempty"`
	// Customer line number of the warranty product in this request for linkage, either hardwareLineLink or warrantyLineLink can be used in a line 
	HardwareLineLink *string `json:"hardwareLineLink,omitempty"`
	// Serial information of the hardware to be associated with the warranty, applicable on post sale orders.
	SerialInfo []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner `json:"serialInfo,omitempty"`
}

// NewOrderCreateRequestLinesInnerWarrantyInfoInner instantiates a new OrderCreateRequestLinesInnerWarrantyInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestLinesInnerWarrantyInfoInner() *OrderCreateRequestLinesInnerWarrantyInfoInner {
	this := OrderCreateRequestLinesInnerWarrantyInfoInner{}
	return &this
}

// NewOrderCreateRequestLinesInnerWarrantyInfoInnerWithDefaults instantiates a new OrderCreateRequestLinesInnerWarrantyInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestLinesInnerWarrantyInfoInnerWithDefaults() *OrderCreateRequestLinesInnerWarrantyInfoInner {
	this := OrderCreateRequestLinesInnerWarrantyInfoInner{}
	return &this
}

// GetDirectLineLink returns the DirectLineLink field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetDirectLineLink() string {
	if o == nil || IsNil(o.DirectLineLink) {
		var ret string
		return ret
	}
	return *o.DirectLineLink
}

// GetDirectLineLinkOk returns a tuple with the DirectLineLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetDirectLineLinkOk() (*string, bool) {
	if o == nil || IsNil(o.DirectLineLink) {
		return nil, false
	}
	return o.DirectLineLink, true
}

// HasDirectLineLink returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasDirectLineLink() bool {
	if o != nil && !IsNil(o.DirectLineLink) {
		return true
	}

	return false
}

// SetDirectLineLink gets a reference to the given string and assigns it to the DirectLineLink field.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetDirectLineLink(v string) {
	o.DirectLineLink = &v
}

// GetWarrantyLineLink returns the WarrantyLineLink field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetWarrantyLineLink() string {
	if o == nil || IsNil(o.WarrantyLineLink) {
		var ret string
		return ret
	}
	return *o.WarrantyLineLink
}

// GetWarrantyLineLinkOk returns a tuple with the WarrantyLineLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetWarrantyLineLinkOk() (*string, bool) {
	if o == nil || IsNil(o.WarrantyLineLink) {
		return nil, false
	}
	return o.WarrantyLineLink, true
}

// HasWarrantyLineLink returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasWarrantyLineLink() bool {
	if o != nil && !IsNil(o.WarrantyLineLink) {
		return true
	}

	return false
}

// SetWarrantyLineLink gets a reference to the given string and assigns it to the WarrantyLineLink field.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetWarrantyLineLink(v string) {
	o.WarrantyLineLink = &v
}

// GetHardwareLineLink returns the HardwareLineLink field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetHardwareLineLink() string {
	if o == nil || IsNil(o.HardwareLineLink) {
		var ret string
		return ret
	}
	return *o.HardwareLineLink
}

// GetHardwareLineLinkOk returns a tuple with the HardwareLineLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetHardwareLineLinkOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareLineLink) {
		return nil, false
	}
	return o.HardwareLineLink, true
}

// HasHardwareLineLink returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasHardwareLineLink() bool {
	if o != nil && !IsNil(o.HardwareLineLink) {
		return true
	}

	return false
}

// SetHardwareLineLink gets a reference to the given string and assigns it to the HardwareLineLink field.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetHardwareLineLink(v string) {
	o.HardwareLineLink = &v
}

// GetSerialInfo returns the SerialInfo field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetSerialInfo() []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner {
	if o == nil || IsNil(o.SerialInfo) {
		var ret []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner
		return ret
	}
	return o.SerialInfo
}

// GetSerialInfoOk returns a tuple with the SerialInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) GetSerialInfoOk() ([]OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner, bool) {
	if o == nil || IsNil(o.SerialInfo) {
		return nil, false
	}
	return o.SerialInfo, true
}

// HasSerialInfo returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) HasSerialInfo() bool {
	if o != nil && !IsNil(o.SerialInfo) {
		return true
	}

	return false
}

// SetSerialInfo gets a reference to the given []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner and assigns it to the SerialInfo field.
func (o *OrderCreateRequestLinesInnerWarrantyInfoInner) SetSerialInfo(v []OrderCreateRequestLinesInnerWarrantyInfoInnerSerialInfoInner) {
	o.SerialInfo = v
}

func (o OrderCreateRequestLinesInnerWarrantyInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestLinesInnerWarrantyInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectLineLink) {
		toSerialize["directLineLink"] = o.DirectLineLink
	}
	if !IsNil(o.WarrantyLineLink) {
		toSerialize["warrantyLineLink"] = o.WarrantyLineLink
	}
	if !IsNil(o.HardwareLineLink) {
		toSerialize["hardwareLineLink"] = o.HardwareLineLink
	}
	if !IsNil(o.SerialInfo) {
		toSerialize["serialInfo"] = o.SerialInfo
	}
	return toSerialize, nil
}

type NullableOrderCreateRequestLinesInnerWarrantyInfoInner struct {
	value *OrderCreateRequestLinesInnerWarrantyInfoInner
	isSet bool
}

func (v NullableOrderCreateRequestLinesInnerWarrantyInfoInner) Get() *OrderCreateRequestLinesInnerWarrantyInfoInner {
	return v.value
}

func (v *NullableOrderCreateRequestLinesInnerWarrantyInfoInner) Set(val *OrderCreateRequestLinesInnerWarrantyInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestLinesInnerWarrantyInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestLinesInnerWarrantyInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestLinesInnerWarrantyInfoInner(val *OrderCreateRequestLinesInnerWarrantyInfoInner) *NullableOrderCreateRequestLinesInnerWarrantyInfoInner {
	return &NullableOrderCreateRequestLinesInnerWarrantyInfoInner{value: val, isSet: true}
}

func (v NullableOrderCreateRequestLinesInnerWarrantyInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestLinesInnerWarrantyInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


