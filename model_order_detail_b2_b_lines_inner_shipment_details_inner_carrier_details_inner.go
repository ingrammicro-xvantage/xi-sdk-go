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

// checks if the OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner{}

// OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner struct for OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner
type OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner struct {
	// The carrier code for the shipment containing the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the carrier of the shipment containing the line item.
	CarrierName *string `json:"carrierName,omitempty"`
	// The quantity shipped of the line item.
	Quantity *int32 `json:"quantity,omitempty"`
	// The actual date when line item shipped.
	ShippedDate *string `json:"shippedDate,omitempty"`
	// The date the line item is expected to be delivered.
	EstimatedDeliveryDate *string `json:"estimatedDeliveryDate,omitempty"`
	// The actual date of delivery of the line item.
	DeliveredDate *string `json:"deliveredDate,omitempty"`
	// The actual date when carrier picked up line item.
	CarrierPickupDate *string `json:"carrierPickupDate,omitempty"`
	// The tracking details for the shipment containing the line item.
	TrackingDetails []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner `json:"trackingDetails,omitempty"`
}

// NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner() *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner {
	this := OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner{}
	return &this
}

// NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerWithDefaults instantiates a new OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerWithDefaults() *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner {
	this := OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetShippedDate returns the ShippedDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetShippedDate() string {
	if o == nil || IsNil(o.ShippedDate) {
		var ret string
		return ret
	}
	return *o.ShippedDate
}

// GetShippedDateOk returns a tuple with the ShippedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetShippedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ShippedDate) {
		return nil, false
	}
	return o.ShippedDate, true
}

// HasShippedDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasShippedDate() bool {
	if o != nil && !IsNil(o.ShippedDate) {
		return true
	}

	return false
}

// SetShippedDate gets a reference to the given string and assigns it to the ShippedDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetShippedDate(v string) {
	o.ShippedDate = &v
}

// GetEstimatedDeliveryDate returns the EstimatedDeliveryDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetEstimatedDeliveryDate() string {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		var ret string
		return ret
	}
	return *o.EstimatedDeliveryDate
}

// GetEstimatedDeliveryDateOk returns a tuple with the EstimatedDeliveryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetEstimatedDeliveryDateOk() (*string, bool) {
	if o == nil || IsNil(o.EstimatedDeliveryDate) {
		return nil, false
	}
	return o.EstimatedDeliveryDate, true
}

// HasEstimatedDeliveryDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasEstimatedDeliveryDate() bool {
	if o != nil && !IsNil(o.EstimatedDeliveryDate) {
		return true
	}

	return false
}

// SetEstimatedDeliveryDate gets a reference to the given string and assigns it to the EstimatedDeliveryDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetEstimatedDeliveryDate(v string) {
	o.EstimatedDeliveryDate = &v
}

// GetDeliveredDate returns the DeliveredDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetDeliveredDate() string {
	if o == nil || IsNil(o.DeliveredDate) {
		var ret string
		return ret
	}
	return *o.DeliveredDate
}

// GetDeliveredDateOk returns a tuple with the DeliveredDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetDeliveredDateOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveredDate) {
		return nil, false
	}
	return o.DeliveredDate, true
}

// HasDeliveredDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasDeliveredDate() bool {
	if o != nil && !IsNil(o.DeliveredDate) {
		return true
	}

	return false
}

// SetDeliveredDate gets a reference to the given string and assigns it to the DeliveredDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetDeliveredDate(v string) {
	o.DeliveredDate = &v
}

// GetCarrierPickupDate returns the CarrierPickupDate field value if set, zero value otherwise.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierPickupDate() string {
	if o == nil || IsNil(o.CarrierPickupDate) {
		var ret string
		return ret
	}
	return *o.CarrierPickupDate
}

// GetCarrierPickupDateOk returns a tuple with the CarrierPickupDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetCarrierPickupDateOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierPickupDate) {
		return nil, false
	}
	return o.CarrierPickupDate, true
}

// HasCarrierPickupDate returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasCarrierPickupDate() bool {
	if o != nil && !IsNil(o.CarrierPickupDate) {
		return true
	}

	return false
}

// SetCarrierPickupDate gets a reference to the given string and assigns it to the CarrierPickupDate field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetCarrierPickupDate(v string) {
	o.CarrierPickupDate = &v
}

// GetTrackingDetails returns the TrackingDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetTrackingDetails() []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner {
	if o == nil {
		var ret []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner
		return ret
	}
	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) GetTrackingDetailsOk() ([]OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner, bool) {
	if o == nil || IsNil(o.TrackingDetails) {
		return nil, false
	}
	return o.TrackingDetails, true
}

// HasTrackingDetails returns a boolean if a field has been set.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) HasTrackingDetails() bool {
	if o != nil && IsNil(o.TrackingDetails) {
		return true
	}

	return false
}

// SetTrackingDetails gets a reference to the given []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner and assigns it to the TrackingDetails field.
func (o *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) SetTrackingDetails(v []OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInnerTrackingDetailsInner) {
	o.TrackingDetails = v
}

func (o OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
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
	if !IsNil(o.CarrierPickupDate) {
		toSerialize["carrierPickupDate"] = o.CarrierPickupDate
	}
	if o.TrackingDetails != nil {
		toSerialize["trackingDetails"] = o.TrackingDetails
	}
	return toSerialize, nil
}

type NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner struct {
	value *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner
	isSet bool
}

func (v NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) Get() *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner {
	return v.value
}

func (v *NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) Set(val *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner(val *OrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) *NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner {
	return &NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner{value: val, isSet: true}
}

func (v NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetailB2BLinesInnerShipmentDetailsInnerCarrierDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


