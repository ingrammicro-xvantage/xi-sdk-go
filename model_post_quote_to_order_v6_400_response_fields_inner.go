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

// checks if the PostQuoteToOrderV6400ResponseFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostQuoteToOrderV6400ResponseFieldsInner{}

// PostQuoteToOrderV6400ResponseFieldsInner struct for PostQuoteToOrderV6400ResponseFieldsInner
type PostQuoteToOrderV6400ResponseFieldsInner struct {
	// Name of the field.
	Field *string `json:"field,omitempty"`
	// A filed level error message.
	Message *string `json:"message,omitempty"`
	// Value of the message.
	Value *string `json:"value,omitempty"`
}

// NewPostQuoteToOrderV6400ResponseFieldsInner instantiates a new PostQuoteToOrderV6400ResponseFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostQuoteToOrderV6400ResponseFieldsInner() *PostQuoteToOrderV6400ResponseFieldsInner {
	this := PostQuoteToOrderV6400ResponseFieldsInner{}
	return &this
}

// NewPostQuoteToOrderV6400ResponseFieldsInnerWithDefaults instantiates a new PostQuoteToOrderV6400ResponseFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostQuoteToOrderV6400ResponseFieldsInnerWithDefaults() *PostQuoteToOrderV6400ResponseFieldsInner {
	this := PostQuoteToOrderV6400ResponseFieldsInner{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) SetMessage(v string) {
	o.Message = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PostQuoteToOrderV6400ResponseFieldsInner) SetValue(v string) {
	o.Value = &v
}

func (o PostQuoteToOrderV6400ResponseFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostQuoteToOrderV6400ResponseFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePostQuoteToOrderV6400ResponseFieldsInner struct {
	value *PostQuoteToOrderV6400ResponseFieldsInner
	isSet bool
}

func (v NullablePostQuoteToOrderV6400ResponseFieldsInner) Get() *PostQuoteToOrderV6400ResponseFieldsInner {
	return v.value
}

func (v *NullablePostQuoteToOrderV6400ResponseFieldsInner) Set(val *PostQuoteToOrderV6400ResponseFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostQuoteToOrderV6400ResponseFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostQuoteToOrderV6400ResponseFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostQuoteToOrderV6400ResponseFieldsInner(val *PostQuoteToOrderV6400ResponseFieldsInner) *NullablePostQuoteToOrderV6400ResponseFieldsInner {
	return &NullablePostQuoteToOrderV6400ResponseFieldsInner{value: val, isSet: true}
}

func (v NullablePostQuoteToOrderV6400ResponseFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostQuoteToOrderV6400ResponseFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

