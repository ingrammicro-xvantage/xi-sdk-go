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

// checks if the QuoteToOrderDetailsDTOEndUserInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteToOrderDetailsDTOEndUserInfoInner{}

// QuoteToOrderDetailsDTOEndUserInfoInner struct for QuoteToOrderDetailsDTOEndUserInfoInner
type QuoteToOrderDetailsDTOEndUserInfoInner struct {
	// The company name for the end user/customer.
	CompanyName *string `json:"companyName,omitempty"`
	// The contact name for the end user/customer.
	Contact *string `json:"contact,omitempty"`
	// The address line 1 for the end user/customer.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The address line 2 for the end user/customer.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The address line 3 for the end user/customer.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The end user/customer's city.
	City *string `json:"city,omitempty"`
	// The end user/customer's state.
	State *string `json:"state,omitempty"`
	// The end user/customer's zip or postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The end user/customer's two character ISO country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// The end user/customer's phone number.
	Email *string `json:"email,omitempty"`
	// The end user/customer's phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// NewQuoteToOrderDetailsDTOEndUserInfoInner instantiates a new QuoteToOrderDetailsDTOEndUserInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteToOrderDetailsDTOEndUserInfoInner() *QuoteToOrderDetailsDTOEndUserInfoInner {
	this := QuoteToOrderDetailsDTOEndUserInfoInner{}
	return &this
}

// NewQuoteToOrderDetailsDTOEndUserInfoInnerWithDefaults instantiates a new QuoteToOrderDetailsDTOEndUserInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteToOrderDetailsDTOEndUserInfoInnerWithDefaults() *QuoteToOrderDetailsDTOEndUserInfoInner {
	this := QuoteToOrderDetailsDTOEndUserInfoInner{}
	return &this
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetContact(v string) {
	o.Contact = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *QuoteToOrderDetailsDTOEndUserInfoInner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o QuoteToOrderDetailsDTOEndUserInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteToOrderDetailsDTOEndUserInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	return toSerialize, nil
}

type NullableQuoteToOrderDetailsDTOEndUserInfoInner struct {
	value *QuoteToOrderDetailsDTOEndUserInfoInner
	isSet bool
}

func (v NullableQuoteToOrderDetailsDTOEndUserInfoInner) Get() *QuoteToOrderDetailsDTOEndUserInfoInner {
	return v.value
}

func (v *NullableQuoteToOrderDetailsDTOEndUserInfoInner) Set(val *QuoteToOrderDetailsDTOEndUserInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteToOrderDetailsDTOEndUserInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteToOrderDetailsDTOEndUserInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteToOrderDetailsDTOEndUserInfoInner(val *QuoteToOrderDetailsDTOEndUserInfoInner) *NullableQuoteToOrderDetailsDTOEndUserInfoInner {
	return &NullableQuoteToOrderDetailsDTOEndUserInfoInner{value: val, isSet: true}
}

func (v NullableQuoteToOrderDetailsDTOEndUserInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteToOrderDetailsDTOEndUserInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

