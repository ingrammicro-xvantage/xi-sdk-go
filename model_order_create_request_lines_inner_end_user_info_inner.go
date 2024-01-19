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

// checks if the OrderCreateRequestLinesInnerEndUserInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestLinesInnerEndUserInfoInner{}

// OrderCreateRequestLinesInnerEndUserInfoInner struct for OrderCreateRequestLinesInnerEndUserInfoInner
type OrderCreateRequestLinesInnerEndUserInfoInner struct {
	// Specifies the type of endUser. It can be endUser or endUserContact for SAP flow.
	EndUserType *string `json:"endUserType,omitempty"`
	// ID for the end user/customer in Ingram Micro's system.
	EndUserId *string `json:"endUserId,omitempty"`
	// The contact name for the end user/customer.
	Contact *string `json:"contact,omitempty"`
	// The company name for the end user/customer.
	CompanyName *string `json:"companyName,omitempty"`
	// name1
	Name1 *string `json:"name1,omitempty"`
	// name2
	Name2 *string `json:"name2,omitempty"`
	// The end user/customer's street address and building or house number.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The end user/customer's apartment number.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Line 3 of the address for the end user/customer.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// Street address4
	AddressLine4 *string `json:"addressLine4,omitempty"`
	// The end user/customer's city.
	City *string `json:"city,omitempty"`
	// The end user/customer's state.
	State *string `json:"state,omitempty"`
	// The end user/customer's zip or postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The end user/customer's two-character ISO country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// The end user/customer's phone number.
	PhoneNumber *float32 `json:"phoneNumber,omitempty"`
	// The end user/customer's email.
	Email *string `json:"email,omitempty"`
}

// NewOrderCreateRequestLinesInnerEndUserInfoInner instantiates a new OrderCreateRequestLinesInnerEndUserInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestLinesInnerEndUserInfoInner() *OrderCreateRequestLinesInnerEndUserInfoInner {
	this := OrderCreateRequestLinesInnerEndUserInfoInner{}
	return &this
}

// NewOrderCreateRequestLinesInnerEndUserInfoInnerWithDefaults instantiates a new OrderCreateRequestLinesInnerEndUserInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestLinesInnerEndUserInfoInnerWithDefaults() *OrderCreateRequestLinesInnerEndUserInfoInner {
	this := OrderCreateRequestLinesInnerEndUserInfoInner{}
	return &this
}

// GetEndUserType returns the EndUserType field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetEndUserType() string {
	if o == nil || IsNil(o.EndUserType) {
		var ret string
		return ret
	}
	return *o.EndUserType
}

// GetEndUserTypeOk returns a tuple with the EndUserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetEndUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserType) {
		return nil, false
	}
	return o.EndUserType, true
}

// HasEndUserType returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasEndUserType() bool {
	if o != nil && !IsNil(o.EndUserType) {
		return true
	}

	return false
}

// SetEndUserType gets a reference to the given string and assigns it to the EndUserType field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetEndUserType(v string) {
	o.EndUserType = &v
}

// GetEndUserId returns the EndUserId field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetEndUserId() string {
	if o == nil || IsNil(o.EndUserId) {
		var ret string
		return ret
	}
	return *o.EndUserId
}

// GetEndUserIdOk returns a tuple with the EndUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetEndUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserId) {
		return nil, false
	}
	return o.EndUserId, true
}

// HasEndUserId returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasEndUserId() bool {
	if o != nil && !IsNil(o.EndUserId) {
		return true
	}

	return false
}

// SetEndUserId gets a reference to the given string and assigns it to the EndUserId field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetEndUserId(v string) {
	o.EndUserId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetName1() string {
	if o == nil || IsNil(o.Name1) {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetName1Ok() (*string, bool) {
	if o == nil || IsNil(o.Name1) {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasName1() bool {
	if o != nil && !IsNil(o.Name1) {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetName2(v string) {
	o.Name2 = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4) {
		var ret string
		return ret
	}
	return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetAddressLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine4) {
		return nil, false
	}
	return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasAddressLine4() bool {
	if o != nil && !IsNil(o.AddressLine4) {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetAddressLine4(v string) {
	o.AddressLine4 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetPhoneNumber() float32 {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret float32
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetPhoneNumberOk() (*float32, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given float32 and assigns it to the PhoneNumber field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetPhoneNumber(v float32) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrderCreateRequestLinesInnerEndUserInfoInner) SetEmail(v string) {
	o.Email = &v
}

func (o OrderCreateRequestLinesInnerEndUserInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestLinesInnerEndUserInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndUserType) {
		toSerialize["endUserType"] = o.EndUserType
	}
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

type NullableOrderCreateRequestLinesInnerEndUserInfoInner struct {
	value *OrderCreateRequestLinesInnerEndUserInfoInner
	isSet bool
}

func (v NullableOrderCreateRequestLinesInnerEndUserInfoInner) Get() *OrderCreateRequestLinesInnerEndUserInfoInner {
	return v.value
}

func (v *NullableOrderCreateRequestLinesInnerEndUserInfoInner) Set(val *OrderCreateRequestLinesInnerEndUserInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestLinesInnerEndUserInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestLinesInnerEndUserInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestLinesInnerEndUserInfoInner(val *OrderCreateRequestLinesInnerEndUserInfoInner) *NullableOrderCreateRequestLinesInnerEndUserInfoInner {
	return &NullableOrderCreateRequestLinesInnerEndUserInfoInner{value: val, isSet: true}
}

func (v NullableOrderCreateRequestLinesInnerEndUserInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestLinesInnerEndUserInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


