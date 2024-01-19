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

// checks if the OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner{}

// OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner struct for OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner
type OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner struct {
	Attributename *string `json:"attributename,omitempty"`
	Attributevalue *string `json:"attributevalue,omitempty"`
}

// NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner {
	this := OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner{}
	return &this
}

// NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInnerWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInnerWithDefaults() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner {
	this := OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner{}
	return &this
}

// GetAttributename returns the Attributename field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) GetAttributename() string {
	if o == nil || IsNil(o.Attributename) {
		var ret string
		return ret
	}
	return *o.Attributename
}

// GetAttributenameOk returns a tuple with the Attributename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) GetAttributenameOk() (*string, bool) {
	if o == nil || IsNil(o.Attributename) {
		return nil, false
	}
	return o.Attributename, true
}

// HasAttributename returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) HasAttributename() bool {
	if o != nil && !IsNil(o.Attributename) {
		return true
	}

	return false
}

// SetAttributename gets a reference to the given string and assigns it to the Attributename field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) SetAttributename(v string) {
	o.Attributename = &v
}

// GetAttributevalue returns the Attributevalue field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) GetAttributevalue() string {
	if o == nil || IsNil(o.Attributevalue) {
		var ret string
		return ret
	}
	return *o.Attributevalue
}

// GetAttributevalueOk returns a tuple with the Attributevalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) GetAttributevalueOk() (*string, bool) {
	if o == nil || IsNil(o.Attributevalue) {
		return nil, false
	}
	return o.Attributevalue, true
}

// HasAttributevalue returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) HasAttributevalue() bool {
	if o != nil && !IsNil(o.Attributevalue) {
		return true
	}

	return false
}

// SetAttributevalue gets a reference to the given string and assigns it to the Attributevalue field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) SetAttributevalue(v string) {
	o.Attributevalue = &v
}

func (o OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributename) {
		toSerialize["attributename"] = o.Attributename
	}
	if !IsNil(o.Attributevalue) {
		toSerialize["attributevalue"] = o.Attributevalue
	}
	return toSerialize, nil
}

type NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner struct {
	value *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner
	isSet bool
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) Get() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner {
	return v.value
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) Set(val *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner(val *OrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner {
	return &NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner{value: val, isSet: true}
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsLinesInnerProductextendedspecsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


