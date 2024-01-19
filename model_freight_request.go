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

// checks if the FreightRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreightRequest{}

// FreightRequest struct for FreightRequest
type FreightRequest struct {
	// Suffix used to identify billing address. Created during onboarding. Resellers are provided with one or more address IDs depending on how many bill to addresses they need for various flooring companies they are using for credit.
	BillToAddressId *string `json:"billToAddressId,omitempty"`
	// The ID references the reseller's address in Ingram Micro's system for shipping. Provided to resellers during the onboarding process.
	ShipToAddressId *string `json:"shipToAddressId,omitempty"`
	// The shipping information.
	ShipToAddress []FreightRequestShipToAddressInner `json:"shipToAddress,omitempty"`
	Lines []FreightRequestLinesInner `json:"lines,omitempty"`
}

// NewFreightRequest instantiates a new FreightRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreightRequest() *FreightRequest {
	this := FreightRequest{}
	return &this
}

// NewFreightRequestWithDefaults instantiates a new FreightRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreightRequestWithDefaults() *FreightRequest {
	this := FreightRequest{}
	return &this
}

// GetBillToAddressId returns the BillToAddressId field value if set, zero value otherwise.
func (o *FreightRequest) GetBillToAddressId() string {
	if o == nil || IsNil(o.BillToAddressId) {
		var ret string
		return ret
	}
	return *o.BillToAddressId
}

// GetBillToAddressIdOk returns a tuple with the BillToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightRequest) GetBillToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.BillToAddressId) {
		return nil, false
	}
	return o.BillToAddressId, true
}

// HasBillToAddressId returns a boolean if a field has been set.
func (o *FreightRequest) HasBillToAddressId() bool {
	if o != nil && !IsNil(o.BillToAddressId) {
		return true
	}

	return false
}

// SetBillToAddressId gets a reference to the given string and assigns it to the BillToAddressId field.
func (o *FreightRequest) SetBillToAddressId(v string) {
	o.BillToAddressId = &v
}

// GetShipToAddressId returns the ShipToAddressId field value if set, zero value otherwise.
func (o *FreightRequest) GetShipToAddressId() string {
	if o == nil || IsNil(o.ShipToAddressId) {
		var ret string
		return ret
	}
	return *o.ShipToAddressId
}

// GetShipToAddressIdOk returns a tuple with the ShipToAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightRequest) GetShipToAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipToAddressId) {
		return nil, false
	}
	return o.ShipToAddressId, true
}

// HasShipToAddressId returns a boolean if a field has been set.
func (o *FreightRequest) HasShipToAddressId() bool {
	if o != nil && !IsNil(o.ShipToAddressId) {
		return true
	}

	return false
}

// SetShipToAddressId gets a reference to the given string and assigns it to the ShipToAddressId field.
func (o *FreightRequest) SetShipToAddressId(v string) {
	o.ShipToAddressId = &v
}

// GetShipToAddress returns the ShipToAddress field value if set, zero value otherwise.
func (o *FreightRequest) GetShipToAddress() []FreightRequestShipToAddressInner {
	if o == nil || IsNil(o.ShipToAddress) {
		var ret []FreightRequestShipToAddressInner
		return ret
	}
	return o.ShipToAddress
}

// GetShipToAddressOk returns a tuple with the ShipToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightRequest) GetShipToAddressOk() ([]FreightRequestShipToAddressInner, bool) {
	if o == nil || IsNil(o.ShipToAddress) {
		return nil, false
	}
	return o.ShipToAddress, true
}

// HasShipToAddress returns a boolean if a field has been set.
func (o *FreightRequest) HasShipToAddress() bool {
	if o != nil && !IsNil(o.ShipToAddress) {
		return true
	}

	return false
}

// SetShipToAddress gets a reference to the given []FreightRequestShipToAddressInner and assigns it to the ShipToAddress field.
func (o *FreightRequest) SetShipToAddress(v []FreightRequestShipToAddressInner) {
	o.ShipToAddress = v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *FreightRequest) GetLines() []FreightRequestLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []FreightRequestLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightRequest) GetLinesOk() ([]FreightRequestLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *FreightRequest) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []FreightRequestLinesInner and assigns it to the Lines field.
func (o *FreightRequest) SetLines(v []FreightRequestLinesInner) {
	o.Lines = v
}

func (o FreightRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreightRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillToAddressId) {
		toSerialize["billToAddressId"] = o.BillToAddressId
	}
	if !IsNil(o.ShipToAddressId) {
		toSerialize["shipToAddressId"] = o.ShipToAddressId
	}
	if !IsNil(o.ShipToAddress) {
		toSerialize["shipToAddress"] = o.ShipToAddress
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	return toSerialize, nil
}

type NullableFreightRequest struct {
	value *FreightRequest
	isSet bool
}

func (v NullableFreightRequest) Get() *FreightRequest {
	return v.value
}

func (v *NullableFreightRequest) Set(val *FreightRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFreightRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFreightRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreightRequest(val *FreightRequest) *NullableFreightRequest {
	return &NullableFreightRequest{value: val, isSet: true}
}

func (v NullableFreightRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreightRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


