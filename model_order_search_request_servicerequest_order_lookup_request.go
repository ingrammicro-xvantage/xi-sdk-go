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

// checks if the OrderSearchRequestServicerequestOrderLookupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchRequestServicerequestOrderLookupRequest{}

// OrderSearchRequestServicerequestOrderLookupRequest struct for OrderSearchRequestServicerequestOrderLookupRequest
type OrderSearchRequestServicerequestOrderLookupRequest struct {
	OrderNumber *OrderSearchRequestServicerequestOrderLookupRequestOrderNumber `json:"orderNumber,omitempty"`
	CustomerOrderNumber *OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber `json:"customerOrderNumber,omitempty"`
}

// NewOrderSearchRequestServicerequestOrderLookupRequest instantiates a new OrderSearchRequestServicerequestOrderLookupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchRequestServicerequestOrderLookupRequest() *OrderSearchRequestServicerequestOrderLookupRequest {
	this := OrderSearchRequestServicerequestOrderLookupRequest{}
	return &this
}

// NewOrderSearchRequestServicerequestOrderLookupRequestWithDefaults instantiates a new OrderSearchRequestServicerequestOrderLookupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchRequestServicerequestOrderLookupRequestWithDefaults() *OrderSearchRequestServicerequestOrderLookupRequest {
	this := OrderSearchRequestServicerequestOrderLookupRequest{}
	return &this
}

// GetOrderNumber returns the OrderNumber field value if set, zero value otherwise.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) GetOrderNumber() OrderSearchRequestServicerequestOrderLookupRequestOrderNumber {
	if o == nil || IsNil(o.OrderNumber) {
		var ret OrderSearchRequestServicerequestOrderLookupRequestOrderNumber
		return ret
	}
	return *o.OrderNumber
}

// GetOrderNumberOk returns a tuple with the OrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) GetOrderNumberOk() (*OrderSearchRequestServicerequestOrderLookupRequestOrderNumber, bool) {
	if o == nil || IsNil(o.OrderNumber) {
		return nil, false
	}
	return o.OrderNumber, true
}

// HasOrderNumber returns a boolean if a field has been set.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) HasOrderNumber() bool {
	if o != nil && !IsNil(o.OrderNumber) {
		return true
	}

	return false
}

// SetOrderNumber gets a reference to the given OrderSearchRequestServicerequestOrderLookupRequestOrderNumber and assigns it to the OrderNumber field.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) SetOrderNumber(v OrderSearchRequestServicerequestOrderLookupRequestOrderNumber) {
	o.OrderNumber = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) GetCustomerOrderNumber() OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) GetCustomerOrderNumberOk() (*OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber and assigns it to the CustomerOrderNumber field.
func (o *OrderSearchRequestServicerequestOrderLookupRequest) SetCustomerOrderNumber(v OrderSearchRequestServicerequestOrderLookupRequestCustomerOrderNumber) {
	o.CustomerOrderNumber = &v
}

func (o OrderSearchRequestServicerequestOrderLookupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchRequestServicerequestOrderLookupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderNumber) {
		toSerialize["orderNumber"] = o.OrderNumber
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	return toSerialize, nil
}

type NullableOrderSearchRequestServicerequestOrderLookupRequest struct {
	value *OrderSearchRequestServicerequestOrderLookupRequest
	isSet bool
}

func (v NullableOrderSearchRequestServicerequestOrderLookupRequest) Get() *OrderSearchRequestServicerequestOrderLookupRequest {
	return v.value
}

func (v *NullableOrderSearchRequestServicerequestOrderLookupRequest) Set(val *OrderSearchRequestServicerequestOrderLookupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchRequestServicerequestOrderLookupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchRequestServicerequestOrderLookupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchRequestServicerequestOrderLookupRequest(val *OrderSearchRequestServicerequestOrderLookupRequest) *NullableOrderSearchRequestServicerequestOrderLookupRequest {
	return &NullableOrderSearchRequestServicerequestOrderLookupRequest{value: val, isSet: true}
}

func (v NullableOrderSearchRequestServicerequestOrderLookupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchRequestServicerequestOrderLookupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


