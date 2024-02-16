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

// checks if the AvailabilityAsyncNotificationRequestResourceInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityAsyncNotificationRequestResourceInner{}

// AvailabilityAsyncNotificationRequestResourceInner struct for AvailabilityAsyncNotificationRequestResourceInner
type AvailabilityAsyncNotificationRequestResourceInner struct {
	// The event name sent in the event request.
	EventType *string `json:"eventType,omitempty"`
	// The Unique IngramMicro part number for the product.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The vendors part number for the product.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// The name of the vendor/manufacturer of the product.
	VendorName *string `json:"vendorName,omitempty"`
	// The UPC code for the product. Consists of 12 numeric digits that are uniquly assigned to each trade item.
	UpcCode *string `json:"upcCode,omitempty"`
	// Status returned saying whether sku is active.
	SkuStatus *string `json:"skuStatus,omitempty"`
	// Backordered Flag.
	BackOrderFlag *string `json:"backOrderFlag,omitempty"`
	// totalAvailability.
	TotalAvailability *string `json:"totalAvailability,omitempty"`
	// Link to Order Details for the order(s).
	Links []AvailabilityAsyncNotificationRequestResourceInnerLinksInner `json:"links,omitempty"`
}

// NewAvailabilityAsyncNotificationRequestResourceInner instantiates a new AvailabilityAsyncNotificationRequestResourceInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityAsyncNotificationRequestResourceInner() *AvailabilityAsyncNotificationRequestResourceInner {
	this := AvailabilityAsyncNotificationRequestResourceInner{}
	return &this
}

// NewAvailabilityAsyncNotificationRequestResourceInnerWithDefaults instantiates a new AvailabilityAsyncNotificationRequestResourceInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityAsyncNotificationRequestResourceInnerWithDefaults() *AvailabilityAsyncNotificationRequestResourceInner {
	this := AvailabilityAsyncNotificationRequestResourceInner{}
	return &this
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetEventType(v string) {
	o.EventType = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetVendorName(v string) {
	o.VendorName = &v
}

// GetUpcCode returns the UpcCode field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetUpcCode() string {
	if o == nil || IsNil(o.UpcCode) {
		var ret string
		return ret
	}
	return *o.UpcCode
}

// GetUpcCodeOk returns a tuple with the UpcCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetUpcCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UpcCode) {
		return nil, false
	}
	return o.UpcCode, true
}

// HasUpcCode returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasUpcCode() bool {
	if o != nil && !IsNil(o.UpcCode) {
		return true
	}

	return false
}

// SetUpcCode gets a reference to the given string and assigns it to the UpcCode field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetUpcCode(v string) {
	o.UpcCode = &v
}

// GetSkuStatus returns the SkuStatus field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetSkuStatus() string {
	if o == nil || IsNil(o.SkuStatus) {
		var ret string
		return ret
	}
	return *o.SkuStatus
}

// GetSkuStatusOk returns a tuple with the SkuStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetSkuStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SkuStatus) {
		return nil, false
	}
	return o.SkuStatus, true
}

// HasSkuStatus returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasSkuStatus() bool {
	if o != nil && !IsNil(o.SkuStatus) {
		return true
	}

	return false
}

// SetSkuStatus gets a reference to the given string and assigns it to the SkuStatus field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetSkuStatus(v string) {
	o.SkuStatus = &v
}

// GetBackOrderFlag returns the BackOrderFlag field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetBackOrderFlag() string {
	if o == nil || IsNil(o.BackOrderFlag) {
		var ret string
		return ret
	}
	return *o.BackOrderFlag
}

// GetBackOrderFlagOk returns a tuple with the BackOrderFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetBackOrderFlagOk() (*string, bool) {
	if o == nil || IsNil(o.BackOrderFlag) {
		return nil, false
	}
	return o.BackOrderFlag, true
}

// HasBackOrderFlag returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasBackOrderFlag() bool {
	if o != nil && !IsNil(o.BackOrderFlag) {
		return true
	}

	return false
}

// SetBackOrderFlag gets a reference to the given string and assigns it to the BackOrderFlag field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetBackOrderFlag(v string) {
	o.BackOrderFlag = &v
}

// GetTotalAvailability returns the TotalAvailability field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetTotalAvailability() string {
	if o == nil || IsNil(o.TotalAvailability) {
		var ret string
		return ret
	}
	return *o.TotalAvailability
}

// GetTotalAvailabilityOk returns a tuple with the TotalAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetTotalAvailabilityOk() (*string, bool) {
	if o == nil || IsNil(o.TotalAvailability) {
		return nil, false
	}
	return o.TotalAvailability, true
}

// HasTotalAvailability returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasTotalAvailability() bool {
	if o != nil && !IsNil(o.TotalAvailability) {
		return true
	}

	return false
}

// SetTotalAvailability gets a reference to the given string and assigns it to the TotalAvailability field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetTotalAvailability(v string) {
	o.TotalAvailability = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetLinks() []AvailabilityAsyncNotificationRequestResourceInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret []AvailabilityAsyncNotificationRequestResourceInnerLinksInner
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) GetLinksOk() ([]AvailabilityAsyncNotificationRequestResourceInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AvailabilityAsyncNotificationRequestResourceInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []AvailabilityAsyncNotificationRequestResourceInnerLinksInner and assigns it to the Links field.
func (o *AvailabilityAsyncNotificationRequestResourceInner) SetLinks(v []AvailabilityAsyncNotificationRequestResourceInnerLinksInner) {
	o.Links = v
}

func (o AvailabilityAsyncNotificationRequestResourceInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailabilityAsyncNotificationRequestResourceInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.UpcCode) {
		toSerialize["upcCode"] = o.UpcCode
	}
	if !IsNil(o.SkuStatus) {
		toSerialize["skuStatus"] = o.SkuStatus
	}
	if !IsNil(o.BackOrderFlag) {
		toSerialize["backOrderFlag"] = o.BackOrderFlag
	}
	if !IsNil(o.TotalAvailability) {
		toSerialize["totalAvailability"] = o.TotalAvailability
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAvailabilityAsyncNotificationRequestResourceInner struct {
	value *AvailabilityAsyncNotificationRequestResourceInner
	isSet bool
}

func (v NullableAvailabilityAsyncNotificationRequestResourceInner) Get() *AvailabilityAsyncNotificationRequestResourceInner {
	return v.value
}

func (v *NullableAvailabilityAsyncNotificationRequestResourceInner) Set(val *AvailabilityAsyncNotificationRequestResourceInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityAsyncNotificationRequestResourceInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityAsyncNotificationRequestResourceInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityAsyncNotificationRequestResourceInner(val *AvailabilityAsyncNotificationRequestResourceInner) *NullableAvailabilityAsyncNotificationRequestResourceInner {
	return &NullableAvailabilityAsyncNotificationRequestResourceInner{value: val, isSet: true}
}

func (v NullableAvailabilityAsyncNotificationRequestResourceInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityAsyncNotificationRequestResourceInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


