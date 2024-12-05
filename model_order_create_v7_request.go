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

// checks if the OrderCreateV7Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateV7Request{}

// OrderCreateV7Request struct for OrderCreateV7Request
type OrderCreateV7Request struct {
	// A unique identifier generated by Ingram Micro's CRM specific to each quote.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// The reseller's order number for reference in their system.
	CustomerOrderNumber *string `json:"customerOrderNumber,omitempty"`
	// The end customer's order number for reference in their system.
	EndCustomerOrderNumber NullableString `json:"endCustomerOrderNumber,omitempty"`
	// Order header level notes.
	Notes *string `json:"notes,omitempty"`
	// Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit.
	BillToAddressId *string `json:"billToAddressId,omitempty"`
	// The bid number is provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers.
	SpecialBidNumber *string `json:"specialBidNumber,omitempty"`
	// ENUM ['true','false'] - accept order if this item is backordered. This field along with shipComplete field decides the value of backorderflag. The value of this field is ignored when shipComplete field is present.
	AcceptBackOrder *bool `json:"acceptBackOrder,omitempty"`
	// Authorization number provided by vendor to Ingram's reseller. Orders will be placed on hold without this value, vendor specific mandatory field - please reach out Ingram Sales team for list of vendor for whom this is mandatory.
	VendAuthNumber *string `json:"vendAuthNumber,omitempty"`
	ResellerInfo *OrderCreateV7RequestResellerInfo `json:"resellerInfo,omitempty"`
	EndUserInfo *OrderCreateV7RequestEndUserInfo `json:"endUserInfo,omitempty"`
	ShipToInfo *OrderCreateV7RequestShipToInfo `json:"shipToInfo,omitempty"`
	ShipmentDetails *OrderCreateV7RequestShipmentDetails `json:"shipmentDetails,omitempty"`
	// Shipment-level additional attributes.
	AdditionalAttributes []OrderCreateV7RequestAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
	// The object containing the list of fields required at a header level by the vendor.
	VmfAdditionalAttributes []OrderCreateV7RequestVmfAdditionalAttributesInner `json:"vmfAdditionalAttributes,omitempty"`
	Lines []OrderCreateV7RequestLinesInner `json:"lines,omitempty"`
}

// NewOrderCreateV7Request instantiates a new OrderCreateV7Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateV7Request() *OrderCreateV7Request {
	this := OrderCreateV7Request{}
	return &this
}

// NewOrderCreateV7RequestWithDefaults instantiates a new OrderCreateV7Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateV7RequestWithDefaults() *OrderCreateV7Request {
	this := OrderCreateV7Request{}
	return &this
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *OrderCreateV7Request) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetCustomerOrderNumber() string {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerOrderNumber) {
		return nil, false
	}
	return o.CustomerOrderNumber, true
}

// HasCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.CustomerOrderNumber) {
		return true
	}

	return false
}

// SetCustomerOrderNumber gets a reference to the given string and assigns it to the CustomerOrderNumber field.
func (o *OrderCreateV7Request) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = &v
}

// GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderCreateV7Request) GetEndCustomerOrderNumber() string {
	if o == nil || IsNil(o.EndCustomerOrderNumber.Get()) {
		var ret string
		return ret
	}
	return *o.EndCustomerOrderNumber.Get()
}

// GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderCreateV7Request) GetEndCustomerOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndCustomerOrderNumber.Get(), o.EndCustomerOrderNumber.IsSet()
}

// HasEndCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasEndCustomerOrderNumber() bool {
	if o != nil && o.EndCustomerOrderNumber.IsSet() {
		return true
	}

	return false
}

// SetEndCustomerOrderNumber gets a reference to the given NullableString and assigns it to the EndCustomerOrderNumber field.
func (o *OrderCreateV7Request) SetEndCustomerOrderNumber(v string) {
	o.EndCustomerOrderNumber.Set(&v)
}
// SetEndCustomerOrderNumberNil sets the value for EndCustomerOrderNumber to be an explicit nil
func (o *OrderCreateV7Request) SetEndCustomerOrderNumberNil() {
	o.EndCustomerOrderNumber.Set(nil)
}

