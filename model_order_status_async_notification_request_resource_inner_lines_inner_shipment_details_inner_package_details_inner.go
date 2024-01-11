/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
)

// checks if the OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner{}

// OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner struct for OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner
type OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner struct {
	// The shipment carton number that contains the line item.
	CartonNumber *string `json:"cartonNumber,omitempty"`
	// The quantity of line items in the box.
	QuantityInbox *string `json:"quantityInbox,omitempty"`
	// The tracking number for the shipment containing the line item.
	TrackingNumber *string `json:"trackingNumber,omitempty"`
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner{}
	return &this
}

// NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInnerWithDefaults instantiates a new OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInnerWithDefaults() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner {
	this := OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner{}
	return &this
}

// GetCartonNumber returns the CartonNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) GetCartonNumber() string {
	if o == nil || IsNil(o.CartonNumber) {
		var ret string
		return ret
	}
	return *o.CartonNumber
}

// GetCartonNumberOk returns a tuple with the CartonNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) GetCartonNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CartonNumber) {
		return nil, false
	}
	return o.CartonNumber, true
}

// HasCartonNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) HasCartonNumber() bool {
	if o != nil && !IsNil(o.CartonNumber) {
		return true
	}

	return false
}

// SetCartonNumber gets a reference to the given string and assigns it to the CartonNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) SetCartonNumber(v string) {
	o.CartonNumber = &v
}

// GetQuantityInbox returns the QuantityInbox field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) GetQuantityInbox() string {
	if o == nil || IsNil(o.QuantityInbox) {
		var ret string
		return ret
	}
	return *o.QuantityInbox
}

// GetQuantityInboxOk returns a tuple with the QuantityInbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) GetQuantityInboxOk() (*string, bool) {
	if o == nil || IsNil(o.QuantityInbox) {
		return nil, false
	}
	return o.QuantityInbox, true
}

// HasQuantityInbox returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) HasQuantityInbox() bool {
	if o != nil && !IsNil(o.QuantityInbox) {
		return true
	}

	return false
}

// SetQuantityInbox gets a reference to the given string and assigns it to the QuantityInbox field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) SetQuantityInbox(v string) {
	o.QuantityInbox = &v
}

// GetTrackingNumber returns the TrackingNumber field value if set, zero value otherwise.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) GetTrackingNumber() string {
	if o == nil || IsNil(o.TrackingNumber) {
		var ret string
		return ret
	}
	return *o.TrackingNumber
}

// GetTrackingNumberOk returns a tuple with the TrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) GetTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumber) {
		return nil, false
	}
	return o.TrackingNumber, true
}

// HasTrackingNumber returns a boolean if a field has been set.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) HasTrackingNumber() bool {
	if o != nil && !IsNil(o.TrackingNumber) {
		return true
	}

	return false
}

// SetTrackingNumber gets a reference to the given string and assigns it to the TrackingNumber field.
func (o *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) SetTrackingNumber(v string) {
	o.TrackingNumber = &v
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CartonNumber) {
		toSerialize["cartonNumber"] = o.CartonNumber
	}
	if !IsNil(o.QuantityInbox) {
		toSerialize["quantityInbox"] = o.QuantityInbox
	}
	if !IsNil(o.TrackingNumber) {
		toSerialize["trackingNumber"] = o.TrackingNumber
	}
	return toSerialize, nil
}

type NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner struct {
	value *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner
	isSet bool
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) Get() *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner {
	return v.value
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) Set(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner(val *OrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner {
	return &NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner{value: val, isSet: true}
}

func (v NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderStatusAsyncNotificationRequestResourceInnerLinesInnerShipmentDetailsInnerPackageDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


