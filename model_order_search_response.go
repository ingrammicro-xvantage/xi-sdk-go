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

// checks if the OrderSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSearchResponse{}

// OrderSearchResponse struct for OrderSearchResponse
type OrderSearchResponse struct {
	// No of recourds found for the search.
	RecordsFound *int32 `json:"recordsFound,omitempty"`
	// No of results per page.(default is 25)
	PageSize *int32 `json:"pageSize,omitempty"`
	// Current page number.(default is 1)
	PageNumber *int32 `json:"pageNumber,omitempty"`
	// The details for the order.
	Orders []OrderSearchResponseOrdersInner `json:"orders,omitempty"`
	// link/URL for accessing next page.
	NextPage *string `json:"nextPage,omitempty"`
	// link/URL for accessing previous page.
	PreviousPage *string `json:"previousPage,omitempty"`
}

// NewOrderSearchResponse instantiates a new OrderSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSearchResponse() *OrderSearchResponse {
	this := OrderSearchResponse{}
	return &this
}

// NewOrderSearchResponseWithDefaults instantiates a new OrderSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSearchResponseWithDefaults() *OrderSearchResponse {
	this := OrderSearchResponse{}
	return &this
}

// GetRecordsFound returns the RecordsFound field value if set, zero value otherwise.
func (o *OrderSearchResponse) GetRecordsFound() int32 {
	if o == nil || IsNil(o.RecordsFound) {
		var ret int32
		return ret
	}
	return *o.RecordsFound
}

// GetRecordsFoundOk returns a tuple with the RecordsFound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponse) GetRecordsFoundOk() (*int32, bool) {
	if o == nil || IsNil(o.RecordsFound) {
		return nil, false
	}
	return o.RecordsFound, true
}

// HasRecordsFound returns a boolean if a field has been set.
func (o *OrderSearchResponse) HasRecordsFound() bool {
	if o != nil && !IsNil(o.RecordsFound) {
		return true
	}

	return false
}

// SetRecordsFound gets a reference to the given int32 and assigns it to the RecordsFound field.
func (o *OrderSearchResponse) SetRecordsFound(v int32) {
	o.RecordsFound = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *OrderSearchResponse) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponse) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *OrderSearchResponse) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *OrderSearchResponse) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *OrderSearchResponse) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponse) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *OrderSearchResponse) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *OrderSearchResponse) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *OrderSearchResponse) GetOrders() []OrderSearchResponseOrdersInner {
	if o == nil || IsNil(o.Orders) {
		var ret []OrderSearchResponseOrdersInner
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponse) GetOrdersOk() ([]OrderSearchResponseOrdersInner, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *OrderSearchResponse) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderSearchResponseOrdersInner and assigns it to the Orders field.
func (o *OrderSearchResponse) SetOrders(v []OrderSearchResponseOrdersInner) {
	o.Orders = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *OrderSearchResponse) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponse) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *OrderSearchResponse) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *OrderSearchResponse) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *OrderSearchResponse) GetPreviousPage() string {
	if o == nil || IsNil(o.PreviousPage) {
		var ret string
		return ret
	}
	return *o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSearchResponse) GetPreviousPageOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousPage) {
		return nil, false
	}
	return o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *OrderSearchResponse) HasPreviousPage() bool {
	if o != nil && !IsNil(o.PreviousPage) {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given string and assigns it to the PreviousPage field.
func (o *OrderSearchResponse) SetPreviousPage(v string) {
	o.PreviousPage = &v
}

func (o OrderSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSearchResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PreviousPage) {
		toSerialize["previousPage"] = o.PreviousPage
	}
	return toSerialize, nil
}

type NullableOrderSearchResponse struct {
	value *OrderSearchResponse
	isSet bool
}

func (v NullableOrderSearchResponse) Get() *OrderSearchResponse {
	return v.value
}

func (v *NullableOrderSearchResponse) Set(val *OrderSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSearchResponse(val *OrderSearchResponse) *NullableOrderSearchResponse {
	return &NullableOrderSearchResponse{value: val, isSet: true}
}

func (v NullableOrderSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


