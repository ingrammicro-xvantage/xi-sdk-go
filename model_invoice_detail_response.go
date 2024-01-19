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

// checks if the InvoiceDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailResponse{}

// InvoiceDetailResponse struct for InvoiceDetailResponse
type InvoiceDetailResponse struct {
	Serviceresponse *InvoiceDetailResponseServiceresponse `json:"serviceresponse,omitempty"`
}

// NewInvoiceDetailResponse instantiates a new InvoiceDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailResponse() *InvoiceDetailResponse {
	this := InvoiceDetailResponse{}
	return &this
}

// NewInvoiceDetailResponseWithDefaults instantiates a new InvoiceDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailResponseWithDefaults() *InvoiceDetailResponse {
	this := InvoiceDetailResponse{}
	return &this
}

// GetServiceresponse returns the Serviceresponse field value if set, zero value otherwise.
func (o *InvoiceDetailResponse) GetServiceresponse() InvoiceDetailResponseServiceresponse {
	if o == nil || IsNil(o.Serviceresponse) {
		var ret InvoiceDetailResponseServiceresponse
		return ret
	}
	return *o.Serviceresponse
}

// GetServiceresponseOk returns a tuple with the Serviceresponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailResponse) GetServiceresponseOk() (*InvoiceDetailResponseServiceresponse, bool) {
	if o == nil || IsNil(o.Serviceresponse) {
		return nil, false
	}
	return o.Serviceresponse, true
}

// HasServiceresponse returns a boolean if a field has been set.
func (o *InvoiceDetailResponse) HasServiceresponse() bool {
	if o != nil && !IsNil(o.Serviceresponse) {
		return true
	}

	return false
}

// SetServiceresponse gets a reference to the given InvoiceDetailResponseServiceresponse and assigns it to the Serviceresponse field.
func (o *InvoiceDetailResponse) SetServiceresponse(v InvoiceDetailResponseServiceresponse) {
	o.Serviceresponse = &v
}

func (o InvoiceDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serviceresponse) {
		toSerialize["serviceresponse"] = o.Serviceresponse
	}
	return toSerialize, nil
}

type NullableInvoiceDetailResponse struct {
	value *InvoiceDetailResponse
	isSet bool
}

func (v NullableInvoiceDetailResponse) Get() *InvoiceDetailResponse {
	return v.value
}

func (v *NullableInvoiceDetailResponse) Set(val *InvoiceDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailResponse(val *InvoiceDetailResponse) *NullableInvoiceDetailResponse {
	return &NullableInvoiceDetailResponse{value: val, isSet: true}
}

func (v NullableInvoiceDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


