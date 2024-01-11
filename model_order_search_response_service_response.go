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

// checks if the OrderSearchResponseServiceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchResponseServiceResponse{}

// OrderSearchResponseServiceResponse struct for OrderSearchResponseServiceResponse
type OrderSearchResponseServiceResponse struct {
	Responsepreamble *OrderSearchResponseServiceResponseResponsepreamble `json:"responsepreamble,omitempty"`
	Ordersearchresponse *OrderSearchResponseServiceResponseOrdersearchresponse `json:"ordersearchresponse,omitempty"`
}

// NewOrderSearchResponseServiceResponse instantiates a new OrderSearchResponseServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchResponseServiceResponse() *OrderSearchResponseServiceResponse {
	this := OrderSearchResponseServiceResponse{}
	return &this
}

// NewOrderSearchResponseServiceResponseWithDefaults instantiates a new OrderSearchResponseServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchResponseServiceResponseWithDefaults() *OrderSearchResponseServiceResponse {
	this := OrderSearchResponseServiceResponse{}
	return &this
}

// GetResponsepreamble returns the Responsepreamble field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponse) GetResponsepreamble() OrderSearchResponseServiceResponseResponsepreamble {
	if o == nil || IsNil(o.Responsepreamble) {
		var ret OrderSearchResponseServiceResponseResponsepreamble
		return ret
	}
	return *o.Responsepreamble
}

// GetResponsepreambleOk returns a tuple with the Responsepreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponse) GetResponsepreambleOk() (*OrderSearchResponseServiceResponseResponsepreamble, bool) {
	if o == nil || IsNil(o.Responsepreamble) {
		return nil, false
	}
	return o.Responsepreamble, true
}

// HasResponsepreamble returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponse) HasResponsepreamble() bool {
	if o != nil && !IsNil(o.Responsepreamble) {
		return true
	}

	return false
}

// SetResponsepreamble gets a reference to the given OrderSearchResponseServiceResponseResponsepreamble and assigns it to the Responsepreamble field.
func (o *OrderSearchResponseServiceResponse) SetResponsepreamble(v OrderSearchResponseServiceResponseResponsepreamble) {
	o.Responsepreamble = &v
}

// GetOrdersearchresponse returns the Ordersearchresponse field value if set, zero value otherwise.
func (o *OrderSearchResponseServiceResponse) GetOrdersearchresponse() OrderSearchResponseServiceResponseOrdersearchresponse {
	if o == nil || IsNil(o.Ordersearchresponse) {
		var ret OrderSearchResponseServiceResponseOrdersearchresponse
		return ret
	}
	return *o.Ordersearchresponse
}

// GetOrdersearchresponseOk returns a tuple with the Ordersearchresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponseServiceResponse) GetOrdersearchresponseOk() (*OrderSearchResponseServiceResponseOrdersearchresponse, bool) {
	if o == nil || IsNil(o.Ordersearchresponse) {
		return nil, false
	}
	return o.Ordersearchresponse, true
}

// HasOrdersearchresponse returns a boolean if a field has been set.
func (o *OrderSearchResponseServiceResponse) HasOrdersearchresponse() bool {
	if o != nil && !IsNil(o.Ordersearchresponse) {
		return true
	}

	return false
}

// SetOrdersearchresponse gets a reference to the given OrderSearchResponseServiceResponseOrdersearchresponse and assigns it to the Ordersearchresponse field.
func (o *OrderSearchResponseServiceResponse) SetOrdersearchresponse(v OrderSearchResponseServiceResponseOrdersearchresponse) {
	o.Ordersearchresponse = &v
}

func (o OrderSearchResponseServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchResponseServiceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Responsepreamble) {
		toSerialize["responsepreamble"] = o.Responsepreamble
	}
	if !IsNil(o.Ordersearchresponse) {
		toSerialize["ordersearchresponse"] = o.Ordersearchresponse
	}
	return toSerialize, nil
}

type NullableOrderSearchResponseServiceResponse struct {
	value *OrderSearchResponseServiceResponse
	isSet bool
}

func (v NullableOrderSearchResponseServiceResponse) Get() *OrderSearchResponseServiceResponse {
	return v.value
}

func (v *NullableOrderSearchResponseServiceResponse) Set(val *OrderSearchResponseServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchResponseServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchResponseServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchResponseServiceResponse(val *OrderSearchResponseServiceResponse) *NullableOrderSearchResponseServiceResponse {
	return &NullableOrderSearchResponseServiceResponse{value: val, isSet: true}
}

func (v NullableOrderSearchResponseServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchResponseServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


