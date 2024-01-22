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

// checks if the OrderCancelResponseServiceresponseResponsepreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCancelResponseServiceresponseResponsepreamble{}

// OrderCancelResponseServiceresponseResponsepreamble struct for OrderCancelResponseServiceresponseResponsepreamble
type OrderCancelResponseServiceresponseResponsepreamble struct {
	RequestStatus *string `json:"requestStatus,omitempty"`
	ReturnCode *string `json:"returnCode,omitempty"`
	ReturnMessage *string `json:"returnMessage,omitempty"`
}

// NewOrderCancelResponseServiceresponseResponsepreamble instantiates a new OrderCancelResponseServiceresponseResponsepreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCancelResponseServiceresponseResponsepreamble() *OrderCancelResponseServiceresponseResponsepreamble {
	this := OrderCancelResponseServiceresponseResponsepreamble{}
	return &this
}

// NewOrderCancelResponseServiceresponseResponsepreambleWithDefaults instantiates a new OrderCancelResponseServiceresponseResponsepreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCancelResponseServiceresponseResponsepreambleWithDefaults() *OrderCancelResponseServiceresponseResponsepreamble {
	this := OrderCancelResponseServiceresponseResponsepreamble{}
	return &this
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *OrderCancelResponseServiceresponseResponsepreamble) GetRequestStatus() string {
	if o == nil || IsNil(o.RequestStatus) {
		var ret string
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelResponseServiceresponseResponsepreamble) GetRequestStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RequestStatus) {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *OrderCancelResponseServiceresponseResponsepreamble) HasRequestStatus() bool {
	if o != nil && !IsNil(o.RequestStatus) {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given string and assigns it to the RequestStatus field.
func (o *OrderCancelResponseServiceresponseResponsepreamble) SetRequestStatus(v string) {
	o.RequestStatus = &v
}

// GetReturnCode returns the ReturnCode field value if set, zero value otherwise.
func (o *OrderCancelResponseServiceresponseResponsepreamble) GetReturnCode() string {
	if o == nil || IsNil(o.ReturnCode) {
		var ret string
		return ret
	}
	return *o.ReturnCode
}

// GetReturnCodeOk returns a tuple with the ReturnCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelResponseServiceresponseResponsepreamble) GetReturnCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnCode) {
		return nil, false
	}
	return o.ReturnCode, true
}

// HasReturnCode returns a boolean if a field has been set.
func (o *OrderCancelResponseServiceresponseResponsepreamble) HasReturnCode() bool {
	if o != nil && !IsNil(o.ReturnCode) {
		return true
	}

	return false
}

// SetReturnCode gets a reference to the given string and assigns it to the ReturnCode field.
func (o *OrderCancelResponseServiceresponseResponsepreamble) SetReturnCode(v string) {
	o.ReturnCode = &v
}

// GetReturnMessage returns the ReturnMessage field value if set, zero value otherwise.
func (o *OrderCancelResponseServiceresponseResponsepreamble) GetReturnMessage() string {
	if o == nil || IsNil(o.ReturnMessage) {
		var ret string
		return ret
	}
	return *o.ReturnMessage
}

// GetReturnMessageOk returns a tuple with the ReturnMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCancelResponseServiceresponseResponsepreamble) GetReturnMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnMessage) {
		return nil, false
	}
	return o.ReturnMessage, true
}

// HasReturnMessage returns a boolean if a field has been set.
func (o *OrderCancelResponseServiceresponseResponsepreamble) HasReturnMessage() bool {
	if o != nil && !IsNil(o.ReturnMessage) {
		return true
	}

	return false
}

// SetReturnMessage gets a reference to the given string and assigns it to the ReturnMessage field.
func (o *OrderCancelResponseServiceresponseResponsepreamble) SetReturnMessage(v string) {
	o.ReturnMessage = &v
}

func (o OrderCancelResponseServiceresponseResponsepreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCancelResponseServiceresponseResponsepreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestStatus) {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if !IsNil(o.ReturnCode) {
		toSerialize["returnCode"] = o.ReturnCode
	}
	if !IsNil(o.ReturnMessage) {
		toSerialize["returnMessage"] = o.ReturnMessage
	}
	return toSerialize, nil
}

type NullableOrderCancelResponseServiceresponseResponsepreamble struct {
	value *OrderCancelResponseServiceresponseResponsepreamble
	isSet bool
}

func (v NullableOrderCancelResponseServiceresponseResponsepreamble) Get() *OrderCancelResponseServiceresponseResponsepreamble {
	return v.value
}

func (v *NullableOrderCancelResponseServiceresponseResponsepreamble) Set(val *OrderCancelResponseServiceresponseResponsepreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCancelResponseServiceresponseResponsepreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCancelResponseServiceresponseResponsepreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCancelResponseServiceresponseResponsepreamble(val *OrderCancelResponseServiceresponseResponsepreamble) *NullableOrderCancelResponseServiceresponseResponsepreamble {
	return &NullableOrderCancelResponseServiceresponseResponsepreamble{value: val, isSet: true}
}

func (v NullableOrderCancelResponseServiceresponseResponsepreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCancelResponseServiceresponseResponsepreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

