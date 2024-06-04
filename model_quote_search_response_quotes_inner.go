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

// checks if the QuoteSearchResponseQuotesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteSearchResponseQuotesInner{}

// QuoteSearchResponseQuotesInner struct for QuoteSearchResponseQuotesInner
type QuoteSearchResponseQuotesInner struct {
	QuoteGuid *string `json:"quoteGuid,omitempty"`
	// Quote Name given to quote by sales team or system generated.  Generally used as a reference to identify the quote.
	QuoteName *string `json:"quoteName,omitempty"`
	// Unique identifier generated by Ingram Micros CRM specific to each quote.  When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// When a quote has been revised and updated, the quote number remains the same throughout the lifecycle of the quote, however, a Revision number is updated for each revision of the quote.  The revision numbers is associated with the Unique Quote Number.
	Revision *string `json:"revision,omitempty"`
	// The country-specific three digit ISO 4217 currency code for the order.
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// End User Name is the end customer name that is associated with a quote in Ingram Micros CRM.
	EndUserContact *string `json:"endUserContact,omitempty"`
	// Special Pricing Bid Number, also refers to as Dart Number relates to a unique pricing deal associated with a vendor for the quote.
	SpecialBidNumber *string `json:"specialBidNumber,omitempty"`
	// Total amount of quoted price for all products in the quote.
	QuoteTotal *float32 `json:"quoteTotal,omitempty"`
	// This refers to the primary status of the quote.
	QuoteStatus *string `json:"quoteStatus,omitempty"`
	// Date the Quote was initially Created.
	IngramQuoteDate *string `json:"ingramQuoteDate,omitempty"`
	// Date the Quote was last updated or modified.
	LastModifiedDate *string `json:"lastModifiedDate,omitempty"`
	// Date when the Quote Expires.
	IngramQuoteExpiryDate *string `json:"ingramQuoteExpiryDate,omitempty"`
	// End User Name
	EndUserName *string `json:"endUserName,omitempty"`
	// Name of the vendor.
	Vendor *string `json:"vendor,omitempty"`
	// Name of the end user/customer who created a quote.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Type of quote
	QuoteType *string `json:"quoteType,omitempty"`
	Links *QuoteSearchResponseQuotesInnerLinks `json:"links,omitempty"`
}

// NewQuoteSearchResponseQuotesInner instantiates a new QuoteSearchResponseQuotesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSearchResponseQuotesInner() *QuoteSearchResponseQuotesInner {
	this := QuoteSearchResponseQuotesInner{}
	return &this
}

// NewQuoteSearchResponseQuotesInnerWithDefaults instantiates a new QuoteSearchResponseQuotesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSearchResponseQuotesInnerWithDefaults() *QuoteSearchResponseQuotesInner {
	this := QuoteSearchResponseQuotesInner{}
	return &this
}

// GetQuoteGuid returns the QuoteGuid field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetQuoteGuid() string {
	if o == nil || IsNil(o.QuoteGuid) {
		var ret string
		return ret
	}
	return *o.QuoteGuid
}

// GetQuoteGuidOk returns a tuple with the QuoteGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetQuoteGuidOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteGuid) {
		return nil, false
	}
	return o.QuoteGuid, true
}

// HasQuoteGuid returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasQuoteGuid() bool {
	if o != nil && !IsNil(o.QuoteGuid) {
		return true
	}

	return false
}

// SetQuoteGuid gets a reference to the given string and assigns it to the QuoteGuid field.
func (o *QuoteSearchResponseQuotesInner) SetQuoteGuid(v string) {
	o.QuoteGuid = &v
}

// GetQuoteName returns the QuoteName field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetQuoteName() string {
	if o == nil || IsNil(o.QuoteName) {
		var ret string
		return ret
	}
	return *o.QuoteName
}

// GetQuoteNameOk returns a tuple with the QuoteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetQuoteNameOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteName) {
		return nil, false
	}
	return o.QuoteName, true
}

// HasQuoteName returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasQuoteName() bool {
	if o != nil && !IsNil(o.QuoteName) {
		return true
	}

	return false
}

// SetQuoteName gets a reference to the given string and assigns it to the QuoteName field.
func (o *QuoteSearchResponseQuotesInner) SetQuoteName(v string) {
	o.QuoteName = &v
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *QuoteSearchResponseQuotesInner) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *QuoteSearchResponseQuotesInner) SetRevision(v string) {
	o.Revision = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *QuoteSearchResponseQuotesInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetEndUserContact returns the EndUserContact field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetEndUserContact() string {
	if o == nil || IsNil(o.EndUserContact) {
		var ret string
		return ret
	}
	return *o.EndUserContact
}

// GetEndUserContactOk returns a tuple with the EndUserContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetEndUserContactOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserContact) {
		return nil, false
	}
	return o.EndUserContact, true
}

