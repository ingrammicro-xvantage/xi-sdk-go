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

// checks if the QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser{}

// QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser struct for QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser
type QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser struct {
	EndUserName *string `json:"endUserName,omitempty"`
	EndUserAddress *string `json:"endUserAddress,omitempty"`
	EndUserAddress2 *string `json:"endUserAddress2,omitempty"`
	EndUserAddress3 *string `json:"endUserAddress3,omitempty"`
	EndUserCity *string `json:"endUserCity,omitempty"`
	EndUserState *string `json:"endUserState,omitempty"`
	EndUserEmail *string `json:"endUserEmail,omitempty"`
	EndUserPhone *string `json:"endUserPhone,omitempty"`
	EndUserZipCode *string `json:"endUserZipCode,omitempty"`
	EndUserContactName *string `json:"endUserContactName,omitempty"`
	EndUserMarketSegment *string `json:"endUserMarketSegment,omitempty"`
}

// NewQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser instantiates a new QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser() *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser {
	this := QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser{}
	return &this
}

// NewQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUserWithDefaults instantiates a new QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUserWithDefaults() *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser {
	this := QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser{}
	return &this
}

// GetEndUserName returns the EndUserName field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserName() string {
	if o == nil || IsNil(o.EndUserName) {
		var ret string
		return ret
	}
	return *o.EndUserName
}

// GetEndUserNameOk returns a tuple with the EndUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserName) {
		return nil, false
	}
	return o.EndUserName, true
}

// HasEndUserName returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserName() bool {
	if o != nil && !IsNil(o.EndUserName) {
		return true
	}

	return false
}

// SetEndUserName gets a reference to the given string and assigns it to the EndUserName field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserName(v string) {
	o.EndUserName = &v
}

// GetEndUserAddress returns the EndUserAddress field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserAddress() string {
	if o == nil || IsNil(o.EndUserAddress) {
		var ret string
		return ret
	}
	return *o.EndUserAddress
}

// GetEndUserAddressOk returns a tuple with the EndUserAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserAddress) {
		return nil, false
	}
	return o.EndUserAddress, true
}

// HasEndUserAddress returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserAddress() bool {
	if o != nil && !IsNil(o.EndUserAddress) {
		return true
	}

	return false
}

// SetEndUserAddress gets a reference to the given string and assigns it to the EndUserAddress field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserAddress(v string) {
	o.EndUserAddress = &v
}

// GetEndUserAddress2 returns the EndUserAddress2 field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserAddress2() string {
	if o == nil || IsNil(o.EndUserAddress2) {
		var ret string
		return ret
	}
	return *o.EndUserAddress2
}

// GetEndUserAddress2Ok returns a tuple with the EndUserAddress2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.EndUserAddress2) {
		return nil, false
	}
	return o.EndUserAddress2, true
}

// HasEndUserAddress2 returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserAddress2() bool {
	if o != nil && !IsNil(o.EndUserAddress2) {
		return true
	}

	return false
}

// SetEndUserAddress2 gets a reference to the given string and assigns it to the EndUserAddress2 field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserAddress2(v string) {
	o.EndUserAddress2 = &v
}

// GetEndUserAddress3 returns the EndUserAddress3 field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserAddress3() string {
	if o == nil || IsNil(o.EndUserAddress3) {
		var ret string
		return ret
	}
	return *o.EndUserAddress3
}

// GetEndUserAddress3Ok returns a tuple with the EndUserAddress3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserAddress3Ok() (*string, bool) {
	if o == nil || IsNil(o.EndUserAddress3) {
		return nil, false
	}
	return o.EndUserAddress3, true
}

// HasEndUserAddress3 returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserAddress3() bool {
	if o != nil && !IsNil(o.EndUserAddress3) {
		return true
	}

	return false
}

// SetEndUserAddress3 gets a reference to the given string and assigns it to the EndUserAddress3 field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserAddress3(v string) {
	o.EndUserAddress3 = &v
}

// GetEndUserCity returns the EndUserCity field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserCity() string {
	if o == nil || IsNil(o.EndUserCity) {
		var ret string
		return ret
	}
	return *o.EndUserCity
}

// GetEndUserCityOk returns a tuple with the EndUserCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserCityOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserCity) {
		return nil, false
	}
	return o.EndUserCity, true
}

// HasEndUserCity returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserCity() bool {
	if o != nil && !IsNil(o.EndUserCity) {
		return true
	}

	return false
}

// SetEndUserCity gets a reference to the given string and assigns it to the EndUserCity field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserCity(v string) {
	o.EndUserCity = &v
}

// GetEndUserState returns the EndUserState field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserState() string {
	if o == nil || IsNil(o.EndUserState) {
		var ret string
		return ret
	}
	return *o.EndUserState
}

// GetEndUserStateOk returns a tuple with the EndUserState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserStateOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserState) {
		return nil, false
	}
	return o.EndUserState, true
}

// HasEndUserState returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserState() bool {
	if o != nil && !IsNil(o.EndUserState) {
		return true
	}

	return false
}

// SetEndUserState gets a reference to the given string and assigns it to the EndUserState field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserState(v string) {
	o.EndUserState = &v
}

