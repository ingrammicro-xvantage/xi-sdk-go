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

// checks if the InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner{}

// InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner struct for InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner
type InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner struct {
	Attributename *string `json:"attributename,omitempty"`
	Attributevalue *string `json:"attributevalue,omitempty"`
}

// NewInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner instantiates a new InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner() *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner {
	this := InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner{}
	return &this
}

// NewInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInnerWithDefaults instantiates a new InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInnerWithDefaults() *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner {
	this := InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner{}
	return &this
}

// GetAttributename returns the Attributename field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) GetAttributename() string {
	if o == nil || IsNil(o.Attributename) {
		var ret string
		return ret
	}
	return *o.Attributename
}

// GetAttributenameOk returns a tuple with the Attributename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) GetAttributenameOk() (*string, bool) {
	if o == nil || IsNil(o.Attributename) {
		return nil, false
	}
	return o.Attributename, true
}

// HasAttributename returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) HasAttributename() bool {
	if o != nil && !IsNil(o.Attributename) {
		return true
	}

	return false
}

// SetAttributename gets a reference to the given string and assigns it to the Attributename field.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) SetAttributename(v string) {
	o.Attributename = &v
}

// GetAttributevalue returns the Attributevalue field value if set, zero value otherwise.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) GetAttributevalue() string {
	if o == nil || IsNil(o.Attributevalue) {
		var ret string
		return ret
	}
	return *o.Attributevalue
}

// GetAttributevalueOk returns a tuple with the Attributevalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) GetAttributevalueOk() (*string, bool) {
	if o == nil || IsNil(o.Attributevalue) {
		return nil, false
	}
	return o.Attributevalue, true
}

// HasAttributevalue returns a boolean if a field has been set.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) HasAttributevalue() bool {
	if o != nil && !IsNil(o.Attributevalue) {
		return true
	}

	return false
}

// SetAttributevalue gets a reference to the given string and assigns it to the Attributevalue field.
func (o *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) SetAttributevalue(v string) {
	o.Attributevalue = &v
}

func (o InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributename) {
		toSerialize["attributename"] = o.Attributename
	}
	if !IsNil(o.Attributevalue) {
		toSerialize["attributevalue"] = o.Attributevalue
	}
	return toSerialize, nil
}

type NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner struct {
	value *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner
	isSet bool
}

func (v NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) Get() *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner {
	return v.value
}

func (v *NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) Set(val *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner(val *InvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) *NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner {
	return &NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner{value: val, isSet: true}
}

func (v NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailResponseServiceresponseInvoicedetailresponseExtendedspecsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