// HasEndUserContact returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasEndUserContact() bool {
	if o != nil && !IsNil(o.EndUserContact) {
		return true
	}

	return false
}

// SetEndUserContact gets a reference to the given string and assigns it to the EndUserContact field.
func (o *QuoteSearchResponseQuotesInner) SetEndUserContact(v string) {
	o.EndUserContact = &v
}

// GetSpecialBidNumber returns the SpecialBidNumber field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetSpecialBidNumber() string {
	if o == nil || IsNil(o.SpecialBidNumber) {
		var ret string
		return ret
	}
	return *o.SpecialBidNumber
}

// GetSpecialBidNumberOk returns a tuple with the SpecialBidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetSpecialBidNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SpecialBidNumber) {
		return nil, false
	}
	return o.SpecialBidNumber, true
}

// HasSpecialBidNumber returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasSpecialBidNumber() bool {
	if o != nil && !IsNil(o.SpecialBidNumber) {
		return true
	}

	return false
}

// SetSpecialBidNumber gets a reference to the given string and assigns it to the SpecialBidNumber field.
func (o *QuoteSearchResponseQuotesInner) SetSpecialBidNumber(v string) {
	o.SpecialBidNumber = &v
}

// GetQuoteTotal returns the QuoteTotal field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetQuoteTotal() float32 {
	if o == nil || IsNil(o.QuoteTotal) {
		var ret float32
		return ret
	}
	return *o.QuoteTotal
}

// GetQuoteTotalOk returns a tuple with the QuoteTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetQuoteTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.QuoteTotal) {
		return nil, false
	}
	return o.QuoteTotal, true
}

// HasQuoteTotal returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasQuoteTotal() bool {
	if o != nil && !IsNil(o.QuoteTotal) {
		return true
	}

	return false
}

// SetQuoteTotal gets a reference to the given float32 and assigns it to the QuoteTotal field.
func (o *QuoteSearchResponseQuotesInner) SetQuoteTotal(v float32) {
	o.QuoteTotal = &v
}

// GetQuoteStatus returns the QuoteStatus field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetQuoteStatus() string {
	if o == nil || IsNil(o.QuoteStatus) {
		var ret string
		return ret
	}
	return *o.QuoteStatus
}

// GetQuoteStatusOk returns a tuple with the QuoteStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetQuoteStatusOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteStatus) {
		return nil, false
	}
	return o.QuoteStatus, true
}

// HasQuoteStatus returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasQuoteStatus() bool {
	if o != nil && !IsNil(o.QuoteStatus) {
		return true
	}

	return false
}

// SetQuoteStatus gets a reference to the given string and assigns it to the QuoteStatus field.
func (o *QuoteSearchResponseQuotesInner) SetQuoteStatus(v string) {
	o.QuoteStatus = &v
}

// GetIngramQuoteDate returns the IngramQuoteDate field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetIngramQuoteDate() string {
	if o == nil || IsNil(o.IngramQuoteDate) {
		var ret string
		return ret
	}
	return *o.IngramQuoteDate
}

// GetIngramQuoteDateOk returns a tuple with the IngramQuoteDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetIngramQuoteDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramQuoteDate) {
		return nil, false
	}
	return o.IngramQuoteDate, true
}

// HasIngramQuoteDate returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasIngramQuoteDate() bool {
	if o != nil && !IsNil(o.IngramQuoteDate) {
		return true
	}

	return false
}

// SetIngramQuoteDate gets a reference to the given string and assigns it to the IngramQuoteDate field.
func (o *QuoteSearchResponseQuotesInner) SetIngramQuoteDate(v string) {
	o.IngramQuoteDate = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetLastModifiedDate() string {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret string
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetLastModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given string and assigns it to the LastModifiedDate field.
func (o *QuoteSearchResponseQuotesInner) SetLastModifiedDate(v string) {
	o.LastModifiedDate = &v
}

// GetIngramQuoteExpiryDate returns the IngramQuoteExpiryDate field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetIngramQuoteExpiryDate() string {
	if o == nil || IsNil(o.IngramQuoteExpiryDate) {
		var ret string
		return ret
	}
	return *o.IngramQuoteExpiryDate
}

// GetIngramQuoteExpiryDateOk returns a tuple with the IngramQuoteExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetIngramQuoteExpiryDateOk() (*string, bool) {
	if o == nil || IsNil(o.IngramQuoteExpiryDate) {
		return nil, false
	}
	return o.IngramQuoteExpiryDate, true
}

