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

// checks if the QuoteDetailsQuoteDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteDetailsQuoteDetailResponse{}

// QuoteDetailsQuoteDetailResponse struct for QuoteDetailsQuoteDetailResponse
type QuoteDetailsQuoteDetailResponse struct {
	ResponsePreamble *QuoteDetailsQuoteDetailResponseResponsePreamble `json:"responsePreamble,omitempty"`
	RetrieveQuoteResponse *QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse `json:"retrieveQuoteResponse,omitempty"`
	QuoteProductList []QuoteProductList `json:"quoteProductList,omitempty"`
	TotalQuoteProductCount *string `json:"totalQuoteProductCount,omitempty"`
	TotalExtendedMsrp *string `json:"totalExtendedMsrp,omitempty"`
	TotalQuantity *int32 `json:"totalQuantity,omitempty"`
	TotalExtendedQuotePrice *string `json:"totalExtendedQuotePrice,omitempty"`
}

// NewQuoteDetailsQuoteDetailResponse instantiates a new QuoteDetailsQuoteDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteDetailsQuoteDetailResponse() *QuoteDetailsQuoteDetailResponse {
	this := QuoteDetailsQuoteDetailResponse{}
	return &this
}

// NewQuoteDetailsQuoteDetailResponseWithDefaults instantiates a new QuoteDetailsQuoteDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteDetailsQuoteDetailResponseWithDefaults() *QuoteDetailsQuoteDetailResponse {
	this := QuoteDetailsQuoteDetailResponse{}
	return &this
}

// GetResponsePreamble returns the ResponsePreamble field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetResponsePreamble() QuoteDetailsQuoteDetailResponseResponsePreamble {
	if o == nil || IsNil(o.ResponsePreamble) {
		var ret QuoteDetailsQuoteDetailResponseResponsePreamble
		return ret
	}
	return *o.ResponsePreamble
}

// GetResponsePreambleOk returns a tuple with the ResponsePreamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetResponsePreambleOk() (*QuoteDetailsQuoteDetailResponseResponsePreamble, bool) {
	if o == nil || IsNil(o.ResponsePreamble) {
		return nil, false
	}
	return o.ResponsePreamble, true
}

// HasResponsePreamble returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasResponsePreamble() bool {
	if o != nil && !IsNil(o.ResponsePreamble) {
		return true
	}

	return false
}

// SetResponsePreamble gets a reference to the given QuoteDetailsQuoteDetailResponseResponsePreamble and assigns it to the ResponsePreamble field.
func (o *QuoteDetailsQuoteDetailResponse) SetResponsePreamble(v QuoteDetailsQuoteDetailResponseResponsePreamble) {
	o.ResponsePreamble = &v
}

// GetRetrieveQuoteResponse returns the RetrieveQuoteResponse field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetRetrieveQuoteResponse() QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse {
	if o == nil || IsNil(o.RetrieveQuoteResponse) {
		var ret QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse
		return ret
	}
	return *o.RetrieveQuoteResponse
}

// GetRetrieveQuoteResponseOk returns a tuple with the RetrieveQuoteResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetRetrieveQuoteResponseOk() (*QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse, bool) {
	if o == nil || IsNil(o.RetrieveQuoteResponse) {
		return nil, false
	}
	return o.RetrieveQuoteResponse, true
}

// HasRetrieveQuoteResponse returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasRetrieveQuoteResponse() bool {
	if o != nil && !IsNil(o.RetrieveQuoteResponse) {
		return true
	}

	return false
}

// SetRetrieveQuoteResponse gets a reference to the given QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse and assigns it to the RetrieveQuoteResponse field.
func (o *QuoteDetailsQuoteDetailResponse) SetRetrieveQuoteResponse(v QuoteDetailsQuoteDetailResponseRetrieveQuoteResponse) {
	o.RetrieveQuoteResponse = &v
}

// GetQuoteProductList returns the QuoteProductList field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetQuoteProductList() []QuoteProductList {
	if o == nil || IsNil(o.QuoteProductList) {
		var ret []QuoteProductList
		return ret
	}
	return o.QuoteProductList
}

// GetQuoteProductListOk returns a tuple with the QuoteProductList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetQuoteProductListOk() ([]QuoteProductList, bool) {
	if o == nil || IsNil(o.QuoteProductList) {
		return nil, false
	}
	return o.QuoteProductList, true
}

// HasQuoteProductList returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasQuoteProductList() bool {
	if o != nil && !IsNil(o.QuoteProductList) {
		return true
	}

	return false
}

// SetQuoteProductList gets a reference to the given []QuoteProductList and assigns it to the QuoteProductList field.
func (o *QuoteDetailsQuoteDetailResponse) SetQuoteProductList(v []QuoteProductList) {
	o.QuoteProductList = v
}

// GetTotalQuoteProductCount returns the TotalQuoteProductCount field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuoteProductCount() string {
	if o == nil || IsNil(o.TotalQuoteProductCount) {
		var ret string
		return ret
	}
	return *o.TotalQuoteProductCount
}

// GetTotalQuoteProductCountOk returns a tuple with the TotalQuoteProductCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuoteProductCountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalQuoteProductCount) {
		return nil, false
	}
	return o.TotalQuoteProductCount, true
}

