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

// checks if the QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest{}

// QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest struct for QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest
type QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest struct {
	// Unique identifier generated by Ingram Micro's CRM specific to each quote. When applying a filter to the quoteNumber and including a partial quote number in the filter, all quotes containing any information included in the filter can be retrieved as a subset of all available customer quotes.
	QuoteNumber *string `json:"quoteNumber,omitempty"`
	// Special Pricing Bid Number, also referred to as a Dart Number by some vendors, is a unique identifier associated with vendor specific products and discounts.
	BidNumber *string `json:"bidNumber,omitempty"`
	// End User Name is the end customer name that is associated with a quote in Ingram Micro's CRM
	EndUserName *string `json:"endUserName,omitempty"`
	// Filter to select the beginning date of a desired date range. The default filter is set to the date the user is logged-in to request quotes. Date format: YYYY-MM-DD - An incorrect date input will result in a message \"Date must be entered as YYYY-MM-DD\"
	FromDate *string `json:"fromDate,omitempty"`
	// Filter to select the end date of a desired date range. The default number of days to request is the previous 30 days from the date user has logged in. Date format: YYYY-MM-DD - An incorrect date input will result in a message \"Date must be entered as YYYY-MM-DD\"
	ToDate *string `json:"toDate,omitempty"`
	// Page index or page number for the list of quotes being returned. When less than 25 quotes are returned, the page number will be \"1\". In cases where more than 25 quotes are returned, and the default quotes per page are 25 (see recordPerPage), then the list will continue on subsequent pages.
	PageIndex *string `json:"pageIndex,omitempty"`
	// Number of records (quotes) to display per page in the quote list. The default is 25, but may be increased using the filter by up to 100 records per page. If more than 100 records are requested a message will be returned \"The number of records requested exceeds the 100 record limit.\" 
	RecordsPerPage *string `json:"recordsPerPage,omitempty"`
	// Sort applies to the selected column (sortingColumnName) and may be specified in Ascending (asc) or Descending (desc) order. The default sort is Descending (desc) - most recent first.
	Sorting *string `json:"sorting,omitempty"`
	// Refers to the column selected to apply the sorting criteria. The default column is dateCreated and will sort by the most recently created quote first with the following in descending order. The default filter retrieves quotes created within the last 30 days. Filtering allows user to select a specific column to sort: quoteNumber, createdDate, lastModifiedDate, expiryDate, and endUserName.
	SortingColumnName *string `json:"sortingColumnName,omitempty"`
	// Unique identifier used to identify the third party source accessing the services.
	ThirdPartySource *string `json:"thirdPartySource,omitempty"`
}

// NewQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest instantiates a new QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest() *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest {
	this := QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest{}
	return &this
}

// NewQuoteListRequestQuoteSearchRequestRetrieveQuoteRequestWithDefaults instantiates a new QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteListRequestQuoteSearchRequestRetrieveQuoteRequestWithDefaults() *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest {
	this := QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest{}
	return &this
}

// GetQuoteNumber returns the QuoteNumber field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetQuoteNumber() string {
	if o == nil || IsNil(o.QuoteNumber) {
		var ret string
		return ret
	}
	return *o.QuoteNumber
}

// GetQuoteNumberOk returns a tuple with the QuoteNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetQuoteNumberOk() (*string, bool) {
	if o == nil || IsNil(o.QuoteNumber) {
		return nil, false
	}
	return o.QuoteNumber, true
}

// HasQuoteNumber returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasQuoteNumber() bool {
	if o != nil && !IsNil(o.QuoteNumber) {
		return true
	}

	return false
}

// SetQuoteNumber gets a reference to the given string and assigns it to the QuoteNumber field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetQuoteNumber(v string) {
	o.QuoteNumber = &v
}

// GetBidNumber returns the BidNumber field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetBidNumber() string {
	if o == nil || IsNil(o.BidNumber) {
		var ret string
		return ret
	}
	return *o.BidNumber
}

// GetBidNumberOk returns a tuple with the BidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetBidNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BidNumber) {
		return nil, false
	}
	return o.BidNumber, true
}

// HasBidNumber returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasBidNumber() bool {
	if o != nil && !IsNil(o.BidNumber) {
		return true
	}

	return false
}

// SetBidNumber gets a reference to the given string and assigns it to the BidNumber field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetBidNumber(v string) {
	o.BidNumber = &v
}

// GetEndUserName returns the EndUserName field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetEndUserName() string {
	if o == nil || IsNil(o.EndUserName) {
		var ret string
		return ret
	}
	return *o.EndUserName
}

// GetEndUserNameOk returns a tuple with the EndUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetEndUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserName) {
		return nil, false
	}
	return o.EndUserName, true
}

// HasEndUserName returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasEndUserName() bool {
	if o != nil && !IsNil(o.EndUserName) {
		return true
	}

	return false
}

