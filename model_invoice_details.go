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

// checks if the InvoiceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetails{}

// InvoiceDetails struct for InvoiceDetails
type InvoiceDetails struct {
	Serviceresponse *InvoiceDetailResponseServiceresponse `json:"serviceresponse,omitempty"`
}

// NewInvoiceDetails instantiates a new InvoiceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetails() *InvoiceDetails {
	this := InvoiceDetails{}
	return &this
}

// NewInvoiceDetailsWithDefaults instantiates a new InvoiceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailsWithDefaults() *InvoiceDetails {
	this := InvoiceDetails{}
	return &this
}

// GetServiceresponse returns the Serviceresponse field value if set, zero value otherwise.
func (o *InvoiceDetails) GetServiceresponse() InvoiceDetailResponseServiceresponse {
	if o == nil || IsNil(o.Serviceresponse) {
		var ret InvoiceDetailResponseServiceresponse
		return ret
	}
	return *o.Serviceresponse
}

// GetServiceresponseOk returns a tuple with the Serviceresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetails) GetServiceresponseOk() (*InvoiceDetailResponseServiceresponse, bool) {
	if o == nil || IsNil(o.Serviceresponse) {
		return nil, false
	}
	return o.Serviceresponse, true
}

// HasServiceresponse returns a boolean if a field has been set.
func (o *InvoiceDetails) HasServiceresponse() bool {
	if o != nil && !IsNil(o.Serviceresponse) {
		return true
	}

	return false
}

// SetServiceresponse gets a reference to the given InvoiceDetailResponseServiceresponse and assigns it to the Serviceresponse field.
func (o *InvoiceDetails) SetServiceresponse(v InvoiceDetailResponseServiceresponse) {
	o.Serviceresponse = &v
}

func (o InvoiceDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serviceresponse) {
		toSerialize["serviceresponse"] = o.Serviceresponse
	}
	return toSerialize, nil
}

type NullableInvoiceDetails struct {
	value *InvoiceDetails
	isSet bool
}

func (v NullableInvoiceDetails) Get() *InvoiceDetails {
	return v.value
}

func (v *NullableInvoiceDetails) Set(val *InvoiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetails(val *InvoiceDetails) *NullableInvoiceDetails {
	return &NullableInvoiceDetails{value: val, isSet: true}
}

func (v NullableInvoiceDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

