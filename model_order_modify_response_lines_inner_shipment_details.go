/*
XI Sdk Resellers

For Resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderModifyResponseLinesInnerShipmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyResponseLinesInnerShipmentDetails{}

// OrderModifyResponseLinesInnerShipmentDetails Shipping details for the order provided by the reseller.
type OrderModifyResponseLinesInnerShipmentDetails struct {
	// The carrier code for the shipment containing the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the carrier of the shipment containing the line item.
	CarrierName *string `json:"carrierName,omitempty"`
	// The reseller's shipping account number with carrier. Used to bill the shipping carrier directly from the reseller's account with the carrier.
	FreightAccountNumber *string `json:"freightAccountNumber,omitempty"`
}

// NewOrderModifyResponseLinesInnerShipmentDetails instantiates a new OrderModifyResponseLinesInnerShipmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyResponseLinesInnerShipmentDetails() *OrderModifyResponseLinesInnerShipmentDetails {
	this := OrderModifyResponseLinesInnerShipmentDetails{}
	return &this
}

// NewOrderModifyResponseLinesInnerShipmentDetailsWithDefaults instantiates a new OrderModifyResponseLinesInnerShipmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyResponseLinesInnerShipmentDetailsWithDefaults() *OrderModifyResponseLinesInnerShipmentDetails {
	this := OrderModifyResponseLinesInnerShipmentDetails{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInnerShipmentDetails) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInnerShipmentDetails) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInnerShipmentDetails) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *OrderModifyResponseLinesInnerShipmentDetails) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInnerShipmentDetails) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInnerShipmentDetails) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInnerShipmentDetails) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *OrderModifyResponseLinesInnerShipmentDetails) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetFreightAccountNumber returns the FreightAccountNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInnerShipmentDetails) GetFreightAccountNumber() string {
	if o == nil || IsNil(o.FreightAccountNumber) {
		var ret string
		return ret
	}
	return *o.FreightAccountNumber
}

// GetFreightAccountNumberOk returns a tuple with the FreightAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInnerShipmentDetails) GetFreightAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FreightAccountNumber) {
		return nil, false
	}
	return o.FreightAccountNumber, true
}

// HasFreightAccountNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInnerShipmentDetails) HasFreightAccountNumber() bool {
	if o != nil && !IsNil(o.FreightAccountNumber) {
		return true
	}

	return false
}

// SetFreightAccountNumber gets a reference to the given string and assigns it to the FreightAccountNumber field.
func (o *OrderModifyResponseLinesInnerShipmentDetails) SetFreightAccountNumber(v string) {
	o.FreightAccountNumber = &v
}

func (o OrderModifyResponseLinesInnerShipmentDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyResponseLinesInnerShipmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.FreightAccountNumber) {
		toSerialize["freightAccountNumber"] = o.FreightAccountNumber
	}
	return toSerialize, nil
}

type NullableOrderModifyResponseLinesInnerShipmentDetails struct {
	value *OrderModifyResponseLinesInnerShipmentDetails
	isSet bool
}

func (v NullableOrderModifyResponseLinesInnerShipmentDetails) Get() *OrderModifyResponseLinesInnerShipmentDetails {
	return v.value
}

func (v *NullableOrderModifyResponseLinesInnerShipmentDetails) Set(val *OrderModifyResponseLinesInnerShipmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyResponseLinesInnerShipmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyResponseLinesInnerShipmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyResponseLinesInnerShipmentDetails(val *OrderModifyResponseLinesInnerShipmentDetails) *NullableOrderModifyResponseLinesInnerShipmentDetails {
	return &NullableOrderModifyResponseLinesInnerShipmentDetails{value: val, isSet: true}
}

func (v NullableOrderModifyResponseLinesInnerShipmentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyResponseLinesInnerShipmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


