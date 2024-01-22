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

// checks if the InvoiceDetailResponseServiceresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailResponseServiceresponse{}

// InvoiceDetailResponseServiceresponse struct for InvoiceDetailResponseServiceresponse
type InvoiceDetailResponseServiceresponse struct {
	Responsepreamble *InvoiceDetailResponseServiceresponseResponsepreamble `json:"responsepreamble,omitempty"`
	Invoicedetailresponse *InvoiceDetailResponseServiceresponseInvoicedetailresponse `json:"invoicedetailresponse,omitempty"`
}

// NewInvoiceDetailResponseServiceresponse instantiates a new InvoiceDetailResponseServiceresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailResponseServiceresponse() *InvoiceDetailResponseServiceresponse {
	this := InvoiceDetailResponseServiceresponse{}
	return &this
}

// NewInvoiceDetailResponseServiceresponseWithDefaults instantiates a new InvoiceDetailResponseServiceresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailResponseServiceresponseWithDefaults() *InvoiceDetailResponseServiceresponse {
	this := InvoiceDetailResponseServiceresponse{}
	return &this
}

// GetResponsepreamble returns the Responsepreamble field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponse) GetResponsepreamble() InvoiceDetailResponseServiceresponseResponsepreamble {
	if o == nil || IsNil(o.Responsepreamble) {
		var ret InvoiceDetailResponseServiceresponseResponsepreamble
		return ret
	}
	return *o.Responsepreamble
}

// GetResponsepreambleOk returns a tuple with the Responsepreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponse) GetResponsepreambleOk() (*InvoiceDetailResponseServiceresponseResponsepreamble, bool) {
	if o == nil || IsNil(o.Responsepreamble) {
		return nil, false
	}
	return o.Responsepreamble, true
}

// HasResponsepreamble returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponse) HasResponsepreamble() bool {
	if o != nil && !IsNil(o.Responsepreamble) {
		return true
	}

	return false
}

// SetResponsepreamble gets a reference to the given InvoiceDetailResponseServiceresponseResponsepreamble and assigns it to the Responsepreamble field.
func (o *InvoiceDetailResponseServiceresponse) SetResponsepreamble(v InvoiceDetailResponseServiceresponseResponsepreamble) {
	o.Responsepreamble = &v
}

// GetInvoicedetailresponse returns the Invoicedetailresponse field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponse) GetInvoicedetailresponse() InvoiceDetailResponseServiceresponseInvoicedetailresponse {
	if o == nil || IsNil(o.Invoicedetailresponse) {
		var ret InvoiceDetailResponseServiceresponseInvoicedetailresponse
		return ret
	}
	return *o.Invoicedetailresponse
}

// GetInvoicedetailresponseOk returns a tuple with the Invoicedetailresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponse) GetInvoicedetailresponseOk() (*InvoiceDetailResponseServiceresponseInvoicedetailresponse, bool) {
	if o == nil || IsNil(o.Invoicedetailresponse) {
		return nil, false
	}
	return o.Invoicedetailresponse, true
}

// HasInvoicedetailresponse returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponse) HasInvoicedetailresponse() bool {
	if o != nil && !IsNil(o.Invoicedetailresponse) {
		return true
	}

	return false
}

// SetInvoicedetailresponse gets a reference to the given InvoiceDetailResponseServiceresponseInvoicedetailresponse and assigns it to the Invoicedetailresponse field.
func (o *InvoiceDetailResponseServiceresponse) SetInvoicedetailresponse(v InvoiceDetailResponseServiceresponseInvoicedetailresponse) {
	o.Invoicedetailresponse = &v
}

func (o InvoiceDetailResponseServiceresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailResponseServiceresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsepreamble) {
		toSerialize["responsepreamble"] = o.Responsepreamble
	}
	if !IsNil(o.Invoicedetailresponse) {
		toSerialize["invoicedetailresponse"] = o.Invoicedetailresponse
	}
	return toSerialize, nil
}

type NullableInvoiceDetailResponseServiceresponse struct {
	value *InvoiceDetailResponseServiceresponse
	isSet bool
}

func (v NullableInvoiceDetailResponseServiceresponse) Get() *InvoiceDetailResponseServiceresponse {
	return v.value
}

func (v *NullableInvoiceDetailResponseServiceresponse) Set(val *InvoiceDetailResponseServiceresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailResponseServiceresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailResponseServiceresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailResponseServiceresponse(val *InvoiceDetailResponseServiceresponse) *NullableInvoiceDetailResponseServiceresponse {
	return &NullableInvoiceDetailResponseServiceresponse{value: val, isSet: true}
}

func (v NullableInvoiceDetailResponseServiceresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailResponseServiceresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

