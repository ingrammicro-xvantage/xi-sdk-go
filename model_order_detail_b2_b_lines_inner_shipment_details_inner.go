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

// checks if the OrderDetailB2BLinesInnerShipmentDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerShipmentDetailsInner{}

// OrderDetailB2BLinesInnerShipmentDetailsInner struct for OrderDetailB2BLinesInnerShipmentDetailsInner
type OrderDetailB2BLinesInnerShipmentDetailsInner struct {
	// The quantity shipped of the line item.
	Quantity NullableInt32 `json:"quantity,omitempty"`
	// The actual date of delivery of the line item.
	DeliveryNumber *string `json:"deliveryNumber,omitempty"`
	// The date the line item is expected to be shipped.
	EstimatedShipDate *string `json:"estimatedShipDate,omitempty"`
	ShippedDate *string `json:"shippedDate,omitempty"`
	EstimatedDeliveryDate *string `json:"estimatedDeliveryDate,omitempty"`
	// The ID of the warehouse the product will ship from.
	ShipFromWarehouseId *string `json:"shipFromWarehouseId,omitempty"`
	// The city and state the line item ships from.
	ShipFromLocation *string `json:"shipFromLocation,omitempty"`
	// The Ingram Micro invoice number for the line item.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// The date the IngramMicro invoice was created for the line item.
	InvoiceDate *string `json:"invoiceDate,omitempty"`
	// The shipment carrier details for the line item.
	CarrierDetails []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner `json:"carrierDetails,omitempty"`
}

// NewOrderDetailB2BLinesInnerShipmentDetailsInner instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerShipmentDetailsInner() *OrderDetailB2BLinesInnerShipmentDetailsInner {
	this := OrderDetailB2BLinesInnerShipmentDetailsInner{}
	return &this
}

// NewOrderDetailB2BLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerShipmentDetailsInnerWithDefaults() *OrderDetailB2BLinesInnerShipmentDetailsInner {
	this := OrderDetailB2BLinesInnerShipmentDetailsInner{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetDeliveryNumber returns the DeliveryNumber field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetDeliveryNumber() string {
	if o == nil || IsNil(o.DeliveryNumber) {
		var ret string
		return ret
	}
	return *o.DeliveryNumber
}

// GetDeliveryNumberOk returns a tuple with the DeliveryNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetDeliveryNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryNumber) {
		return nil, false
	}
	return o.DeliveryNumber, true
}

// HasDeliveryNumber returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasDeliveryNumber() bool {
	if o != nil && !IsNil(o.DeliveryNumber) {
		return true
	}

	return false
}

// SetDeliveryNumber gets a reference to the given string and assigns it to the DeliveryNumber field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetDeliveryNumber(v string) {
	o.DeliveryNumber = &v
}

// GetEstimatedShipDate returns the EstimatedShipDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedShipDate() string {
	if o == nil || IsNil(o.EstimatedShipDate) {
		var ret string
		return ret
	}
	return *o.EstimatedShipDate
}

// GetEstimatedShipDateOk returns a tuple with the EstimatedShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedShipDateOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedShipDate) {
		return nil, false
	}
	return o.EstimatedShipDate, true
}

// HasEstimatedShipDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasEstimatedShipDate() bool {
	if o != nil && !IsNil(o.EstimatedShipDate) {
		return true
	}

	return false
}

// SetEstimatedShipDate gets a reference to the given string and assigns it to the EstimatedShipDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetEstimatedShipDate(v string) {
	o.EstimatedShipDate = &v
}

// GetShippedDate returns the ShippedDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShippedDate() string {
	if o == nil || IsNil(o.ShippedDate) {
		var ret string
		return ret
	}
	return *o.ShippedDate
}

// GetShippedDateOk returns a tuple with the ShippedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShippedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShippedDate) {
		return nil, false
	}
	return o.ShippedDate, true
}

// HasShippedDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasShippedDate() bool {
	if o != nil && !IsNil(o.ShippedDate) {
		return true
	}

	return false
}

// SetShippedDate gets a reference to the given string and assigns it to the ShippedDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetShippedDate(v string) {
	o.ShippedDate = &v
}

// GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDate() string {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.EstimatedDeliveryDate
}

// GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		return nil, false
	}
	return o.EstimatedDeliveryDate, true
}

// HasEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDate gets a reference to the given string and assigns it to the EstimatedDeliveryDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetEstimatedDeliveryDate(v string) {
	o.EstimatedDeliveryDate = &v
}

// GetShipFromWarehouseId returns the ShipFromWarehouseId field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		var ret string
		return ret
	}
	return *o.ShipFromWarehouseId
}

// GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		return nil, false
	}
	return o.ShipFromWarehouseId, true
}

// HasShipFromWarehouseId returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool {
	if o != nil && !IsNil(o.ShipFromWarehouseId) {
		return true
	}

	return false
}

// SetShipFromWarehouseId gets a reference to the given string and assigns it to the ShipFromWarehouseId field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string) {
	o.ShipFromWarehouseId = &v
}

// GetShipFromLocation returns the ShipFromLocation field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromLocation() string {
	if o == nil || IsNil(o.ShipFromLocation) {
		var ret string
		return ret
	}
	return *o.ShipFromLocation
}

// GetShipFromLocationOk returns a tuple with the ShipFromLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetShipFromLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromLocation) {
		return nil, false
	}
	return o.ShipFromLocation, true
}

// HasShipFromLocation returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasShipFromLocation() bool {
	if o != nil && !IsNil(o.ShipFromLocation) {
		return true
	}

	return false
}

// SetShipFromLocation gets a reference to the given string and assigns it to the ShipFromLocation field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetShipFromLocation(v string) {
	o.ShipFromLocation = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetCarrierDetails returns the CarrierDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetCarrierDetails() []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner {
	if o == nil {
		var ret []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner
		return ret
	}
	return o.CarrierDetails
}

// GetCarrierDetailsOk returns a tuple with the CarrierDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) GetCarrierDetailsOk() ([]OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner, bool) {
	if o == nil || IsNil(o.CarrierDetails) {
		return nil, false
	}
	return o.CarrierDetails, true
}

// HasCarrierDetails returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) HasCarrierDetails() bool {
	if o != nil && !IsNil(o.CarrierDetails) {
		return true
	}

	return false
}

// SetCarrierDetails gets a reference to the given []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner and assigns it to the CarrierDetails field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInner) SetCarrierDetails(v []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) {
	o.CarrierDetails = v
}

func (o OrderDetailB2BLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerShipmentDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
	if !IsNil(o.DeliveryNumber) {
		toSerialize["deliveryNumber"] = o.DeliveryNumber
	}
	if !IsNil(o.EstimatedShipDate) {
		toSerialize["estimatedShipDate"] = o.EstimatedShipDate
	}
	if !IsNil(o.ShippedDate) {
		toSerialize["shippedDate"] = o.ShippedDate
	}
	if !IsNil(o.EstimatedDeliveryDate) {
		toSerialize["estimatedDeliveryDate"] = o.EstimatedDeliveryDate
	}
	if !IsNil(o.ShipFromWarehouseId) {
		toSerialize["shipFromWarehouseId"] = o.ShipFromWarehouseId
	}
	if !IsNil(o.ShipFromLocation) {
		toSerialize["shipFromLocation"] = o.ShipFromLocation
	}
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoiceDate"] = o.InvoiceDate
	}
	if o.CarrierDetails != nil {
		toSerialize["carrierDetails"] = o.CarrierDetails
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerShipmentDetailsInner struct {
	value *OrderDetailB2BLinesInnerShipmentDetailsInner
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerShipmentDetailsInner) Get() *OrderDetailB2BLinesInnerShipmentDetailsInner {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerShipmentDetailsInner) Set(val *OrderDetailB2BLinesInnerShipmentDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerShipmentDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerShipmentDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerShipmentDetailsInner(val *OrderDetailB2BLinesInnerShipmentDetailsInner) *NullableOrderDetailB2BLinesInnerShipmentDetailsInner {
	return &NullableOrderDetailB2BLinesInnerShipmentDetailsInner{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerShipmentDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


