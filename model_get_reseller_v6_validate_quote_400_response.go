/*
XI Sdk Resellers

For Resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the GetResellerV6ValidateQuote400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetResellerV6ValidateQuote400Response{}

// GetResellerV6ValidateQuote400Response struct for GetResellerV6ValidateQuote400Response
type GetResellerV6ValidateQuote400Response struct {
	// Unique Id to identify error.
	Traceid *string `json:"traceid,omitempty"`
	// Describes the type of the error.
	Type *string `json:"type,omitempty"`
	// A detailed error message.
	Message *string `json:"message,omitempty"`
	Fields []GetResellerV6ValidateQuote400ResponseFieldsInner `json:"fields,omitempty"`
}

// NewGetResellerV6ValidateQuote400Response instantiates a new GetResellerV6ValidateQuote400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetResellerV6ValidateQuote400Response() *GetResellerV6ValidateQuote400Response {
	this := GetResellerV6ValidateQuote400Response{}
	return &this
}

// NewGetResellerV6ValidateQuote400ResponseWithDefaults instantiates a new GetResellerV6ValidateQuote400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetResellerV6ValidateQuote400ResponseWithDefaults() *GetResellerV6ValidateQuote400Response {
	this := GetResellerV6ValidateQuote400Response{}
	return &this
}

// GetTraceid returns the Traceid field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400Response) GetTraceid() string {
	if o == nil || IsNil(o.Traceid) {
		var ret string
		return ret
	}
	return *o.Traceid
}

// GetTraceidOk returns a tuple with the Traceid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400Response) GetTraceidOk() (*string, bool) {
	if o == nil || IsNil(o.Traceid) {
		return nil, false
	}
	return o.Traceid, true
}

// HasTraceid returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400Response) HasTraceid() bool {
	if o != nil && !IsNil(o.Traceid) {
		return true
	}

	return false
}

// SetTraceid gets a reference to the given string and assigns it to the Traceid field.
func (o *GetResellerV6ValidateQuote400Response) SetTraceid(v string) {
	o.Traceid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetResellerV6ValidateQuote400Response) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetResellerV6ValidateQuote400Response) SetMessage(v string) {
	o.Message = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *GetResellerV6ValidateQuote400Response) GetFields() []GetResellerV6ValidateQuote400ResponseFieldsInner {
	if o == nil || IsNil(o.Fields) {
		var ret []GetResellerV6ValidateQuote400ResponseFieldsInner
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetResellerV6ValidateQuote400Response) GetFieldsOk() ([]GetResellerV6ValidateQuote400ResponseFieldsInner, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *GetResellerV6ValidateQuote400Response) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []GetResellerV6ValidateQuote400ResponseFieldsInner and assigns it to the Fields field.
func (o *GetResellerV6ValidateQuote400Response) SetFields(v []GetResellerV6ValidateQuote400ResponseFieldsInner) {
	o.Fields = v
}

func (o GetResellerV6ValidateQuote400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetResellerV6ValidateQuote400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Traceid) {
		toSerialize["traceid"] = o.Traceid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableGetResellerV6ValidateQuote400Response struct {
	value *GetResellerV6ValidateQuote400Response
	isSet bool
}

func (v NullableGetResellerV6ValidateQuote400Response) Get() *GetResellerV6ValidateQuote400Response {
	return v.value
}

func (v *NullableGetResellerV6ValidateQuote400Response) Set(val *GetResellerV6ValidateQuote400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetResellerV6ValidateQuote400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetResellerV6ValidateQuote400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetResellerV6ValidateQuote400Response(val *GetResellerV6ValidateQuote400Response) *NullableGetResellerV6ValidateQuote400Response {
	return &NullableGetResellerV6ValidateQuote400Response{value: val, isSet: true}
}

func (v NullableGetResellerV6ValidateQuote400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetResellerV6ValidateQuote400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


