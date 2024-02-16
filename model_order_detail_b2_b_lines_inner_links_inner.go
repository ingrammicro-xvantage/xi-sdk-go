/*
XI SDK Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderDetailB2BLinesInnerLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerLinksInner{}

// OrderDetailB2BLinesInnerLinksInner Link to Order Details for the line item.
type OrderDetailB2BLinesInnerLinksInner struct {
	// Provides the details of the line item.
	Topic *string `json:"topic,omitempty"`
	// The API endpoint for accessing the relevant data.
	Href *string `json:"href,omitempty"`
	// The type of call that can be made to the href link(GET,POST etc).
	Type *string `json:"type,omitempty"`
}

// NewOrderDetailB2BLinesInnerLinksInner instantiates a new OrderDetailB2BLinesInnerLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerLinksInner() *OrderDetailB2BLinesInnerLinksInner {
	this := OrderDetailB2BLinesInnerLinksInner{}
	return &this
}

// NewOrderDetailB2BLinesInnerLinksInnerWithDefaults instantiates a new OrderDetailB2BLinesInnerLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerLinksInnerWithDefaults() *OrderDetailB2BLinesInnerLinksInner {
	this := OrderDetailB2BLinesInnerLinksInner{}
	return &this
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerLinksInner) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerLinksInner) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerLinksInner) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *OrderDetailB2BLinesInnerLinksInner) SetTopic(v string) {
	o.Topic = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerLinksInner) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerLinksInner) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerLinksInner) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *OrderDetailB2BLinesInnerLinksInner) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerLinksInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerLinksInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerLinksInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderDetailB2BLinesInnerLinksInner) SetType(v string) {
	o.Type = &v
}

func (o OrderDetailB2BLinesInnerLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerLinksInner) ToMap() (map[string]interface{}, error) {
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

type NullableOrderDetailB2BLinesInnerLinksInner struct {
	value *OrderDetailB2BLinesInnerLinksInner
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerLinksInner) Get() *OrderDetailB2BLinesInnerLinksInner {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerLinksInner) Set(val *OrderDetailB2BLinesInnerLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerLinksInner(val *OrderDetailB2BLinesInnerLinksInner) *NullableOrderDetailB2BLinesInnerLinksInner {
	return &NullableOrderDetailB2BLinesInnerLinksInner{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


