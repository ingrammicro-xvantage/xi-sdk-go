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

// checks if the DealsSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DealsSearchResponse{}

// DealsSearchResponse struct for DealsSearchResponse
type DealsSearchResponse struct {
	// Number of records found.
	RecordsFound *int32 `json:"recordsFound,omitempty"`
	// Number of records in a page.
	PageSize *int32 `json:"pageSize,omitempty"`
	// Number of page.
	PageNumber *int32 `json:"pageNumber,omitempty"`
	Deals []DealsSearchResponseDealsInner `json:"deals,omitempty"`
	// URL for the next page.
	NextPage *string `json:"nextPage,omitempty"`
	// URL for the previous page.
	PreviousPage *string `json:"previousPage,omitempty"`
}

// NewDealsSearchResponse instantiates a new DealsSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDealsSearchResponse() *DealsSearchResponse {
	this := DealsSearchResponse{}
	return &this
}

// NewDealsSearchResponseWithDefaults instantiates a new DealsSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDealsSearchResponseWithDefaults() *DealsSearchResponse {
	this := DealsSearchResponse{}
	return &this
}

// GetRecordsFound returns the RecordsFound field value if set, zero value otherwise.
func (o *DealsSearchResponse) GetRecordsFound() int32 {
	if o == nil || IsNil(o.RecordsFound) {
		var ret int32
		return ret
	}
	return *o.RecordsFound
}

// GetRecordsFoundOk returns a tuple with the RecordsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponse) GetRecordsFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsFound) {
		return nil, false
	}
	return o.RecordsFound, true
}

// HasRecordsFound returns a boolean if a field has been set.
func (o *DealsSearchResponse) HasRecordsFound() bool {
	if o != nil && !IsNil(o.RecordsFound) {
		return true
	}

	return false
}

// SetRecordsFound gets a reference to the given int32 and assigns it to the RecordsFound field.
func (o *DealsSearchResponse) SetRecordsFound(v int32) {
	o.RecordsFound = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DealsSearchResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DealsSearchResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *DealsSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DealsSearchResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DealsSearchResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *DealsSearchResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetDeals returns the Deals field value if set, zero value otherwise.
func (o *DealsSearchResponse) GetDeals() []DealsSearchResponseDealsInner {
	if o == nil || IsNil(o.Deals) {
		var ret []DealsSearchResponseDealsInner
		return ret
	}
	return o.Deals
}

// GetDealsOk returns a tuple with the Deals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponse) GetDealsOk() ([]DealsSearchResponseDealsInner, bool) {
	if o == nil || IsNil(o.Deals) {
		return nil, false
	}
	return o.Deals, true
}

// HasDeals returns a boolean if a field has been set.
func (o *DealsSearchResponse) HasDeals() bool {
	if o != nil && !IsNil(o.Deals) {
		return true
	}

	return false
}

// SetDeals gets a reference to the given []DealsSearchResponseDealsInner and assigns it to the Deals field.
func (o *DealsSearchResponse) SetDeals(v []DealsSearchResponseDealsInner) {
	o.Deals = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *DealsSearchResponse) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponse) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *DealsSearchResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *DealsSearchResponse) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *DealsSearchResponse) GetPreviousPage() string {
	if o == nil || IsNil(o.PreviousPage) {
		var ret string
		return ret
	}
	return *o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DealsSearchResponse) GetPreviousPageOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousPage) {
		return nil, false
	}
	return o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *DealsSearchResponse) HasPreviousPage() bool {
	if o != nil && !IsNil(o.PreviousPage) {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given string and assigns it to the PreviousPage field.
func (o *DealsSearchResponse) SetPreviousPage(v string) {
	o.PreviousPage = &v
}

func (o DealsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DealsSearchResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Deals) {
		toSerialize["deals"] = o.Deals
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PreviousPage) {
		toSerialize["previousPage"] = o.PreviousPage
	}
	return toSerialize, nil
}

type NullableDealsSearchResponse struct {
	value *DealsSearchResponse
	isSet bool
}

func (v NullableDealsSearchResponse) Get() *DealsSearchResponse {
	return v.value
}

func (v *NullableDealsSearchResponse) Set(val *DealsSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDealsSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDealsSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDealsSearchResponse(val *DealsSearchResponse) *NullableDealsSearchResponse {
	return &NullableDealsSearchResponse{value: val, isSet: true}
}

func (v NullableDealsSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDealsSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


