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

// checks if the OrderCreateV7ResponseResourceOrdersInnerLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateV7ResponseResourceOrdersInnerLinesInner{}

// OrderCreateV7ResponseResourceOrdersInnerLinesInner struct for OrderCreateV7ResponseResourceOrdersInnerLinesInner
type OrderCreateV7ResponseResourceOrdersInnerLinesInner struct {
	// The sub order number. The two-digit prefix is the warehouse code of the warehouse nearest the reseller. The middle number is the order number. The two-digit suffix is the sub order number.
	SubOrderNumber *string `json:"subOrderNumber,omitempty"`
	// The Ingram Micro line number for the product.
	IngramLineNumber *string `json:"ingramLineNumber,omitempty"`
	// The reseller's line number for reference in their system.
	CustomerLineNumber *string `json:"customerLineNumber,omitempty"`
	// The status for the line item in the order. One of: Backordered, Open
	LineStatus *string `json:"lineStatus,omitempty"`
	// The Ingram Micro part number for the line item.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The unit price for the line item.
	UnitPrice *float32 `json:"unitPrice,omitempty"`
	// The extended list price (unit price X quantity) for the line item.
	ExtendedUnitPrice *float32 `json:"extendedUnitPrice,omitempty"`
	// The quantity of the line item ordered.
	QuantityOrdered *int32 `json:"quantityOrdered,omitempty"`
	// The quantity of the line item that has been confirmed.
	QuantityConfirmed *int32 `json:"quantityConfirmed,omitempty"`
	// The quantity of the line item that is backordered.
	QuantityBackOrdered *int32 `json:"quantityBackOrdered,omitempty"`
	// Line-level notes.
	Notes *string `json:"notes,omitempty"`
	// The shipment details for the line item.
	ShipmentDetails []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner `json:"shipmentDetails,omitempty"`
}

// NewOrderCreateV7ResponseResourceOrdersInnerLinesInner instantiates a new OrderCreateV7ResponseResourceOrdersInnerLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateV7ResponseResourceOrdersInnerLinesInner() *OrderCreateV7ResponseResourceOrdersInnerLinesInner {
	this := OrderCreateV7ResponseResourceOrdersInnerLinesInner{}
	return &this
}

// NewOrderCreateV7ResponseResourceOrdersInnerLinesInnerWithDefaults instantiates a new OrderCreateV7ResponseResourceOrdersInnerLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateV7ResponseResourceOrdersInnerLinesInnerWithDefaults() *OrderCreateV7ResponseResourceOrdersInnerLinesInner {
	this := OrderCreateV7ResponseResourceOrdersInnerLinesInner{}
	return &this
}

// GetSubOrderNumber returns the SubOrderNumber field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetSubOrderNumber() string {
	if o == nil || IsNil(o.SubOrderNumber) {
		var ret string
		return ret
	}
	return *o.SubOrderNumber
}

// GetSubOrderNumberOk returns a tuple with the SubOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetSubOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SubOrderNumber) {
		return nil, false
	}
	return o.SubOrderNumber, true
}

// HasSubOrderNumber returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasSubOrderNumber() bool {
	if o != nil && !IsNil(o.SubOrderNumber) {
		return true
	}

	return false
}

// SetSubOrderNumber gets a reference to the given string and assigns it to the SubOrderNumber field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetSubOrderNumber(v string) {
	o.SubOrderNumber = &v
}

// GetIngramLineNumber returns the IngramLineNumber field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramLineNumber() string {
	if o == nil || IsNil(o.IngramLineNumber) {
		var ret string
		return ret
	}
	return *o.IngramLineNumber
}

// GetIngramLineNumberOk returns a tuple with the IngramLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramLineNumber) {
		return nil, false
	}
	return o.IngramLineNumber, true
}

// HasIngramLineNumber returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasIngramLineNumber() bool {
	if o != nil && !IsNil(o.IngramLineNumber) {
		return true
	}

	return false
}

// SetIngramLineNumber gets a reference to the given string and assigns it to the IngramLineNumber field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetIngramLineNumber(v string) {
	o.IngramLineNumber = &v
}

// GetCustomerLineNumber returns the CustomerLineNumber field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetCustomerLineNumber() string {
	if o == nil || IsNil(o.CustomerLineNumber) {
		var ret string
		return ret
	}
	return *o.CustomerLineNumber
}

// GetCustomerLineNumberOk returns a tuple with the CustomerLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetCustomerLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerLineNumber) {
		return nil, false
	}
	return o.CustomerLineNumber, true
}

// HasCustomerLineNumber returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasCustomerLineNumber() bool {
	if o != nil && !IsNil(o.CustomerLineNumber) {
		return true
	}

	return false
}

// SetCustomerLineNumber gets a reference to the given string and assigns it to the CustomerLineNumber field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetCustomerLineNumber(v string) {
	o.CustomerLineNumber = &v
}

// GetLineStatus returns the LineStatus field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetLineStatus() string {
	if o == nil || IsNil(o.LineStatus) {
		var ret string
		return ret
	}
	return *o.LineStatus
}

// GetLineStatusOk returns a tuple with the LineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetLineStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LineStatus) {
		return nil, false
	}
	return o.LineStatus, true
}

// HasLineStatus returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasLineStatus() bool {
	if o != nil && !IsNil(o.LineStatus) {
		return true
	}

	return false
}

// SetLineStatus gets a reference to the given string and assigns it to the LineStatus field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetLineStatus(v string) {
	o.LineStatus = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetUnitPrice() float32 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

