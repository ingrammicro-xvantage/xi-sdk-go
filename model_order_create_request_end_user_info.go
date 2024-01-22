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

// checks if the OrderCreateRequestEndUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestEndUserInfo{}

// OrderCreateRequestEndUserInfo The contact information for the end user/customer provided by the reseller. Used to determine pricing and discounts.
type OrderCreateRequestEndUserInfo struct {
	// ID for the end user/customer in Ingram Micro's system.
	EndUserId *string `json:"endUserId,omitempty"`
	// The contact name for the end user/customer.
	Contact *string `json:"contact,omitempty"`
	// The company name for the end user/customer. Required for Impulse countries.
	CompanyName *string `json:"companyName,omitempty"`
	// name1
	Name1 *string `json:"name1,omitempty"`
	// name2
	Name2 *string `json:"name2,omitempty"`
	// The end user/customer's street address and building or house number. Required for Impulse countries.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The end user/customer's apartment number.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Line 3 of the address for the end user/customer.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// Street address4
	AddressLine4 *string `json:"addressLine4,omitempty"`
	// The end user/customer's city. Required for Impulse countries.
	City *string `json:"city,omitempty"`
	// The end user/customer's state. Required for Impulse countries but optional for EMEA countries.
	State *string `json:"state,omitempty"`
	// The end user/customer's zip or postal code. Required for Impulse countries.
	PostalCode *string `json:"postalCode,omitempty"`
	// The end user/customer's two-character ISO country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// The end user/customer's phone number.
	PhoneNumber *int32 `json:"phoneNumber,omitempty"`
	// The end user/customer's email.
	Email *string `json:"email,omitempty"`
}

// NewOrderCreateRequestEndUserInfo instantiates a new OrderCreateRequestEndUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestEndUserInfo() *OrderCreateRequestEndUserInfo {
	this := OrderCreateRequestEndUserInfo{}
	return &this
}

// NewOrderCreateRequestEndUserInfoWithDefaults instantiates a new OrderCreateRequestEndUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestEndUserInfoWithDefaults() *OrderCreateRequestEndUserInfo {
	this := OrderCreateRequestEndUserInfo{}
	return &this
}

// GetEndUserId returns the EndUserId field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetEndUserId() string {
	if o == nil || IsNil(o.EndUserId) {
		var ret string
		return ret
	}
	return *o.EndUserId
}

// GetEndUserIdOk returns a tuple with the EndUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetEndUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserId) {
		return nil, false
	}
	return o.EndUserId, true
}

// HasEndUserId returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasEndUserId() bool {
	if o != nil && !IsNil(o.EndUserId) {
		return true
	}

	return false
}

// SetEndUserId gets a reference to the given string and assigns it to the EndUserId field.
func (o *OrderCreateRequestEndUserInfo) SetEndUserId(v string) {
	o.EndUserId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *OrderCreateRequestEndUserInfo) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrderCreateRequestEndUserInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetName1() string {
	if o == nil || IsNil(o.Name1) {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetName1Ok() (*string, bool) {
	if o == nil || IsNil(o.Name1) {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasName1() bool {
	if o != nil && !IsNil(o.Name1) {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *OrderCreateRequestEndUserInfo) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *OrderCreateRequestEndUserInfo) SetName2(v string) {
	o.Name2 = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *OrderCreateRequestEndUserInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *OrderCreateRequestEndUserInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *OrderCreateRequestEndUserInfo) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4) {
		var ret string
		return ret
	}
	return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetAddressLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine4) {
		return nil, false
	}
	return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasAddressLine4() bool {
	if o != nil && !IsNil(o.AddressLine4) {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *OrderCreateRequestEndUserInfo) SetAddressLine4(v string) {
	o.AddressLine4 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderCreateRequestEndUserInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderCreateRequestEndUserInfo) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrderCreateRequestEndUserInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *OrderCreateRequestEndUserInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetPhoneNumber() int32 {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret int32
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetPhoneNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given int32 and assigns it to the PhoneNumber field.
func (o *OrderCreateRequestEndUserInfo) SetPhoneNumber(v int32) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderCreateRequestEndUserInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestEndUserInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderCreateRequestEndUserInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderCreateRequestEndUserInfo) SetEmail(v string) {
	o.Email = &v
}

func (o OrderCreateRequestEndUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestEndUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndUserId) {
		toSerialize["endUserId"] = o.EndUserId
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Name1) {
		toSerialize["name1"] = o.Name1
	}
	if !IsNil(o.Name2) {
		toSerialize["name2"] = o.Name2
	}
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
	}
	if !IsNil(o.AddressLine4) {
		toSerialize["addressLine4"] = o.AddressLine4
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableOrderCreateRequestEndUserInfo struct {
	value *OrderCreateRequestEndUserInfo
	isSet bool
}

func (v NullableOrderCreateRequestEndUserInfo) Get() *OrderCreateRequestEndUserInfo {
	return v.value
}

func (v *NullableOrderCreateRequestEndUserInfo) Set(val *OrderCreateRequestEndUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestEndUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestEndUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestEndUserInfo(val *OrderCreateRequestEndUserInfo) *NullableOrderCreateRequestEndUserInfo {
	return &NullableOrderCreateRequestEndUserInfo{value: val, isSet: true}
}

func (v NullableOrderCreateRequestEndUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestEndUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

