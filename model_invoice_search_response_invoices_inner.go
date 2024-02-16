/*
XI Sdk Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the InvoiceSearchResponseInvoicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceSearchResponseInvoicesInner{}

// InvoiceSearchResponseInvoicesInner struct for InvoiceSearchResponseInvoicesInner
type InvoiceSearchResponseInvoicesInner struct {
	// Payment Terms Due date.
	PaymentTermsDueDate *string `json:"paymentTermsDueDate,omitempty"`
	// Order number
	ErpOrderNumber *string `json:"erpOrderNumber,omitempty"`
	// Invoice no.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// Invoice Status.
	InvoiceStatus *string `json:"invoiceStatus,omitempty"`
	// Invoice Date.
	InvoiceDate *string `json:"invoiceDate,omitempty"`
	// Invoice Due Date.
	InvoiceDueDate *string `json:"invoiceDueDate,omitempty"`
	// Invoice Amount.
	InvoicedAmountDue *float32 `json:"invoicedAmountDue,omitempty"`
	// Customer Order No.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// Order Create Date.
	OrderCreateDate *string `json:"orderCreateDate,omitempty"`
	// End Customer Order number.
	EndCustomerOrderNumber *string `json:"endCustomerOrderNumber,omitempty"`
	// Invoice Amount Inclusive of Taxes
	InvoiceAmountInclTax *float32 `json:"invoiceAmountInclTax,omitempty"`
}

// NewInvoiceSearchResponseInvoicesInner instantiates a new InvoiceSearchResponseInvoicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceSearchResponseInvoicesInner() *InvoiceSearchResponseInvoicesInner {
	this := InvoiceSearchResponseInvoicesInner{}
	return &this
}

// NewInvoiceSearchResponseInvoicesInnerWithDefaults instantiates a new InvoiceSearchResponseInvoicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceSearchResponseInvoicesInnerWithDefaults() *InvoiceSearchResponseInvoicesInner {
	this := InvoiceSearchResponseInvoicesInner{}
	return &this
}

// GetPaymentTermsDueDate returns the PaymentTermsDueDate field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetPaymentTermsDueDate() string {
	if o == nil || IsNil(o.PaymentTermsDueDate) {
		var ret string
		return ret
	}
	return *o.PaymentTermsDueDate
}

// GetPaymentTermsDueDateOk returns a tuple with the PaymentTermsDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetPaymentTermsDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTermsDueDate) {
		return nil, false
	}
	return o.PaymentTermsDueDate, true
}

// HasPaymentTermsDueDate returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasPaymentTermsDueDate() bool {
	if o != nil && !IsNil(o.PaymentTermsDueDate) {
		return true
	}

	return false
}

// SetPaymentTermsDueDate gets a reference to the given string and assigns it to the PaymentTermsDueDate field.
func (o *InvoiceSearchResponseInvoicesInner) SetPaymentTermsDueDate(v string) {
	o.PaymentTermsDueDate = &v
}

// GetErpOrderNumber returns the ErpOrderNumber field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetErpOrderNumber() string {
	if o == nil || IsNil(o.ErpOrderNumber) {
		var ret string
		return ret
	}
	return *o.ErpOrderNumber
}

// GetErpOrderNumberOk returns a tuple with the ErpOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetErpOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ErpOrderNumber) {
		return nil, false
	}
	return o.ErpOrderNumber, true
}

// HasErpOrderNumber returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasErpOrderNumber() bool {
	if o != nil && !IsNil(o.ErpOrderNumber) {
		return true
	}

	return false
}

// SetErpOrderNumber gets a reference to the given string and assigns it to the ErpOrderNumber field.
func (o *InvoiceSearchResponseInvoicesInner) SetErpOrderNumber(v string) {
	o.ErpOrderNumber = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *InvoiceSearchResponseInvoicesInner) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetInvoiceStatus returns the InvoiceStatus field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceStatus() string {
	if o == nil || IsNil(o.InvoiceStatus) {
		var ret string
		return ret
	}
	return *o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceStatusOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceStatus) {
		return nil, false
	}
	return o.InvoiceStatus, true
}

// HasInvoiceStatus returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasInvoiceStatus() bool {
	if o != nil && !IsNil(o.InvoiceStatus) {
		return true
	}

	return false
}

// SetInvoiceStatus gets a reference to the given string and assigns it to the InvoiceStatus field.
func (o *InvoiceSearchResponseInvoicesInner) SetInvoiceStatus(v string) {
	o.InvoiceStatus = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *InvoiceSearchResponseInvoicesInner) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetInvoiceDueDate returns the InvoiceDueDate field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceDueDate() string {
	if o == nil || IsNil(o.InvoiceDueDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDueDate
}

// GetInvoiceDueDateOk returns a tuple with the InvoiceDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDueDate) {
		return nil, false
	}
	return o.InvoiceDueDate, true
}

// HasInvoiceDueDate returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasInvoiceDueDate() bool {
	if o != nil && !IsNil(o.InvoiceDueDate) {
		return true
	}

	return false
}

// SetInvoiceDueDate gets a reference to the given string and assigns it to the InvoiceDueDate field.
func (o *InvoiceSearchResponseInvoicesInner) SetInvoiceDueDate(v string) {
	o.InvoiceDueDate = &v
}

// GetInvoicedAmountDue returns the InvoicedAmountDue field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoicedAmountDue() float32 {
	if o == nil || IsNil(o.InvoicedAmountDue) {
		var ret float32
		return ret
	}
	return *o.InvoicedAmountDue
}

// GetInvoicedAmountDueOk returns a tuple with the InvoicedAmountDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoicedAmountDueOk() (*float32, bool) {
	if o == nil || IsNil(o.InvoicedAmountDue) {
		return nil, false
	}
	return o.InvoicedAmountDue, true
}

// HasInvoicedAmountDue returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasInvoicedAmountDue() bool {
	if o != nil && !IsNil(o.InvoicedAmountDue) {
		return true
	}

	return false
}

// SetInvoicedAmountDue gets a reference to the given float32 and assigns it to the InvoicedAmountDue field.
func (o *InvoiceSearchResponseInvoicesInner) SetInvoicedAmountDue(v float32) {
	o.InvoicedAmountDue = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *InvoiceSearchResponseInvoicesInner) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetOrderCreateDate returns the OrderCreateDate field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetOrderCreateDate() string {
	if o == nil || IsNil(o.OrderCreateDate) {
		var ret string
		return ret
	}
	return *o.OrderCreateDate
}

// GetOrderCreateDateOk returns a tuple with the OrderCreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetOrderCreateDateOk() (*string, bool) {
	if o == nil || IsNil(o.OrderCreateDate) {
		return nil, false
	}
	return o.OrderCreateDate, true
}

// HasOrderCreateDate returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasOrderCreateDate() bool {
	if o != nil && !IsNil(o.OrderCreateDate) {
		return true
	}

	return false
}

// SetOrderCreateDate gets a reference to the given string and assigns it to the OrderCreateDate field.
func (o *InvoiceSearchResponseInvoicesInner) SetOrderCreateDate(v string) {
	o.OrderCreateDate = &v
}

// GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetEndCustomerOrderNumber() string {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.EndCustomerOrderNumber
}

// GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetEndCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		return nil, false
	}
	return o.EndCustomerOrderNumber, true
}

// HasEndCustomerOrderNumber returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasEndCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.EndCustomerOrderNumber) {
		return true
	}

	return false
}

// SetEndCustomerOrderNumber gets a reference to the given string and assigns it to the EndCustomerOrderNumber field.
func (o *InvoiceSearchResponseInvoicesInner) SetEndCustomerOrderNumber(v string) {
	o.EndCustomerOrderNumber = &v
}

// GetInvoiceAmountInclTax returns the InvoiceAmountInclTax field value if set, zero value otherwise.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceAmountInclTax() float32 {
	if o == nil || IsNil(o.InvoiceAmountInclTax) {
		var ret float32
		return ret
	}
	return *o.InvoiceAmountInclTax
}

// GetInvoiceAmountInclTaxOk returns a tuple with the InvoiceAmountInclTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceSearchResponseInvoicesInner) GetInvoiceAmountInclTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.InvoiceAmountInclTax) {
		return nil, false
	}
	return o.InvoiceAmountInclTax, true
}

// HasInvoiceAmountInclTax returns a boolean if a field has been set.
func (o *InvoiceSearchResponseInvoicesInner) HasInvoiceAmountInclTax() bool {
	if o != nil && !IsNil(o.InvoiceAmountInclTax) {
		return true
	}

	return false
}

// SetInvoiceAmountInclTax gets a reference to the given float32 and assigns it to the InvoiceAmountInclTax field.
func (o *InvoiceSearchResponseInvoicesInner) SetInvoiceAmountInclTax(v float32) {
	o.InvoiceAmountInclTax = &v
}

func (o InvoiceSearchResponseInvoicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceSearchResponseInvoicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PaymentTermsDueDate) {
		toSerialize["paymentTermsDueDate"] = o.PaymentTermsDueDate
	}
	if !IsNil(o.ErpOrderNumber) {
		toSerialize["erpOrderNumber"] = o.ErpOrderNumber
	}
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if !IsNil(o.InvoiceStatus) {
		toSerialize["invoiceStatus"] = o.InvoiceStatus
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoiceDate"] = o.InvoiceDate
	}
	if !IsNil(o.InvoiceDueDate) {
		toSerialize["invoiceDueDate"] = o.InvoiceDueDate
	}
	if !IsNil(o.InvoicedAmountDue) {
		toSerialize["invoicedAmountDue"] = o.InvoicedAmountDue
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if !IsNil(o.OrderCreateDate) {
		toSerialize["orderCreateDate"] = o.OrderCreateDate
	}
	if !IsNil(o.EndCustomerOrderNumber) {
		toSerialize["endCustomerOrderNumber"] = o.EndCustomerOrderNumber
	}
	if !IsNil(o.InvoiceAmountInclTax) {
		toSerialize["invoiceAmountInclTax"] = o.InvoiceAmountInclTax
	}
	return toSerialize, nil
}

type NullableInvoiceSearchResponseInvoicesInner struct {
	value *InvoiceSearchResponseInvoicesInner
	isSet bool
}

func (v NullableInvoiceSearchResponseInvoicesInner) Get() *InvoiceSearchResponseInvoicesInner {
	return v.value
}

func (v *NullableInvoiceSearchResponseInvoicesInner) Set(val *InvoiceSearchResponseInvoicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceSearchResponseInvoicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceSearchResponseInvoicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceSearchResponseInvoicesInner(val *InvoiceSearchResponseInvoicesInner) *NullableInvoiceSearchResponseInvoicesInner {
	return &NullableInvoiceSearchResponseInvoicesInner{value: val, isSet: true}
}

func (v NullableInvoiceSearchResponseInvoicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceSearchResponseInvoicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


