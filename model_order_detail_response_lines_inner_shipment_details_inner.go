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

// checks if the OrderDetailResponseLinesInnerShipmentDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailResponseLinesInnerShipmentDetailsInner{}

// OrderDetailResponseLinesInnerShipmentDetailsInner Shipping details for the line item.
type OrderDetailResponseLinesInnerShipmentDetailsInner struct {
	// The quantity shipped of the line item.
	Quantity *int32 `json:"quantity,omitempty"`
	// The estimated ship date for the line item.
	EstimatedShipDate *string `json:"estimatedShipDate,omitempty"`
	// The date the line item was shipped.
	ShippedDate *string `json:"shippedDate,omitempty"`
	// The date the line item is expected to be delivered.
	EstimatedDeliveryDate *string `json:"estimatedDeliveryDate,omitempty"`
	// The actual date of delivery of the line item.
	DeliveredDate *string `json:"deliveredDate,omitempty"`
	// The ID of the warehouse the product will ship from.
	ShipFromWarehouseId *string `json:"shipFromWarehouseId,omitempty"`
	// The city and state the line item ships from.
	ShipFromLocation *string `json:"shipFromLocation,omitempty"`
	// The Ingram Micro invoice number for the line item.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// The date the IngramMicro invoice was created for the line item.
	InvoiceDate *string `json:"invoiceDate,omitempty"`
	CarrierDetails *OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails `json:"carrierDetails,omitempty"`
}

// NewOrderDetailResponseLinesInnerShipmentDetailsInner instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailResponseLinesInnerShipmentDetailsInner() *OrderDetailResponseLinesInnerShipmentDetailsInner {
	this := OrderDetailResponseLinesInnerShipmentDetailsInner{}
	return &this
}

// NewOrderDetailResponseLinesInnerShipmentDetailsInnerWithDefaults instantiates a new OrderDetailResponseLinesInnerShipmentDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailResponseLinesInnerShipmentDetailsInnerWithDefaults() *OrderDetailResponseLinesInnerShipmentDetailsInner {
	this := OrderDetailResponseLinesInnerShipmentDetailsInner{}
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetEstimatedShipDate returns the EstimatedShipDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedShipDate() string {
	if o == nil || IsNil(o.EstimatedShipDate) {
		var ret string
		return ret
	}
	return *o.EstimatedShipDate
}

// GetEstimatedShipDateOk returns a tuple with the EstimatedShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedShipDateOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedShipDate) {
		return nil, false
	}
	return o.EstimatedShipDate, true
}

// HasEstimatedShipDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasEstimatedShipDate() bool {
	if o != nil && !IsNil(o.EstimatedShipDate) {
		return true
	}

	return false
}

// SetEstimatedShipDate gets a reference to the given string and assigns it to the EstimatedShipDate field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetEstimatedShipDate(v string) {
	o.EstimatedShipDate = &v
}

// GetShippedDate returns the ShippedDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShippedDate() string {
	if o == nil || IsNil(o.ShippedDate) {
		var ret string
		return ret
	}
	return *o.ShippedDate
}

// GetShippedDateOk returns a tuple with the ShippedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShippedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShippedDate) {
		return nil, false
	}
	return o.ShippedDate, true
}

// HasShippedDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasShippedDate() bool {
	if o != nil && !IsNil(o.ShippedDate) {
		return true
	}

	return false
}

// SetShippedDate gets a reference to the given string and assigns it to the ShippedDate field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetShippedDate(v string) {
	o.ShippedDate = &v
}

// GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDate() string {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.EstimatedDeliveryDate
}

// GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetEstimatedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		return nil, false
	}
	return o.EstimatedDeliveryDate, true
}

// HasEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDate gets a reference to the given string and assigns it to the EstimatedDeliveryDate field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetEstimatedDeliveryDate(v string) {
	o.EstimatedDeliveryDate = &v
}

// GetDeliveredDate returns the DeliveredDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetDeliveredDate() string {
	if o == nil || IsNil(o.DeliveredDate) {
		var ret string
		return ret
	}
	return *o.DeliveredDate
}

// GetDeliveredDateOk returns a tuple with the DeliveredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetDeliveredDateOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveredDate) {
		return nil, false
	}
	return o.DeliveredDate, true
}

// HasDeliveredDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasDeliveredDate() bool {
	if o != nil && !IsNil(o.DeliveredDate) {
		return true
	}

	return false
}

// SetDeliveredDate gets a reference to the given string and assigns it to the DeliveredDate field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetDeliveredDate(v string) {
	o.DeliveredDate = &v
}

// GetShipFromWarehouseId returns the ShipFromWarehouseId field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromWarehouseId() string {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		var ret string
		return ret
	}
	return *o.ShipFromWarehouseId
}

// GetShipFromWarehouseIdOk returns a tuple with the ShipFromWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromWarehouseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromWarehouseId) {
		return nil, false
	}
	return o.ShipFromWarehouseId, true
}

// HasShipFromWarehouseId returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasShipFromWarehouseId() bool {
	if o != nil && !IsNil(o.ShipFromWarehouseId) {
		return true
	}

	return false
}

// SetShipFromWarehouseId gets a reference to the given string and assigns it to the ShipFromWarehouseId field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetShipFromWarehouseId(v string) {
	o.ShipFromWarehouseId = &v
}

// GetShipFromLocation returns the ShipFromLocation field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromLocation() string {
	if o == nil || IsNil(o.ShipFromLocation) {
		var ret string
		return ret
	}
	return *o.ShipFromLocation
}

// GetShipFromLocationOk returns a tuple with the ShipFromLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetShipFromLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromLocation) {
		return nil, false
	}
	return o.ShipFromLocation, true
}

// HasShipFromLocation returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasShipFromLocation() bool {
	if o != nil && !IsNil(o.ShipFromLocation) {
		return true
	}

	return false
}

// SetShipFromLocation gets a reference to the given string and assigns it to the ShipFromLocation field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetShipFromLocation(v string) {
	o.ShipFromLocation = &v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetCarrierDetails returns the CarrierDetails field value if set, zero value otherwise.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetCarrierDetails() OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails {
	if o == nil || IsNil(o.CarrierDetails) {
		var ret OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails
		return ret
	}
	return *o.CarrierDetails
}

// GetCarrierDetailsOk returns a tuple with the CarrierDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) GetCarrierDetailsOk() (*OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails, bool) {
	if o == nil || IsNil(o.CarrierDetails) {
		return nil, false
	}
	return o.CarrierDetails, true
}

// HasCarrierDetails returns a boolean if a field has been set.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) HasCarrierDetails() bool {
	if o != nil && !IsNil(o.CarrierDetails) {
		return true
	}

	return false
}

// SetCarrierDetails gets a reference to the given OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails and assigns it to the CarrierDetails field.
func (o *OrderDetailResponseLinesInnerShipmentDetailsInner) SetCarrierDetails(v OrderDetailResponseLinesInnerShipmentDetailsInnerCarrierDetails) {
	o.CarrierDetails = &v
}

func (o OrderDetailResponseLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailResponseLinesInnerShipmentDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
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
	if !IsNil(o.DeliveredDate) {
		toSerialize["deliveredDate"] = o.DeliveredDate
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
	if !IsNil(o.CarrierDetails) {
		toSerialize["carrierDetails"] = o.CarrierDetails
	}
	return toSerialize, nil
}

type NullableOrderDetailResponseLinesInnerShipmentDetailsInner struct {
	value *OrderDetailResponseLinesInnerShipmentDetailsInner
	isSet bool
}

func (v NullableOrderDetailResponseLinesInnerShipmentDetailsInner) Get() *OrderDetailResponseLinesInnerShipmentDetailsInner {
	return v.value
}

func (v *NullableOrderDetailResponseLinesInnerShipmentDetailsInner) Set(val *OrderDetailResponseLinesInnerShipmentDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailResponseLinesInnerShipmentDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailResponseLinesInnerShipmentDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailResponseLinesInnerShipmentDetailsInner(val *OrderDetailResponseLinesInnerShipmentDetailsInner) *NullableOrderDetailResponseLinesInnerShipmentDetailsInner {
	return &NullableOrderDetailResponseLinesInnerShipmentDetailsInner{value: val, isSet: true}
}

func (v NullableOrderDetailResponseLinesInnerShipmentDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailResponseLinesInnerShipmentDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


