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

// checks if the DealsSearchResponseDealsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DealsSearchResponseDealsInner{}

// DealsSearchResponseDealsInner struct for DealsSearchResponseDealsInner
type DealsSearchResponseDealsInner struct {
	// Deal/Special bid number.
	DealId *string `json:"dealId,omitempty"`
	// Most recent version number of the deal.
	Version *string `json:"version,omitempty"`
	// The end user/customer's name.
	EndUser *string `json:"endUser,omitempty"`
	// The vendor's name.
	Vendor *string `json:"vendor,omitempty"`
	// Expiration date of the deal/Special bid.
	DealExpiryDate *string `json:"dealExpiryDate,omitempty"`
	Links *RenewalsSearchResponseRenewalsInnerLinksInner `json:"links,omitempty"`
}

// NewDealsSearchResponseDealsInner instantiates a new DealsSearchResponseDealsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealsSearchResponseDealsInner() *DealsSearchResponseDealsInner {
	this := DealsSearchResponseDealsInner{}
	return &this
}

// NewDealsSearchResponseDealsInnerWithDefaults instantiates a new DealsSearchResponseDealsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealsSearchResponseDealsInnerWithDefaults() *DealsSearchResponseDealsInner {
	this := DealsSearchResponseDealsInner{}
	return &this
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *DealsSearchResponseDealsInner) GetDealId() string {
	if o == nil || IsNil(o.DealId) {
		var ret string
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponseDealsInner) GetDealIdOk() (*string, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *DealsSearchResponseDealsInner) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given string and assigns it to the DealId field.
func (o *DealsSearchResponseDealsInner) SetDealId(v string) {
	o.DealId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DealsSearchResponseDealsInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponseDealsInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DealsSearchResponseDealsInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DealsSearchResponseDealsInner) SetVersion(v string) {
	o.Version = &v
}

// GetEndUser returns the EndUser field value if set, zero value otherwise.
func (o *DealsSearchResponseDealsInner) GetEndUser() string {
	if o == nil || IsNil(o.EndUser) {
		var ret string
		return ret
	}
	return *o.EndUser
}

// GetEndUserOk returns a tuple with the EndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponseDealsInner) GetEndUserOk() (*string, bool) {
	if o == nil || IsNil(o.EndUser) {
		return nil, false
	}
	return o.EndUser, true
}

// HasEndUser returns a boolean if a field has been set.
func (o *DealsSearchResponseDealsInner) HasEndUser() bool {
	if o != nil && !IsNil(o.EndUser) {
		return true
	}

	return false
}

// SetEndUser gets a reference to the given string and assigns it to the EndUser field.
func (o *DealsSearchResponseDealsInner) SetEndUser(v string) {
	o.EndUser = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *DealsSearchResponseDealsInner) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponseDealsInner) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *DealsSearchResponseDealsInner) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *DealsSearchResponseDealsInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetDealExpiryDate returns the DealExpiryDate field value if set, zero value otherwise.
func (o *DealsSearchResponseDealsInner) GetDealExpiryDate() string {
	if o == nil || IsNil(o.DealExpiryDate) {
		var ret string
		return ret
	}
	return *o.DealExpiryDate
}

// GetDealExpiryDateOk returns a tuple with the DealExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponseDealsInner) GetDealExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.DealExpiryDate) {
		return nil, false
	}
	return o.DealExpiryDate, true
}

// HasDealExpiryDate returns a boolean if a field has been set.
func (o *DealsSearchResponseDealsInner) HasDealExpiryDate() bool {
	if o != nil && !IsNil(o.DealExpiryDate) {
		return true
	}

	return false
}

// SetDealExpiryDate gets a reference to the given string and assigns it to the DealExpiryDate field.
func (o *DealsSearchResponseDealsInner) SetDealExpiryDate(v string) {
	o.DealExpiryDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DealsSearchResponseDealsInner) GetLinks() RenewalsSearchResponseRenewalsInnerLinksInner {
	if o == nil || IsNil(o.Links) {
		var ret RenewalsSearchResponseRenewalsInnerLinksInner
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponseDealsInner) GetLinksOk() (*RenewalsSearchResponseRenewalsInnerLinksInner, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DealsSearchResponseDealsInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RenewalsSearchResponseRenewalsInnerLinksInner and assigns it to the Links field.
func (o *DealsSearchResponseDealsInner) SetLinks(v RenewalsSearchResponseRenewalsInnerLinksInner) {
	o.Links = &v
}

func (o DealsSearchResponseDealsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DealsSearchResponseDealsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DealId) {
		toSerialize["dealId"] = o.DealId
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.EndUser) {
		toSerialize["endUser"] = o.EndUser
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.DealExpiryDate) {
		toSerialize["dealExpiryDate"] = o.DealExpiryDate
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableDealsSearchResponseDealsInner struct {
	value *DealsSearchResponseDealsInner
	isSet bool
}

func (v NullableDealsSearchResponseDealsInner) Get() *DealsSearchResponseDealsInner {
	return v.value
}

func (v *NullableDealsSearchResponseDealsInner) Set(val *DealsSearchResponseDealsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDealsSearchResponseDealsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDealsSearchResponseDealsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealsSearchResponseDealsInner(val *DealsSearchResponseDealsInner) *NullableDealsSearchResponseDealsInner {
	return &NullableDealsSearchResponseDealsInner{value: val, isSet: true}
}

func (v NullableDealsSearchResponseDealsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealsSearchResponseDealsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