// HasTotalQuoteProductCount returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasTotalQuoteProductCount() bool {
	if o != nil && !IsNil(o.TotalQuoteProductCount) {
		return true
	}

	return false
}

// SetTotalQuoteProductCount gets a reference to the given string and assigns it to the TotalQuoteProductCount field.
func (o *QuoteDetailsQuoteDetailResponse) SetTotalQuoteProductCount(v string) {
	o.TotalQuoteProductCount = &v
}

// GetTotalExtendedMsrp returns the TotalExtendedMsrp field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedMsrp() string {
	if o == nil || IsNil(o.TotalExtendedMsrp) {
		var ret string
		return ret
	}
	return *o.TotalExtendedMsrp
}

// GetTotalExtendedMsrpOk returns a tuple with the TotalExtendedMsrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedMsrpOk() (*string, bool) {
	if o == nil || IsNil(o.TotalExtendedMsrp) {
		return nil, false
	}
	return o.TotalExtendedMsrp, true
}

// HasTotalExtendedMsrp returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasTotalExtendedMsrp() bool {
	if o != nil && !IsNil(o.TotalExtendedMsrp) {
		return true
	}

	return false
}

// SetTotalExtendedMsrp gets a reference to the given string and assigns it to the TotalExtendedMsrp field.
func (o *QuoteDetailsQuoteDetailResponse) SetTotalExtendedMsrp(v string) {
	o.TotalExtendedMsrp = &v
}

// GetTotalQuantity returns the TotalQuantity field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuantity() int32 {
	if o == nil || IsNil(o.TotalQuantity) {
		var ret int32
		return ret
	}
	return *o.TotalQuantity
}

// GetTotalQuantityOk returns a tuple with the TotalQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalQuantity) {
		return nil, false
	}
	return o.TotalQuantity, true
}

// HasTotalQuantity returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasTotalQuantity() bool {
	if o != nil && !IsNil(o.TotalQuantity) {
		return true
	}

	return false
}

// SetTotalQuantity gets a reference to the given int32 and assigns it to the TotalQuantity field.
func (o *QuoteDetailsQuoteDetailResponse) SetTotalQuantity(v int32) {
	o.TotalQuantity = &v
}

// GetTotalExtendedQuotePrice returns the TotalExtendedQuotePrice field value if set, zero value otherwise.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedQuotePrice() string {
	if o == nil || IsNil(o.TotalExtendedQuotePrice) {
		var ret string
		return ret
	}
	return *o.TotalExtendedQuotePrice
}

// GetTotalExtendedQuotePriceOk returns a tuple with the TotalExtendedQuotePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteDetailsQuoteDetailResponse) GetTotalExtendedQuotePriceOk() (*string, bool) {
	if o == nil || IsNil(o.TotalExtendedQuotePrice) {
		return nil, false
	}
	return o.TotalExtendedQuotePrice, true
}

// HasTotalExtendedQuotePrice returns a boolean if a field has been set.
func (o *QuoteDetailsQuoteDetailResponse) HasTotalExtendedQuotePrice() bool {
	if o != nil && !IsNil(o.TotalExtendedQuotePrice) {
		return true
	}

	return false
}

// SetTotalExtendedQuotePrice gets a reference to the given string and assigns it to the TotalExtendedQuotePrice field.
func (o *QuoteDetailsQuoteDetailResponse) SetTotalExtendedQuotePrice(v string) {
	o.TotalExtendedQuotePrice = &v
}

func (o QuoteDetailsQuoteDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteDetailsQuoteDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResponsePreamble) {
		toSerialize["responsePreamble"] = o.ResponsePreamble
	}
	if !IsNil(o.RetrieveQuoteResponse) {
		toSerialize["retrieveQuoteResponse"] = o.RetrieveQuoteResponse
	}
	if !IsNil(o.QuoteProductList) {
		toSerialize["quoteProductList"] = o.QuoteProductList
	}
	if !IsNil(o.TotalQuoteProductCount) {
		toSerialize["totalQuoteProductCount"] = o.TotalQuoteProductCount
	}
	if !IsNil(o.TotalExtendedMsrp) {
		toSerialize["totalExtendedMsrp"] = o.TotalExtendedMsrp
	}
	if !IsNil(o.TotalQuantity) {
		toSerialize["totalQuantity"] = o.TotalQuantity
	}
	if !IsNil(o.TotalExtendedQuotePrice) {
		toSerialize["totalExtendedQuotePrice"] = o.TotalExtendedQuotePrice
	}
	return toSerialize, nil
}

type NullableQuoteDetailsQuoteDetailResponse struct {
	value *QuoteDetailsQuoteDetailResponse
	isSet bool
}

func (v NullableQuoteDetailsQuoteDetailResponse) Get() *QuoteDetailsQuoteDetailResponse {
	return v.value
}

func (v *NullableQuoteDetailsQuoteDetailResponse) Set(val *QuoteDetailsQuoteDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteDetailsQuoteDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteDetailsQuoteDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteDetailsQuoteDetailResponse(val *QuoteDetailsQuoteDetailResponse) *NullableQuoteDetailsQuoteDetailResponse {
	return &NullableQuoteDetailsQuoteDetailResponse{value: val, isSet: true}
}

func (v NullableQuoteDetailsQuoteDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteDetailsQuoteDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

