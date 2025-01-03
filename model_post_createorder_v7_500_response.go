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

// checks if the PostCreateorderV7500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCreateorderV7500Response{}

// PostCreateorderV7500Response struct for PostCreateorderV7500Response
type PostCreateorderV7500Response struct {
	// Unique Id to identify error.
	Traceid *string `json:"traceid,omitempty"`
	// Describes the type of the error.
	Type *string `json:"type,omitempty"`
	// Describes the error message.
	Message *string `json:"message,omitempty"`
	Fields []map[string]interface{} `json:"fields,omitempty"`
}

// NewPostCreateorderV7500Response instantiates a new PostCreateorderV7500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCreateorderV7500Response() *PostCreateorderV7500Response {
	this := PostCreateorderV7500Response{}
	return &this
}

// NewPostCreateorderV7500ResponseWithDefaults instantiates a new PostCreateorderV7500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCreateorderV7500ResponseWithDefaults() *PostCreateorderV7500Response {
	this := PostCreateorderV7500Response{}
	return &this
}

// GetTraceid returns the Traceid field value if set, zero value otherwise.
func (o *PostCreateorderV7500Response) GetTraceid() string {
	if o == nil || IsNil(o.Traceid) {
		var ret string
		return ret
	}
	return *o.Traceid
}

// GetTraceidOk returns a tuple with the Traceid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateorderV7500Response) GetTraceidOk() (*string, bool) {
	if o == nil || IsNil(o.Traceid) {
		return nil, false
	}
	return o.Traceid, true
}

// HasTraceid returns a boolean if a field has been set.
func (o *PostCreateorderV7500Response) HasTraceid() bool {
	if o != nil && !IsNil(o.Traceid) {
		return true
	}

	return false
}

// SetTraceid gets a reference to the given string and assigns it to the Traceid field.
func (o *PostCreateorderV7500Response) SetTraceid(v string) {
	o.Traceid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PostCreateorderV7500Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateorderV7500Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PostCreateorderV7500Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PostCreateorderV7500Response) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostCreateorderV7500Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateorderV7500Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostCreateorderV7500Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostCreateorderV7500Response) SetMessage(v string) {
	o.Message = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *PostCreateorderV7500Response) GetFields() []map[string]interface{} {
	if o == nil || IsNil(o.Fields) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCreateorderV7500Response) GetFieldsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *PostCreateorderV7500Response) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []map[string]interface{} and assigns it to the Fields field.
func (o *PostCreateorderV7500Response) SetFields(v []map[string]interface{}) {
	o.Fields = v
}

func (o PostCreateorderV7500Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCreateorderV7500Response) ToMap() (map[string]interface{}, error) {
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

type NullablePostCreateorderV7500Response struct {
	value *PostCreateorderV7500Response
	isSet bool
}

func (v NullablePostCreateorderV7500Response) Get() *PostCreateorderV7500Response {
	return v.value
}

func (v *NullablePostCreateorderV7500Response) Set(val *PostCreateorderV7500Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCreateorderV7500Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCreateorderV7500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCreateorderV7500Response(val *PostCreateorderV7500Response) *NullablePostCreateorderV7500Response {
	return &NullablePostCreateorderV7500Response{value: val, isSet: true}
}

func (v NullablePostCreateorderV7500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCreateorderV7500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


