/*
XI SDK Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MultiSKUPriceAndStockRequestServicerequestRequestpreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiSKUPriceAndStockRequestServicerequestRequestpreamble{}

// MultiSKUPriceAndStockRequestServicerequestRequestpreamble struct for MultiSKUPriceAndStockRequestServicerequestRequestpreamble
type MultiSKUPriceAndStockRequestServicerequestRequestpreamble struct {
	// 2 Digit code “US”-United States “CA”-Canada
	Isocountrycode string `json:"isocountrycode"`
	// Ingram Micro customer number 10-12389
	Customernumber string `json:"customernumber"`
}

type _MultiSKUPriceAndStockRequestServicerequestRequestpreamble MultiSKUPriceAndStockRequestServicerequestRequestpreamble

// NewMultiSKUPriceAndStockRequestServicerequestRequestpreamble instantiates a new MultiSKUPriceAndStockRequestServicerequestRequestpreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSKUPriceAndStockRequestServicerequestRequestpreamble(isocountrycode string, customernumber string) *MultiSKUPriceAndStockRequestServicerequestRequestpreamble {
	this := MultiSKUPriceAndStockRequestServicerequestRequestpreamble{}
	this.Isocountrycode = isocountrycode
	this.Customernumber = customernumber
	return &this
}

// NewMultiSKUPriceAndStockRequestServicerequestRequestpreambleWithDefaults instantiates a new MultiSKUPriceAndStockRequestServicerequestRequestpreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSKUPriceAndStockRequestServicerequestRequestpreambleWithDefaults() *MultiSKUPriceAndStockRequestServicerequestRequestpreamble {
	this := MultiSKUPriceAndStockRequestServicerequestRequestpreamble{}
	return &this
}

// GetIsocountrycode returns the Isocountrycode field value
func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) GetIsocountrycode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isocountrycode
}

// GetIsocountrycodeOk returns a tuple with the Isocountrycode field value
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) GetIsocountrycodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isocountrycode, true
}

// SetIsocountrycode sets field value
func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) SetIsocountrycode(v string) {
	o.Isocountrycode = v
}

// GetCustomernumber returns the Customernumber field value
func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) GetCustomernumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Customernumber
}

// GetCustomernumberOk returns a tuple with the Customernumber field value
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) GetCustomernumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customernumber, true
}

// SetCustomernumber sets field value
func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) SetCustomernumber(v string) {
	o.Customernumber = v
}

func (o MultiSKUPriceAndStockRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiSKUPriceAndStockRequestServicerequestRequestpreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isocountrycode"] = o.Isocountrycode
	toSerialize["customernumber"] = o.Customernumber
	return toSerialize, nil
}

func (o *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) UnmarshalJSON(data []byte) (err error) {
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

	varMultiSKUPriceAndStockRequestServicerequestRequestpreamble := _MultiSKUPriceAndStockRequestServicerequestRequestpreamble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMultiSKUPriceAndStockRequestServicerequestRequestpreamble)

	if err != nil {
		return err
	}

	*o = MultiSKUPriceAndStockRequestServicerequestRequestpreamble(varMultiSKUPriceAndStockRequestServicerequestRequestpreamble)

	return err
}

type NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble struct {
	value *MultiSKUPriceAndStockRequestServicerequestRequestpreamble
	isSet bool
}

func (v NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble) Get() *MultiSKUPriceAndStockRequestServicerequestRequestpreamble {
	return v.value
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble) Set(val *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble(val *MultiSKUPriceAndStockRequestServicerequestRequestpreamble) *NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble {
	return &NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble{value: val, isSet: true}
}

func (v NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequestRequestpreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


