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

// checks if the OrderCreateResponseOrdersInnerLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseOrdersInnerLinksInner{}

// OrderCreateResponseOrdersInnerLinksInner struct for OrderCreateResponseOrdersInnerLinksInner
type OrderCreateResponseOrdersInnerLinksInner struct {
	// Provides the details of the orders.
	Topic *string `json:"topic,omitempty"`
	// The URL endpoint for accessing the relevant data.
	Href *string `json:"href,omitempty"`
	// The type of call that can be made to the href link (GET, POST, Etc.).
	Type *string `json:"type,omitempty"`
}

// NewOrderCreateResponseOrdersInnerLinksInner instantiates a new OrderCreateResponseOrdersInnerLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseOrdersInnerLinksInner() *OrderCreateResponseOrdersInnerLinksInner {
	this := OrderCreateResponseOrdersInnerLinksInner{}
	return &this
}

// NewOrderCreateResponseOrdersInnerLinksInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseOrdersInnerLinksInnerWithDefaults() *OrderCreateResponseOrdersInnerLinksInner {
	this := OrderCreateResponseOrdersInnerLinksInner{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinksInner) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinksInner) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinksInner) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *OrderCreateResponseOrdersInnerLinksInner) SetTopic(v string) {
	o.Topic = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinksInner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinksInner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinksInner) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *OrderCreateResponseOrdersInnerLinksInner) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinksInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinksInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinksInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderCreateResponseOrdersInnerLinksInner) SetType(v string) {
	o.Type = &v
}

func (o OrderCreateResponseOrdersInnerLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseOrdersInnerLinksInner) ToMap() (map[string]interface{}, error) {
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

type NullableOrderCreateResponseOrdersInnerLinksInner struct {
	value *OrderCreateResponseOrdersInnerLinksInner
	isSet bool
}

func (v NullableOrderCreateResponseOrdersInnerLinksInner) Get() *OrderCreateResponseOrdersInnerLinksInner {
	return v.value
}

func (v *NullableOrderCreateResponseOrdersInnerLinksInner) Set(val *OrderCreateResponseOrdersInnerLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseOrdersInnerLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseOrdersInnerLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseOrdersInnerLinksInner(val *OrderCreateResponseOrdersInnerLinksInner) *NullableOrderCreateResponseOrdersInnerLinksInner {
	return &NullableOrderCreateResponseOrdersInnerLinksInner{value: val, isSet: true}
}

func (v NullableOrderCreateResponseOrdersInnerLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseOrdersInnerLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