// GetEndUserEmail returns the EndUserEmail field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserEmail() string {
	if o == nil || IsNil(o.EndUserEmail) {
		var ret string
		return ret
	}
	return *o.EndUserEmail
}

// GetEndUserEmailOk returns a tuple with the EndUserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserEmail) {
		return nil, false
	}
	return o.EndUserEmail, true
}

// HasEndUserEmail returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserEmail() bool {
	if o != nil && !IsNil(o.EndUserEmail) {
		return true
	}

	return false
}

// SetEndUserEmail gets a reference to the given string and assigns it to the EndUserEmail field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserEmail(v string) {
	o.EndUserEmail = &v
}

// GetEndUserPhone returns the EndUserPhone field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserPhone() string {
	if o == nil || IsNil(o.EndUserPhone) {
		var ret string
		return ret
	}
	return *o.EndUserPhone
}

// GetEndUserPhoneOk returns a tuple with the EndUserPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserPhone) {
		return nil, false
	}
	return o.EndUserPhone, true
}

// HasEndUserPhone returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserPhone() bool {
	if o != nil && !IsNil(o.EndUserPhone) {
		return true
	}

	return false
}

// SetEndUserPhone gets a reference to the given string and assigns it to the EndUserPhone field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserPhone(v string) {
	o.EndUserPhone = &v
}

// GetEndUserZipCode returns the EndUserZipCode field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserZipCode() string {
	if o == nil || IsNil(o.EndUserZipCode) {
		var ret string
		return ret
	}
	return *o.EndUserZipCode
}

// GetEndUserZipCodeOk returns a tuple with the EndUserZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserZipCodeOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserZipCode) {
		return nil, false
	}
	return o.EndUserZipCode, true
}

// HasEndUserZipCode returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserZipCode() bool {
	if o != nil && !IsNil(o.EndUserZipCode) {
		return true
	}

	return false
}

// SetEndUserZipCode gets a reference to the given string and assigns it to the EndUserZipCode field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserZipCode(v string) {
	o.EndUserZipCode = &v
}

// GetEndUserContactName returns the EndUserContactName field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserContactName() string {
	if o == nil || IsNil(o.EndUserContactName) {
		var ret string
		return ret
	}
	return *o.EndUserContactName
}

// GetEndUserContactNameOk returns a tuple with the EndUserContactName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserContactNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserContactName) {
		return nil, false
	}
	return o.EndUserContactName, true
}

// HasEndUserContactName returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserContactName() bool {
	if o != nil && !IsNil(o.EndUserContactName) {
		return true
	}

	return false
}

// SetEndUserContactName gets a reference to the given string and assigns it to the EndUserContactName field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserContactName(v string) {
	o.EndUserContactName = &v
}

// GetEndUserMarketSegment returns the EndUserMarketSegment field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserMarketSegment() string {
	if o == nil || IsNil(o.EndUserMarketSegment) {
		var ret string
		return ret
	}
	return *o.EndUserMarketSegment
}

// GetEndUserMarketSegmentOk returns a tuple with the EndUserMarketSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) GetEndUserMarketSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserMarketSegment) {
		return nil, false
	}
	return o.EndUserMarketSegment, true
}

// HasEndUserMarketSegment returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) HasEndUserMarketSegment() bool {
	if o != nil && !IsNil(o.EndUserMarketSegment) {
		return true
	}

	return false
}

// SetEndUserMarketSegment gets a reference to the given string and assigns it to the EndUserMarketSegment field.
func (o *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) SetEndUserMarketSegment(v string) {
	o.EndUserMarketSegment = &v
}

func (o QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndUserName) {
		toSerialize["endUserName"] = o.EndUserName
	}
	if !IsNil(o.EndUserAddress) {
		toSerialize["endUserAddress"] = o.EndUserAddress
	}
	if !IsNil(o.EndUserAddress2) {
		toSerialize["endUserAddress2"] = o.EndUserAddress2
	}
	if !IsNil(o.EndUserAddress3) {
		toSerialize["endUserAddress3"] = o.EndUserAddress3
	}
	if !IsNil(o.EndUserCity) {
		toSerialize["endUserCity"] = o.EndUserCity
	}
	if !IsNil(o.EndUserState) {
		toSerialize["endUserState"] = o.EndUserState
	}
	if !IsNil(o.EndUserEmail) {
		toSerialize["endUserEmail"] = o.EndUserEmail
	}
	if !IsNil(o.EndUserPhone) {
		toSerialize["endUserPhone"] = o.EndUserPhone
	}
	if !IsNil(o.EndUserZipCode) {
		toSerialize["endUserZipCode"] = o.EndUserZipCode
	}
	if !IsNil(o.EndUserContactName) {
		toSerialize["endUserContactName"] = o.EndUserContactName
	}
	if !IsNil(o.EndUserMarketSegment) {
		toSerialize["endUserMarketSegment"] = o.EndUserMarketSegment
	}
	return toSerialize, nil
}

type NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser struct {
	value *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser
	isSet bool
}

func (v NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) Get() *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser {
	return v.value
}

func (v *NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) Set(val *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser(val *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) *NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser {
	return &NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser{value: val, isSet: true}
}

func (v NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsQuoteDetailResponseRetrieveQuoteResponseEndUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


