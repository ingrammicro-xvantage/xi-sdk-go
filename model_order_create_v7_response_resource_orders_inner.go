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

// checks if the OrderCreateV7ResponseResourceOrdersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateV7ResponseResourceOrdersInner{}

// OrderCreateV7ResponseResourceOrdersInner struct for OrderCreateV7ResponseResourceOrdersInner
type OrderCreateV7ResponseResourceOrdersInner struct {
	// The number of lines in the order that were successful.
	NumberOfLinesWithSuccess *int32 `json:"numberOfLinesWithSuccess,omitempty"`
	// The number of lines in the order that have errors.
	NumberOfLinesWithError *int32 `json:"numberOfLinesWithError,omitempty"`
	// The number of lines in the order that have a warning.
	NumberOfLinesWithWarning *int32 `json:"numberOfLinesWithWarning,omitempty"`
	// The Ingram Micro order number.
	IngramOrderNumber *string `json:"ingramOrderNumber,omitempty"`
	// The date in UTC format that the order was created in Ingram Micro's system.
	IngramOrderDate *string `json:"ingramOrderDate,omitempty"`
	// Order-level notes.
	Notes *string `json:"notes,omitempty"`
	// The order typer. One of: S=Stocked PO D=Direct Ship PO
	OrderType *string `json:"orderType,omitempty"`
	// The total price for the order.
	OrderTotal *float32 `json:"orderTotal,omitempty"`
	// The total freight charges for the order.
	FreightCharges *float32 `json:"freightCharges,omitempty"`
	// The total tax for the order.
	TotalTax *float32 `json:"totalTax,omitempty"`
	// The country-specific three character ISO 4217 currency code used for the order.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// The line-level details for the order.
	Lines []OrderCreateV7ResponseResourceOrdersInnerLinesInner `json:"lines,omitempty"`
	// Link to Order Details for the order(s).
	Links []OrderCreateResponseOrdersInnerLinksInner `json:"links,omitempty"`
}

// NewOrderCreateV7ResponseResourceOrdersInner instantiates a new OrderCreateV7ResponseResourceOrdersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateV7ResponseResourceOrdersInner() *OrderCreateV7ResponseResourceOrdersInner {
	this := OrderCreateV7ResponseResourceOrdersInner{}
	return &this
}

// NewOrderCreateV7ResponseResourceOrdersInnerWithDefaults instantiates a new OrderCreateV7ResponseResourceOrdersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateV7ResponseResourceOrdersInnerWithDefaults() *OrderCreateV7ResponseResourceOrdersInner {
	this := OrderCreateV7ResponseResourceOrdersInner{}
	return &this
}

// GetNumberOfLinesWithSuccess returns the NumberOfLinesWithSuccess field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithSuccess() int32 {
	if o == nil || IsNil(o.NumberOfLinesWithSuccess) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinesWithSuccess
}

// GetNumberOfLinesWithSuccessOk returns a tuple with the NumberOfLinesWithSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithSuccessOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinesWithSuccess) {
		return nil, false
	}
	return o.NumberOfLinesWithSuccess, true
}

// HasNumberOfLinesWithSuccess returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasNumberOfLinesWithSuccess() bool {
	if o != nil && !IsNil(o.NumberOfLinesWithSuccess) {
		return true
	}

	return false
}

// SetNumberOfLinesWithSuccess gets a reference to the given int32 and assigns it to the NumberOfLinesWithSuccess field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetNumberOfLinesWithSuccess(v int32) {
	o.NumberOfLinesWithSuccess = &v
}

// GetNumberOfLinesWithError returns the NumberOfLinesWithError field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithError() int32 {
	if o == nil || IsNil(o.NumberOfLinesWithError) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinesWithError
}

// GetNumberOfLinesWithErrorOk returns a tuple with the NumberOfLinesWithError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithErrorOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinesWithError) {
		return nil, false
	}
	return o.NumberOfLinesWithError, true
}

// HasNumberOfLinesWithError returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasNumberOfLinesWithError() bool {
	if o != nil && !IsNil(o.NumberOfLinesWithError) {
		return true
	}

	return false
}

