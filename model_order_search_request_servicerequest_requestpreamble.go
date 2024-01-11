/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderSearchRequestServicerequestRequestpreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchRequestServicerequestRequestpreamble{}

// OrderSearchRequestServicerequestRequestpreamble struct for OrderSearchRequestServicerequestRequestpreamble
type OrderSearchRequestServicerequestRequestpreamble struct {
	IsoCountryCode string `json:"isoCountryCode"`
	CustomerNumber string `json:"customerNumber"`
}

type _OrderSearchRequestServicerequestRequestpreamble OrderSearchRequestServicerequestRequestpreamble

// NewOrderSearchRequestServicerequestRequestpreamble instantiates a new OrderSearchRequestServicerequestRequestpreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchRequestServicerequestRequestpreamble(isoCountryCode string, customerNumber string) *OrderSearchRequestServicerequestRequestpreamble {
	this := OrderSearchRequestServicerequestRequestpreamble{}
	this.IsoCountryCode = isoCountryCode
	this.CustomerNumber = customerNumber
	return &this
}

// NewOrderSearchRequestServicerequestRequestpreambleWithDefaults instantiates a new OrderSearchRequestServicerequestRequestpreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchRequestServicerequestRequestpreambleWithDefaults() *OrderSearchRequestServicerequestRequestpreamble {
	this := OrderSearchRequestServicerequestRequestpreamble{}
	return &this
}

// GetIsoCountryCode returns the IsoCountryCode field value
func (o *OrderSearchRequestServicerequestRequestpreamble) GetIsoCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsoCountryCode
}

// GetIsoCountryCodeOk returns a tuple with the IsoCountryCode field value
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequestRequestpreamble) GetIsoCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsoCountryCode, true
}

// SetIsoCountryCode sets field value
func (o *OrderSearchRequestServicerequestRequestpreamble) SetIsoCountryCode(v string) {
	o.IsoCountryCode = v
}

// GetCustomerNumber returns the CustomerNumber field value
func (o *OrderSearchRequestServicerequestRequestpreamble) GetCustomerNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequestRequestpreamble) GetCustomerNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerNumber, true
}

// SetCustomerNumber sets field value
func (o *OrderSearchRequestServicerequestRequestpreamble) SetCustomerNumber(v string) {
	o.CustomerNumber = v
}

func (o OrderSearchRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchRequestServicerequestRequestpreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isoCountryCode"] = o.IsoCountryCode
	toSerialize["customerNumber"] = o.CustomerNumber
	return toSerialize, nil
}

func (o *OrderSearchRequestServicerequestRequestpreamble) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"isoCountryCode",
		"customerNumber",
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

	varOrderSearchRequestServicerequestRequestpreamble := _OrderSearchRequestServicerequestRequestpreamble{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderSearchRequestServicerequestRequestpreamble)

	if err != nil {
		return err
	}

	*o = OrderSearchRequestServicerequestRequestpreamble(varOrderSearchRequestServicerequestRequestpreamble)

	return err
}

type NullableOrderSearchRequestServicerequestRequestpreamble struct {
	value *OrderSearchRequestServicerequestRequestpreamble
	isSet bool
}

func (v NullableOrderSearchRequestServicerequestRequestpreamble) Get() *OrderSearchRequestServicerequestRequestpreamble {
	return v.value
}

func (v *NullableOrderSearchRequestServicerequestRequestpreamble) Set(val *OrderSearchRequestServicerequestRequestpreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchRequestServicerequestRequestpreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchRequestServicerequestRequestpreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchRequestServicerequestRequestpreamble(val *OrderSearchRequestServicerequestRequestpreamble) *NullableOrderSearchRequestServicerequestRequestpreamble {
	return &NullableOrderSearchRequestServicerequestRequestpreamble{value: val, isSet: true}
}

func (v NullableOrderSearchRequestServicerequestRequestpreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchRequestServicerequestRequestpreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


