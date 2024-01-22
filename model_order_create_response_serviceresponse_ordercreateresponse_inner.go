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

// checks if the OrderCreateResponseServiceresponseOrdercreateresponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseServiceresponseOrdercreateresponseInner{}

// OrderCreateResponseServiceresponseOrdercreateresponseInner struct for OrderCreateResponseServiceresponseOrdercreateresponseInner
type OrderCreateResponseServiceresponseOrdercreateresponseInner struct {
	// Number of line items that were successful
	Numberoflineswithsuccess *string `json:"numberoflineswithsuccess,omitempty"`
	// Number of line items with error
	Numberoflineswitherror *string `json:"numberoflineswitherror,omitempty"`
	// Number of line items with warnings
	Numberoflineswithwarning *string `json:"numberoflineswithwarning,omitempty"`
	// Ingram sales order number
	Globalorderid *string `json:"globalorderid,omitempty"`
	// S=Stocked PO D=Direct Ship PO
	Ordertype *string `json:"ordertype,omitempty"`
	// Time order received
	Ordertimestamp *string `json:"ordertimestamp,omitempty"`
	// Ingram Micro generated order number
	Invoicingsystemorderid *string `json:"invoicingsystemorderid,omitempty"`
	Taxamount *float32 `json:"taxamount,omitempty"`
	// Freight amount customer pays for freight
	Freightamount *float32 `json:"freightamount,omitempty"`
	// Total amount of order with freight and taxes
	Orderamount *float32 `json:"orderamount,omitempty"`
	// Collection of lines
	Lines []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner `json:"Lines,omitempty"`
}

// NewOrderCreateResponseServiceresponseOrdercreateresponseInner instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseServiceresponseOrdercreateresponseInner() *OrderCreateResponseServiceresponseOrdercreateresponseInner {
	this := OrderCreateResponseServiceresponseOrdercreateresponseInner{}
	return &this
}

// NewOrderCreateResponseServiceresponseOrdercreateresponseInnerWithDefaults instantiates a new OrderCreateResponseServiceresponseOrdercreateresponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseServiceresponseOrdercreateresponseInnerWithDefaults() *OrderCreateResponseServiceresponseOrdercreateresponseInner {
	this := OrderCreateResponseServiceresponseOrdercreateresponseInner{}
	return &this
}

// GetNumberoflineswithsuccess returns the Numberoflineswithsuccess field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithsuccess() string {
	if o == nil || IsNil(o.Numberoflineswithsuccess) {
		var ret string
		return ret
	}
	return *o.Numberoflineswithsuccess
}

// GetNumberoflineswithsuccessOk returns a tuple with the Numberoflineswithsuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithsuccessOk() (*string, bool) {
	if o == nil || IsNil(o.Numberoflineswithsuccess) {
		return nil, false
	}
	return o.Numberoflineswithsuccess, true
}

// HasNumberoflineswithsuccess returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasNumberoflineswithsuccess() bool {
	if o != nil && !IsNil(o.Numberoflineswithsuccess) {
		return true
	}

	return false
}

// SetNumberoflineswithsuccess gets a reference to the given string and assigns it to the Numberoflineswithsuccess field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetNumberoflineswithsuccess(v string) {
	o.Numberoflineswithsuccess = &v
}

// GetNumberoflineswitherror returns the Numberoflineswitherror field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswitherror() string {
	if o == nil || IsNil(o.Numberoflineswitherror) {
		var ret string
		return ret
	}
	return *o.Numberoflineswitherror
}

// GetNumberoflineswitherrorOk returns a tuple with the Numberoflineswitherror field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswitherrorOk() (*string, bool) {
	if o == nil || IsNil(o.Numberoflineswitherror) {
		return nil, false
	}
	return o.Numberoflineswitherror, true
}

// HasNumberoflineswitherror returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasNumberoflineswitherror() bool {
	if o != nil && !IsNil(o.Numberoflineswitherror) {
		return true
	}

	return false
}

// SetNumberoflineswitherror gets a reference to the given string and assigns it to the Numberoflineswitherror field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetNumberoflineswitherror(v string) {
	o.Numberoflineswitherror = &v
}

// GetNumberoflineswithwarning returns the Numberoflineswithwarning field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithwarning() string {
	if o == nil || IsNil(o.Numberoflineswithwarning) {
		var ret string
		return ret
	}
	return *o.Numberoflineswithwarning
}

