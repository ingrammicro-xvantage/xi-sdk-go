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

// checks if the OrderModifyResponseRejectedLineItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyResponseRejectedLineItemsInner{}

// OrderModifyResponseRejectedLineItemsInner struct for OrderModifyResponseRejectedLineItemsInner
type OrderModifyResponseRejectedLineItemsInner struct {
	// The IngramMicro line number for the failed line item.
	IngramLineNumber *string `json:"ingramLineNumber,omitempty"`
	// The reseller's line number of the failed line item for reference in their system.
	CustomerLineNumber *string `json:"customerLineNumber,omitempty"`
	// The IngramMicro part number for the failed line item.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The vendor's part number for the failed line item.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// The quantity ordered of the failed line item.
	QuantityOrdered *int32 `json:"quantityOrdered,omitempty"`
	// The rejection code for the failed line item.
	RejectCode *string `json:"rejectCode,omitempty"`
	// The rejection reason for the failed line item.
	RejectReason *string `json:"rejectReason,omitempty"`
}

// NewOrderModifyResponseRejectedLineItemsInner instantiates a new OrderModifyResponseRejectedLineItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyResponseRejectedLineItemsInner() *OrderModifyResponseRejectedLineItemsInner {
	this := OrderModifyResponseRejectedLineItemsInner{}
	return &this
}

// NewOrderModifyResponseRejectedLineItemsInnerWithDefaults instantiates a new OrderModifyResponseRejectedLineItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyResponseRejectedLineItemsInnerWithDefaults() *OrderModifyResponseRejectedLineItemsInner {
	this := OrderModifyResponseRejectedLineItemsInner{}
	return &this
}

// GetIngramLineNumber returns the IngramLineNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramLineNumber() string {
	if o == nil || IsNil(o.IngramLineNumber) {
		var ret string
		return ret
	}
	return *o.IngramLineNumber
}

// GetIngramLineNumberOk returns a tuple with the IngramLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramLineNumber) {
		return nil, false
	}
	return o.IngramLineNumber, true
}

// HasIngramLineNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasIngramLineNumber() bool {
	if o != nil && !IsNil(o.IngramLineNumber) {
		return true
	}

	return false
}

// SetIngramLineNumber gets a reference to the given string and assigns it to the IngramLineNumber field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetIngramLineNumber(v string) {
	o.IngramLineNumber = &v
}

// GetCustomerLineNumber returns the CustomerLineNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetCustomerLineNumber() string {
	if o == nil || IsNil(o.CustomerLineNumber) {
		var ret string
		return ret
	}
	return *o.CustomerLineNumber
}

// GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetCustomerLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerLineNumber) {
		return nil, false
	}
	return o.CustomerLineNumber, true
}

// HasCustomerLineNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasCustomerLineNumber() bool {
	if o != nil && !IsNil(o.CustomerLineNumber) {
		return true
	}

	return false
}

// SetCustomerLineNumber gets a reference to the given string and assigns it to the CustomerLineNumber field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetCustomerLineNumber(v string) {
	o.CustomerLineNumber = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetQuantityOrdered() int32 {
	if o == nil || IsNil(o.QuantityOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityOrdered) {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasQuantityOrdered() bool {
	if o != nil && !IsNil(o.QuantityOrdered) {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given int32 and assigns it to the QuantityOrdered field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = &v
}

// GetRejectCode returns the RejectCode field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectCode() string {
	if o == nil || IsNil(o.RejectCode) {
		var ret string
		return ret
	}
	return *o.RejectCode
}

// GetRejectCodeOk returns a tuple with the RejectCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectCodeOk() (*string, bool) {
	if o == nil || IsNil(o.RejectCode) {
		return nil, false
	}
	return o.RejectCode, true
}

// HasRejectCode returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasRejectCode() bool {
	if o != nil && !IsNil(o.RejectCode) {
		return true
	}

	return false
}

// SetRejectCode gets a reference to the given string and assigns it to the RejectCode field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetRejectCode(v string) {
	o.RejectCode = &v
}

// GetRejectReason returns the RejectReason field value if set, zero value otherwise.
func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectReason() string {
	if o == nil || IsNil(o.RejectReason) {
		var ret string
		return ret
	}
	return *o.RejectReason
}

// GetRejectReasonOk returns a tuple with the RejectReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) GetRejectReasonOk() (*string, bool) {
	if o == nil || IsNil(o.RejectReason) {
		return nil, false
	}
	return o.RejectReason, true
}

// HasRejectReason returns a boolean if a field has been set.
func (o *OrderModifyResponseRejectedLineItemsInner) HasRejectReason() bool {
	if o != nil && !IsNil(o.RejectReason) {
		return true
	}

	return false
}

// SetRejectReason gets a reference to the given string and assigns it to the RejectReason field.
func (o *OrderModifyResponseRejectedLineItemsInner) SetRejectReason(v string) {
	o.RejectReason = &v
}

func (o OrderModifyResponseRejectedLineItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyResponseRejectedLineItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramLineNumber) {
		toSerialize["ingramLineNumber"] = o.IngramLineNumber
	}
	if !IsNil(o.CustomerLineNumber) {
		toSerialize["customerLineNumber"] = o.CustomerLineNumber
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.QuantityOrdered) {
		toSerialize["quantityOrdered"] = o.QuantityOrdered
	}
	if !IsNil(o.RejectCode) {
		toSerialize["rejectCode"] = o.RejectCode
	}
	if !IsNil(o.RejectReason) {
		toSerialize["rejectReason"] = o.RejectReason
	}
	return toSerialize, nil
}

type NullableOrderModifyResponseRejectedLineItemsInner struct {
	value *OrderModifyResponseRejectedLineItemsInner
	isSet bool
}

func (v NullableOrderModifyResponseRejectedLineItemsInner) Get() *OrderModifyResponseRejectedLineItemsInner {
	return v.value
}

func (v *NullableOrderModifyResponseRejectedLineItemsInner) Set(val *OrderModifyResponseRejectedLineItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyResponseRejectedLineItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyResponseRejectedLineItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyResponseRejectedLineItemsInner(val *OrderModifyResponseRejectedLineItemsInner) *NullableOrderModifyResponseRejectedLineItemsInner {
	return &NullableOrderModifyResponseRejectedLineItemsInner{value: val, isSet: true}
}

func (v NullableOrderModifyResponseRejectedLineItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyResponseRejectedLineItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


