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

// checks if the OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner{}

// OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner struct for OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner
type OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner struct {
	// The serial number for the line item.                  
	SerialNumber *string `json:"serialNumber,omitempty"`
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner{}
	return &this
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	return toSerialize, nil
}

type NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner struct {
	value *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner
	isSet bool
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) Get() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner {
	return v.value
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) Set(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner {
	return &NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner{value: val, isSet: true}
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerSerialNumberDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