// GetExtendedUnitPrice returns the ExtendedUnitPrice field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetExtendedUnitPrice() float32 {
	if o == nil || IsNil(o.ExtendedUnitPrice) {
		var ret float32
		return ret
	}
	return *o.ExtendedUnitPrice
}

// GetExtendedUnitPriceOk returns a tuple with the ExtendedUnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetExtendedUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.ExtendedUnitPrice) {
		return nil, false
	}
	return o.ExtendedUnitPrice, true
}

// HasExtendedUnitPrice returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasExtendedUnitPrice() bool {
	if o != nil && !IsNil(o.ExtendedUnitPrice) {
		return true
	}

	return false
}

// SetExtendedUnitPrice gets a reference to the given float32 and assigns it to the ExtendedUnitPrice field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetExtendedUnitPrice(v float32) {
	o.ExtendedUnitPrice = &v
}

// GetQuantityOrdered returns the QuantityOrdered field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityOrdered() int32 {
	if o == nil || IsNil(o.QuantityOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityOrdered
}

// GetQuantityOrderedOk returns a tuple with the QuantityOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityOrdered) {
		return nil, false
	}
	return o.QuantityOrdered, true
}

// HasQuantityOrdered returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasQuantityOrdered() bool {
	if o != nil && !IsNil(o.QuantityOrdered) {
		return true
	}

	return false
}

// SetQuantityOrdered gets a reference to the given int32 and assigns it to the QuantityOrdered field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetQuantityOrdered(v int32) {
	o.QuantityOrdered = &v
}

// GetQuantityConfirmed returns the QuantityConfirmed field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityConfirmed() int32 {
	if o == nil || IsNil(o.QuantityConfirmed) {
		var ret int32
		return ret
	}
	return *o.QuantityConfirmed
}

// GetQuantityConfirmedOk returns a tuple with the QuantityConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityConfirmedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityConfirmed) {
		return nil, false
	}
	return o.QuantityConfirmed, true
}

// HasQuantityConfirmed returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasQuantityConfirmed() bool {
	if o != nil && !IsNil(o.QuantityConfirmed) {
		return true
	}

	return false
}

// SetQuantityConfirmed gets a reference to the given int32 and assigns it to the QuantityConfirmed field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetQuantityConfirmed(v int32) {
	o.QuantityConfirmed = &v
}

// GetQuantityBackOrdered returns the QuantityBackOrdered field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityBackOrdered() int32 {
	if o == nil || IsNil(o.QuantityBackOrdered) {
		var ret int32
		return ret
	}
	return *o.QuantityBackOrdered
}

// GetQuantityBackOrderedOk returns a tuple with the QuantityBackOrdered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetQuantityBackOrderedOk() (*int32, bool) {
	if o == nil || IsNil(o.QuantityBackOrdered) {
		return nil, false
	}
	return o.QuantityBackOrdered, true
}

// HasQuantityBackOrdered returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasQuantityBackOrdered() bool {
	if o != nil && !IsNil(o.QuantityBackOrdered) {
		return true
	}

	return false
}

// SetQuantityBackOrdered gets a reference to the given int32 and assigns it to the QuantityBackOrdered field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetQuantityBackOrdered(v int32) {
	o.QuantityBackOrdered = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetNotes(v string) {
	o.Notes = &v
}

// GetShipmentDetails returns the ShipmentDetails field value if set, zero value otherwise.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetShipmentDetails() []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner {
	if o == nil || IsNil(o.ShipmentDetails) {
		var ret []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner
		return ret
	}
	return o.ShipmentDetails
}

// GetShipmentDetailsOk returns a tuple with the ShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) GetShipmentDetailsOk() ([]OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner, bool) {
	if o == nil || IsNil(o.ShipmentDetails) {
		return nil, false
	}
	return o.ShipmentDetails, true
}

// HasShipmentDetails returns a boolean if a field has been set.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) HasShipmentDetails() bool {
	if o != nil && !IsNil(o.ShipmentDetails) {
		return true
	}

	return false
}

// SetShipmentDetails gets a reference to the given []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner and assigns it to the ShipmentDetails field.
func (o *OrderCreateV7ResponseResourceOrdersInnerLinesInner) SetShipmentDetails(v []OrderCreateV7ResponseResourceOrdersInnerLinesInnerShipmentDetailsInner) {
	o.ShipmentDetails = v
}

func (o OrderCreateV7ResponseResourceOrdersInnerLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateV7ResponseResourceOrdersInnerLinesInner) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.LineStatus) {
		toSerialize["lineStatus"] = o.LineStatus
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if !IsNil(o.ExtendedUnitPrice) {
		toSerialize["extendedUnitPrice"] = o.ExtendedUnitPrice
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
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ShipmentDetails) {
		toSerialize["shipmentDetails"] = o.ShipmentDetails
	}
	return toSerialize, nil
}

type NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner struct {
	value *OrderCreateV7ResponseResourceOrdersInnerLinesInner
	isSet bool
}

func (v NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner) Get() *OrderCreateV7ResponseResourceOrdersInnerLinesInner {
	return v.value
}

func (v *NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner) Set(val *OrderCreateV7ResponseResourceOrdersInnerLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateV7ResponseResourceOrdersInnerLinesInner(val *OrderCreateV7ResponseResourceOrdersInnerLinesInner) *NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner {
	return &NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner{value: val, isSet: true}
}

func (v NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateV7ResponseResourceOrdersInnerLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


