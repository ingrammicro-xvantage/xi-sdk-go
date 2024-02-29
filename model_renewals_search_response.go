/*
XI Sdk Resellers

For Ingram Micro Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the RenewalsSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsSearchResponse{}

// RenewalsSearchResponse struct for RenewalsSearchResponse
type RenewalsSearchResponse struct {
	// Number of records found.
	RecordsFound *int32 `json:"recordsFound,omitempty"`
	// Number of records in a page.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Number of page.
	PageNumber *int32 `json:"pageNumber,omitempty"`
	Renewals []RenewalsSearchResponseRenewalsInner `json:"renewals,omitempty"`
	// URL for the next page.
	NextPage *string `json:"nextPage,omitempty"`
	// URL for the previous page.
	PreviousPage *string `json:"previousPage,omitempty"`
}

// NewRenewalsSearchResponse instantiates a new RenewalsSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsSearchResponse() *RenewalsSearchResponse {
	this := RenewalsSearchResponse{}
	return &this
}

// NewRenewalsSearchResponseWithDefaults instantiates a new RenewalsSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsSearchResponseWithDefaults() *RenewalsSearchResponse {
	this := RenewalsSearchResponse{}
	return &this
}

// GetRecordsFound returns the RecordsFound field value if set, zero value otherwise.
func (o *RenewalsSearchResponse) GetRecordsFound() int32 {
	if o == nil || IsNil(o.RecordsFound) {
		var ret int32
		return ret
	}
	return *o.RecordsFound
}

// GetRecordsFoundOk returns a tuple with the RecordsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponse) GetRecordsFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsFound) {
		return nil, false
	}
	return o.RecordsFound, true
}

// HasRecordsFound returns a boolean if a field has been set.
func (o *RenewalsSearchResponse) HasRecordsFound() bool {
	if o != nil && !IsNil(o.RecordsFound) {
		return true
	}

	return false
}

// SetRecordsFound gets a reference to the given int32 and assigns it to the RecordsFound field.
func (o *RenewalsSearchResponse) SetRecordsFound(v int32) {
	o.RecordsFound = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *RenewalsSearchResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *RenewalsSearchResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *RenewalsSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *RenewalsSearchResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *RenewalsSearchResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *RenewalsSearchResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetRenewals returns the Renewals field value if set, zero value otherwise.
func (o *RenewalsSearchResponse) GetRenewals() []RenewalsSearchResponseRenewalsInner {
	if o == nil || IsNil(o.Renewals) {
		var ret []RenewalsSearchResponseRenewalsInner
		return ret
	}
	return o.Renewals
}

// GetRenewalsOk returns a tuple with the Renewals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponse) GetRenewalsOk() ([]RenewalsSearchResponseRenewalsInner, bool) {
	if o == nil || IsNil(o.Renewals) {
		return nil, false
	}
	return o.Renewals, true
}

// HasRenewals returns a boolean if a field has been set.
func (o *RenewalsSearchResponse) HasRenewals() bool {
	if o != nil && !IsNil(o.Renewals) {
		return true
	}

	return false
}

// SetRenewals gets a reference to the given []RenewalsSearchResponseRenewalsInner and assigns it to the Renewals field.
func (o *RenewalsSearchResponse) SetRenewals(v []RenewalsSearchResponseRenewalsInner) {
	o.Renewals = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *RenewalsSearchResponse) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponse) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *RenewalsSearchResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *RenewalsSearchResponse) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *RenewalsSearchResponse) GetPreviousPage() string {
	if o == nil || IsNil(o.PreviousPage) {
		var ret string
		return ret
	}
	return *o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsSearchResponse) GetPreviousPageOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousPage) {
		return nil, false
	}
	return o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *RenewalsSearchResponse) HasPreviousPage() bool {
	if o != nil && !IsNil(o.PreviousPage) {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given string and assigns it to the PreviousPage field.
func (o *RenewalsSearchResponse) SetPreviousPage(v string) {
	o.PreviousPage = &v
}

func (o RenewalsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordsFound) {
		toSerialize["recordsFound"] = o.RecordsFound
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.Renewals) {
		toSerialize["renewals"] = o.Renewals
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PreviousPage) {
		toSerialize["previousPage"] = o.PreviousPage
	}
	return toSerialize, nil
}

type NullableRenewalsSearchResponse struct {
	value *RenewalsSearchResponse
	isSet bool
}

func (v NullableRenewalsSearchResponse) Get() *RenewalsSearchResponse {
	return v.value
}

func (v *NullableRenewalsSearchResponse) Set(val *RenewalsSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsSearchResponse(val *RenewalsSearchResponse) *NullableRenewalsSearchResponse {
	return &NullableRenewalsSearchResponse{value: val, isSet: true}
}

func (v NullableRenewalsSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


