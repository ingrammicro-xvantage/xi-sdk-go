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

// checks if the OrderCreateV7RequestShipToInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateV7RequestShipToInfo{}

// OrderCreateV7RequestShipToInfo The shipping information provided by the reseller.
type OrderCreateV7RequestShipToInfo struct {
	// The ID references the resellers address in Ingram Micro's system for shipping. Provided to resellers during the onboarding process.
	AddressId *string `json:"addressId,omitempty"`
	// The company contact provided by the reseller.
	Contact *string `json:"contact,omitempty"`
	// The reseller's company name or the End-User's Name
	CompanyName *string `json:"companyName,omitempty"`
	// The street address and building or house number the order will be shipped to.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The apartment number the order will be shipped to.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The city the order will be shipped to.
	City *string `json:"city,omitempty"`
	// The state the order will be shipped to.
	State *string `json:"state,omitempty"`
	// End User Name
	PostalCode *string `json:"postalCode,omitempty"`
	// The zip or postal code the order will be shipped to.
	CountryCode *string `json:"countryCode,omitempty"`
	// The company contact email address.
	Email *string `json:"email,omitempty"`
	// 
	ShippingNotes *string `json:"shippingNotes,omitempty"`
	// The company contact phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// NewOrderCreateV7RequestShipToInfo instantiates a new OrderCreateV7RequestShipToInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateV7RequestShipToInfo() *OrderCreateV7RequestShipToInfo {
	this := OrderCreateV7RequestShipToInfo{}
	return &this
}

// NewOrderCreateV7RequestShipToInfoWithDefaults instantiates a new OrderCreateV7RequestShipToInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateV7RequestShipToInfoWithDefaults() *OrderCreateV7RequestShipToInfo {
	this := OrderCreateV7RequestShipToInfo{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *OrderCreateV7RequestShipToInfo) SetAddressId(v string) {
	o.AddressId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *OrderCreateV7RequestShipToInfo) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrderCreateV7RequestShipToInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *OrderCreateV7RequestShipToInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *OrderCreateV7RequestShipToInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderCreateV7RequestShipToInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderCreateV7RequestShipToInfo) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrderCreateV7RequestShipToInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *OrderCreateV7RequestShipToInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderCreateV7RequestShipToInfo) SetEmail(v string) {
	o.Email = &v
}

// GetShippingNotes returns the ShippingNotes field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetShippingNotes() string {
	if o == nil || IsNil(o.ShippingNotes) {
		var ret string
		return ret
	}
	return *o.ShippingNotes
}

// GetShippingNotesOk returns a tuple with the ShippingNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetShippingNotesOk() (*string, bool) {
	if o == nil || IsNil(o.ShippingNotes) {
		return nil, false
	}
	return o.ShippingNotes, true
}

// HasShippingNotes returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasShippingNotes() bool {
	if o != nil && !IsNil(o.ShippingNotes) {
		return true
	}

	return false
}

// SetShippingNotes gets a reference to the given string and assigns it to the ShippingNotes field.
func (o *OrderCreateV7RequestShipToInfo) SetShippingNotes(v string) {
	o.ShippingNotes = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrderCreateV7RequestShipToInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateV7RequestShipToInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrderCreateV7RequestShipToInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OrderCreateV7RequestShipToInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o OrderCreateV7RequestShipToInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateV7RequestShipToInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AddressLine1) {
		toSerialize["addressLine1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ShippingNotes) {
		toSerialize["shippingNotes"] = o.ShippingNotes
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableOrderCreateV7RequestShipToInfo struct {
	value *OrderCreateV7RequestShipToInfo
	isSet bool
}

func (v NullableOrderCreateV7RequestShipToInfo) Get() *OrderCreateV7RequestShipToInfo {
	return v.value
}

func (v *NullableOrderCreateV7RequestShipToInfo) Set(val *OrderCreateV7RequestShipToInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateV7RequestShipToInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateV7RequestShipToInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateV7RequestShipToInfo(val *OrderCreateV7RequestShipToInfo) *NullableOrderCreateV7RequestShipToInfo {
	return &NullableOrderCreateV7RequestShipToInfo{value: val, isSet: true}
}

func (v NullableOrderCreateV7RequestShipToInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateV7RequestShipToInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

