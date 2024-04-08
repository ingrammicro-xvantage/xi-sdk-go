/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner{}

// OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner struct for OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner
type OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner struct {
	// For orders or invoices. For orders the link provides details of the order. For invoices the link provides details of the invoice.
	Topic *string `json:"topic,omitempty"`
	// The URL endpoint for accessing the relevant data.
	Href *string `json:"href,omitempty"`
	// The type of call that can be made to the href link (GET, POST, Etc.).
	Type *string `json:"type,omitempty"`
}

// NewOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner instantiates a new OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner() *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner {
	this := OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner{}
	return &this
}

// NewOrderSearchResponseOrdersInnerSubOrdersInnerLinksInnerWithDefaults instantiates a new OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchResponseOrdersInnerSubOrdersInnerLinksInnerWithDefaults() *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner {
	this := OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) SetTopic(v string) {
	o.Topic = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) SetType(v string) {
	o.Type = &v
}

func (o OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner struct {
	value *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner
	isSet bool
}

func (v NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) Get() *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner {
	return v.value
}

func (v *NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) Set(val *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner(val *OrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) *NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner {
	return &NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner{value: val, isSet: true}
}

func (v NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchResponseOrdersInnerSubOrdersInnerLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