// UnsetEndCustomerOrderNumber ensures that no value is present for EndCustomerOrderNumber, not even an explicit nil
func (o *OrderCreateV7Request) UnsetEndCustomerOrderNumber() {
	o.EndCustomerOrderNumber.Unset()
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderCreateV7Request) SetNotes(v string) {
	o.Notes = &v
}

// GetBillToAddressId returns the BillToAddressId field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetBillToAddressId() string {
	if o == nil || IsNil(o.BillToAddressId) {
		var ret string
		return ret
	}
	return *o.BillToAddressId
}

// GetBillToAddressIdOk returns a tuple with the BillToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetBillToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillToAddressId) {
		return nil, false
	}
	return o.BillToAddressId, true
}

// HasBillToAddressId returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasBillToAddressId() bool {
	if o != nil && !IsNil(o.BillToAddressId) {
		return true
	}

	return false
}

// SetBillToAddressId gets a reference to the given string and assigns it to the BillToAddressId field.
func (o *OrderCreateV7Request) SetBillToAddressId(v string) {
	o.BillToAddressId = &v
}

// GetSpecialBidNumber returns the SpecialBidNumber field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetSpecialBidNumber() string {
	if o == nil || IsNil(o.SpecialBidNumber) {
		var ret string
		return ret
	}
	return *o.SpecialBidNumber
}

// GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetSpecialBidNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidNumber) {
		return nil, false
	}
	return o.SpecialBidNumber, true
}

// HasSpecialBidNumber returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasSpecialBidNumber() bool {
	if o != nil && !IsNil(o.SpecialBidNumber) {
		return true
	}

	return false
}

// SetSpecialBidNumber gets a reference to the given string and assigns it to the SpecialBidNumber field.
func (o *OrderCreateV7Request) SetSpecialBidNumber(v string) {
	o.SpecialBidNumber = &v
}

// GetAcceptBackOrder returns the AcceptBackOrder field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetAcceptBackOrder() bool {
	if o == nil || IsNil(o.AcceptBackOrder) {
		var ret bool
		return ret
	}
	return *o.AcceptBackOrder
}

// GetAcceptBackOrderOk returns a tuple with the AcceptBackOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetAcceptBackOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptBackOrder) {
		return nil, false
	}
	return o.AcceptBackOrder, true
}

// HasAcceptBackOrder returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasAcceptBackOrder() bool {
	if o != nil && !IsNil(o.AcceptBackOrder) {
		return true
	}

	return false
}

// SetAcceptBackOrder gets a reference to the given bool and assigns it to the AcceptBackOrder field.
func (o *OrderCreateV7Request) SetAcceptBackOrder(v bool) {
	o.AcceptBackOrder = &v
}

// GetVendAuthNumber returns the VendAuthNumber field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetVendAuthNumber() string {
	if o == nil || IsNil(o.VendAuthNumber) {
		var ret string
		return ret
	}
	return *o.VendAuthNumber
}

// GetVendAuthNumberOk returns a tuple with the VendAuthNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetVendAuthNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendAuthNumber) {
		return nil, false
	}
	return o.VendAuthNumber, true
}

// HasVendAuthNumber returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasVendAuthNumber() bool {
	if o != nil && !IsNil(o.VendAuthNumber) {
		return true
	}

	return false
}

// SetVendAuthNumber gets a reference to the given string and assigns it to the VendAuthNumber field.
func (o *OrderCreateV7Request) SetVendAuthNumber(v string) {
	o.VendAuthNumber = &v
}

