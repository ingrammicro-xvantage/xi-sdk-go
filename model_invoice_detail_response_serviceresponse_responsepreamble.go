/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the InvoiceDetailResponseServiceresponseResponsepreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailResponseServiceresponseResponsepreamble{}

// InvoiceDetailResponseServiceresponseResponsepreamble struct for InvoiceDetailResponseServiceresponseResponsepreamble
type InvoiceDetailResponseServiceresponseResponsepreamble struct {
	Responsestatus *string `json:"responsestatus,omitempty"`
	Statuscode *string `json:"statuscode,omitempty"`
	Responsemessage *string `json:"responsemessage,omitempty"`
}

// NewInvoiceDetailResponseServiceresponseResponsepreamble instantiates a new InvoiceDetailResponseServiceresponseResponsepreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailResponseServiceresponseResponsepreamble() *InvoiceDetailResponseServiceresponseResponsepreamble {
	this := InvoiceDetailResponseServiceresponseResponsepreamble{}
	return &this
}

// NewInvoiceDetailResponseServiceresponseResponsepreambleWithDefaults instantiates a new InvoiceDetailResponseServiceresponseResponsepreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailResponseServiceresponseResponsepreambleWithDefaults() *InvoiceDetailResponseServiceresponseResponsepreamble {
	this := InvoiceDetailResponseServiceresponseResponsepreamble{}
	return &this
}

// GetResponsestatus returns the Responsestatus field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) GetResponsestatus() string {
	if o == nil || IsNil(o.Responsestatus) {
		var ret string
		return ret
	}
	return *o.Responsestatus
}

// GetResponsestatusOk returns a tuple with the Responsestatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) GetResponsestatusOk() (*string, bool) {
	if o == nil || IsNil(o.Responsestatus) {
		return nil, false
	}
	return o.Responsestatus, true
}

// HasResponsestatus returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) HasResponsestatus() bool {
	if o != nil && !IsNil(o.Responsestatus) {
		return true
	}

	return false
}

// SetResponsestatus gets a reference to the given string and assigns it to the Responsestatus field.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) SetResponsestatus(v string) {
	o.Responsestatus = &v
}

// GetStatuscode returns the Statuscode field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) GetStatuscode() string {
	if o == nil || IsNil(o.Statuscode) {
		var ret string
		return ret
	}
	return *o.Statuscode
}

// GetStatuscodeOk returns a tuple with the Statuscode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) GetStatuscodeOk() (*string, bool) {
	if o == nil || IsNil(o.Statuscode) {
		return nil, false
	}
	return o.Statuscode, true
}

// HasStatuscode returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) HasStatuscode() bool {
	if o != nil && !IsNil(o.Statuscode) {
		return true
	}

	return false
}

// SetStatuscode gets a reference to the given string and assigns it to the Statuscode field.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) SetStatuscode(v string) {
	o.Statuscode = &v
}

// GetResponsemessage returns the Responsemessage field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) GetResponsemessage() string {
	if o == nil || IsNil(o.Responsemessage) {
		var ret string
		return ret
	}
	return *o.Responsemessage
}

// GetResponsemessageOk returns a tuple with the Responsemessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) GetResponsemessageOk() (*string, bool) {
	if o == nil || IsNil(o.Responsemessage) {
		return nil, false
	}
	return o.Responsemessage, true
}

// HasResponsemessage returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) HasResponsemessage() bool {
	if o != nil && !IsNil(o.Responsemessage) {
		return true
	}

	return false
}

// SetResponsemessage gets a reference to the given string and assigns it to the Responsemessage field.
func (o *InvoiceDetailResponseServiceresponseResponsepreamble) SetResponsemessage(v string) {
	o.Responsemessage = &v
}

func (o InvoiceDetailResponseServiceresponseResponsepreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailResponseServiceresponseResponsepreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsestatus) {
		toSerialize["responsestatus"] = o.Responsestatus
	}
	if !IsNil(o.Statuscode) {
		toSerialize["statuscode"] = o.Statuscode
	}
	if !IsNil(o.Responsemessage) {
		toSerialize["responsemessage"] = o.Responsemessage
	}
	return toSerialize, nil
}

type NullableInvoiceDetailResponseServiceresponseResponsepreamble struct {
	value *InvoiceDetailResponseServiceresponseResponsepreamble
	isSet bool
}

func (v NullableInvoiceDetailResponseServiceresponseResponsepreamble) Get() *InvoiceDetailResponseServiceresponseResponsepreamble {
	return v.value
}

func (v *NullableInvoiceDetailResponseServiceresponseResponsepreamble) Set(val *InvoiceDetailResponseServiceresponseResponsepreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailResponseServiceresponseResponsepreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailResponseServiceresponseResponsepreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailResponseServiceresponseResponsepreamble(val *InvoiceDetailResponseServiceresponseResponsepreamble) *NullableInvoiceDetailResponseServiceresponseResponsepreamble {
	return &NullableInvoiceDetailResponseServiceresponseResponsepreamble{value: val, isSet: true}
}

func (v NullableInvoiceDetailResponseServiceresponseResponsepreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailResponseServiceresponseResponsepreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