// HasIngramQuoteExpiryDate returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasIngramQuoteExpiryDate() bool {
	if o != nil && !IsNil(o.IngramQuoteExpiryDate) {
		return true
	}

	return false
}

// SetIngramQuoteExpiryDate gets a reference to the given string and assigns it to the IngramQuoteExpiryDate field.
func (o *QuoteSearchResponseQuotesInner) SetIngramQuoteExpiryDate(v string) {
	o.IngramQuoteExpiryDate = &v
}

// GetEndUserName returns the EndUserName field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetEndUserName() string {
	if o == nil || IsNil(o.EndUserName) {
		var ret string
		return ret
	}
	return *o.EndUserName
}

// GetEndUserNameOk returns a tuple with the EndUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetEndUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserName) {
		return nil, false
	}
	return o.EndUserName, true
}

// HasEndUserName returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasEndUserName() bool {
	if o != nil && !IsNil(o.EndUserName) {
		return true
	}

	return false
}

// SetEndUserName gets a reference to the given string and assigns it to the EndUserName field.
func (o *QuoteSearchResponseQuotesInner) SetEndUserName(v string) {
	o.EndUserName = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetVendor() string {
	if o == nil || IsNil(o.Vendor) {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetVendorOk() (*string, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *QuoteSearchResponseQuotesInner) SetVendor(v string) {
	o.Vendor = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *QuoteSearchResponseQuotesInner) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetQuoteType returns the QuoteType field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetQuoteType() string {
	if o == nil || IsNil(o.QuoteType) {
		var ret string
		return ret
	}
	return *o.QuoteType
}

// GetQuoteTypeOk returns a tuple with the QuoteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetQuoteTypeOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteType) {
		return nil, false
	}
	return o.QuoteType, true
}

// HasQuoteType returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasQuoteType() bool {
	if o != nil && !IsNil(o.QuoteType) {
		return true
	}

	return false
}

// SetQuoteType gets a reference to the given string and assigns it to the QuoteType field.
func (o *QuoteSearchResponseQuotesInner) SetQuoteType(v string) {
	o.QuoteType = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *QuoteSearchResponseQuotesInner) GetLinks() QuoteSearchResponseQuotesInnerLinks {
	if o == nil || IsNil(o.Links) {
		var ret QuoteSearchResponseQuotesInnerLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSearchResponseQuotesInner) GetLinksOk() (*QuoteSearchResponseQuotesInnerLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *QuoteSearchResponseQuotesInner) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given QuoteSearchResponseQuotesInnerLinks and assigns it to the Links field.
func (o *QuoteSearchResponseQuotesInner) SetLinks(v QuoteSearchResponseQuotesInnerLinks) {
	o.Links = &v
}

func (o QuoteSearchResponseQuotesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteSearchResponseQuotesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteGuid) {
		toSerialize["quoteGuid"] = o.QuoteGuid
	}
	if !IsNil(o.QuoteName) {
		toSerialize["quoteName"] = o.QuoteName
	}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.EndUserContact) {
		toSerialize["endUserContact"] = o.EndUserContact
	}
	if !IsNil(o.SpecialBidNumber) {
		toSerialize["specialBidNumber"] = o.SpecialBidNumber
	}
	if !IsNil(o.QuoteTotal) {
		toSerialize["quoteTotal"] = o.QuoteTotal
	}
	if !IsNil(o.QuoteStatus) {
		toSerialize["quoteStatus"] = o.QuoteStatus
	}
	if !IsNil(o.IngramQuoteDate) {
		toSerialize["ingramQuoteDate"] = o.IngramQuoteDate
	}
	if !IsNil(o.LastModifiedDate) {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	if !IsNil(o.IngramQuoteExpiryDate) {
		toSerialize["ingramQuoteExpiryDate"] = o.IngramQuoteExpiryDate
	}
	if !IsNil(o.EndUserName) {
		toSerialize["endUserName"] = o.EndUserName
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.QuoteType) {
		toSerialize["quoteType"] = o.QuoteType
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableQuoteSearchResponseQuotesInner struct {
	value *QuoteSearchResponseQuotesInner
	isSet bool
}

func (v NullableQuoteSearchResponseQuotesInner) Get() *QuoteSearchResponseQuotesInner {
	return v.value
}

func (v *NullableQuoteSearchResponseQuotesInner) Set(val *QuoteSearchResponseQuotesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSearchResponseQuotesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSearchResponseQuotesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSearchResponseQuotesInner(val *QuoteSearchResponseQuotesInner) *NullableQuoteSearchResponseQuotesInner {
	return &NullableQuoteSearchResponseQuotesInner{value: val, isSet: true}
}

func (v NullableQuoteSearchResponseQuotesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSearchResponseQuotesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


