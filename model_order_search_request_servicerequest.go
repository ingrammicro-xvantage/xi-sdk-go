/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderSearchRequestServicerequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchRequestServicerequest{}

// OrderSearchRequestServicerequest struct for OrderSearchRequestServicerequest
type OrderSearchRequestServicerequest struct {
	Requestpreamble OrderSearchRequestServicerequestRequestpreamble `json:"requestpreamble"`
	OrderLookupRequest *OrderSearchRequestServicerequestOrderLookupRequest `json:"orderLookupRequest,omitempty"`
}

type _OrderSearchRequestServicerequest OrderSearchRequestServicerequest

// NewOrderSearchRequestServicerequest instantiates a new OrderSearchRequestServicerequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchRequestServicerequest(requestpreamble OrderSearchRequestServicerequestRequestpreamble) *OrderSearchRequestServicerequest {
	this := OrderSearchRequestServicerequest{}
	this.Requestpreamble = requestpreamble
	return &this
}

// NewOrderSearchRequestServicerequestWithDefaults instantiates a new OrderSearchRequestServicerequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchRequestServicerequestWithDefaults() *OrderSearchRequestServicerequest {
	this := OrderSearchRequestServicerequest{}
	return &this
}

// GetRequestpreamble returns the Requestpreamble field value
func (o *OrderSearchRequestServicerequest) GetRequestpreamble() OrderSearchRequestServicerequestRequestpreamble {
	if o == nil {
		var ret OrderSearchRequestServicerequestRequestpreamble
		return ret
	}

	return o.Requestpreamble
}

// GetRequestpreambleOk returns a tuple with the Requestpreamble field value
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequest) GetRequestpreambleOk() (*OrderSearchRequestServicerequestRequestpreamble, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requestpreamble, true
}

// SetRequestpreamble sets field value
func (o *OrderSearchRequestServicerequest) SetRequestpreamble(v OrderSearchRequestServicerequestRequestpreamble) {
	o.Requestpreamble = v
}

// GetOrderLookupRequest returns the OrderLookupRequest field value if set, zero value otherwise.
func (o *OrderSearchRequestServicerequest) GetOrderLookupRequest() OrderSearchRequestServicerequestOrderLookupRequest {
	if o == nil || IsNil(o.OrderLookupRequest) {
		var ret OrderSearchRequestServicerequestOrderLookupRequest
		return ret
	}
	return *o.OrderLookupRequest
}

// GetOrderLookupRequestOk returns a tuple with the OrderLookupRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequest) GetOrderLookupRequestOk() (*OrderSearchRequestServicerequestOrderLookupRequest, bool) {
	if o == nil || IsNil(o.OrderLookupRequest) {
		return nil, false
	}
	return o.OrderLookupRequest, true
}

// HasOrderLookupRequest returns a boolean if a field has been set.
func (o *OrderSearchRequestServicerequest) HasOrderLookupRequest() bool {
	if o != nil && !IsNil(o.OrderLookupRequest) {
		return true
	}

	return false
}

// SetOrderLookupRequest gets a reference to the given OrderSearchRequestServicerequestOrderLookupRequest and assigns it to the OrderLookupRequest field.
func (o *OrderSearchRequestServicerequest) SetOrderLookupRequest(v OrderSearchRequestServicerequestOrderLookupRequest) {
	o.OrderLookupRequest = &v
}

func (o OrderSearchRequestServicerequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchRequestServicerequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["requestpreamble"] = o.Requestpreamble
	if !IsNil(o.OrderLookupRequest) {
		toSerialize["orderLookupRequest"] = o.OrderLookupRequest
	}
	return toSerialize, nil
}

func (o *OrderSearchRequestServicerequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requestpreamble",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderSearchRequestServicerequest := _OrderSearchRequestServicerequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderSearchRequestServicerequest)

	if err != nil {
		return err
	}

	*o = OrderSearchRequestServicerequest(varOrderSearchRequestServicerequest)

	return err
}

type NullableOrderSearchRequestServicerequest struct {
	value *OrderSearchRequestServicerequest
	isSet bool
}

func (v NullableOrderSearchRequestServicerequest) Get() *OrderSearchRequestServicerequest {
	return v.value
}

func (v *NullableOrderSearchRequestServicerequest) Set(val *OrderSearchRequestServicerequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchRequestServicerequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchRequestServicerequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchRequestServicerequest(val *OrderSearchRequestServicerequest) *NullableOrderSearchRequestServicerequest {
	return &NullableOrderSearchRequestServicerequest{value: val, isSet: true}
}

func (v NullableOrderSearchRequestServicerequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchRequestServicerequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


