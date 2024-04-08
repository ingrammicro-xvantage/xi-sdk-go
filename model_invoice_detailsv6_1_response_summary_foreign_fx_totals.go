/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the InvoiceDetailsv61ResponseSummaryForeignFxTotals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailsv61ResponseSummaryForeignFxTotals{}

// InvoiceDetailsv61ResponseSummaryForeignFxTotals struct for InvoiceDetailsv61ResponseSummaryForeignFxTotals
type InvoiceDetailsv61ResponseSummaryForeignFxTotals struct {
	// Foreign Currency Code.
	ForeignCurrencyCode *string `json:"foreignCurrencyCode,omitempty"`
	// Foreign rate.
	ForeignCurrencyFxRate *float64 `json:"foreignCurrencyFxRate,omitempty"`
	// Foreign amount.
	ForeignTotalTaxableAmount *string `json:"foreignTotalTaxableAmount,omitempty"`
	// Foreign amount.
	ForeignTotalTaxAmount *float64 `json:"foreignTotalTaxAmount,omitempty"`
	// Foreign due.
	ForeignInvoiceAmountDue *string `json:"foreignInvoiceAmountDue,omitempty"`
}

// NewInvoiceDetailsv61ResponseSummaryForeignFxTotals instantiates a new InvoiceDetailsv61ResponseSummaryForeignFxTotals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailsv61ResponseSummaryForeignFxTotals() *InvoiceDetailsv61ResponseSummaryForeignFxTotals {
	this := InvoiceDetailsv61ResponseSummaryForeignFxTotals{}
	return &this
}

// NewInvoiceDetailsv61ResponseSummaryForeignFxTotalsWithDefaults instantiates a new InvoiceDetailsv61ResponseSummaryForeignFxTotals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailsv61ResponseSummaryForeignFxTotalsWithDefaults() *InvoiceDetailsv61ResponseSummaryForeignFxTotals {
	this := InvoiceDetailsv61ResponseSummaryForeignFxTotals{}
	return &this
}

// GetForeignCurrencyCode returns the ForeignCurrencyCode field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignCurrencyCode() string {
	if o == nil || IsNil(o.ForeignCurrencyCode) {
		var ret string
		return ret
	}
	return *o.ForeignCurrencyCode
}

// GetForeignCurrencyCodeOk returns a tuple with the ForeignCurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ForeignCurrencyCode) {
		return nil, false
	}
	return o.ForeignCurrencyCode, true
}

// HasForeignCurrencyCode returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) HasForeignCurrencyCode() bool {
	if o != nil && !IsNil(o.ForeignCurrencyCode) {
		return true
	}

	return false
}

// SetForeignCurrencyCode gets a reference to the given string and assigns it to the ForeignCurrencyCode field.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) SetForeignCurrencyCode(v string) {
	o.ForeignCurrencyCode = &v
}

// GetForeignCurrencyFxRate returns the ForeignCurrencyFxRate field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignCurrencyFxRate() float64 {
	if o == nil || IsNil(o.ForeignCurrencyFxRate) {
		var ret float64
		return ret
	}
	return *o.ForeignCurrencyFxRate
}

// GetForeignCurrencyFxRateOk returns a tuple with the ForeignCurrencyFxRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignCurrencyFxRateOk() (*float64, bool) {
	if o == nil || IsNil(o.ForeignCurrencyFxRate) {
		return nil, false
	}
	return o.ForeignCurrencyFxRate, true
}

// HasForeignCurrencyFxRate returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) HasForeignCurrencyFxRate() bool {
	if o != nil && !IsNil(o.ForeignCurrencyFxRate) {
		return true
	}

	return false
}

// SetForeignCurrencyFxRate gets a reference to the given float64 and assigns it to the ForeignCurrencyFxRate field.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) SetForeignCurrencyFxRate(v float64) {
	o.ForeignCurrencyFxRate = &v
}

// GetForeignTotalTaxableAmount returns the ForeignTotalTaxableAmount field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignTotalTaxableAmount() string {
	if o == nil || IsNil(o.ForeignTotalTaxableAmount) {
		var ret string
		return ret
	}
	return *o.ForeignTotalTaxableAmount
}

