/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
)

// checks if the PostRenewalssearch400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRenewalssearch400Response{}

// PostRenewalssearch400Response struct for PostRenewalssearch400Response
type PostRenewalssearch400Response struct {
	// Unique Id to identify error.
	Traceid *string `json:"traceid,omitempty"`
	// Describes the type of the error.
	Type *string `json:"type,omitempty"`
	Fields []GetResellerV6ValidateQuote400ResponseFieldsInner `json:"fields,omitempty"`
}

// NewPostRenewalssearch400Response instantiates a new PostRenewalssearch400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRenewalssearch400Response() *PostRenewalssearch400Response {
	this := PostRenewalssearch400Response{}
	return &this
}

// NewPostRenewalssearch400ResponseWithDefaults instantiates a new PostRenewalssearch400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRenewalssearch400ResponseWithDefaults() *PostRenewalssearch400Response {
	this := PostRenewalssearch400Response{}
	return &this
}

// GetTraceid returns the Traceid field value if set, zero value otherwise.
func (o *PostRenewalssearch400Response) GetTraceid() string {
	if o == nil || IsNil(o.Traceid) {
		var ret string
		return ret
	}
	return *o.Traceid
}

// GetTraceidOk returns a tuple with the Traceid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRenewalssearch400Response) GetTraceidOk() (*string, bool) {
	if o == nil || IsNil(o.Traceid) {
		return nil, false
	}
	return o.Traceid, true
}

// HasTraceid returns a boolean if a field has been set.
func (o *PostRenewalssearch400Response) HasTraceid() bool {
	if o != nil && !IsNil(o.Traceid) {
		return true
	}

	return false
}

// SetTraceid gets a reference to the given string and assigns it to the Traceid field.
func (o *PostRenewalssearch400Response) SetTraceid(v string) {
	o.Traceid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostRenewalssearch400Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRenewalssearch400Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostRenewalssearch400Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostRenewalssearch400Response) SetType(v string) {
	o.Type = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PostRenewalssearch400Response) GetFields() []GetResellerV6ValidateQuote400ResponseFieldsInner {
	if o == nil || IsNil(o.Fields) {
		var ret []GetResellerV6ValidateQuote400ResponseFieldsInner
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRenewalssearch400Response) GetFieldsOk() ([]GetResellerV6ValidateQuote400ResponseFieldsInner, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PostRenewalssearch400Response) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []GetResellerV6ValidateQuote400ResponseFieldsInner and assigns it to the Fields field.
func (o *PostRenewalssearch400Response) SetFields(v []GetResellerV6ValidateQuote400ResponseFieldsInner) {
	o.Fields = v
}

func (o PostRenewalssearch400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRenewalssearch400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Traceid) {
		toSerialize["traceid"] = o.Traceid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullablePostRenewalssearch400Response struct {
	value *PostRenewalssearch400Response
	isSet bool
}

func (v NullablePostRenewalssearch400Response) Get() *PostRenewalssearch400Response {
	return v.value
}

func (v *NullablePostRenewalssearch400Response) Set(val *PostRenewalssearch400Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRenewalssearch400Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRenewalssearch400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRenewalssearch400Response(val *PostRenewalssearch400Response) *NullablePostRenewalssearch400Response {
	return &NullablePostRenewalssearch400Response{value: val, isSet: true}
}

func (v NullablePostRenewalssearch400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRenewalssearch400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


