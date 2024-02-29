/*
XI Sdk Resellers

For Ingram Micro Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the ErrorResponseErrorsInnerFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseErrorsInnerFieldsInner{}

// ErrorResponseErrorsInnerFieldsInner struct for ErrorResponseErrorsInnerFieldsInner
type ErrorResponseErrorsInnerFieldsInner struct {
	// Contains the name of the field.
	Field *string `json:"field,omitempty"`
	// Value sent in the input for the specific field.
	Value *string `json:"value,omitempty"`
	// Gives the description of the field message.
	Message *string `json:"message,omitempty"`
}

// NewErrorResponseErrorsInnerFieldsInner instantiates a new ErrorResponseErrorsInnerFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseErrorsInnerFieldsInner() *ErrorResponseErrorsInnerFieldsInner {
	this := ErrorResponseErrorsInnerFieldsInner{}
	return &this
}

// NewErrorResponseErrorsInnerFieldsInnerWithDefaults instantiates a new ErrorResponseErrorsInnerFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseErrorsInnerFieldsInnerWithDefaults() *ErrorResponseErrorsInnerFieldsInner {
	this := ErrorResponseErrorsInnerFieldsInner{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInnerFieldsInner) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInnerFieldsInner) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInnerFieldsInner) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ErrorResponseErrorsInnerFieldsInner) SetField(v string) {
	o.Field = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInnerFieldsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInnerFieldsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInnerFieldsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ErrorResponseErrorsInnerFieldsInner) SetValue(v string) {
	o.Value = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorResponseErrorsInnerFieldsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseErrorsInnerFieldsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponseErrorsInnerFieldsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorResponseErrorsInnerFieldsInner) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorResponseErrorsInnerFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseErrorsInnerFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableErrorResponseErrorsInnerFieldsInner struct {
	value *ErrorResponseErrorsInnerFieldsInner
	isSet bool
}

func (v NullableErrorResponseErrorsInnerFieldsInner) Get() *ErrorResponseErrorsInnerFieldsInner {
	return v.value
}

func (v *NullableErrorResponseErrorsInnerFieldsInner) Set(val *ErrorResponseErrorsInnerFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseErrorsInnerFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseErrorsInnerFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseErrorsInnerFieldsInner(val *ErrorResponseErrorsInnerFieldsInner) *NullableErrorResponseErrorsInnerFieldsInner {
	return &NullableErrorResponseErrorsInnerFieldsInner{value: val, isSet: true}
}

func (v NullableErrorResponseErrorsInnerFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseErrorsInnerFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


