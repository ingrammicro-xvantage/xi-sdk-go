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

// checks if the AsyncOrderCreateDTOLinesInnerEndUserInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsyncOrderCreateDTOLinesInnerEndUserInfoInner{}

// AsyncOrderCreateDTOLinesInnerEndUserInfoInner struct for AsyncOrderCreateDTOLinesInnerEndUserInfoInner
type AsyncOrderCreateDTOLinesInnerEndUserInfoInner struct {
	// ID for the end user/customer in Ingram Micro's system.
	EndUserId *string `json:"endUserId,omitempty"`
	// End user type
	EndUserType *string `json:"endUserType,omitempty"`
	// The company name for the end user/customer.
	CompanyName *string `json:"companyName,omitempty"`
	Name1 *string `json:"name1,omitempty"`
	Name2 *string `json:"name2,omitempty"`
	// The contact Id for the end user/customer.
	ContactId *string `json:"contactId,omitempty"`
	// The address line 1 for the end user/customer.
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// The address line 2 for the end user/customer.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The address line 3 for the end user/customer.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// The contact name for the end user/customer.
	Contact *string `json:"contact,omitempty"`
	// The end user/customer's city.
	City *string `json:"city,omitempty"`
	// The end user/customer's state.
	State *string `json:"state,omitempty"`
	// The end user/customer's zip or postal code.
	PostalCode *string `json:"postalCode,omitempty"`
	// The address line 4 for the end user/customer.
	AddressLine4 *string `json:"addressLine4,omitempty"`
	// The end user/customer's two character ISO country code.
	CountryCode *string `json:"countryCode,omitempty"`
	// The end user/customer's phone number.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The end user/customer's phone number.
	Email *string `json:"email,omitempty"`
}

// NewAsyncOrderCreateDTOLinesInnerEndUserInfoInner instantiates a new AsyncOrderCreateDTOLinesInnerEndUserInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncOrderCreateDTOLinesInnerEndUserInfoInner() *AsyncOrderCreateDTOLinesInnerEndUserInfoInner {
	this := AsyncOrderCreateDTOLinesInnerEndUserInfoInner{}
	return &this
}

// NewAsyncOrderCreateDTOLinesInnerEndUserInfoInnerWithDefaults instantiates a new AsyncOrderCreateDTOLinesInnerEndUserInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncOrderCreateDTOLinesInnerEndUserInfoInnerWithDefaults() *AsyncOrderCreateDTOLinesInnerEndUserInfoInner {
	this := AsyncOrderCreateDTOLinesInnerEndUserInfoInner{}
	return &this
}

// GetEndUserId returns the EndUserId field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetEndUserId() string {
	if o == nil || IsNil(o.EndUserId) {
		var ret string
		return ret
	}
	return *o.EndUserId
}

// GetEndUserIdOk returns a tuple with the EndUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetEndUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserId) {
		return nil, false
	}
	return o.EndUserId, true
}

// HasEndUserId returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasEndUserId() bool {
	if o != nil && !IsNil(o.EndUserId) {
		return true
	}

	return false
}

// SetEndUserId gets a reference to the given string and assigns it to the EndUserId field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetEndUserId(v string) {
	o.EndUserId = &v
}

// GetEndUserType returns the EndUserType field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetEndUserType() string {
	if o == nil || IsNil(o.EndUserType) {
		var ret string
		return ret
	}
	return *o.EndUserType
}

// GetEndUserTypeOk returns a tuple with the EndUserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetEndUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserType) {
		return nil, false
	}
	return o.EndUserType, true
}

// HasEndUserType returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasEndUserType() bool {
	if o != nil && !IsNil(o.EndUserType) {
		return true
	}

	return false
}

// SetEndUserType gets a reference to the given string and assigns it to the EndUserType field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetEndUserType(v string) {
	o.EndUserType = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetName1() string {
	if o == nil || IsNil(o.Name1) {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetName1Ok() (*string, bool) {
	if o == nil || IsNil(o.Name1) {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasName1() bool {
	if o != nil && !IsNil(o.Name1) {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetName2() string {
	if o == nil || IsNil(o.Name2) {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetName2Ok() (*string, bool) {
	if o == nil || IsNil(o.Name2) {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasName2() bool {
	if o != nil && !IsNil(o.Name2) {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetName2(v string) {
	o.Name2 = &v
}

// GetContactId returns the ContactId field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetContactId() string {
	if o == nil || IsNil(o.ContactId) {
		var ret string
		return ret
	}
	return *o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContactId) {
		return nil, false
	}
	return o.ContactId, true
}

// HasContactId returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasContactId() bool {
	if o != nil && !IsNil(o.ContactId) {
		return true
	}

	return false
}

// SetContactId gets a reference to the given string and assigns it to the ContactId field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetContactId(v string) {
	o.ContactId = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetContact(v string) {
	o.Contact = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4) {
		var ret string
		return ret
	}
	return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetAddressLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine4) {
		return nil, false
	}
	return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasAddressLine4() bool {
	if o != nil && !IsNil(o.AddressLine4) {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetAddressLine4(v string) {
	o.AddressLine4 = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) SetEmail(v string) {
	o.Email = &v
}

func (o AsyncOrderCreateDTOLinesInnerEndUserInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsyncOrderCreateDTOLinesInnerEndUserInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndUserId) {
		toSerialize["endUserId"] = o.EndUserId
	}
	if !IsNil(o.EndUserType) {
		toSerialize["endUserType"] = o.EndUserType
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
	if !IsNil(o.ContactId) {
		toSerialize["contactId"] = o.ContactId
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
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
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
	if !IsNil(o.AddressLine4) {
		toSerialize["addressLine4"] = o.AddressLine4
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

type NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner struct {
	value *AsyncOrderCreateDTOLinesInnerEndUserInfoInner
	isSet bool
}

func (v NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner) Get() *AsyncOrderCreateDTOLinesInnerEndUserInfoInner {
	return v.value
}

func (v *NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner) Set(val *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner(val *AsyncOrderCreateDTOLinesInnerEndUserInfoInner) *NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner {
	return &NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner{value: val, isSet: true}
}

func (v NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncOrderCreateDTOLinesInnerEndUserInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

