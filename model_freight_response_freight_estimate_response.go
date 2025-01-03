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

// checks if the FreightResponseFreightEstimateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreightResponseFreightEstimateResponse{}

// FreightResponseFreightEstimateResponse struct for FreightResponseFreightEstimateResponse
type FreightResponseFreightEstimateResponse struct {
	// The country-specific three-character ISO 4217 currency code used for the order.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Total freight amount.
	TotalFreightAmount *float32 `json:"totalFreightAmount,omitempty"`
	// Total tax amount.
	TotalTaxAmount *float32 `json:"totalTaxAmount,omitempty"`
	// Total fees.
	TotalFees *float32 `json:"totalFees,omitempty"`
	// Total net amount.
	TotalNetAmount *float32 `json:"totalNetAmount,omitempty"`
	// Gross amount.
	GrossAmount *float32 `json:"grossAmount,omitempty"`
	Distribution []FreightResponseFreightEstimateResponseDistributionInner `json:"distribution,omitempty"`
	Lines []FreightResponseFreightEstimateResponseLinesInner `json:"lines,omitempty"`
}

// NewFreightResponseFreightEstimateResponse instantiates a new FreightResponseFreightEstimateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreightResponseFreightEstimateResponse() *FreightResponseFreightEstimateResponse {
	this := FreightResponseFreightEstimateResponse{}
	return &this
}

// NewFreightResponseFreightEstimateResponseWithDefaults instantiates a new FreightResponseFreightEstimateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreightResponseFreightEstimateResponseWithDefaults() *FreightResponseFreightEstimateResponse {
	this := FreightResponseFreightEstimateResponse{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *FreightResponseFreightEstimateResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetTotalFreightAmount returns the TotalFreightAmount field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetTotalFreightAmount() float32 {
	if o == nil || IsNil(o.TotalFreightAmount) {
		var ret float32
		return ret
	}
	return *o.TotalFreightAmount
}

// GetTotalFreightAmountOk returns a tuple with the TotalFreightAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetTotalFreightAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalFreightAmount) {
		return nil, false
	}
	return o.TotalFreightAmount, true
}

// HasTotalFreightAmount returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasTotalFreightAmount() bool {
	if o != nil && !IsNil(o.TotalFreightAmount) {
		return true
	}

	return false
}

// SetTotalFreightAmount gets a reference to the given float32 and assigns it to the TotalFreightAmount field.
func (o *FreightResponseFreightEstimateResponse) SetTotalFreightAmount(v float32) {
	o.TotalFreightAmount = &v
}

// GetTotalTaxAmount returns the TotalTaxAmount field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetTotalTaxAmount() float32 {
	if o == nil || IsNil(o.TotalTaxAmount) {
		var ret float32
		return ret
	}
	return *o.TotalTaxAmount
}

// GetTotalTaxAmountOk returns a tuple with the TotalTaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetTotalTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTaxAmount) {
		return nil, false
	}
	return o.TotalTaxAmount, true
}

// HasTotalTaxAmount returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasTotalTaxAmount() bool {
	if o != nil && !IsNil(o.TotalTaxAmount) {
		return true
	}

	return false
}

// SetTotalTaxAmount gets a reference to the given float32 and assigns it to the TotalTaxAmount field.
func (o *FreightResponseFreightEstimateResponse) SetTotalTaxAmount(v float32) {
	o.TotalTaxAmount = &v
}

// GetTotalFees returns the TotalFees field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetTotalFees() float32 {
	if o == nil || IsNil(o.TotalFees) {
		var ret float32
		return ret
	}
	return *o.TotalFees
}

// GetTotalFeesOk returns a tuple with the TotalFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetTotalFeesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalFees) {
		return nil, false
	}
	return o.TotalFees, true
}

// HasTotalFees returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasTotalFees() bool {
	if o != nil && !IsNil(o.TotalFees) {
		return true
	}

	return false
}

// SetTotalFees gets a reference to the given float32 and assigns it to the TotalFees field.
func (o *FreightResponseFreightEstimateResponse) SetTotalFees(v float32) {
	o.TotalFees = &v
}

// GetTotalNetAmount returns the TotalNetAmount field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetTotalNetAmount() float32 {
	if o == nil || IsNil(o.TotalNetAmount) {
		var ret float32
		return ret
	}
	return *o.TotalNetAmount
}

