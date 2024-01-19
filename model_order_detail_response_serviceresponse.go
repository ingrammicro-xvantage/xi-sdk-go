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

// checks if the OrderDetailResponseServiceresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseServiceresponse{}

// OrderDetailResponseServiceresponse struct for OrderDetailResponseServiceresponse
type OrderDetailResponseServiceresponse struct {
	Responsepreamble *InvoiceDetailResponseServiceresponseResponsepreamble `json:"responsepreamble,omitempty"`
	Orderdetailresponse *OrderDetailResponseServiceresponseOrderdetailresponse `json:"orderdetailresponse,omitempty"`
}

// NewOrderDetailResponseServiceresponse instantiates a new OrderDetailResponseServiceresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseServiceresponse() *OrderDetailResponseServiceresponse {
	this := OrderDetailResponseServiceresponse{}
	return &this
}

// NewOrderDetailResponseServiceresponseWithDefaults instantiates a new OrderDetailResponseServiceresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseServiceresponseWithDefaults() *OrderDetailResponseServiceresponse {
	this := OrderDetailResponseServiceresponse{}
	return &this
}

// GetResponsepreamble returns the Responsepreamble field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponse) GetResponsepreamble() InvoiceDetailResponseServiceresponseResponsepreamble {
	if o == nil || IsNil(o.Responsepreamble) {
		var ret InvoiceDetailResponseServiceresponseResponsepreamble
		return ret
	}
	return *o.Responsepreamble
}

// GetResponsepreambleOk returns a tuple with the Responsepreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponse) GetResponsepreambleOk() (*InvoiceDetailResponseServiceresponseResponsepreamble, bool) {
	if o == nil || IsNil(o.Responsepreamble) {
		return nil, false
	}
	return o.Responsepreamble, true
}

// HasResponsepreamble returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponse) HasResponsepreamble() bool {
	if o != nil && !IsNil(o.Responsepreamble) {
		return true
	}

	return false
}

// SetResponsepreamble gets a reference to the given InvoiceDetailResponseServiceresponseResponsepreamble and assigns it to the Responsepreamble field.
func (o *OrderDetailResponseServiceresponse) SetResponsepreamble(v InvoiceDetailResponseServiceresponseResponsepreamble) {
	o.Responsepreamble = &v
}

// GetOrderdetailresponse returns the Orderdetailresponse field value if set, zero value otherwise.
func (o *OrderDetailResponseServiceresponse) GetOrderdetailresponse() OrderDetailResponseServiceresponseOrderdetailresponse {
	if o == nil || IsNil(o.Orderdetailresponse) {
		var ret OrderDetailResponseServiceresponseOrderdetailresponse
		return ret
	}
	return *o.Orderdetailresponse
}

// GetOrderdetailresponseOk returns a tuple with the Orderdetailresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseServiceresponse) GetOrderdetailresponseOk() (*OrderDetailResponseServiceresponseOrderdetailresponse, bool) {
	if o == nil || IsNil(o.Orderdetailresponse) {
		return nil, false
	}
	return o.Orderdetailresponse, true
}

// HasOrderdetailresponse returns a boolean if a field has been set.
func (o *OrderDetailResponseServiceresponse) HasOrderdetailresponse() bool {
	if o != nil && !IsNil(o.Orderdetailresponse) {
		return true
	}

	return false
}

// SetOrderdetailresponse gets a reference to the given OrderDetailResponseServiceresponseOrderdetailresponse and assigns it to the Orderdetailresponse field.
func (o *OrderDetailResponseServiceresponse) SetOrderdetailresponse(v OrderDetailResponseServiceresponseOrderdetailresponse) {
	o.Orderdetailresponse = &v
}

func (o OrderDetailResponseServiceresponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseServiceresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsepreamble) {
		toSerialize["responsepreamble"] = o.Responsepreamble
	}
	if !IsNil(o.Orderdetailresponse) {
		toSerialize["orderdetailresponse"] = o.Orderdetailresponse
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseServiceresponse struct {
	value *OrderDetailResponseServiceresponse
	isSet bool
}

func (v NullableOrderDetailResponseServiceresponse) Get() *OrderDetailResponseServiceresponse {
	return v.value
}

func (v *NullableOrderDetailResponseServiceresponse) Set(val *OrderDetailResponseServiceresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseServiceresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseServiceresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseServiceresponse(val *OrderDetailResponseServiceresponse) *NullableOrderDetailResponseServiceresponse {
	return &NullableOrderDetailResponseServiceresponse{value: val, isSet: true}
}

func (v NullableOrderDetailResponseServiceresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseServiceresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


