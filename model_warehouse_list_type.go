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

// checks if the WarehouseListType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WarehouseListType{}

// WarehouseListType struct for WarehouseListType
type WarehouseListType struct {
	// 
	Warehouseid *string `json:"warehouseid,omitempty"`
	// City of the Ingram Micro warehouse location
	Warehousedescription *string `json:"warehousedescription,omitempty"`
	// On hand available quantity
	Availablequantity *int32 `json:"availablequantity,omitempty"`
	// On Order quantity
	Onorderquantity *int32 `json:"onorderquantity,omitempty"`
	// On hold quantity
	Onholdquantity *int32 `json:"onholdquantity,omitempty"`
	Etadate *string `json:"etadate,omitempty"`
}

// NewWarehouseListType instantiates a new WarehouseListType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseListType() *WarehouseListType {
	this := WarehouseListType{}
	return &this
}

// NewWarehouseListTypeWithDefaults instantiates a new WarehouseListType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseListTypeWithDefaults() *WarehouseListType {
	this := WarehouseListType{}
	return &this
}

// GetWarehouseid returns the Warehouseid field value if set, zero value otherwise.
func (o *WarehouseListType) GetWarehouseid() string {
	if o == nil || IsNil(o.Warehouseid) {
		var ret string
		return ret
	}
	return *o.Warehouseid
}

// GetWarehouseidOk returns a tuple with the Warehouseid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseListType) GetWarehouseidOk() (*string, bool) {
	if o == nil || IsNil(o.Warehouseid) {
		return nil, false
	}
	return o.Warehouseid, true
}

// HasWarehouseid returns a boolean if a field has been set.
func (o *WarehouseListType) HasWarehouseid() bool {
	if o != nil && !IsNil(o.Warehouseid) {
		return true
	}

	return false
}

// SetWarehouseid gets a reference to the given string and assigns it to the Warehouseid field.
func (o *WarehouseListType) SetWarehouseid(v string) {
	o.Warehouseid = &v
}

// GetWarehousedescription returns the Warehousedescription field value if set, zero value otherwise.
func (o *WarehouseListType) GetWarehousedescription() string {
	if o == nil || IsNil(o.Warehousedescription) {
		var ret string
		return ret
	}
	return *o.Warehousedescription
}

// GetWarehousedescriptionOk returns a tuple with the Warehousedescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseListType) GetWarehousedescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Warehousedescription) {
		return nil, false
	}
	return o.Warehousedescription, true
}

// HasWarehousedescription returns a boolean if a field has been set.
func (o *WarehouseListType) HasWarehousedescription() bool {
	if o != nil && !IsNil(o.Warehousedescription) {
		return true
	}

	return false
}

// SetWarehousedescription gets a reference to the given string and assigns it to the Warehousedescription field.
func (o *WarehouseListType) SetWarehousedescription(v string) {
	o.Warehousedescription = &v
}

// GetAvailablequantity returns the Availablequantity field value if set, zero value otherwise.
func (o *WarehouseListType) GetAvailablequantity() int32 {
	if o == nil || IsNil(o.Availablequantity) {
		var ret int32
		return ret
	}
	return *o.Availablequantity
}

// GetAvailablequantityOk returns a tuple with the Availablequantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseListType) GetAvailablequantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Availablequantity) {
		return nil, false
	}
	return o.Availablequantity, true
}

// HasAvailablequantity returns a boolean if a field has been set.
func (o *WarehouseListType) HasAvailablequantity() bool {
	if o != nil && !IsNil(o.Availablequantity) {
		return true
	}

	return false
}

// SetAvailablequantity gets a reference to the given int32 and assigns it to the Availablequantity field.
func (o *WarehouseListType) SetAvailablequantity(v int32) {
	o.Availablequantity = &v
}

// GetOnorderquantity returns the Onorderquantity field value if set, zero value otherwise.
func (o *WarehouseListType) GetOnorderquantity() int32 {
	if o == nil || IsNil(o.Onorderquantity) {
		var ret int32
		return ret
	}
	return *o.Onorderquantity
}

// GetOnorderquantityOk returns a tuple with the Onorderquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseListType) GetOnorderquantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Onorderquantity) {
		return nil, false
	}
	return o.Onorderquantity, true
}

// HasOnorderquantity returns a boolean if a field has been set.
func (o *WarehouseListType) HasOnorderquantity() bool {
	if o != nil && !IsNil(o.Onorderquantity) {
		return true
	}

	return false
}

// SetOnorderquantity gets a reference to the given int32 and assigns it to the Onorderquantity field.
func (o *WarehouseListType) SetOnorderquantity(v int32) {
	o.Onorderquantity = &v
}

// GetOnholdquantity returns the Onholdquantity field value if set, zero value otherwise.
func (o *WarehouseListType) GetOnholdquantity() int32 {
	if o == nil || IsNil(o.Onholdquantity) {
		var ret int32
		return ret
	}
	return *o.Onholdquantity
}

// GetOnholdquantityOk returns a tuple with the Onholdquantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseListType) GetOnholdquantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Onholdquantity) {
		return nil, false
	}
	return o.Onholdquantity, true
}

// HasOnholdquantity returns a boolean if a field has been set.
func (o *WarehouseListType) HasOnholdquantity() bool {
	if o != nil && !IsNil(o.Onholdquantity) {
		return true
	}

	return false
}

// SetOnholdquantity gets a reference to the given int32 and assigns it to the Onholdquantity field.
func (o *WarehouseListType) SetOnholdquantity(v int32) {
	o.Onholdquantity = &v
}

// GetEtadate returns the Etadate field value if set, zero value otherwise.
func (o *WarehouseListType) GetEtadate() string {
	if o == nil || IsNil(o.Etadate) {
		var ret string
		return ret
	}
	return *o.Etadate
}

// GetEtadateOk returns a tuple with the Etadate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseListType) GetEtadateOk() (*string, bool) {
	if o == nil || IsNil(o.Etadate) {
		return nil, false
	}
	return o.Etadate, true
}

// HasEtadate returns a boolean if a field has been set.
func (o *WarehouseListType) HasEtadate() bool {
	if o != nil && !IsNil(o.Etadate) {
		return true
	}

	return false
}

// SetEtadate gets a reference to the given string and assigns it to the Etadate field.
func (o *WarehouseListType) SetEtadate(v string) {
	o.Etadate = &v
}

func (o WarehouseListType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WarehouseListType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Warehouseid) {
		toSerialize["warehouseid"] = o.Warehouseid
	}
	if !IsNil(o.Warehousedescription) {
		toSerialize["warehousedescription"] = o.Warehousedescription
	}
	if !IsNil(o.Availablequantity) {
		toSerialize["availablequantity"] = o.Availablequantity
	}
	if !IsNil(o.Onorderquantity) {
		toSerialize["onorderquantity"] = o.Onorderquantity
	}
	if !IsNil(o.Onholdquantity) {
		toSerialize["onholdquantity"] = o.Onholdquantity
	}
	if !IsNil(o.Etadate) {
		toSerialize["etadate"] = o.Etadate
	}
	return toSerialize, nil
}

type NullableWarehouseListType struct {
	value *WarehouseListType
	isSet bool
}

func (v NullableWarehouseListType) Get() *WarehouseListType {
	return v.value
}

func (v *NullableWarehouseListType) Set(val *WarehouseListType) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseListType) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseListType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseListType(val *WarehouseListType) *NullableWarehouseListType {
	return &NullableWarehouseListType{value: val, isSet: true}
}

func (v NullableWarehouseListType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseListType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