// GetTotalNetAmountOk returns a tuple with the TotalNetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetTotalNetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalNetAmount) {
		return nil, false
	}
	return o.TotalNetAmount, true
}

// HasTotalNetAmount returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasTotalNetAmount() bool {
	if o != nil && !IsNil(o.TotalNetAmount) {
		return true
	}

	return false
}

// SetTotalNetAmount gets a reference to the given float32 and assigns it to the TotalNetAmount field.
func (o *FreightResponseFreightEstimateResponse) SetTotalNetAmount(v float32) {
	o.TotalNetAmount = &v
}

// GetGrossAmount returns the GrossAmount field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetGrossAmount() float32 {
	if o == nil || IsNil(o.GrossAmount) {
		var ret float32
		return ret
	}
	return *o.GrossAmount
}

// GetGrossAmountOk returns a tuple with the GrossAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetGrossAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.GrossAmount) {
		return nil, false
	}
	return o.GrossAmount, true
}

// HasGrossAmount returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasGrossAmount() bool {
	if o != nil && !IsNil(o.GrossAmount) {
		return true
	}

	return false
}

// SetGrossAmount gets a reference to the given float32 and assigns it to the GrossAmount field.
func (o *FreightResponseFreightEstimateResponse) SetGrossAmount(v float32) {
	o.GrossAmount = &v
}

// GetDistribution returns the Distribution field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetDistribution() []FreightResponseFreightEstimateResponseDistributionInner {
	if o == nil || IsNil(o.Distribution) {
		var ret []FreightResponseFreightEstimateResponseDistributionInner
		return ret
	}
	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetDistributionOk() ([]FreightResponseFreightEstimateResponseDistributionInner, bool) {
	if o == nil || IsNil(o.Distribution) {
		return nil, false
	}
	return o.Distribution, true
}

// HasDistribution returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasDistribution() bool {
	if o != nil && !IsNil(o.Distribution) {
		return true
	}

	return false
}

// SetDistribution gets a reference to the given []FreightResponseFreightEstimateResponseDistributionInner and assigns it to the Distribution field.
func (o *FreightResponseFreightEstimateResponse) SetDistribution(v []FreightResponseFreightEstimateResponseDistributionInner) {
	o.Distribution = v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponse) GetLines() []FreightResponseFreightEstimateResponseLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []FreightResponseFreightEstimateResponseLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponse) GetLinesOk() ([]FreightResponseFreightEstimateResponseLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponse) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []FreightResponseFreightEstimateResponseLinesInner and assigns it to the Lines field.
func (o *FreightResponseFreightEstimateResponse) SetLines(v []FreightResponseFreightEstimateResponseLinesInner) {
	o.Lines = v
}

func (o FreightResponseFreightEstimateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreightResponseFreightEstimateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.TotalFreightAmount) {
		toSerialize["totalFreightAmount"] = o.TotalFreightAmount
	}
	if !IsNil(o.TotalTaxAmount) {
		toSerialize["totalTaxAmount"] = o.TotalTaxAmount
	}
	if !IsNil(o.TotalFees) {
		toSerialize["totalFees"] = o.TotalFees
	}
	if !IsNil(o.TotalNetAmount) {
		toSerialize["totalNetAmount"] = o.TotalNetAmount
	}
	if !IsNil(o.GrossAmount) {
		toSerialize["grossAmount"] = o.GrossAmount
	}
	if !IsNil(o.Distribution) {
		toSerialize["distribution"] = o.Distribution
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	return toSerialize, nil
}

type NullableFreightResponseFreightEstimateResponse struct {
	value *FreightResponseFreightEstimateResponse
	isSet bool
}

func (v NullableFreightResponseFreightEstimateResponse) Get() *FreightResponseFreightEstimateResponse {
	return v.value
}

func (v *NullableFreightResponseFreightEstimateResponse) Set(val *FreightResponseFreightEstimateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFreightResponseFreightEstimateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFreightResponseFreightEstimateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreightResponseFreightEstimateResponse(val *FreightResponseFreightEstimateResponse) *NullableFreightResponseFreightEstimateResponse {
	return &NullableFreightResponseFreightEstimateResponse{value: val, isSet: true}
}

func (v NullableFreightResponseFreightEstimateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreightResponseFreightEstimateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


