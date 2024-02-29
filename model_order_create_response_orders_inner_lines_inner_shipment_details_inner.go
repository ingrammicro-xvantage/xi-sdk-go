/*
XI Sdk Resellers

For Ingram Micro Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner{}

// OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner struct for OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner
type OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner struct {
	// The code for the shipping carrier for the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the shipping carrier for the line item.
	CarrierName *string `json:"carrierName,omitempty"`
	// The ID of the warehouse the line item will ship from.
	ShipFromWarehouseId *string `json:"shipFromWarehouseId,omitempty"`
	// Location from which order is shipped.
	ShipFromLocation *string `json:"shipFromLocation,omitempty"`
	// The reseller's shipping account number with carrier. Used to bill the shipping carrier directly from the reseller's account with the carrier.
	FreightAccountNumber *string `json:"freightAccountNumber,omitempty"`
	// Specifies whether a signature is required for delivery. Default is False.
	SignatureRequired *string `json:"signatureRequired,omitempty"`
	// The shipping instructions for the order.
	ShippingInstructions *string `json:"shippingInstructions,omitempty"`
}

// NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner instantiates a new OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner() *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner {
	this := OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner{}
	return &this
}

// NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInnerWithDefaults() *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner {
	this := OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetShipFromWarehouseId returns the ShipFromWarehouseId field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		var ret string
		return ret
	}
	return *o.ShipFromWarehouseId
}

// GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		return nil, false
	}
	return o.ShipFromWarehouseId, true
}

// HasShipFromWarehouseId returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool {
	if o != nil && !IsNil(o.ShipFromWarehouseId) {
		return true
	}

	return false
}

// SetShipFromWarehouseId gets a reference to the given string and assigns it to the ShipFromWarehouseId field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string) {
	o.ShipFromWarehouseId = &v
}

// GetShipFromLocation returns the ShipFromLocation field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromLocation() string {
	if o == nil || IsNil(o.ShipFromLocation) {
		var ret string
		return ret
	}
	return *o.ShipFromLocation
}

// GetShipFromLocationOk returns a tuple with the ShipFromLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShipFromLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromLocation) {
		return nil, false
	}
	return o.ShipFromLocation, true
}

// HasShipFromLocation returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasShipFromLocation() bool {
	if o != nil && !IsNil(o.ShipFromLocation) {
		return true
	}

	return false
}

// SetShipFromLocation gets a reference to the given string and assigns it to the ShipFromLocation field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetShipFromLocation(v string) {
	o.ShipFromLocation = &v
}

// GetFreightAccountNumber returns the FreightAccountNumber field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetFreightAccountNumber() string {
	if o == nil || IsNil(o.FreightAccountNumber) {
		var ret string
		return ret
	}
	return *o.FreightAccountNumber
}

// GetFreightAccountNumberOk returns a tuple with the FreightAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetFreightAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FreightAccountNumber) {
		return nil, false
	}
	return o.FreightAccountNumber, true
}

// HasFreightAccountNumber returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasFreightAccountNumber() bool {
	if o != nil && !IsNil(o.FreightAccountNumber) {
		return true
	}

	return false
}

// SetFreightAccountNumber gets a reference to the given string and assigns it to the FreightAccountNumber field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetFreightAccountNumber(v string) {
	o.FreightAccountNumber = &v
}

// GetSignatureRequired returns the SignatureRequired field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetSignatureRequired() string {
	if o == nil || IsNil(o.SignatureRequired) {
		var ret string
		return ret
	}
	return *o.SignatureRequired
}

// GetSignatureRequiredOk returns a tuple with the SignatureRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetSignatureRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureRequired) {
		return nil, false
	}
	return o.SignatureRequired, true
}

// HasSignatureRequired returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasSignatureRequired() bool {
	if o != nil && !IsNil(o.SignatureRequired) {
		return true
	}

	return false
}

// SetSignatureRequired gets a reference to the given string and assigns it to the SignatureRequired field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetSignatureRequired(v string) {
	o.SignatureRequired = &v
}

// GetShippingInstructions returns the ShippingInstructions field value if set, zero value otherwise.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShippingInstructions() string {
	if o == nil || IsNil(o.ShippingInstructions) {
		var ret string
		return ret
	}
	return *o.ShippingInstructions
}

// GetShippingInstructionsOk returns a tuple with the ShippingInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) GetShippingInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingInstructions) {
		return nil, false
	}
	return o.ShippingInstructions, true
}

// HasShippingInstructions returns a boolean if a field has been set.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) HasShippingInstructions() bool {
	if o != nil && !IsNil(o.ShippingInstructions) {
		return true
	}

	return false
}

// SetShippingInstructions gets a reference to the given string and assigns it to the ShippingInstructions field.
func (o *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) SetShippingInstructions(v string) {
	o.ShippingInstructions = &v
}

func (o OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.ShipFromWarehouseId) {
		toSerialize["shipFromWarehouseId"] = o.ShipFromWarehouseId
	}
	if !IsNil(o.ShipFromLocation) {
		toSerialize["shipFromLocation"] = o.ShipFromLocation
	}
	if !IsNil(o.FreightAccountNumber) {
		toSerialize["freightAccountNumber"] = o.FreightAccountNumber
	}
	if !IsNil(o.SignatureRequired) {
		toSerialize["signatureRequired"] = o.SignatureRequired
	}
	if !IsNil(o.ShippingInstructions) {
		toSerialize["shippingInstructions"] = o.ShippingInstructions
	}
	return toSerialize, nil
}

type NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner struct {
	value *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner
	isSet bool
}

func (v NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) Get() *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner {
	return v.value
}

func (v *NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) Set(val *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner(val *OrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) *NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner {
	return &NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner{value: val, isSet: true}
}

func (v NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateResponseOrdersInnerLinesInnerShipmentDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