// GetResellerInfo returns the ResellerInfo field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetResellerInfo() OrderCreateV7RequestResellerInfo {
	if o == nil || IsNil(o.ResellerInfo) {
		var ret OrderCreateV7RequestResellerInfo
		return ret
	}
	return *o.ResellerInfo
}

// GetResellerInfoOk returns a tuple with the ResellerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetResellerInfoOk() (*OrderCreateV7RequestResellerInfo, bool) {
	if o == nil || IsNil(o.ResellerInfo) {
		return nil, false
	}
	return o.ResellerInfo, true
}

// HasResellerInfo returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasResellerInfo() bool {
	if o != nil && !IsNil(o.ResellerInfo) {
		return true
	}

	return false
}

// SetResellerInfo gets a reference to the given OrderCreateV7RequestResellerInfo and assigns it to the ResellerInfo field.
func (o *OrderCreateV7Request) SetResellerInfo(v OrderCreateV7RequestResellerInfo) {
	o.ResellerInfo = &v
}

// GetEndUserInfo returns the EndUserInfo field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetEndUserInfo() OrderCreateV7RequestEndUserInfo {
	if o == nil || IsNil(o.EndUserInfo) {
		var ret OrderCreateV7RequestEndUserInfo
		return ret
	}
	return *o.EndUserInfo
}

// GetEndUserInfoOk returns a tuple with the EndUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetEndUserInfoOk() (*OrderCreateV7RequestEndUserInfo, bool) {
	if o == nil || IsNil(o.EndUserInfo) {
		return nil, false
	}
	return o.EndUserInfo, true
}

// HasEndUserInfo returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasEndUserInfo() bool {
	if o != nil && !IsNil(o.EndUserInfo) {
		return true
	}

	return false
}

// SetEndUserInfo gets a reference to the given OrderCreateV7RequestEndUserInfo and assigns it to the EndUserInfo field.
func (o *OrderCreateV7Request) SetEndUserInfo(v OrderCreateV7RequestEndUserInfo) {
	o.EndUserInfo = &v
}

// GetShipToInfo returns the ShipToInfo field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetShipToInfo() OrderCreateV7RequestShipToInfo {
	if o == nil || IsNil(o.ShipToInfo) {
		var ret OrderCreateV7RequestShipToInfo
		return ret
	}
	return *o.ShipToInfo
}

// GetShipToInfoOk returns a tuple with the ShipToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetShipToInfoOk() (*OrderCreateV7RequestShipToInfo, bool) {
	if o == nil || IsNil(o.ShipToInfo) {
		return nil, false
	}
	return o.ShipToInfo, true
}

// HasShipToInfo returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasShipToInfo() bool {
	if o != nil && !IsNil(o.ShipToInfo) {
		return true
	}

	return false
}

// SetShipToInfo gets a reference to the given OrderCreateV7RequestShipToInfo and assigns it to the ShipToInfo field.
func (o *OrderCreateV7Request) SetShipToInfo(v OrderCreateV7RequestShipToInfo) {
	o.ShipToInfo = &v
}

// GetShipmentDetails returns the ShipmentDetails field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetShipmentDetails() OrderCreateV7RequestShipmentDetails {
	if o == nil || IsNil(o.ShipmentDetails) {
		var ret OrderCreateV7RequestShipmentDetails
		return ret
	}
	return *o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetShipmentDetailsOk() (*OrderCreateV7RequestShipmentDetails, bool) {
	if o == nil || IsNil(o.ShipmentDetails) {
		return nil, false
	}
	return o.ShipmentDetails, true
}

// HasShipmentDetails returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasShipmentDetails() bool {
	if o != nil && !IsNil(o.ShipmentDetails) {
		return true
	}

	return false
}

