/*
Reseller API

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi.sdk.resellers

import (
	"encoding/json"
)

// checks if the RenewalsDetailsResponseReferenceNumberInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsDetailsResponseReferenceNumberInner{}

// RenewalsDetailsResponseReferenceNumberInner struct for RenewalsDetailsResponseReferenceNumberInner
type RenewalsDetailsResponseReferenceNumberInner struct {
	// Notification id of the communication sent from Ingram.
	NotificationId *string `json:"notificationId,omitempty"`
	// Quote number for the renewal.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
}

// NewRenewalsDetailsResponseReferenceNumberInner instantiates a new RenewalsDetailsResponseReferenceNumberInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsDetailsResponseReferenceNumberInner() *RenewalsDetailsResponseReferenceNumberInner {
	this := RenewalsDetailsResponseReferenceNumberInner{}
	return &this
}

// NewRenewalsDetailsResponseReferenceNumberInnerWithDefaults instantiates a new RenewalsDetailsResponseReferenceNumberInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsDetailsResponseReferenceNumberInnerWithDefaults() *RenewalsDetailsResponseReferenceNumberInner {
	this := RenewalsDetailsResponseReferenceNumberInner{}
	return &this
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseReferenceNumberInner) GetNotificationId() string {
	if o == nil || IsNil(o.NotificationId) {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseReferenceNumberInner) GetNotificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationId) {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseReferenceNumberInner) HasNotificationId() bool {
	if o != nil && !IsNil(o.NotificationId) {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *RenewalsDetailsResponseReferenceNumberInner) SetNotificationId(v string) {
	o.NotificationId = &v
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseReferenceNumberInner) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseReferenceNumberInner) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseReferenceNumberInner) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *RenewalsDetailsResponseReferenceNumberInner) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

func (o RenewalsDetailsResponseReferenceNumberInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsDetailsResponseReferenceNumberInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationId) {
		toSerialize["notificationId"] = o.NotificationId
	}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	return toSerialize, nil
}

type NullableRenewalsDetailsResponseReferenceNumberInner struct {
	value *RenewalsDetailsResponseReferenceNumberInner
	isSet bool
}

func (v NullableRenewalsDetailsResponseReferenceNumberInner) Get() *RenewalsDetailsResponseReferenceNumberInner {
	return v.value
}

func (v *NullableRenewalsDetailsResponseReferenceNumberInner) Set(val *RenewalsDetailsResponseReferenceNumberInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsDetailsResponseReferenceNumberInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsDetailsResponseReferenceNumberInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsDetailsResponseReferenceNumberInner(val *RenewalsDetailsResponseReferenceNumberInner) *NullableRenewalsDetailsResponseReferenceNumberInner {
	return &NullableRenewalsDetailsResponseReferenceNumberInner{value: val, isSet: true}
}

func (v NullableRenewalsDetailsResponseReferenceNumberInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsDetailsResponseReferenceNumberInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