// SetEndUserName gets a reference to the given string and assigns it to the EndUserName field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetEndUserName(v string) {
	o.EndUserName = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetFromDate() string {
	if o == nil || IsNil(o.FromDate) {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetFromDateOk() (*string, bool) {
	if o == nil || IsNil(o.FromDate) {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasFromDate() bool {
	if o != nil && !IsNil(o.FromDate) {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetToDate() string {
	if o == nil || IsNil(o.ToDate) {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetToDateOk() (*string, bool) {
	if o == nil || IsNil(o.ToDate) {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasToDate() bool {
	if o != nil && !IsNil(o.ToDate) {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetToDate(v string) {
	o.ToDate = &v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetPageIndex() string {
	if o == nil || IsNil(o.PageIndex) {
		var ret string
		return ret
	}
	return *o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetPageIndexOk() (*string, bool) {
	if o == nil || IsNil(o.PageIndex) {
		return nil, false
	}
	return o.PageIndex, true
}

// HasPageIndex returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasPageIndex() bool {
	if o != nil && !IsNil(o.PageIndex) {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given string and assigns it to the PageIndex field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetPageIndex(v string) {
	o.PageIndex = &v
}

// GetRecordsPerPage returns the RecordsPerPage field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetRecordsPerPage() string {
	if o == nil || IsNil(o.RecordsPerPage) {
		var ret string
		return ret
	}
	return *o.RecordsPerPage
}

// GetRecordsPerPageOk returns a tuple with the RecordsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetRecordsPerPageOk() (*string, bool) {
	if o == nil || IsNil(o.RecordsPerPage) {
		return nil, false
	}
	return o.RecordsPerPage, true
}

// HasRecordsPerPage returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasRecordsPerPage() bool {
	if o != nil && !IsNil(o.RecordsPerPage) {
		return true
	}

	return false
}

// SetRecordsPerPage gets a reference to the given string and assigns it to the RecordsPerPage field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetRecordsPerPage(v string) {
	o.RecordsPerPage = &v
}

// GetSorting returns the Sorting field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetSorting() string {
	if o == nil || IsNil(o.Sorting) {
		var ret string
		return ret
	}
	return *o.Sorting
}

// GetSortingOk returns a tuple with the Sorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetSortingOk() (*string, bool) {
	if o == nil || IsNil(o.Sorting) {
		return nil, false
	}
	return o.Sorting, true
}

// HasSorting returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasSorting() bool {
	if o != nil && !IsNil(o.Sorting) {
		return true
	}

	return false
}

// SetSorting gets a reference to the given string and assigns it to the Sorting field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetSorting(v string) {
	o.Sorting = &v
}

// GetSortingColumnName returns the SortingColumnName field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetSortingColumnName() string {
	if o == nil || IsNil(o.SortingColumnName) {
		var ret string
		return ret
	}
	return *o.SortingColumnName
}

// GetSortingColumnNameOk returns a tuple with the SortingColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetSortingColumnNameOk() (*string, bool) {
	if o == nil || IsNil(o.SortingColumnName) {
		return nil, false
	}
	return o.SortingColumnName, true
}

// HasSortingColumnName returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasSortingColumnName() bool {
	if o != nil && !IsNil(o.SortingColumnName) {
		return true
	}

	return false
}

// SetSortingColumnName gets a reference to the given string and assigns it to the SortingColumnName field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetSortingColumnName(v string) {
	o.SortingColumnName = &v
}

// GetThirdPartySource returns the ThirdPartySource field value if set, zero value otherwise.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetThirdPartySource() string {
	if o == nil || IsNil(o.ThirdPartySource) {
		var ret string
		return ret
	}
	return *o.ThirdPartySource
}

// GetThirdPartySourceOk returns a tuple with the ThirdPartySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) GetThirdPartySourceOk() (*string, bool) {
	if o == nil || IsNil(o.ThirdPartySource) {
		return nil, false
	}
	return o.ThirdPartySource, true
}

// HasThirdPartySource returns a boolean if a field has been set.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) HasThirdPartySource() bool {
	if o != nil && !IsNil(o.ThirdPartySource) {
		return true
	}

	return false
}

// SetThirdPartySource gets a reference to the given string and assigns it to the ThirdPartySource field.
func (o *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) SetThirdPartySource(v string) {
	o.ThirdPartySource = &v
}

func (o QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QuoteNumber) {
		toSerialize["quoteNumber"] = o.QuoteNumber
	}
	if !IsNil(o.BidNumber) {
		toSerialize["bidNumber"] = o.BidNumber
	}
	if !IsNil(o.EndUserName) {
		toSerialize["endUserName"] = o.EndUserName
	}
	if !IsNil(o.FromDate) {
		toSerialize["fromDate"] = o.FromDate
	}
	if !IsNil(o.ToDate) {
		toSerialize["toDate"] = o.ToDate
	}
	if !IsNil(o.PageIndex) {
		toSerialize["pageIndex"] = o.PageIndex
	}
	if !IsNil(o.RecordsPerPage) {
		toSerialize["recordsPerPage"] = o.RecordsPerPage
	}
	if !IsNil(o.Sorting) {
		toSerialize["sorting"] = o.Sorting
	}
	if !IsNil(o.SortingColumnName) {
		toSerialize["sortingColumnName"] = o.SortingColumnName
	}
	if !IsNil(o.ThirdPartySource) {
		toSerialize["thirdPartySource"] = o.ThirdPartySource
	}
	return toSerialize, nil
}

type NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest struct {
	value *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest
	isSet bool
}

func (v NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) Get() *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest {
	return v.value
}

func (v *NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) Set(val *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest(val *QuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) *NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest {
	return &NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest{value: val, isSet: true}
}

func (v NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteListRequestQuoteSearchRequestRetrieveQuoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

