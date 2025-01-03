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

// checks if the ErrorResponseErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseErrorsInner{}

// ErrorResponseErrorsInner struct for ErrorResponseErrorsInner
type ErrorResponseErrorsInner struct {
	// Unique Id to identify error.
	Id *string `json:"id,omitempty"`
	// Describes the type of the error.
	Type *string `json:"type,omitempty"`
	// Describes the error message.
	Message *string `json:"message,omitempty"`
	Fields []ErrorResponseErrorsInnerFieldsInner `json:"fields,omitempty"`
}

// NewErrorResponseErrorsInner instantiates a new ErrorResponseErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseErrorsInner() *ErrorResponseErrorsInner {
	this := ErrorResponseErrorsInner{}
	return &this
}

// NewErrorResponseErrorsInnerWithDefaults instantiates a new ErrorResponseErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseErrorsInnerWithDefaults() *ErrorResponseErrorsInner {
	this := ErrorResponseErrorsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ErrorResponseErrorsInner) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ErrorResponseErrorsInner) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorResponseErrorsInner) SetMessage(v string) {
	o.Message = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInner) GetFields() []ErrorResponseErrorsInnerFieldsInner {
	if o == nil || IsNil(o.Fields) {
		var ret []ErrorResponseErrorsInnerFieldsInner
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInner) GetFieldsOk() ([]ErrorResponseErrorsInnerFieldsInner, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInner) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ErrorResponseErrorsInnerFieldsInner and assigns it to the Fields field.
func (o *ErrorResponseErrorsInner) SetFields(v []ErrorResponseErrorsInnerFieldsInner) {
	o.Fields = v
}

func (o ErrorResponseErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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

type NullableErrorResponseErrorsInner struct {
	value *ErrorResponseErrorsInner
	isSet bool
}

func (v NullableErrorResponseErrorsInner) Get() *ErrorResponseErrorsInner {
	return v.value
}

func (v *NullableErrorResponseErrorsInner) Set(val *ErrorResponseErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseErrorsInner(val *ErrorResponseErrorsInner) *NullableErrorResponseErrorsInner {
	return &NullableErrorResponseErrorsInner{value: val, isSet: true}
}

func (v NullableErrorResponseErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