// GetNumberoflineswithwarningOk returns a tuple with the Numberoflineswithwarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetNumberoflineswithwarningOk() (*string, bool) {
	if o == nil || IsNil(o.Numberoflineswithwarning) {
		return nil, false
	}
	return o.Numberoflineswithwarning, true
}

// HasNumberoflineswithwarning returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasNumberoflineswithwarning() bool {
	if o != nil && !IsNil(o.Numberoflineswithwarning) {
		return true
	}

	return false
}

// SetNumberoflineswithwarning gets a reference to the given string and assigns it to the Numberoflineswithwarning field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetNumberoflineswithwarning(v string) {
	o.Numberoflineswithwarning = &v
}

// GetGlobalorderid returns the Globalorderid field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetGlobalorderid() string {
	if o == nil || IsNil(o.Globalorderid) {
		var ret string
		return ret
	}
	return *o.Globalorderid
}

// GetGlobalorderidOk returns a tuple with the Globalorderid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetGlobalorderidOk() (*string, bool) {
	if o == nil || IsNil(o.Globalorderid) {
		return nil, false
	}
	return o.Globalorderid, true
}

// HasGlobalorderid returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasGlobalorderid() bool {
	if o != nil && !IsNil(o.Globalorderid) {
		return true
	}

	return false
}

// SetGlobalorderid gets a reference to the given string and assigns it to the Globalorderid field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetGlobalorderid(v string) {
	o.Globalorderid = &v
}

// GetOrdertype returns the Ordertype field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertype() string {
	if o == nil || IsNil(o.Ordertype) {
		var ret string
		return ret
	}
	return *o.Ordertype
}

// GetOrdertypeOk returns a tuple with the Ordertype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertypeOk() (*string, bool) {
	if o == nil || IsNil(o.Ordertype) {
		return nil, false
	}
	return o.Ordertype, true
}

// HasOrdertype returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasOrdertype() bool {
	if o != nil && !IsNil(o.Ordertype) {
		return true
	}

	return false
}

// SetOrdertype gets a reference to the given string and assigns it to the Ordertype field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetOrdertype(v string) {
	o.Ordertype = &v
}

// GetOrdertimestamp returns the Ordertimestamp field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertimestamp() string {
	if o == nil || IsNil(o.Ordertimestamp) {
		var ret string
		return ret
	}
	return *o.Ordertimestamp
}

// GetOrdertimestampOk returns a tuple with the Ordertimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrdertimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Ordertimestamp) {
		return nil, false
	}
	return o.Ordertimestamp, true
}

// HasOrdertimestamp returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasOrdertimestamp() bool {
	if o != nil && !IsNil(o.Ordertimestamp) {
		return true
	}

	return false
}

// SetOrdertimestamp gets a reference to the given string and assigns it to the Ordertimestamp field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetOrdertimestamp(v string) {
	o.Ordertimestamp = &v
}

// GetInvoicingsystemorderid returns the Invoicingsystemorderid field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetInvoicingsystemorderid() string {
	if o == nil || IsNil(o.Invoicingsystemorderid) {
		var ret string
		return ret
	}
	return *o.Invoicingsystemorderid
}

// GetInvoicingsystemorderidOk returns a tuple with the Invoicingsystemorderid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetInvoicingsystemorderidOk() (*string, bool) {
	if o == nil || IsNil(o.Invoicingsystemorderid) {
		return nil, false
	}
	return o.Invoicingsystemorderid, true
}

// HasInvoicingsystemorderid returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasInvoicingsystemorderid() bool {
	if o != nil && !IsNil(o.Invoicingsystemorderid) {
		return true
	}

	return false
}

// SetInvoicingsystemorderid gets a reference to the given string and assigns it to the Invoicingsystemorderid field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetInvoicingsystemorderid(v string) {
	o.Invoicingsystemorderid = &v
}

// GetTaxamount returns the Taxamount field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetTaxamount() float32 {
	if o == nil || IsNil(o.Taxamount) {
		var ret float32
		return ret
	}
	return *o.Taxamount
}

// GetTaxamountOk returns a tuple with the Taxamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetTaxamountOk() (*float32, bool) {
	if o == nil || IsNil(o.Taxamount) {
		return nil, false
	}
	return o.Taxamount, true
}

// HasTaxamount returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasTaxamount() bool {
	if o != nil && !IsNil(o.Taxamount) {
		return true
	}

	return false
}

