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

// checks if the QuoteDetailsResponseEndUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsResponseEndUserInfo{}

// QuoteDetailsResponseEndUserInfo struct for QuoteDetailsResponseEndUserInfo
type QuoteDetailsResponseEndUserInfo struct {
	// End User Name
	Contact *string `json:"contact,omitempty"`
	// Contact name  of end user associated with the quote.
	CompanyName *string `json:"companyName,omitempty"`
	// Address line 1 for end user associated with the quote
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// Address line 2 for end user associated with the quote.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// Address line 3 for end user associated with the quote.
	AddressLine3 *string `json:"addressLine3,omitempty"`
	// City for end user associated with the quote
	City *string `json:"city,omitempty"`
	// Two letter state abreviation for end user associated with the quote
	State *string `json:"state,omitempty"`
	// Email of end user the quote associated with the quote.
	Email *string `json:"email,omitempty"`
	// Phone number of end user associated with the quote.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Zip code of end user associated with the quote.
	PostalCode *string `json:"postalCode,omitempty"`
	// Market Segment of end user associated with the quote. End user market segment is included when end user is included in specific market segments like Educational, Government, Military, Medical - that may receive special pricing due to their segmentation.
	MarketSegment *string `json:"marketSegment,omitempty"`
}

// NewQuoteDetailsResponseEndUserInfo instantiates a new QuoteDetailsResponseEndUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsResponseEndUserInfo() *QuoteDetailsResponseEndUserInfo {
	this := QuoteDetailsResponseEndUserInfo{}
	return &this
}

// NewQuoteDetailsResponseEndUserInfoWithDefaults instantiates a new QuoteDetailsResponseEndUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsResponseEndUserInfoWithDefaults() *QuoteDetailsResponseEndUserInfo {
	this := QuoteDetailsResponseEndUserInfo{}
	return &this
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetContact() string {
	if o == nil || IsNil(o.Contact) {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetContactOk() (*string, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *QuoteDetailsResponseEndUserInfo) SetContact(v string) {
	o.Contact = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *QuoteDetailsResponseEndUserInfo) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *QuoteDetailsResponseEndUserInfo) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *QuoteDetailsResponseEndUserInfo) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *QuoteDetailsResponseEndUserInfo) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *QuoteDetailsResponseEndUserInfo) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *QuoteDetailsResponseEndUserInfo) SetState(v string) {
	o.State = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *QuoteDetailsResponseEndUserInfo) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *QuoteDetailsResponseEndUserInfo) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *QuoteDetailsResponseEndUserInfo) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetMarketSegment returns the MarketSegment field value if set, zero value otherwise.
func (o *QuoteDetailsResponseEndUserInfo) GetMarketSegment() string {
	if o == nil || IsNil(o.MarketSegment) {
		var ret string
		return ret
	}
	return *o.MarketSegment
}

// GetMarketSegmentOk returns a tuple with the MarketSegment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsResponseEndUserInfo) GetMarketSegmentOk() (*string, bool) {
	if o == nil || IsNil(o.MarketSegment) {
		return nil, false
	}
	return o.MarketSegment, true
}

// HasMarketSegment returns a boolean if a field has been set.
func (o *QuoteDetailsResponseEndUserInfo) HasMarketSegment() bool {
	if o != nil && !IsNil(o.MarketSegment) {
		return true
	}

	return false
}

// SetMarketSegment gets a reference to the given string and assigns it to the MarketSegment field.
func (o *QuoteDetailsResponseEndUserInfo) SetMarketSegment(v string) {
	o.MarketSegment = &v
}

func (o QuoteDetailsResponseEndUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsResponseEndUserInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.MarketSegment) {
		toSerialize["marketSegment"] = o.MarketSegment
	}
	return toSerialize, nil
}

type NullableQuoteDetailsResponseEndUserInfo struct {
	value *QuoteDetailsResponseEndUserInfo
	isSet bool
}

func (v NullableQuoteDetailsResponseEndUserInfo) Get() *QuoteDetailsResponseEndUserInfo {
	return v.value
}

func (v *NullableQuoteDetailsResponseEndUserInfo) Set(val *QuoteDetailsResponseEndUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsResponseEndUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsResponseEndUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsResponseEndUserInfo(val *QuoteDetailsResponseEndUserInfo) *NullableQuoteDetailsResponseEndUserInfo {
	return &NullableQuoteDetailsResponseEndUserInfo{value: val, isSet: true}
}

func (v NullableQuoteDetailsResponseEndUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsResponseEndUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


