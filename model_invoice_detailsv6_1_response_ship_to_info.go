/*
XI Sdk Resellers

For Resellers seeking to innovate with Ingram Micro's API solutions, automate your eCommerce experience with our array of API's and webhooks to craft a seamless journey for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the InvoiceDetailsv61ResponseShipToInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceDetailsv61ResponseShipToInfo{}

// InvoiceDetailsv61ResponseShipToInfo struct for InvoiceDetailsv61ResponseShipToInfo
type InvoiceDetailsv61ResponseShipToInfo struct {
	// Ship to Name.
	Contact *string `json:"contact,omitempty"`
	// Ship to company.
	CompanyName *string `json:"companyName,omitempty"`
	// Ship to Address Line1.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// Ship to Address Line2.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Ship to Address Line3.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// Ship to City.
	City *string `json:"city,omitempty"`
	// Ship to State code
	State *string `json:"state,omitempty"`
	// Ship to Postalcode code.
	PostalCode *string `json:"postalCode,omitempty"`
	// Ship to Country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// Phone number of the Ship to company.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Email address of the Ship to company.
	Email *string `json:"email,omitempty"`
}

// NewInvoiceDetailsv61ResponseShipToInfo instantiates a new InvoiceDetailsv61ResponseShipToInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceDetailsv61ResponseShipToInfo() *InvoiceDetailsv61ResponseShipToInfo {
	this := InvoiceDetailsv61ResponseShipToInfo{}
	return &this
}

// NewInvoiceDetailsv61ResponseShipToInfoWithDefaults instantiates a new InvoiceDetailsv61ResponseShipToInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceDetailsv61ResponseShipToInfoWithDefaults() *InvoiceDetailsv61ResponseShipToInfo {
	this := InvoiceDetailsv61ResponseShipToInfo{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *InvoiceDetailsv61ResponseShipToInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *InvoiceDetailsv61ResponseShipToInfo) SetEmail(v string) {
	o.Email = &v
}

func (o InvoiceDetailsv61ResponseShipToInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceDetailsv61ResponseShipToInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.AddressLine3) {
		toSerialize["addressLine3"] = o.AddressLine3
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

type NullableInvoiceDetailsv61ResponseShipToInfo struct {
	value *InvoiceDetailsv61ResponseShipToInfo
	isSet bool
}

func (v NullableInvoiceDetailsv61ResponseShipToInfo) Get() *InvoiceDetailsv61ResponseShipToInfo {
	return v.value
}

func (v *NullableInvoiceDetailsv61ResponseShipToInfo) Set(val *InvoiceDetailsv61ResponseShipToInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceDetailsv61ResponseShipToInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceDetailsv61ResponseShipToInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceDetailsv61ResponseShipToInfo(val *InvoiceDetailsv61ResponseShipToInfo) *NullableInvoiceDetailsv61ResponseShipToInfo {
	return &NullableInvoiceDetailsv61ResponseShipToInfo{value: val, isSet: true}
}

func (v NullableInvoiceDetailsv61ResponseShipToInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceDetailsv61ResponseShipToInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


