/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderCreateRequestOrdercreaterequestRequestpreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestOrdercreaterequestRequestpreamble{}

// OrderCreateRequestOrdercreaterequestRequestpreamble struct for OrderCreateRequestOrdercreaterequestRequestpreamble
type OrderCreateRequestOrdercreaterequestRequestpreamble struct {
	// 2 digit ISO country code
	Isocountrycode string `json:"isocountrycode"`
	// Your unique Ingram Micro customer number
	Customernumber string `json:"customernumber"`
}

type _OrderCreateRequestOrdercreaterequestRequestpreamble OrderCreateRequestOrdercreaterequestRequestpreamble

// NewOrderCreateRequestOrdercreaterequestRequestpreamble instantiates a new OrderCreateRequestOrdercreaterequestRequestpreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestOrdercreaterequestRequestpreamble(isocountrycode string, customernumber string) *OrderCreateRequestOrdercreaterequestRequestpreamble {
	this := OrderCreateRequestOrdercreaterequestRequestpreamble{}
	this.Isocountrycode = isocountrycode
	this.Customernumber = customernumber
	return &this
}

// NewOrderCreateRequestOrdercreaterequestRequestpreambleWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestRequestpreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestOrdercreaterequestRequestpreambleWithDefaults() *OrderCreateRequestOrdercreaterequestRequestpreamble {
	this := OrderCreateRequestOrdercreaterequestRequestpreamble{}
	return &this
}

// GetIsocountrycode returns the Isocountrycode field value
func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) GetIsocountrycode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isocountrycode
}

// GetIsocountrycodeOk returns a tuple with the Isocountrycode field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) GetIsocountrycodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isocountrycode, true
}

// SetIsocountrycode sets field value
func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) SetIsocountrycode(v string) {
	o.Isocountrycode = v
}

// GetCustomernumber returns the Customernumber field value
func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) GetCustomernumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Customernumber
}

// GetCustomernumberOk returns a tuple with the Customernumber field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) GetCustomernumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customernumber, true
}

// SetCustomernumber sets field value
func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) SetCustomernumber(v string) {
	o.Customernumber = v
}

func (o OrderCreateRequestOrdercreaterequestRequestpreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestOrdercreaterequestRequestpreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isocountrycode"] = o.Isocountrycode
	toSerialize["customernumber"] = o.Customernumber
	return toSerialize, nil
}

func (o *OrderCreateRequestOrdercreaterequestRequestpreamble) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isocountrycode",
		"customernumber",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderCreateRequestOrdercreaterequestRequestpreamble := _OrderCreateRequestOrdercreaterequestRequestpreamble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderCreateRequestOrdercreaterequestRequestpreamble)

	if err != nil {
		return err
	}

	*o = OrderCreateRequestOrdercreaterequestRequestpreamble(varOrderCreateRequestOrdercreaterequestRequestpreamble)

	return err
}

type NullableOrderCreateRequestOrdercreaterequestRequestpreamble struct {
	value *OrderCreateRequestOrdercreaterequestRequestpreamble
	isSet bool
}

func (v NullableOrderCreateRequestOrdercreaterequestRequestpreamble) Get() *OrderCreateRequestOrdercreaterequestRequestpreamble {
	return v.value
}

func (v *NullableOrderCreateRequestOrdercreaterequestRequestpreamble) Set(val *OrderCreateRequestOrdercreaterequestRequestpreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestOrdercreaterequestRequestpreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestOrdercreaterequestRequestpreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestOrdercreaterequestRequestpreamble(val *OrderCreateRequestOrdercreaterequestRequestpreamble) *NullableOrderCreateRequestOrdercreaterequestRequestpreamble {
	return &NullableOrderCreateRequestOrdercreaterequestRequestpreamble{value: val, isSet: true}
}

func (v NullableOrderCreateRequestOrdercreaterequestRequestpreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestOrdercreaterequestRequestpreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


