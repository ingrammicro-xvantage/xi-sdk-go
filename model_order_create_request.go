/*
XI Sdk Resellers

For resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequest{}

// OrderCreateRequest struct for OrderCreateRequest
type OrderCreateRequest struct {
	// The reseller's unique PO/Order number.
	CustomerOrderNumber string `json:"customerOrderNumber"`
	// The end user/customer's Purchase Order number.
	EndCustomerOrderNumber *string `json:"endCustomerOrderNumber,omitempty"`
	// Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit.
	BillToAddressId *string `json:"billToAddressId,omitempty"`
	// The bid number provided to the reseller by the vendor for special pricing and discounts. Line-level bid numbers take precedence over header-level bid numbers.
	SpecialBidNumber *string `json:"specialBidNumber,omitempty"`
	// Order level notes.
	Notes *string `json:"notes,omitempty"`
	// ENUM [\"true\",\"false\"] - accept order if this item is backordered. This field along with shipComplete field decides the value of backorderflag. The value of this field is ignored when shipComplete field is present.
	AcceptBackOrder *bool `json:"acceptBackOrder,omitempty"`
	ResellerInfo *OrderCreateRequestResellerInfo `json:"resellerInfo,omitempty"`
	Vmf *OrderCreateRequestVmf `json:"vmf,omitempty"`
	ShipToInfo *OrderCreateRequestShipToInfo `json:"shipToInfo,omitempty"`
	EndUserInfo *OrderCreateRequestEndUserInfo `json:"endUserInfo,omitempty"`
	// The line-level details of the order.
	Lines []OrderCreateRequestLinesInner `json:"lines,omitempty"`
	ShipmentDetails *OrderCreateRequestShipmentDetails `json:"shipmentDetails,omitempty"`
	// Shipment-level additional attributes.
	AdditionalAttributes []OrderCreateRequestAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

type _OrderCreateRequest OrderCreateRequest

// NewOrderCreateRequest instantiates a new OrderCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequest(customerOrderNumber string) *OrderCreateRequest {
	this := OrderCreateRequest{}
	this.CustomerOrderNumber = customerOrderNumber
	return &this
}

// NewOrderCreateRequestWithDefaults instantiates a new OrderCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestWithDefaults() *OrderCreateRequest {
	this := OrderCreateRequest{}
	return &this
}

// GetCustomerOrderNumber returns the CustomerOrderNumber field value
func (o *OrderCreateRequest) GetCustomerOrderNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerOrderNumber
}

// GetCustomerOrderNumberOk returns a tuple with the CustomerOrderNumber field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetCustomerOrderNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerOrderNumber, true
}

// SetCustomerOrderNumber sets field value
func (o *OrderCreateRequest) SetCustomerOrderNumber(v string) {
	o.CustomerOrderNumber = v
}

// GetEndCustomerOrderNumber returns the EndCustomerOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetEndCustomerOrderNumber() string {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		var ret string
		return ret
	}
	return *o.EndCustomerOrderNumber
}

// GetEndCustomerOrderNumberOk returns a tuple with the EndCustomerOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetEndCustomerOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.EndCustomerOrderNumber) {
		return nil, false
	}
	return o.EndCustomerOrderNumber, true
}

// HasEndCustomerOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasEndCustomerOrderNumber() bool {
	if o != nil && !IsNil(o.EndCustomerOrderNumber) {
		return true
	}

	return false
}

// SetEndCustomerOrderNumber gets a reference to the given string and assigns it to the EndCustomerOrderNumber field.
func (o *OrderCreateRequest) SetEndCustomerOrderNumber(v string) {
	o.EndCustomerOrderNumber = &v
}

// GetBillToAddressId returns the BillToAddressId field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetBillToAddressId() string {
	if o == nil || IsNil(o.BillToAddressId) {
		var ret string
		return ret
	}
	return *o.BillToAddressId
}

// GetBillToAddressIdOk returns a tuple with the BillToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetBillToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillToAddressId) {
		return nil, false
	}
	return o.BillToAddressId, true
}

// HasBillToAddressId returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasBillToAddressId() bool {
	if o != nil && !IsNil(o.BillToAddressId) {
		return true
	}

	return false
}

// SetBillToAddressId gets a reference to the given string and assigns it to the BillToAddressId field.
func (o *OrderCreateRequest) SetBillToAddressId(v string) {
	o.BillToAddressId = &v
}

// GetSpecialBidNumber returns the SpecialBidNumber field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetSpecialBidNumber() string {
	if o == nil || IsNil(o.SpecialBidNumber) {
		var ret string
		return ret
	}
	return *o.SpecialBidNumber
}

// GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetSpecialBidNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidNumber) {
		return nil, false
	}
	return o.SpecialBidNumber, true
}

// HasSpecialBidNumber returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasSpecialBidNumber() bool {
	if o != nil && !IsNil(o.SpecialBidNumber) {
		return true
	}

	return false
}

// SetSpecialBidNumber gets a reference to the given string and assigns it to the SpecialBidNumber field.
func (o *OrderCreateRequest) SetSpecialBidNumber(v string) {
	o.SpecialBidNumber = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderCreateRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetAcceptBackOrder returns the AcceptBackOrder field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetAcceptBackOrder() bool {
	if o == nil || IsNil(o.AcceptBackOrder) {
		var ret bool
		return ret
	}
	return *o.AcceptBackOrder
}

// GetAcceptBackOrderOk returns a tuple with the AcceptBackOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetAcceptBackOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptBackOrder) {
		return nil, false
	}
	return o.AcceptBackOrder, true
}

// HasAcceptBackOrder returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasAcceptBackOrder() bool {
	if o != nil && !IsNil(o.AcceptBackOrder) {
		return true
	}

	return false
}

// SetAcceptBackOrder gets a reference to the given bool and assigns it to the AcceptBackOrder field.
func (o *OrderCreateRequest) SetAcceptBackOrder(v bool) {
	o.AcceptBackOrder = &v
}

// GetResellerInfo returns the ResellerInfo field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetResellerInfo() OrderCreateRequestResellerInfo {
	if o == nil || IsNil(o.ResellerInfo) {
		var ret OrderCreateRequestResellerInfo
		return ret
	}
	return *o.ResellerInfo
}

// GetResellerInfoOk returns a tuple with the ResellerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetResellerInfoOk() (*OrderCreateRequestResellerInfo, bool) {
	if o == nil || IsNil(o.ResellerInfo) {
		return nil, false
	}
	return o.ResellerInfo, true
}

// HasResellerInfo returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasResellerInfo() bool {
	if o != nil && !IsNil(o.ResellerInfo) {
		return true
	}

	return false
}

// SetResellerInfo gets a reference to the given OrderCreateRequestResellerInfo and assigns it to the ResellerInfo field.
func (o *OrderCreateRequest) SetResellerInfo(v OrderCreateRequestResellerInfo) {
	o.ResellerInfo = &v
}

// GetVmf returns the Vmf field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetVmf() OrderCreateRequestVmf {
	if o == nil || IsNil(o.Vmf) {
		var ret OrderCreateRequestVmf
		return ret
	}
	return *o.Vmf
}

// GetVmfOk returns a tuple with the Vmf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetVmfOk() (*OrderCreateRequestVmf, bool) {
	if o == nil || IsNil(o.Vmf) {
		return nil, false
	}
	return o.Vmf, true
}

// HasVmf returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasVmf() bool {
	if o != nil && !IsNil(o.Vmf) {
		return true
	}

	return false
}

// SetVmf gets a reference to the given OrderCreateRequestVmf and assigns it to the Vmf field.
func (o *OrderCreateRequest) SetVmf(v OrderCreateRequestVmf) {
	o.Vmf = &v
}

// GetShipToInfo returns the ShipToInfo field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetShipToInfo() OrderCreateRequestShipToInfo {
	if o == nil || IsNil(o.ShipToInfo) {
		var ret OrderCreateRequestShipToInfo
		return ret
	}
	return *o.ShipToInfo
}

// GetShipToInfoOk returns a tuple with the ShipToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetShipToInfoOk() (*OrderCreateRequestShipToInfo, bool) {
	if o == nil || IsNil(o.ShipToInfo) {
		return nil, false
	}
	return o.ShipToInfo, true
}

// HasShipToInfo returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasShipToInfo() bool {
	if o != nil && !IsNil(o.ShipToInfo) {
		return true
	}

	return false
}

// SetShipToInfo gets a reference to the given OrderCreateRequestShipToInfo and assigns it to the ShipToInfo field.
func (o *OrderCreateRequest) SetShipToInfo(v OrderCreateRequestShipToInfo) {
	o.ShipToInfo = &v
}

// GetEndUserInfo returns the EndUserInfo field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetEndUserInfo() OrderCreateRequestEndUserInfo {
	if o == nil || IsNil(o.EndUserInfo) {
		var ret OrderCreateRequestEndUserInfo
		return ret
	}
	return *o.EndUserInfo
}

// GetEndUserInfoOk returns a tuple with the EndUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetEndUserInfoOk() (*OrderCreateRequestEndUserInfo, bool) {
	if o == nil || IsNil(o.EndUserInfo) {
		return nil, false
	}
	return o.EndUserInfo, true
}

// HasEndUserInfo returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasEndUserInfo() bool {
	if o != nil && !IsNil(o.EndUserInfo) {
		return true
	}

	return false
}

// SetEndUserInfo gets a reference to the given OrderCreateRequestEndUserInfo and assigns it to the EndUserInfo field.
func (o *OrderCreateRequest) SetEndUserInfo(v OrderCreateRequestEndUserInfo) {
	o.EndUserInfo = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetLines() []OrderCreateRequestLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderCreateRequestLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetLinesOk() ([]OrderCreateRequestLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderCreateRequestLinesInner and assigns it to the Lines field.
func (o *OrderCreateRequest) SetLines(v []OrderCreateRequestLinesInner) {
	o.Lines = v
}

// GetShipmentDetails returns the ShipmentDetails field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetShipmentDetails() OrderCreateRequestShipmentDetails {
	if o == nil || IsNil(o.ShipmentDetails) {
		var ret OrderCreateRequestShipmentDetails
		return ret
	}
	return *o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetShipmentDetailsOk() (*OrderCreateRequestShipmentDetails, bool) {
	if o == nil || IsNil(o.ShipmentDetails) {
		return nil, false
	}
	return o.ShipmentDetails, true
}

// HasShipmentDetails returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasShipmentDetails() bool {
	if o != nil && !IsNil(o.ShipmentDetails) {
		return true
	}

	return false
}

// SetShipmentDetails gets a reference to the given OrderCreateRequestShipmentDetails and assigns it to the ShipmentDetails field.
func (o *OrderCreateRequest) SetShipmentDetails(v OrderCreateRequestShipmentDetails) {
	o.ShipmentDetails = &v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderCreateRequest) GetAdditionalAttributes() []OrderCreateRequestAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderCreateRequestAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequest) GetAdditionalAttributesOk() ([]OrderCreateRequestAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderCreateRequest) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderCreateRequestAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderCreateRequest) SetAdditionalAttributes(v []OrderCreateRequestAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o OrderCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerOrderNumber"] = o.CustomerOrderNumber
	if !IsNil(o.EndCustomerOrderNumber) {
		toSerialize["endCustomerOrderNumber"] = o.EndCustomerOrderNumber
	}
	if !IsNil(o.BillToAddressId) {
		toSerialize["billToAddressId"] = o.BillToAddressId
	}
	if !IsNil(o.SpecialBidNumber) {
		toSerialize["specialBidNumber"] = o.SpecialBidNumber
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.AcceptBackOrder) {
		toSerialize["acceptBackOrder"] = o.AcceptBackOrder
	}
	if !IsNil(o.ResellerInfo) {
		toSerialize["resellerInfo"] = o.ResellerInfo
	}
	if !IsNil(o.Vmf) {
		toSerialize["vmf"] = o.Vmf
	}
	if !IsNil(o.ShipToInfo) {
		toSerialize["shipToInfo"] = o.ShipToInfo
	}
	if !IsNil(o.EndUserInfo) {
		toSerialize["endUserInfo"] = o.EndUserInfo
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.ShipmentDetails) {
		toSerialize["shipmentDetails"] = o.ShipmentDetails
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

func (o *OrderCreateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerOrderNumber",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderCreateRequest := _OrderCreateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderCreateRequest)

	if err != nil {
		return err
	}

	*o = OrderCreateRequest(varOrderCreateRequest)

	return err
}

type NullableOrderCreateRequest struct {
	value *OrderCreateRequest
	isSet bool
}

func (v NullableOrderCreateRequest) Get() *OrderCreateRequest {
	return v.value
}

func (v *NullableOrderCreateRequest) Set(val *OrderCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequest(val *OrderCreateRequest) *NullableOrderCreateRequest {
	return &NullableOrderCreateRequest{value: val, isSet: true}
}

func (v NullableOrderCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