// SetTaxamount gets a reference to the given float32 and assigns it to the Taxamount field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetTaxamount(v float32) {
	o.Taxamount = &v
}

// GetFreightamount returns the Freightamount field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetFreightamount() float32 {
	if o == nil || IsNil(o.Freightamount) {
		var ret float32
		return ret
	}
	return *o.Freightamount
}

// GetFreightamountOk returns a tuple with the Freightamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetFreightamountOk() (*float32, bool) {
	if o == nil || IsNil(o.Freightamount) {
		return nil, false
	}
	return o.Freightamount, true
}

// HasFreightamount returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasFreightamount() bool {
	if o != nil && !IsNil(o.Freightamount) {
		return true
	}

	return false
}

// SetFreightamount gets a reference to the given float32 and assigns it to the Freightamount field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetFreightamount(v float32) {
	o.Freightamount = &v
}

// GetOrderamount returns the Orderamount field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrderamount() float32 {
	if o == nil || IsNil(o.Orderamount) {
		var ret float32
		return ret
	}
	return *o.Orderamount
}

// GetOrderamountOk returns a tuple with the Orderamount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetOrderamountOk() (*float32, bool) {
	if o == nil || IsNil(o.Orderamount) {
		return nil, false
	}
	return o.Orderamount, true
}

// HasOrderamount returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasOrderamount() bool {
	if o != nil && !IsNil(o.Orderamount) {
		return true
	}

	return false
}

// SetOrderamount gets a reference to the given float32 and assigns it to the Orderamount field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetOrderamount(v float32) {
	o.Orderamount = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetLines() []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) GetLinesOk() ([]OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner and assigns it to the Lines field.
func (o *OrderCreateResponseServiceresponseOrdercreateresponseInner) SetLines(v []OrderCreateResponseServiceresponseOrdercreateresponseInnerLinesInner) {
	o.Lines = v
}

func (o OrderCreateResponseServiceresponseOrdercreateresponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseServiceresponseOrdercreateresponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Numberoflineswithsuccess) {
		toSerialize["numberoflineswithsuccess"] = o.Numberoflineswithsuccess
	}
	if !IsNil(o.Numberoflineswitherror) {
		toSerialize["numberoflineswitherror"] = o.Numberoflineswitherror
	}
	if !IsNil(o.Numberoflineswithwarning) {
		toSerialize["numberoflineswithwarning"] = o.Numberoflineswithwarning
	}
	if !IsNil(o.Globalorderid) {
		toSerialize["globalorderid"] = o.Globalorderid
	}
	if !IsNil(o.Ordertype) {
		toSerialize["ordertype"] = o.Ordertype
	}
	if !IsNil(o.Ordertimestamp) {
		toSerialize["ordertimestamp"] = o.Ordertimestamp
	}
	if !IsNil(o.Invoicingsystemorderid) {
		toSerialize["invoicingsystemorderid"] = o.Invoicingsystemorderid
	}
	if !IsNil(o.Taxamount) {
		toSerialize["taxamount"] = o.Taxamount
	}
	if !IsNil(o.Freightamount) {
		toSerialize["freightamount"] = o.Freightamount
	}
	if !IsNil(o.Orderamount) {
		toSerialize["orderamount"] = o.Orderamount
	}
	if !IsNil(o.Lines) {
		toSerialize["Lines"] = o.Lines
	}
	return toSerialize, nil
}

type NullableOrderCreateResponseServiceresponseOrdercreateresponseInner struct {
	value *OrderCreateResponseServiceresponseOrdercreateresponseInner
	isSet bool
}

func (v NullableOrderCreateResponseServiceresponseOrdercreateresponseInner) Get() *OrderCreateResponseServiceresponseOrdercreateresponseInner {
	return v.value
}

func (v *NullableOrderCreateResponseServiceresponseOrdercreateresponseInner) Set(val *OrderCreateResponseServiceresponseOrdercreateresponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseServiceresponseOrdercreateresponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseServiceresponseOrdercreateresponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseServiceresponseOrdercreateresponseInner(val *OrderCreateResponseServiceresponseOrdercreateresponseInner) *NullableOrderCreateResponseServiceresponseOrdercreateresponseInner {
	return &NullableOrderCreateResponseServiceresponseOrdercreateresponseInner{value: val, isSet: true}
}

func (v NullableOrderCreateResponseServiceresponseOrdercreateresponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseServiceresponseOrdercreateresponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