// SetNumberOfLinesWithError gets a reference to the given int32 and assigns it to the NumberOfLinesWithError field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetNumberOfLinesWithError(v int32) {
	o.NumberOfLinesWithError = &v
}

// GetNumberOfLinesWithWarning returns the NumberOfLinesWithWarning field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithWarning() int32 {
	if o == nil || IsNil(o.NumberOfLinesWithWarning) {
		var ret int32
		return ret
	}
	return *o.NumberOfLinesWithWarning
}

// GetNumberOfLinesWithWarningOk returns a tuple with the NumberOfLinesWithWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNumberOfLinesWithWarningOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfLinesWithWarning) {
		return nil, false
	}
	return o.NumberOfLinesWithWarning, true
}

// HasNumberOfLinesWithWarning returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasNumberOfLinesWithWarning() bool {
	if o != nil && !IsNil(o.NumberOfLinesWithWarning) {
		return true
	}

	return false
}

// SetNumberOfLinesWithWarning gets a reference to the given int32 and assigns it to the NumberOfLinesWithWarning field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetNumberOfLinesWithWarning(v int32) {
	o.NumberOfLinesWithWarning = &v
}

// GetIngramOrderNumber returns the IngramOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderNumber() string {
	if o == nil || IsNil(o.IngramOrderNumber) {
		var ret string
		return ret
	}
	return *o.IngramOrderNumber
}

// GetIngramOrderNumberOk returns a tuple with the IngramOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderNumber) {
		return nil, false
	}
	return o.IngramOrderNumber, true
}

// HasIngramOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasIngramOrderNumber() bool {
	if o != nil && !IsNil(o.IngramOrderNumber) {
		return true
	}

	return false
}

// SetIngramOrderNumber gets a reference to the given string and assigns it to the IngramOrderNumber field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetIngramOrderNumber(v string) {
	o.IngramOrderNumber = &v
}

// GetIngramOrderDate returns the IngramOrderDate field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderDate() string {
	if o == nil || IsNil(o.IngramOrderDate) {
		var ret string
		return ret
	}
	return *o.IngramOrderDate
}

// GetIngramOrderDateOk returns a tuple with the IngramOrderDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetIngramOrderDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramOrderDate) {
		return nil, false
	}
	return o.IngramOrderDate, true
}

// HasIngramOrderDate returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasIngramOrderDate() bool {
	if o != nil && !IsNil(o.IngramOrderDate) {
		return true
	}

	return false
}

// SetIngramOrderDate gets a reference to the given string and assigns it to the IngramOrderDate field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetIngramOrderDate(v string) {
	o.IngramOrderDate = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetNotes(v string) {
	o.Notes = &v
}

// GetOrderType returns the OrderType field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderType() string {
	if o == nil || IsNil(o.OrderType) {
		var ret string
		return ret
	}
	return *o.OrderType
}

// GetOrderTypeOk returns a tuple with the OrderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OrderType) {
		return nil, false
	}
	return o.OrderType, true
}

// HasOrderType returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasOrderType() bool {
	if o != nil && !IsNil(o.OrderType) {
		return true
	}

	return false
}

// SetOrderType gets a reference to the given string and assigns it to the OrderType field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetOrderType(v string) {
	o.OrderType = &v
}

// GetOrderTotal returns the OrderTotal field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderTotal() float32 {
	if o == nil || IsNil(o.OrderTotal) {
		var ret float32
		return ret
	}
	return *o.OrderTotal
}

// GetOrderTotalOk returns a tuple with the OrderTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetOrderTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderTotal) {
		return nil, false
	}
	return o.OrderTotal, true
}

// HasOrderTotal returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasOrderTotal() bool {
	if o != nil && !IsNil(o.OrderTotal) {
		return true
	}

	return false
}

// SetOrderTotal gets a reference to the given float32 and assigns it to the OrderTotal field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetOrderTotal(v float32) {
	o.OrderTotal = &v
}

// GetFreightCharges returns the FreightCharges field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetFreightCharges() float32 {
	if o == nil || IsNil(o.FreightCharges) {
		var ret float32
		return ret
	}
	return *o.FreightCharges
}

