/*
XI SDK Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderModifyResponseLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyResponseLinesInner{}

// OrderModifyResponseLinesInner struct for OrderModifyResponseLinesInner
type OrderModifyResponseLinesInner struct {
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// The IngramMicro line number.
	IngramLineNumber *string `json:"ingramLineNumber,omitempty"`
	// The reseller's line number for reference in their system.
	CustomerLineNumber *string `json:"customerLineNumber,omitempty"`
	// The unique IngramMicro part number for the line item.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The vendor's part number for the line item.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// The quantity ordered of the line item.
	QuantityOrdered *int32 `json:"quantityOrdered,omitempty"`
	// The quantity confirmed of the line item.
	QuantityConfirmed *int32 `json:"quantityConfirmed,omitempty"`
	// The quantity backordered of the line item.
	QuantityBackOrdered *int32 `json:"quantityBackOrdered,omitempty"`
	ShipmentDetails *OrderModifyResponseLinesInnerShipmentDetails `json:"shipmentDetails,omitempty"`
	// SAP requested and country-specific line level details.
	AdditionalAttributes []OrderModifyResponseLinesInnerAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
	// Line-level notes for the order.
	Notes *string `json:"notes,omitempty"`
}

// NewOrderModifyResponseLinesInner instantiates a new OrderModifyResponseLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyResponseLinesInner() *OrderModifyResponseLinesInner {
	this := OrderModifyResponseLinesInner{}
	return &this
}

// NewOrderModifyResponseLinesInnerWithDefaults instantiates a new OrderModifyResponseLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyResponseLinesInnerWithDefaults() *OrderModifyResponseLinesInner {
	this := OrderModifyResponseLinesInner{}
	return &this
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderModifyResponseLinesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetIngramLineNumber returns the IngramLineNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetIngramLineNumber() string {
	if o == nil || IsNil(o.IngramLineNumber) {
		var ret string
		return ret
	}
	return *o.IngramLineNumber
}

// GetIngramLineNumberOk returns a tuple with the IngramLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetIngramLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramLineNumber) {
		return nil, false
	}
	return o.IngramLineNumber, true
}

// HasIngramLineNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasIngramLineNumber() bool {
	if o != nil && !IsNil(o.IngramLineNumber) {
		return true
	}

	return false
}

// SetIngramLineNumber gets a reference to the given string and assigns it to the IngramLineNumber field.
func (o *OrderModifyResponseLinesInner) SetIngramLineNumber(v string) {
	o.IngramLineNumber = &v
}

// GetCustomerLineNumber returns the CustomerLineNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetCustomerLineNumber() string {
	if o == nil || IsNil(o.CustomerLineNumber) {
		var ret string
		return ret
	}
	return *o.CustomerLineNumber
}

// GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetCustomerLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerLineNumber) {
		return nil, false
	}
	return o.CustomerLineNumber, true
}

// HasCustomerLineNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasCustomerLineNumber() bool {
	if o != nil && !IsNil(o.CustomerLineNumber) {
		return true
	}

	return false
}

// SetCustomerLineNumber gets a reference to the given string and assigns it to the CustomerLineNumber field.
func (o *OrderModifyResponseLinesInner) SetCustomerLineNumber(v string) {
	o.CustomerLineNumber = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *OrderModifyResponseLinesInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *OrderModifyResponseLinesInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetQuantityOrdered() int32 {
	if o == nil || IsNil(o.QuantityOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityOrdered) {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasQuantityOrdered() bool {
	if o != nil && !IsNil(o.QuantityOrdered) {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given int32 and assigns it to the QuantityOrdered field.
func (o *OrderModifyResponseLinesInner) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = &v
}

// GetQuantityConfirmed returns the QuantityConfirmed field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetQuantityConfirmed() int32 {
	if o == nil || IsNil(o.QuantityConfirmed) {
		var ret int32
		return ret
	}
	return *o.QuantityConfirmed
}

// GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetQuantityConfirmedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityConfirmed) {
		return nil, false
	}
	return o.QuantityConfirmed, true
}

// HasQuantityConfirmed returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasQuantityConfirmed() bool {
	if o != nil && !IsNil(o.QuantityConfirmed) {
		return true
	}

	return false
}

// SetQuantityConfirmed gets a reference to the given int32 and assigns it to the QuantityConfirmed field.
func (o *OrderModifyResponseLinesInner) SetQuantityConfirmed(v int32) {
	o.QuantityConfirmed = &v
}

// GetQuantityBackOrdered returns the QuantityBackOrdered field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetQuantityBackOrdered() int32 {
	if o == nil || IsNil(o.QuantityBackOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityBackOrdered
}

// GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetQuantityBackOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityBackOrdered) {
		return nil, false
	}
	return o.QuantityBackOrdered, true
}

// HasQuantityBackOrdered returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasQuantityBackOrdered() bool {
	if o != nil && !IsNil(o.QuantityBackOrdered) {
		return true
	}

	return false
}

// SetQuantityBackOrdered gets a reference to the given int32 and assigns it to the QuantityBackOrdered field.
func (o *OrderModifyResponseLinesInner) SetQuantityBackOrdered(v int32) {
	o.QuantityBackOrdered = &v
}

// GetShipmentDetails returns the ShipmentDetails field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetShipmentDetails() OrderModifyResponseLinesInnerShipmentDetails {
	if o == nil || IsNil(o.ShipmentDetails) {
		var ret OrderModifyResponseLinesInnerShipmentDetails
		return ret
	}
	return *o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetShipmentDetailsOk() (*OrderModifyResponseLinesInnerShipmentDetails, bool) {
	if o == nil || IsNil(o.ShipmentDetails) {
		return nil, false
	}
	return o.ShipmentDetails, true
}

// HasShipmentDetails returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasShipmentDetails() bool {
	if o != nil && !IsNil(o.ShipmentDetails) {
		return true
	}

	return false
}

// SetShipmentDetails gets a reference to the given OrderModifyResponseLinesInnerShipmentDetails and assigns it to the ShipmentDetails field.
func (o *OrderModifyResponseLinesInner) SetShipmentDetails(v OrderModifyResponseLinesInnerShipmentDetails) {
	o.ShipmentDetails = &v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetAdditionalAttributes() []OrderModifyResponseLinesInnerAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderModifyResponseLinesInnerAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetAdditionalAttributesOk() ([]OrderModifyResponseLinesInnerAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderModifyResponseLinesInnerAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderModifyResponseLinesInner) SetAdditionalAttributes(v []OrderModifyResponseLinesInnerAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderModifyResponseLinesInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyResponseLinesInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderModifyResponseLinesInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderModifyResponseLinesInner) SetNotes(v string) {
	o.Notes = &v
}

func (o OrderModifyResponseLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyResponseLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubOrderNumber) {
		toSerialize["subOrderNumber"] = o.SubOrderNumber
	}
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
	if !IsNil(o.QuantityConfirmed) {
		toSerialize["quantityConfirmed"] = o.QuantityConfirmed
	}
	if !IsNil(o.QuantityBackOrdered) {
		toSerialize["quantityBackOrdered"] = o.QuantityBackOrdered
	}
	if !IsNil(o.ShipmentDetails) {
		toSerialize["shipmentDetails"] = o.ShipmentDetails
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return toSerialize, nil
}

type NullableOrderModifyResponseLinesInner struct {
	value *OrderModifyResponseLinesInner
	isSet bool
}

func (v NullableOrderModifyResponseLinesInner) Get() *OrderModifyResponseLinesInner {
	return v.value
}

func (v *NullableOrderModifyResponseLinesInner) Set(val *OrderModifyResponseLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyResponseLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyResponseLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyResponseLinesInner(val *OrderModifyResponseLinesInner) *NullableOrderModifyResponseLinesInner {
	return &NullableOrderModifyResponseLinesInner{value: val, isSet: true}
}

func (v NullableOrderModifyResponseLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyResponseLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


