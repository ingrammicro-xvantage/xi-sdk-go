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

// checks if the OrderModifyResponseServiceresponseResponsepreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyResponseServiceresponseResponsepreamble{}

// OrderModifyResponseServiceresponseResponsepreamble struct for OrderModifyResponseServiceresponseResponsepreamble
type OrderModifyResponseServiceresponseResponsepreamble struct {
	Responsestatus *string `json:"responsestatus,omitempty"`
	Responsemessage *string `json:"responsemessage,omitempty"`
}

// NewOrderModifyResponseServiceresponseResponsepreamble instantiates a new OrderModifyResponseServiceresponseResponsepreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyResponseServiceresponseResponsepreamble() *OrderModifyResponseServiceresponseResponsepreamble {
	this := OrderModifyResponseServiceresponseResponsepreamble{}
	return &this
}

// NewOrderModifyResponseServiceresponseResponsepreambleWithDefaults instantiates a new OrderModifyResponseServiceresponseResponsepreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyResponseServiceresponseResponsepreambleWithDefaults() *OrderModifyResponseServiceresponseResponsepreamble {
	this := OrderModifyResponseServiceresponseResponsepreamble{}
	return &this
}

// GetResponsestatus returns the Responsestatus field value if set, zero value otherwise.
func (o *OrderModifyResponseServiceresponseResponsepreamble) GetResponsestatus() string {
	if o == nil || IsNil(o.Responsestatus) {
		var ret string
		return ret
	}
	return *o.Responsestatus
}

// GetResponsestatusOk returns a tuple with the Responsestatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseServiceresponseResponsepreamble) GetResponsestatusOk() (*string, bool) {
	if o == nil || IsNil(o.Responsestatus) {
		return nil, false
	}
	return o.Responsestatus, true
}

// HasResponsestatus returns a boolean if a field has been set.
func (o *OrderModifyResponseServiceresponseResponsepreamble) HasResponsestatus() bool {
	if o != nil && !IsNil(o.Responsestatus) {
		return true
	}

	return false
}

// SetResponsestatus gets a reference to the given string and assigns it to the Responsestatus field.
func (o *OrderModifyResponseServiceresponseResponsepreamble) SetResponsestatus(v string) {
	o.Responsestatus = &v
}

// GetResponsemessage returns the Responsemessage field value if set, zero value otherwise.
func (o *OrderModifyResponseServiceresponseResponsepreamble) GetResponsemessage() string {
	if o == nil || IsNil(o.Responsemessage) {
		var ret string
		return ret
	}
	return *o.Responsemessage
}

// GetResponsemessageOk returns a tuple with the Responsemessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseServiceresponseResponsepreamble) GetResponsemessageOk() (*string, bool) {
	if o == nil || IsNil(o.Responsemessage) {
		return nil, false
	}
	return o.Responsemessage, true
}

// HasResponsemessage returns a boolean if a field has been set.
func (o *OrderModifyResponseServiceresponseResponsepreamble) HasResponsemessage() bool {
	if o != nil && !IsNil(o.Responsemessage) {
		return true
	}

	return false
}

// SetResponsemessage gets a reference to the given string and assigns it to the Responsemessage field.
func (o *OrderModifyResponseServiceresponseResponsepreamble) SetResponsemessage(v string) {
	o.Responsemessage = &v
}

func (o OrderModifyResponseServiceresponseResponsepreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyResponseServiceresponseResponsepreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsestatus) {
		toSerialize["responsestatus"] = o.Responsestatus
	}
	if !IsNil(o.Responsemessage) {
		toSerialize["responsemessage"] = o.Responsemessage
	}
	return toSerialize, nil
}

type NullableOrderModifyResponseServiceresponseResponsepreamble struct {
	value *OrderModifyResponseServiceresponseResponsepreamble
	isSet bool
}

func (v NullableOrderModifyResponseServiceresponseResponsepreamble) Get() *OrderModifyResponseServiceresponseResponsepreamble {
	return v.value
}

func (v *NullableOrderModifyResponseServiceresponseResponsepreamble) Set(val *OrderModifyResponseServiceresponseResponsepreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyResponseServiceresponseResponsepreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyResponseServiceresponseResponsepreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyResponseServiceresponseResponsepreamble(val *OrderModifyResponseServiceresponseResponsepreamble) *NullableOrderModifyResponseServiceresponseResponsepreamble {
	return &NullableOrderModifyResponseServiceresponseResponsepreamble{value: val, isSet: true}
}

func (v NullableOrderModifyResponseServiceresponseResponsepreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyResponseServiceresponseResponsepreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


