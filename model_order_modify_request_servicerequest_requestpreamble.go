/*
XI SDK Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderModifyRequestServicerequestRequestpreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyRequestServicerequestRequestpreamble{}

// OrderModifyRequestServicerequestRequestpreamble struct for OrderModifyRequestServicerequestRequestpreamble
type OrderModifyRequestServicerequestRequestpreamble struct {
	Isocountrycode *string `json:"isocountrycode,omitempty"`
	Customernumber *string `json:"customernumber,omitempty"`
}

// NewOrderModifyRequestServicerequestRequestpreamble instantiates a new OrderModifyRequestServicerequestRequestpreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyRequestServicerequestRequestpreamble() *OrderModifyRequestServicerequestRequestpreamble {
	this := OrderModifyRequestServicerequestRequestpreamble{}
	return &this
}

// NewOrderModifyRequestServicerequestRequestpreambleWithDefaults instantiates a new OrderModifyRequestServicerequestRequestpreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyRequestServicerequestRequestpreambleWithDefaults() *OrderModifyRequestServicerequestRequestpreamble {
	this := OrderModifyRequestServicerequestRequestpreamble{}
	return &this
}

// GetIsocountrycode returns the Isocountrycode field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestRequestpreamble) GetIsocountrycode() string {
	if o == nil || IsNil(o.Isocountrycode) {
		var ret string
		return ret
	}
	return *o.Isocountrycode
}

// GetIsocountrycodeOk returns a tuple with the Isocountrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestRequestpreamble) GetIsocountrycodeOk() (*string, bool) {
	if o == nil || IsNil(o.Isocountrycode) {
		return nil, false
	}
	return o.Isocountrycode, true
}

// HasIsocountrycode returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestRequestpreamble) HasIsocountrycode() bool {
	if o != nil && !IsNil(o.Isocountrycode) {
		return true
	}

	return false
}

// SetIsocountrycode gets a reference to the given string and assigns it to the Isocountrycode field.
func (o *OrderModifyRequestServicerequestRequestpreamble) SetIsocountrycode(v string) {
	o.Isocountrycode = &v
}

// GetCustomernumber returns the Customernumber field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestRequestpreamble) GetCustomernumber() string {
	if o == nil || IsNil(o.Customernumber) {
		var ret string
		return ret
	}
	return *o.Customernumber
}

// GetCustomernumberOk returns a tuple with the Customernumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestRequestpreamble) GetCustomernumberOk() (*string, bool) {
	if o == nil || IsNil(o.Customernumber) {
		return nil, false
	}
	return o.Customernumber, true
}

// HasCustomernumber returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestRequestpreamble) HasCustomernumber() bool {
	if o != nil && !IsNil(o.Customernumber) {
		return true
	}

	return false
}

// SetCustomernumber gets a reference to the given string and assigns it to the Customernumber field.
func (o *OrderModifyRequestServicerequestRequestpreamble) SetCustomernumber(v string) {
	o.Customernumber = &v
}

func (o OrderModifyRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyRequestServicerequestRequestpreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Isocountrycode) {
		toSerialize["isocountrycode"] = o.Isocountrycode
	}
	if !IsNil(o.Customernumber) {
		toSerialize["customernumber"] = o.Customernumber
	}
	return toSerialize, nil
}

type NullableOrderModifyRequestServicerequestRequestpreamble struct {
	value *OrderModifyRequestServicerequestRequestpreamble
	isSet bool
}

func (v NullableOrderModifyRequestServicerequestRequestpreamble) Get() *OrderModifyRequestServicerequestRequestpreamble {
	return v.value
}

func (v *NullableOrderModifyRequestServicerequestRequestpreamble) Set(val *OrderModifyRequestServicerequestRequestpreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyRequestServicerequestRequestpreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyRequestServicerequestRequestpreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyRequestServicerequestRequestpreamble(val *OrderModifyRequestServicerequestRequestpreamble) *NullableOrderModifyRequestServicerequestRequestpreamble {
	return &NullableOrderModifyRequestServicerequestRequestpreamble{value: val, isSet: true}
}

func (v NullableOrderModifyRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyRequestServicerequestRequestpreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


