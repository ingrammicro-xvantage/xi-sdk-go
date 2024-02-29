/*
XI Sdk Resellers

For Ingram Micro Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the InvoiceDetailsv61ResponsePaymentTermsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailsv61ResponsePaymentTermsInfo{}

// InvoiceDetailsv61ResponsePaymentTermsInfo Payment terms is the agreement between Ingram and the customer by what period they should pay the invoice by
type InvoiceDetailsv61ResponsePaymentTermsInfo struct {
	// Code of the payment terms.
	PaymentTermsCode *string `json:"paymentTermsCode,omitempty"`
	// Description of the payment terms.
	PaymentTermsDescription *string `json:"paymentTermsDescription,omitempty"`
	// Due date of the payment terms.
	PaymentTermsDueDate *string `json:"paymentTermsDueDate,omitempty"`
}

// NewInvoiceDetailsv61ResponsePaymentTermsInfo instantiates a new InvoiceDetailsv61ResponsePaymentTermsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailsv61ResponsePaymentTermsInfo() *InvoiceDetailsv61ResponsePaymentTermsInfo {
	this := InvoiceDetailsv61ResponsePaymentTermsInfo{}
	return &this
}

// NewInvoiceDetailsv61ResponsePaymentTermsInfoWithDefaults instantiates a new InvoiceDetailsv61ResponsePaymentTermsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailsv61ResponsePaymentTermsInfoWithDefaults() *InvoiceDetailsv61ResponsePaymentTermsInfo {
	this := InvoiceDetailsv61ResponsePaymentTermsInfo{}
	return &this
}

// GetPaymentTermsCode returns the PaymentTermsCode field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) GetPaymentTermsCode() string {
	if o == nil || IsNil(o.PaymentTermsCode) {
		var ret string
		return ret
	}
	return *o.PaymentTermsCode
}

// GetPaymentTermsCodeOk returns a tuple with the PaymentTermsCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) GetPaymentTermsCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTermsCode) {
		return nil, false
	}
	return o.PaymentTermsCode, true
}

// HasPaymentTermsCode returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) HasPaymentTermsCode() bool {
	if o != nil && !IsNil(o.PaymentTermsCode) {
		return true
	}

	return false
}

// SetPaymentTermsCode gets a reference to the given string and assigns it to the PaymentTermsCode field.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) SetPaymentTermsCode(v string) {
	o.PaymentTermsCode = &v
}

// GetPaymentTermsDescription returns the PaymentTermsDescription field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) GetPaymentTermsDescription() string {
	if o == nil || IsNil(o.PaymentTermsDescription) {
		var ret string
		return ret
	}
	return *o.PaymentTermsDescription
}

// GetPaymentTermsDescriptionOk returns a tuple with the PaymentTermsDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) GetPaymentTermsDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTermsDescription) {
		return nil, false
	}
	return o.PaymentTermsDescription, true
}

// HasPaymentTermsDescription returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) HasPaymentTermsDescription() bool {
	if o != nil && !IsNil(o.PaymentTermsDescription) {
		return true
	}

	return false
}

// SetPaymentTermsDescription gets a reference to the given string and assigns it to the PaymentTermsDescription field.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) SetPaymentTermsDescription(v string) {
	o.PaymentTermsDescription = &v
}

// GetPaymentTermsDueDate returns the PaymentTermsDueDate field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) GetPaymentTermsDueDate() string {
	if o == nil || IsNil(o.PaymentTermsDueDate) {
		var ret string
		return ret
	}
	return *o.PaymentTermsDueDate
}

// GetPaymentTermsDueDateOk returns a tuple with the PaymentTermsDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) GetPaymentTermsDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTermsDueDate) {
		return nil, false
	}
	return o.PaymentTermsDueDate, true
}

// HasPaymentTermsDueDate returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) HasPaymentTermsDueDate() bool {
	if o != nil && !IsNil(o.PaymentTermsDueDate) {
		return true
	}

	return false
}

// SetPaymentTermsDueDate gets a reference to the given string and assigns it to the PaymentTermsDueDate field.
func (o *InvoiceDetailsv61ResponsePaymentTermsInfo) SetPaymentTermsDueDate(v string) {
	o.PaymentTermsDueDate = &v
}

func (o InvoiceDetailsv61ResponsePaymentTermsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailsv61ResponsePaymentTermsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentTermsCode) {
		toSerialize["paymentTermsCode"] = o.PaymentTermsCode
	}
	if !IsNil(o.PaymentTermsDescription) {
		toSerialize["paymentTermsDescription"] = o.PaymentTermsDescription
	}
	if !IsNil(o.PaymentTermsDueDate) {
		toSerialize["paymentTermsDueDate"] = o.PaymentTermsDueDate
	}
	return toSerialize, nil
}

type NullableInvoiceDetailsv61ResponsePaymentTermsInfo struct {
	value *InvoiceDetailsv61ResponsePaymentTermsInfo
	isSet bool
}

func (v NullableInvoiceDetailsv61ResponsePaymentTermsInfo) Get() *InvoiceDetailsv61ResponsePaymentTermsInfo {
	return v.value
}

func (v *NullableInvoiceDetailsv61ResponsePaymentTermsInfo) Set(val *InvoiceDetailsv61ResponsePaymentTermsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailsv61ResponsePaymentTermsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailsv61ResponsePaymentTermsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailsv61ResponsePaymentTermsInfo(val *InvoiceDetailsv61ResponsePaymentTermsInfo) *NullableInvoiceDetailsv61ResponsePaymentTermsInfo {
	return &NullableInvoiceDetailsv61ResponsePaymentTermsInfo{value: val, isSet: true}
}

func (v NullableInvoiceDetailsv61ResponsePaymentTermsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailsv61ResponsePaymentTermsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


