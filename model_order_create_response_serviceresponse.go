/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
)

// checks if the OrderCreateResponseServiceresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseServiceresponse{}

// OrderCreateResponseServiceresponse struct for OrderCreateResponseServiceresponse
type OrderCreateResponseServiceresponse struct {
	Responsepreamble *InvoiceDetailResponseServiceresponseResponsepreamble `json:"responsepreamble,omitempty"`
	Ordersummary *OrderCreateResponseServiceresponseOrdersummary `json:"ordersummary,omitempty"`
	// Collection of orders
	Ordercreateresponse []OrderCreateResponseServiceresponseOrdercreateresponseInner `json:"ordercreateresponse,omitempty"`
}

// NewOrderCreateResponseServiceresponse instantiates a new OrderCreateResponseServiceresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseServiceresponse() *OrderCreateResponseServiceresponse {
	this := OrderCreateResponseServiceresponse{}
	return &this
}

// NewOrderCreateResponseServiceresponseWithDefaults instantiates a new OrderCreateResponseServiceresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseServiceresponseWithDefaults() *OrderCreateResponseServiceresponse {
	this := OrderCreateResponseServiceresponse{}
	return &this
}

// GetResponsepreamble returns the Responsepreamble field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponse) GetResponsepreamble() InvoiceDetailResponseServiceresponseResponsepreamble {
	if o == nil || IsNil(o.Responsepreamble) {
		var ret InvoiceDetailResponseServiceresponseResponsepreamble
		return ret
	}
	return *o.Responsepreamble
}

// GetResponsepreambleOk returns a tuple with the Responsepreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponse) GetResponsepreambleOk() (*InvoiceDetailResponseServiceresponseResponsepreamble, bool) {
	if o == nil || IsNil(o.Responsepreamble) {
		return nil, false
	}
	return o.Responsepreamble, true
}

// HasResponsepreamble returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponse) HasResponsepreamble() bool {
	if o != nil && !IsNil(o.Responsepreamble) {
		return true
	}

	return false
}

// SetResponsepreamble gets a reference to the given InvoiceDetailResponseServiceresponseResponsepreamble and assigns it to the Responsepreamble field.
func (o *OrderCreateResponseServiceresponse) SetResponsepreamble(v InvoiceDetailResponseServiceresponseResponsepreamble) {
	o.Responsepreamble = &v
}

// GetOrdersummary returns the Ordersummary field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponse) GetOrdersummary() OrderCreateResponseServiceresponseOrdersummary {
	if o == nil || IsNil(o.Ordersummary) {
		var ret OrderCreateResponseServiceresponseOrdersummary
		return ret
	}
	return *o.Ordersummary
}

// GetOrdersummaryOk returns a tuple with the Ordersummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponse) GetOrdersummaryOk() (*OrderCreateResponseServiceresponseOrdersummary, bool) {
	if o == nil || IsNil(o.Ordersummary) {
		return nil, false
	}
	return o.Ordersummary, true
}

// HasOrdersummary returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponse) HasOrdersummary() bool {
	if o != nil && !IsNil(o.Ordersummary) {
		return true
	}

	return false
}

// SetOrdersummary gets a reference to the given OrderCreateResponseServiceresponseOrdersummary and assigns it to the Ordersummary field.
func (o *OrderCreateResponseServiceresponse) SetOrdersummary(v OrderCreateResponseServiceresponseOrdersummary) {
	o.Ordersummary = &v
}

// GetOrdercreateresponse returns the Ordercreateresponse field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponse) GetOrdercreateresponse() []OrderCreateResponseServiceresponseOrdercreateresponseInner {
	if o == nil || IsNil(o.Ordercreateresponse) {
		var ret []OrderCreateResponseServiceresponseOrdercreateresponseInner
		return ret
	}
	return o.Ordercreateresponse
}

// GetOrdercreateresponseOk returns a tuple with the Ordercreateresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponse) GetOrdercreateresponseOk() ([]OrderCreateResponseServiceresponseOrdercreateresponseInner, bool) {
	if o == nil || IsNil(o.Ordercreateresponse) {
		return nil, false
	}
	return o.Ordercreateresponse, true
}

// HasOrdercreateresponse returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponse) HasOrdercreateresponse() bool {
	if o != nil && !IsNil(o.Ordercreateresponse) {
		return true
	}

	return false
}

// SetOrdercreateresponse gets a reference to the given []OrderCreateResponseServiceresponseOrdercreateresponseInner and assigns it to the Ordercreateresponse field.
func (o *OrderCreateResponseServiceresponse) SetOrdercreateresponse(v []OrderCreateResponseServiceresponseOrdercreateresponseInner) {
	o.Ordercreateresponse = v
}

func (o OrderCreateResponseServiceresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseServiceresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsepreamble) {
		toSerialize["responsepreamble"] = o.Responsepreamble
	}
	if !IsNil(o.Ordersummary) {
		toSerialize["ordersummary"] = o.Ordersummary
	}
	if !IsNil(o.Ordercreateresponse) {
		toSerialize["ordercreateresponse"] = o.Ordercreateresponse
	}
	return toSerialize, nil
}

type NullableOrderCreateResponseServiceresponse struct {
	value *OrderCreateResponseServiceresponse
	isSet bool
}

func (v NullableOrderCreateResponseServiceresponse) Get() *OrderCreateResponseServiceresponse {
	return v.value
}

func (v *NullableOrderCreateResponseServiceresponse) Set(val *OrderCreateResponseServiceresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseServiceresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseServiceresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseServiceresponse(val *OrderCreateResponseServiceresponse) *NullableOrderCreateResponseServiceresponse {
	return &NullableOrderCreateResponseServiceresponse{value: val, isSet: true}
}

func (v NullableOrderCreateResponseServiceresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseServiceresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


