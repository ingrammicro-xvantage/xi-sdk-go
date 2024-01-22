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

// checks if the QuoteListResponseQuoteSearchResponseResponsePreamble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuoteListResponseQuoteSearchResponseResponsePreamble{}

// QuoteListResponseQuoteSearchResponseResponsePreamble struct for QuoteListResponseQuoteSearchResponseResponsePreamble
type QuoteListResponseQuoteSearchResponseResponsePreamble struct {
	// Status of the Request - \"Passed\", \"Failed\"
	ResponseStatus *string `json:"responseStatus,omitempty"`
	// responseStatusCode is the code returned in response to a request. The following Codes are returned: 200 400 500
	ResponseStatusCode *string `json:"responseStatusCode,omitempty"`
	// 200 = Action was successfully received, understood and accepted. 400 = The request contains bad syntax or can not be fullfilled. This means there is a problem with the request. 500 = The server failed to fulfill an apparently valid request. This is a temporary problem, the request should be resubmitted.
	ResponseMessage *string `json:"responseMessage,omitempty"`
}

// NewQuoteListResponseQuoteSearchResponseResponsePreamble instantiates a new QuoteListResponseQuoteSearchResponseResponsePreamble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteListResponseQuoteSearchResponseResponsePreamble() *QuoteListResponseQuoteSearchResponseResponsePreamble {
	this := QuoteListResponseQuoteSearchResponseResponsePreamble{}
	return &this
}

// NewQuoteListResponseQuoteSearchResponseResponsePreambleWithDefaults instantiates a new QuoteListResponseQuoteSearchResponseResponsePreamble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteListResponseQuoteSearchResponseResponsePreambleWithDefaults() *QuoteListResponseQuoteSearchResponseResponsePreamble {
	this := QuoteListResponseQuoteSearchResponseResponsePreamble{}
	return &this
}

// GetResponseStatus returns the ResponseStatus field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatus() string {
	if o == nil || IsNil(o.ResponseStatus) {
		var ret string
		return ret
	}
	return *o.ResponseStatus
}

// GetResponseStatusOk returns a tuple with the ResponseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseStatus) {
		return nil, false
	}
	return o.ResponseStatus, true
}

// HasResponseStatus returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) HasResponseStatus() bool {
	if o != nil && !IsNil(o.ResponseStatus) {
		return true
	}

	return false
}

// SetResponseStatus gets a reference to the given string and assigns it to the ResponseStatus field.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) SetResponseStatus(v string) {
	o.ResponseStatus = &v
}

// GetResponseStatusCode returns the ResponseStatusCode field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatusCode() string {
	if o == nil || IsNil(o.ResponseStatusCode) {
		var ret string
		return ret
	}
	return *o.ResponseStatusCode
}

// GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseStatusCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseStatusCode) {
		return nil, false
	}
	return o.ResponseStatusCode, true
}

// HasResponseStatusCode returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) HasResponseStatusCode() bool {
	if o != nil && !IsNil(o.ResponseStatusCode) {
		return true
	}

	return false
}

// SetResponseStatusCode gets a reference to the given string and assigns it to the ResponseStatusCode field.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) SetResponseStatusCode(v string) {
	o.ResponseStatusCode = &v
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseMessage() string {
	if o == nil || IsNil(o.ResponseMessage) {
		var ret string
		return ret
	}
	return *o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) GetResponseMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseMessage) {
		return nil, false
	}
	return o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) HasResponseMessage() bool {
	if o != nil && !IsNil(o.ResponseMessage) {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given string and assigns it to the ResponseMessage field.
func (o *QuoteListResponseQuoteSearchResponseResponsePreamble) SetResponseMessage(v string) {
	o.ResponseMessage = &v
}

func (o QuoteListResponseQuoteSearchResponseResponsePreamble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuoteListResponseQuoteSearchResponseResponsePreamble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResponseStatus) {
		toSerialize["responseStatus"] = o.ResponseStatus
	}
	if !IsNil(o.ResponseStatusCode) {
		toSerialize["responseStatusCode"] = o.ResponseStatusCode
	}
	if !IsNil(o.ResponseMessage) {
		toSerialize["responseMessage"] = o.ResponseMessage
	}
	return toSerialize, nil
}

type NullableQuoteListResponseQuoteSearchResponseResponsePreamble struct {
	value *QuoteListResponseQuoteSearchResponseResponsePreamble
	isSet bool
}

func (v NullableQuoteListResponseQuoteSearchResponseResponsePreamble) Get() *QuoteListResponseQuoteSearchResponseResponsePreamble {
	return v.value
}

func (v *NullableQuoteListResponseQuoteSearchResponseResponsePreamble) Set(val *QuoteListResponseQuoteSearchResponseResponsePreamble) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteListResponseQuoteSearchResponseResponsePreamble) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteListResponseQuoteSearchResponseResponsePreamble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteListResponseQuoteSearchResponseResponsePreamble(val *QuoteListResponseQuoteSearchResponseResponsePreamble) *NullableQuoteListResponseQuoteSearchResponseResponsePreamble {
	return &NullableQuoteListResponseQuoteSearchResponseResponsePreamble{value: val, isSet: true}
}

func (v NullableQuoteListResponseQuoteSearchResponseResponsePreamble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteListResponseQuoteSearchResponseResponsePreamble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