// GetForeignTotalTaxableAmountOk returns a tuple with the ForeignTotalTaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignTotalTaxableAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ForeignTotalTaxableAmount) {
		return nil, false
	}
	return o.ForeignTotalTaxableAmount, true
}

// HasForeignTotalTaxableAmount returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) HasForeignTotalTaxableAmount() bool {
	if o != nil && !IsNil(o.ForeignTotalTaxableAmount) {
		return true
	}

	return false
}

// SetForeignTotalTaxableAmount gets a reference to the given string and assigns it to the ForeignTotalTaxableAmount field.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) SetForeignTotalTaxableAmount(v string) {
	o.ForeignTotalTaxableAmount = &v
}

// GetForeignTotalTaxAmount returns the ForeignTotalTaxAmount field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignTotalTaxAmount() float64 {
	if o == nil || IsNil(o.ForeignTotalTaxAmount) {
		var ret float64
		return ret
	}
	return *o.ForeignTotalTaxAmount
}

// GetForeignTotalTaxAmountOk returns a tuple with the ForeignTotalTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignTotalTaxAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.ForeignTotalTaxAmount) {
		return nil, false
	}
	return o.ForeignTotalTaxAmount, true
}

// HasForeignTotalTaxAmount returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) HasForeignTotalTaxAmount() bool {
	if o != nil && !IsNil(o.ForeignTotalTaxAmount) {
		return true
	}

	return false
}

// SetForeignTotalTaxAmount gets a reference to the given float64 and assigns it to the ForeignTotalTaxAmount field.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) SetForeignTotalTaxAmount(v float64) {
	o.ForeignTotalTaxAmount = &v
}

// GetForeignInvoiceAmountDue returns the ForeignInvoiceAmountDue field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignInvoiceAmountDue() string {
	if o == nil || IsNil(o.ForeignInvoiceAmountDue) {
		var ret string
		return ret
	}
	return *o.ForeignInvoiceAmountDue
}

// GetForeignInvoiceAmountDueOk returns a tuple with the ForeignInvoiceAmountDue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) GetForeignInvoiceAmountDueOk() (*string, bool) {
	if o == nil || IsNil(o.ForeignInvoiceAmountDue) {
		return nil, false
	}
	return o.ForeignInvoiceAmountDue, true
}

// HasForeignInvoiceAmountDue returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) HasForeignInvoiceAmountDue() bool {
	if o != nil && !IsNil(o.ForeignInvoiceAmountDue) {
		return true
	}

	return false
}

// SetForeignInvoiceAmountDue gets a reference to the given string and assigns it to the ForeignInvoiceAmountDue field.
func (o *InvoiceDetailsv61ResponseSummaryForeignFxTotals) SetForeignInvoiceAmountDue(v string) {
	o.ForeignInvoiceAmountDue = &v
}

func (o InvoiceDetailsv61ResponseSummaryForeignFxTotals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailsv61ResponseSummaryForeignFxTotals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForeignCurrencyCode) {
		toSerialize["foreignCurrencyCode"] = o.ForeignCurrencyCode
	}
	if !IsNil(o.ForeignCurrencyFxRate) {
		toSerialize["foreignCurrencyFxRate"] = o.ForeignCurrencyFxRate
	}
	if !IsNil(o.ForeignTotalTaxableAmount) {
		toSerialize["foreignTotalTaxableAmount"] = o.ForeignTotalTaxableAmount
	}
	if !IsNil(o.ForeignTotalTaxAmount) {
		toSerialize["foreignTotalTaxAmount"] = o.ForeignTotalTaxAmount
	}
	if !IsNil(o.ForeignInvoiceAmountDue) {
		toSerialize["foreignInvoiceAmountDue"] = o.ForeignInvoiceAmountDue
	}
	return toSerialize, nil
}

type NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals struct {
	value *InvoiceDetailsv61ResponseSummaryForeignFxTotals
	isSet bool
}

func (v NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals) Get() *InvoiceDetailsv61ResponseSummaryForeignFxTotals {
	return v.value
}

func (v *NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals) Set(val *InvoiceDetailsv61ResponseSummaryForeignFxTotals) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailsv61ResponseSummaryForeignFxTotals(val *InvoiceDetailsv61ResponseSummaryForeignFxTotals) *NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals {
	return &NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals{value: val, isSet: true}
}

func (v NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailsv61ResponseSummaryForeignFxTotals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


