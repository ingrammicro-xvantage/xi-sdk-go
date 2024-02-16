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

// checks if the ProductSearchRequestServicerequestRequestpreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductSearchRequestServicerequestRequestpreamble{}

// ProductSearchRequestServicerequestRequestpreamble struct for ProductSearchRequestServicerequestRequestpreamble
type ProductSearchRequestServicerequestRequestpreamble struct {
	Isocountrycode string `json:"isocountrycode"`
	Customernumber string `json:"customernumber"`
	Vendornumber *string `json:"vendornumber,omitempty"`
}

type _ProductSearchRequestServicerequestRequestpreamble ProductSearchRequestServicerequestRequestpreamble

// NewProductSearchRequestServicerequestRequestpreamble instantiates a new ProductSearchRequestServicerequestRequestpreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductSearchRequestServicerequestRequestpreamble(isocountrycode string, customernumber string) *ProductSearchRequestServicerequestRequestpreamble {
	this := ProductSearchRequestServicerequestRequestpreamble{}
	this.Isocountrycode = isocountrycode
	this.Customernumber = customernumber
	return &this
}

// NewProductSearchRequestServicerequestRequestpreambleWithDefaults instantiates a new ProductSearchRequestServicerequestRequestpreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductSearchRequestServicerequestRequestpreambleWithDefaults() *ProductSearchRequestServicerequestRequestpreamble {
	this := ProductSearchRequestServicerequestRequestpreamble{}
	return &this
}

// GetIsocountrycode returns the Isocountrycode field value
func (o *ProductSearchRequestServicerequestRequestpreamble) GetIsocountrycode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Isocountrycode
}

// GetIsocountrycodeOk returns a tuple with the Isocountrycode field value
// and a boolean to check if the value has been set.
func (o *ProductSearchRequestServicerequestRequestpreamble) GetIsocountrycodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Isocountrycode, true
}

// SetIsocountrycode sets field value
func (o *ProductSearchRequestServicerequestRequestpreamble) SetIsocountrycode(v string) {
	o.Isocountrycode = v
}

// GetCustomernumber returns the Customernumber field value
func (o *ProductSearchRequestServicerequestRequestpreamble) GetCustomernumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Customernumber
}

// GetCustomernumberOk returns a tuple with the Customernumber field value
// and a boolean to check if the value has been set.
func (o *ProductSearchRequestServicerequestRequestpreamble) GetCustomernumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customernumber, true
}

// SetCustomernumber sets field value
func (o *ProductSearchRequestServicerequestRequestpreamble) SetCustomernumber(v string) {
	o.Customernumber = v
}

// GetVendornumber returns the Vendornumber field value if set, zero value otherwise.
func (o *ProductSearchRequestServicerequestRequestpreamble) GetVendornumber() string {
	if o == nil || IsNil(o.Vendornumber) {
		var ret string
		return ret
	}
	return *o.Vendornumber
}

// GetVendornumberOk returns a tuple with the Vendornumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductSearchRequestServicerequestRequestpreamble) GetVendornumberOk() (*string, bool) {
	if o == nil || IsNil(o.Vendornumber) {
		return nil, false
	}
	return o.Vendornumber, true
}

// HasVendornumber returns a boolean if a field has been set.
func (o *ProductSearchRequestServicerequestRequestpreamble) HasVendornumber() bool {
	if o != nil && !IsNil(o.Vendornumber) {
		return true
	}

	return false
}

// SetVendornumber gets a reference to the given string and assigns it to the Vendornumber field.
func (o *ProductSearchRequestServicerequestRequestpreamble) SetVendornumber(v string) {
	o.Vendornumber = &v
}

func (o ProductSearchRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductSearchRequestServicerequestRequestpreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isocountrycode"] = o.Isocountrycode
	toSerialize["customernumber"] = o.Customernumber
	if !IsNil(o.Vendornumber) {
		toSerialize["vendornumber"] = o.Vendornumber
	}
	return toSerialize, nil
}

func (o *ProductSearchRequestServicerequestRequestpreamble) UnmarshalJSON(data []byte) (err error) {
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

	varProductSearchRequestServicerequestRequestpreamble := _ProductSearchRequestServicerequestRequestpreamble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductSearchRequestServicerequestRequestpreamble)

	if err != nil {
		return err
	}

	*o = ProductSearchRequestServicerequestRequestpreamble(varProductSearchRequestServicerequestRequestpreamble)

	return err
}

type NullableProductSearchRequestServicerequestRequestpreamble struct {
	value *ProductSearchRequestServicerequestRequestpreamble
	isSet bool
}

func (v NullableProductSearchRequestServicerequestRequestpreamble) Get() *ProductSearchRequestServicerequestRequestpreamble {
	return v.value
}

func (v *NullableProductSearchRequestServicerequestRequestpreamble) Set(val *ProductSearchRequestServicerequestRequestpreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSearchRequestServicerequestRequestpreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSearchRequestServicerequestRequestpreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSearchRequestServicerequestRequestpreamble(val *ProductSearchRequestServicerequestRequestpreamble) *NullableProductSearchRequestServicerequestRequestpreamble {
	return &NullableProductSearchRequestServicerequestRequestpreamble{value: val, isSet: true}
}

func (v NullableProductSearchRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSearchRequestServicerequestRequestpreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


