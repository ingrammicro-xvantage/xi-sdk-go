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

// checks if the OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner{}

// OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner struct for OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner
type OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner struct {
	// Handling charges/Miscellaneous Fee description
	Description *string `json:"description,omitempty"`
	// Handling charges/ Miscelaneous fee amount
	Chargeamount *string `json:"chargeamount,omitempty"`
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner() *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner {
	this := OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner{}
	return &this
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInnerWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInnerWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner {
	this := OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) SetDescription(v string) {
	o.Description = &v
}

// GetChargeamount returns the Chargeamount field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) GetChargeamount() string {
	if o == nil || IsNil(o.Chargeamount) {
		var ret string
		return ret
	}
	return *o.Chargeamount
}

// GetChargeamountOk returns a tuple with the Chargeamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) GetChargeamountOk() (*string, bool) {
	if o == nil || IsNil(o.Chargeamount) {
		return nil, false
	}
	return o.Chargeamount, true
}

// HasChargeamount returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) HasChargeamount() bool {
	if o != nil && !IsNil(o.Chargeamount) {
		return true
	}

	return false
}

// SetChargeamount gets a reference to the given string and assigns it to the Chargeamount field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) SetChargeamount(v string) {
	o.Chargeamount = &v
}

func (o OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Chargeamount) {
		toSerialize["chargeamount"] = o.Chargeamount
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner struct {
	value *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner
	isSet bool
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) Get() *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner {
	return v.value
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) Set(val *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner(val *OrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) *NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner {
	return &NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner{value: val, isSet: true}
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseMiscfeelineInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