// SetShipmentDetails gets a reference to the given OrderCreateV7RequestShipmentDetails and assigns it to the ShipmentDetails field.
func (o *OrderCreateV7Request) SetShipmentDetails(v OrderCreateV7RequestShipmentDetails) {
	o.ShipmentDetails = &v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetAdditionalAttributes() []OrderCreateV7RequestAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderCreateV7RequestAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetAdditionalAttributesOk() ([]OrderCreateV7RequestAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderCreateV7RequestAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderCreateV7Request) SetAdditionalAttributes(v []OrderCreateV7RequestAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

// GetVmfAdditionalAttributes returns the VmfAdditionalAttributes field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetVmfAdditionalAttributes() []OrderCreateV7RequestVmfAdditionalAttributesInner {
	if o == nil || IsNil(o.VmfAdditionalAttributes) {
		var ret []OrderCreateV7RequestVmfAdditionalAttributesInner
		return ret
	}
	return o.VmfAdditionalAttributes
}

// GetVmfAdditionalAttributesOk returns a tuple with the VmfAdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetVmfAdditionalAttributesOk() ([]OrderCreateV7RequestVmfAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.VmfAdditionalAttributes) {
		return nil, false
	}
	return o.VmfAdditionalAttributes, true
}

// HasVmfAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasVmfAdditionalAttributes() bool {
	if o != nil && !IsNil(o.VmfAdditionalAttributes) {
		return true
	}

	return false
}

// SetVmfAdditionalAttributes gets a reference to the given []OrderCreateV7RequestVmfAdditionalAttributesInner and assigns it to the VmfAdditionalAttributes field.
func (o *OrderCreateV7Request) SetVmfAdditionalAttributes(v []OrderCreateV7RequestVmfAdditionalAttributesInner) {
	o.VmfAdditionalAttributes = v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderCreateV7Request) GetLines() []OrderCreateV7RequestLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderCreateV7RequestLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7Request) GetLinesOk() ([]OrderCreateV7RequestLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderCreateV7Request) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderCreateV7RequestLinesInner and assigns it to the Lines field.
func (o *OrderCreateV7Request) SetLines(v []OrderCreateV7RequestLinesInner) {
	o.Lines = v
}

func (o OrderCreateV7Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateV7Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.CustomerOrderNumber) {
		toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	}
	if o.EndCustomerOrderNumber.IsSet() {
		toSerialize["endCustomerOrderNumber"] = o.EndCustomerOrderNumber.Get()
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.BillToAddressId) {
		toSerialize["billToAddressId"] = o.BillToAddressId
	}
	if !IsNil(o.SpecialBidNumber) {
		toSerialize["specialBidNumber"] = o.SpecialBidNumber
	}
	if !IsNil(o.AcceptBackOrder) {
		toSerialize["acceptBackOrder"] = o.AcceptBackOrder
	}
	if !IsNil(o.VendAuthNumber) {
		toSerialize["vendAuthNumber"] = o.VendAuthNumber
	}
	if !IsNil(o.ResellerInfo) {
		toSerialize["resellerInfo"] = o.ResellerInfo
	}
	if !IsNil(o.EndUserInfo) {
		toSerialize["endUserInfo"] = o.EndUserInfo
	}
	if !IsNil(o.ShipToInfo) {
		toSerialize["shipToInfo"] = o.ShipToInfo
	}
	if !IsNil(o.ShipmentDetails) {
		toSerialize["shipmentDetails"] = o.ShipmentDetails
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	if !IsNil(o.VmfAdditionalAttributes) {
		toSerialize["vmfAdditionalAttributes"] = o.VmfAdditionalAttributes
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	return toSerialize, nil
}

type NullableOrderCreateV7Request struct {
	value *OrderCreateV7Request
	isSet bool
}

func (v NullableOrderCreateV7Request) Get() *OrderCreateV7Request {
	return v.value
}

func (v *NullableOrderCreateV7Request) Set(val *OrderCreateV7Request) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateV7Request) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateV7Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateV7Request(val *OrderCreateV7Request) *NullableOrderCreateV7Request {
	return &NullableOrderCreateV7Request{value: val, isSet: true}
}

func (v NullableOrderCreateV7Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateV7Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