// GetFreightChargesOk returns a tuple with the FreightCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetFreightChargesOk() (*float32, bool) {
	if o == nil || IsNil(o.FreightCharges) {
		return nil, false
	}
	return o.FreightCharges, true
}

// HasFreightCharges returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasFreightCharges() bool {
	if o != nil && !IsNil(o.FreightCharges) {
		return true
	}

	return false
}

// SetFreightCharges gets a reference to the given float32 and assigns it to the FreightCharges field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetFreightCharges(v float32) {
	o.FreightCharges = &v
}

// GetTotalTax returns the TotalTax field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetTotalTax() float32 {
	if o == nil || IsNil(o.TotalTax) {
		var ret float32
		return ret
	}
	return *o.TotalTax
}

// GetTotalTaxOk returns a tuple with the TotalTax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetTotalTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTax) {
		return nil, false
	}
	return o.TotalTax, true
}

// HasTotalTax returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasTotalTax() bool {
	if o != nil && !IsNil(o.TotalTax) {
		return true
	}

	return false
}

// SetTotalTax gets a reference to the given float32 and assigns it to the TotalTax field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetTotalTax(v float32) {
	o.TotalTax = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetLines() []OrderCreateV7ResponseResourceOrdersInnerLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderCreateV7ResponseResourceOrdersInnerLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetLinesOk() ([]OrderCreateV7ResponseResourceOrdersInnerLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderCreateV7ResponseResourceOrdersInnerLinesInner and assigns it to the Lines field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetLines(v []OrderCreateV7ResponseResourceOrdersInnerLinesInner) {
	o.Lines = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetLinks() []OrderCreateResponseOrdersInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []OrderCreateResponseOrdersInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) GetLinksOk() ([]OrderCreateResponseOrdersInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []OrderCreateResponseOrdersInnerLinksInner and assigns it to the Links field.
func (o *OrderCreateV7ResponseResourceOrdersInner) SetLinks(v []OrderCreateResponseOrdersInnerLinksInner) {
	o.Links = v
}

func (o OrderCreateV7ResponseResourceOrdersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateV7ResponseResourceOrdersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfLinesWithSuccess) {
		toSerialize["numberOfLinesWithSuccess"] = o.NumberOfLinesWithSuccess
	}
	if !IsNil(o.NumberOfLinesWithError) {
		toSerialize["numberOfLinesWithError"] = o.NumberOfLinesWithError
	}
	if !IsNil(o.NumberOfLinesWithWarning) {
		toSerialize["numberOfLinesWithWarning"] = o.NumberOfLinesWithWarning
	}
	if !IsNil(o.IngramOrderNumber) {
		toSerialize["ingramOrderNumber"] = o.IngramOrderNumber
	}
	if !IsNil(o.IngramOrderDate) {
		toSerialize["ingramOrderDate"] = o.IngramOrderDate
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.OrderType) {
		toSerialize["orderType"] = o.OrderType
	}
	if !IsNil(o.OrderTotal) {
		toSerialize["orderTotal"] = o.OrderTotal
	}
	if !IsNil(o.FreightCharges) {
		toSerialize["freightCharges"] = o.FreightCharges
	}
	if !IsNil(o.TotalTax) {
		toSerialize["totalTax"] = o.TotalTax
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableOrderCreateV7ResponseResourceOrdersInner struct {
	value *OrderCreateV7ResponseResourceOrdersInner
	isSet bool
}

func (v NullableOrderCreateV7ResponseResourceOrdersInner) Get() *OrderCreateV7ResponseResourceOrdersInner {
	return v.value
}

func (v *NullableOrderCreateV7ResponseResourceOrdersInner) Set(val *OrderCreateV7ResponseResourceOrdersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateV7ResponseResourceOrdersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateV7ResponseResourceOrdersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateV7ResponseResourceOrdersInner(val *OrderCreateV7ResponseResourceOrdersInner) *NullableOrderCreateV7ResponseResourceOrdersInner {
	return &NullableOrderCreateV7ResponseResourceOrdersInner{value: val, isSet: true}
}

func (v NullableOrderCreateV7ResponseResourceOrdersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateV7ResponseResourceOrdersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

