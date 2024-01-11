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

// checks if the OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner{}

// OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner struct for OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner
type OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner struct {
	Commenttext1 *string `json:"commenttext1,omitempty"`
	Commenttext2 *string `json:"commenttext2,omitempty"`
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner() *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner {
	this := OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner{}
	return &this
}

// NewOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInnerWithDefaults instantiates a new OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInnerWithDefaults() *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner {
	this := OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner{}
	return &this
}

// GetCommenttext1 returns the Commenttext1 field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) GetCommenttext1() string {
	if o == nil || IsNil(o.Commenttext1) {
		var ret string
		return ret
	}
	return *o.Commenttext1
}

// GetCommenttext1Ok returns a tuple with the Commenttext1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) GetCommenttext1Ok() (*string, bool) {
	if o == nil || IsNil(o.Commenttext1) {
		return nil, false
	}
	return o.Commenttext1, true
}

// HasCommenttext1 returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) HasCommenttext1() bool {
	if o != nil && !IsNil(o.Commenttext1) {
		return true
	}

	return false
}

// SetCommenttext1 gets a reference to the given string and assigns it to the Commenttext1 field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) SetCommenttext1(v string) {
	o.Commenttext1 = &v
}

// GetCommenttext2 returns the Commenttext2 field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) GetCommenttext2() string {
	if o == nil || IsNil(o.Commenttext2) {
		var ret string
		return ret
	}
	return *o.Commenttext2
}

// GetCommenttext2Ok returns a tuple with the Commenttext2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) GetCommenttext2Ok() (*string, bool) {
	if o == nil || IsNil(o.Commenttext2) {
		return nil, false
	}
	return o.Commenttext2, true
}

// HasCommenttext2 returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) HasCommenttext2() bool {
	if o != nil && !IsNil(o.Commenttext2) {
		return true
	}

	return false
}

// SetCommenttext2 gets a reference to the given string and assigns it to the Commenttext2 field.
func (o *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) SetCommenttext2(v string) {
	o.Commenttext2 = &v
}

func (o OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commenttext1) {
		toSerialize["commenttext1"] = o.Commenttext1
	}
	if !IsNil(o.Commenttext2) {
		toSerialize["commenttext2"] = o.Commenttext2
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner struct {
	value *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner
	isSet bool
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) Get() *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner {
	return v.value
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) Set(val *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner(val *OrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) *NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner {
	return &NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner{value: val, isSet: true}
}

func (v NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseServiceresponseOrderdetailresponseCommentlinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


