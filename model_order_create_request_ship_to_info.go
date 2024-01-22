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

// checks if the OrderCreateRequestShipToInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestShipToInfo{}

// OrderCreateRequestShipToInfo The shipping information provided by the reseller.
type OrderCreateRequestShipToInfo struct {
	// The ID references the resellers address in Ingram Micro's system for shipping. Provided to resellers during the onboarding process.
	AddressId *string `json:"addressId,omitempty"`
	// The company contact provided by the reseller.
	Contact *string `json:"contact,omitempty"`
	// The name of the company the order will be shipped to.
	CompanyName *string `json:"companyName,omitempty"`
	// name1
	Name1 *string `json:"name1,omitempty"`
	// name2
	Name2 *string `json:"name2,omitempty"`
	// The street address and building or house number the order will be shipped to.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The apartment number the order will be shipped to.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Line 3 of the address the order will be shipped to.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// Street address4
	AddressLine4 *string `json:"addressLine4,omitempty"`
	// The city the order will be shipped to.
	City *string `json:"city,omitempty"`
	// The state the order will be shipped to.
	State *string `json:"state,omitempty"`
	// The zip or postal code the order will be shipped to.
	PostalCode *string `json:"postalCode,omitempty"`
	// The two-character ISO country code the order will be shipped to.
	CountryCode *string `json:"countryCode,omitempty"`
	// The company contact phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The company contact email address.
	Email *string `json:"email,omitempty"`
}

// NewOrderCreateRequestShipToInfo instantiates a new OrderCreateRequestShipToInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestShipToInfo() *OrderCreateRequestShipToInfo {
	this := OrderCreateRequestShipToInfo{}
	return &this
}

// NewOrderCreateRequestShipToInfoWithDefaults instantiates a new OrderCreateRequestShipToInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestShipToInfoWithDefaults() *OrderCreateRequestShipToInfo {
	this := OrderCreateRequestShipToInfo{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *OrderCreateRequestShipToInfo) SetAddressId(v string) {
	o.AddressId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *OrderCreateRequestShipToInfo) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrderCreateRequestShipToInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetName1() string {
	if o == nil || IsNil(o.Name1) {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetName1Ok() (*string, bool) {
	if o == nil || IsNil(o.Name1) {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasName1() bool {
	if o != nil && !IsNil(o.Name1) {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *OrderCreateRequestShipToInfo) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *OrderCreateRequestShipToInfo) SetName2(v string) {
	o.Name2 = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *OrderCreateRequestShipToInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *OrderCreateRequestShipToInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *OrderCreateRequestShipToInfo) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4) {
		var ret string
		return ret
	}
	return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetAddressLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine4) {
		return nil, false
	}
	return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasAddressLine4() bool {
	if o != nil && !IsNil(o.AddressLine4) {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *OrderCreateRequestShipToInfo) SetAddressLine4(v string) {
	o.AddressLine4 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderCreateRequestShipToInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderCreateRequestShipToInfo) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrderCreateRequestShipToInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *OrderCreateRequestShipToInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OrderCreateRequestShipToInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderCreateRequestShipToInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestShipToInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderCreateRequestShipToInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderCreateRequestShipToInfo) SetEmail(v string) {
	o.Email = &v
}

func (o OrderCreateRequestShipToInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestShipToInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["addressId"] = o.AddressId
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

type NullableOrderCreateRequestShipToInfo struct {
	value *OrderCreateRequestShipToInfo
	isSet bool
}

func (v NullableOrderCreateRequestShipToInfo) Get() *OrderCreateRequestShipToInfo {
	return v.value
}

func (v *NullableOrderCreateRequestShipToInfo) Set(val *OrderCreateRequestShipToInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestShipToInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestShipToInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestShipToInfo(val *OrderCreateRequestShipToInfo) *NullableOrderCreateRequestShipToInfo {
	return &NullableOrderCreateRequestShipToInfo{value: val, isSet: true}
}

func (v NullableOrderCreateRequestShipToInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestShipToInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

