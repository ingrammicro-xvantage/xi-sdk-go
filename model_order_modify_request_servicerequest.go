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

// checks if the OrderModifyRequestServicerequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyRequestServicerequest{}

// OrderModifyRequestServicerequest struct for OrderModifyRequestServicerequest
type OrderModifyRequestServicerequest struct {
	Requestpreamble *OrderModifyRequestServicerequestRequestpreamble `json:"requestpreamble,omitempty"`
	Ordermodifyrequest *OrderModifyRequestServicerequestOrdermodifyrequest `json:"ordermodifyrequest,omitempty"`
}

// NewOrderModifyRequestServicerequest instantiates a new OrderModifyRequestServicerequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyRequestServicerequest() *OrderModifyRequestServicerequest {
	this := OrderModifyRequestServicerequest{}
	return &this
}

// NewOrderModifyRequestServicerequestWithDefaults instantiates a new OrderModifyRequestServicerequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyRequestServicerequestWithDefaults() *OrderModifyRequestServicerequest {
	this := OrderModifyRequestServicerequest{}
	return &this
}

// GetRequestpreamble returns the Requestpreamble field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequest) GetRequestpreamble() OrderModifyRequestServicerequestRequestpreamble {
	if o == nil || IsNil(o.Requestpreamble) {
		var ret OrderModifyRequestServicerequestRequestpreamble
		return ret
	}
	return *o.Requestpreamble
}

// GetRequestpreambleOk returns a tuple with the Requestpreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequest) GetRequestpreambleOk() (*OrderModifyRequestServicerequestRequestpreamble, bool) {
	if o == nil || IsNil(o.Requestpreamble) {
		return nil, false
	}
	return o.Requestpreamble, true
}

// HasRequestpreamble returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequest) HasRequestpreamble() bool {
	if o != nil && !IsNil(o.Requestpreamble) {
		return true
	}

	return false
}

// SetRequestpreamble gets a reference to the given OrderModifyRequestServicerequestRequestpreamble and assigns it to the Requestpreamble field.
func (o *OrderModifyRequestServicerequest) SetRequestpreamble(v OrderModifyRequestServicerequestRequestpreamble) {
	o.Requestpreamble = &v
}

// GetOrdermodifyrequest returns the Ordermodifyrequest field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequest) GetOrdermodifyrequest() OrderModifyRequestServicerequestOrdermodifyrequest {
	if o == nil || IsNil(o.Ordermodifyrequest) {
		var ret OrderModifyRequestServicerequestOrdermodifyrequest
		return ret
	}
	return *o.Ordermodifyrequest
}

// GetOrdermodifyrequestOk returns a tuple with the Ordermodifyrequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequest) GetOrdermodifyrequestOk() (*OrderModifyRequestServicerequestOrdermodifyrequest, bool) {
	if o == nil || IsNil(o.Ordermodifyrequest) {
		return nil, false
	}
	return o.Ordermodifyrequest, true
}

// HasOrdermodifyrequest returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequest) HasOrdermodifyrequest() bool {
	if o != nil && !IsNil(o.Ordermodifyrequest) {
		return true
	}

	return false
}

// SetOrdermodifyrequest gets a reference to the given OrderModifyRequestServicerequestOrdermodifyrequest and assigns it to the Ordermodifyrequest field.
func (o *OrderModifyRequestServicerequest) SetOrdermodifyrequest(v OrderModifyRequestServicerequestOrdermodifyrequest) {
	o.Ordermodifyrequest = &v
}

func (o OrderModifyRequestServicerequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyRequestServicerequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requestpreamble) {
		toSerialize["requestpreamble"] = o.Requestpreamble
	}
	if !IsNil(o.Ordermodifyrequest) {
		toSerialize["ordermodifyrequest"] = o.Ordermodifyrequest
	}
	return toSerialize, nil
}

type NullableOrderModifyRequestServicerequest struct {
	value *OrderModifyRequestServicerequest
	isSet bool
}

func (v NullableOrderModifyRequestServicerequest) Get() *OrderModifyRequestServicerequest {
	return v.value
}

func (v *NullableOrderModifyRequestServicerequest) Set(val *OrderModifyRequestServicerequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyRequestServicerequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyRequestServicerequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyRequestServicerequest(val *OrderModifyRequestServicerequest) *NullableOrderModifyRequestServicerequest {
	return &NullableOrderModifyRequestServicerequest{value: val, isSet: true}
}

func (v NullableOrderModifyRequestServicerequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyRequestServicerequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


