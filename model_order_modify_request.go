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

// checks if the OrderModifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyRequest{}

// OrderModifyRequest struct for OrderModifyRequest
type OrderModifyRequest struct {
	// Shipment-level notes.
	Notes *string `json:"notes,omitempty"`
	ShipToInfo *OrderModifyRequestShipToInfo `json:"shipToInfo,omitempty"`
	// The order line items.
	Lines []OrderModifyRequestLinesInner `json:"lines,omitempty"`
	// Header-level additional attributes.
	AdditionalAttributes []OrderModifyRequestAdditionalAttributesInner `json:"additionalAttributes,omitempty"`
}

// NewOrderModifyRequest instantiates a new OrderModifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyRequest() *OrderModifyRequest {
	this := OrderModifyRequest{}
	return &this
}

// NewOrderModifyRequestWithDefaults instantiates a new OrderModifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyRequestWithDefaults() *OrderModifyRequest {
	this := OrderModifyRequest{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OrderModifyRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OrderModifyRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OrderModifyRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetShipToInfo returns the ShipToInfo field value if set, zero value otherwise.
func (o *OrderModifyRequest) GetShipToInfo() OrderModifyRequestShipToInfo {
	if o == nil || IsNil(o.ShipToInfo) {
		var ret OrderModifyRequestShipToInfo
		return ret
	}
	return *o.ShipToInfo
}

// GetShipToInfoOk returns a tuple with the ShipToInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequest) GetShipToInfoOk() (*OrderModifyRequestShipToInfo, bool) {
	if o == nil || IsNil(o.ShipToInfo) {
		return nil, false
	}
	return o.ShipToInfo, true
}

// HasShipToInfo returns a boolean if a field has been set.
func (o *OrderModifyRequest) HasShipToInfo() bool {
	if o != nil && !IsNil(o.ShipToInfo) {
		return true
	}

	return false
}

// SetShipToInfo gets a reference to the given OrderModifyRequestShipToInfo and assigns it to the ShipToInfo field.
func (o *OrderModifyRequest) SetShipToInfo(v OrderModifyRequestShipToInfo) {
	o.ShipToInfo = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *OrderModifyRequest) GetLines() []OrderModifyRequestLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []OrderModifyRequestLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequest) GetLinesOk() ([]OrderModifyRequestLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *OrderModifyRequest) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []OrderModifyRequestLinesInner and assigns it to the Lines field.
func (o *OrderModifyRequest) SetLines(v []OrderModifyRequestLinesInner) {
	o.Lines = v
}

// GetAdditionalAttributes returns the AdditionalAttributes field value if set, zero value otherwise.
func (o *OrderModifyRequest) GetAdditionalAttributes() []OrderModifyRequestAdditionalAttributesInner {
	if o == nil || IsNil(o.AdditionalAttributes) {
		var ret []OrderModifyRequestAdditionalAttributesInner
		return ret
	}
	return o.AdditionalAttributes
}

// GetAdditionalAttributesOk returns a tuple with the AdditionalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequest) GetAdditionalAttributesOk() ([]OrderModifyRequestAdditionalAttributesInner, bool) {
	if o == nil || IsNil(o.AdditionalAttributes) {
		return nil, false
	}
	return o.AdditionalAttributes, true
}

// HasAdditionalAttributes returns a boolean if a field has been set.
func (o *OrderModifyRequest) HasAdditionalAttributes() bool {
	if o != nil && !IsNil(o.AdditionalAttributes) {
		return true
	}

	return false
}

// SetAdditionalAttributes gets a reference to the given []OrderModifyRequestAdditionalAttributesInner and assigns it to the AdditionalAttributes field.
func (o *OrderModifyRequest) SetAdditionalAttributes(v []OrderModifyRequestAdditionalAttributesInner) {
	o.AdditionalAttributes = v
}

func (o OrderModifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.ShipToInfo) {
		toSerialize["shipToInfo"] = o.ShipToInfo
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.AdditionalAttributes) {
		toSerialize["additionalAttributes"] = o.AdditionalAttributes
	}
	return toSerialize, nil
}

type NullableOrderModifyRequest struct {
	value *OrderModifyRequest
	isSet bool
}

func (v NullableOrderModifyRequest) Get() *OrderModifyRequest {
	return v.value
}

func (v *NullableOrderModifyRequest) Set(val *OrderModifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyRequest(val *OrderModifyRequest) *NullableOrderModifyRequest {
	return &NullableOrderModifyRequest{value: val, isSet: true}
}

func (v NullableOrderModifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


