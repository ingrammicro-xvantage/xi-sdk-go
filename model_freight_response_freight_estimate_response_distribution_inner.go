/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
)

// checks if the FreightResponseFreightEstimateResponseDistributionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreightResponseFreightEstimateResponseDistributionInner{}

// FreightResponseFreightEstimateResponseDistributionInner struct for FreightResponseFreightEstimateResponseDistributionInner
type FreightResponseFreightEstimateResponseDistributionInner struct {
	// The ID of the warehouse the line item will ship from.
	ShipFromBranchNumber *string `json:"shipFromBranchNumber,omitempty"`
	// The code for the shipping carrier for the line item.
	CarrierCode *string `json:"carrierCode,omitempty"`
	// The name of the shipping carrier.
	ShipVia *string `json:"shipVia,omitempty"`
	// Estimated freight charge.
	FreightRate *float32 `json:"freightRate,omitempty"`
	// Total weight.
	TotalWeight *float32 `json:"totalWeight,omitempty"`
	// Number of transit days.
	TransitDays *int32 `json:"transitDays,omitempty"`
	CarrierList []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner `json:"carrierList,omitempty"`
}

// NewFreightResponseFreightEstimateResponseDistributionInner instantiates a new FreightResponseFreightEstimateResponseDistributionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreightResponseFreightEstimateResponseDistributionInner() *FreightResponseFreightEstimateResponseDistributionInner {
	this := FreightResponseFreightEstimateResponseDistributionInner{}
	return &this
}

// NewFreightResponseFreightEstimateResponseDistributionInnerWithDefaults instantiates a new FreightResponseFreightEstimateResponseDistributionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreightResponseFreightEstimateResponseDistributionInnerWithDefaults() *FreightResponseFreightEstimateResponseDistributionInner {
	this := FreightResponseFreightEstimateResponseDistributionInner{}
	return &this
}

// GetShipFromBranchNumber returns the ShipFromBranchNumber field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipFromBranchNumber() string {
	if o == nil || IsNil(o.ShipFromBranchNumber) {
		var ret string
		return ret
	}
	return *o.ShipFromBranchNumber
}

// GetShipFromBranchNumberOk returns a tuple with the ShipFromBranchNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipFromBranchNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ShipFromBranchNumber) {
		return nil, false
	}
	return o.ShipFromBranchNumber, true
}

// HasShipFromBranchNumber returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasShipFromBranchNumber() bool {
	if o != nil && !IsNil(o.ShipFromBranchNumber) {
		return true
	}

	return false
}

// SetShipFromBranchNumber gets a reference to the given string and assigns it to the ShipFromBranchNumber field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetShipFromBranchNumber(v string) {
	o.ShipFromBranchNumber = &v
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierCode() string {
	if o == nil || IsNil(o.CarrierCode) {
		var ret string
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given string and assigns it to the CarrierCode field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetCarrierCode(v string) {
	o.CarrierCode = &v
}

// GetShipVia returns the ShipVia field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipVia() string {
	if o == nil || IsNil(o.ShipVia) {
		var ret string
		return ret
	}
	return *o.ShipVia
}

// GetShipViaOk returns a tuple with the ShipVia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetShipViaOk() (*string, bool) {
	if o == nil || IsNil(o.ShipVia) {
		return nil, false
	}
	return o.ShipVia, true
}

// HasShipVia returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasShipVia() bool {
	if o != nil && !IsNil(o.ShipVia) {
		return true
	}

	return false
}

// SetShipVia gets a reference to the given string and assigns it to the ShipVia field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetShipVia(v string) {
	o.ShipVia = &v
}

// GetFreightRate returns the FreightRate field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetFreightRate() float32 {
	if o == nil || IsNil(o.FreightRate) {
		var ret float32
		return ret
	}
	return *o.FreightRate
}

// GetFreightRateOk returns a tuple with the FreightRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetFreightRateOk() (*float32, bool) {
	if o == nil || IsNil(o.FreightRate) {
		return nil, false
	}
	return o.FreightRate, true
}

// HasFreightRate returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasFreightRate() bool {
	if o != nil && !IsNil(o.FreightRate) {
		return true
	}

	return false
}

// SetFreightRate gets a reference to the given float32 and assigns it to the FreightRate field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetFreightRate(v float32) {
	o.FreightRate = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTotalWeight() float32 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret float32
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTotalWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given float32 and assigns it to the TotalWeight field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetTotalWeight(v float32) {
	o.TotalWeight = &v
}

// GetTransitDays returns the TransitDays field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTransitDays() int32 {
	if o == nil || IsNil(o.TransitDays) {
		var ret int32
		return ret
	}
	return *o.TransitDays
}

// GetTransitDaysOk returns a tuple with the TransitDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetTransitDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.TransitDays) {
		return nil, false
	}
	return o.TransitDays, true
}

// HasTransitDays returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasTransitDays() bool {
	if o != nil && !IsNil(o.TransitDays) {
		return true
	}

	return false
}

// SetTransitDays gets a reference to the given int32 and assigns it to the TransitDays field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetTransitDays(v int32) {
	o.TransitDays = &v
}

// GetCarrierList returns the CarrierList field value if set, zero value otherwise.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierList() []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner {
	if o == nil || IsNil(o.CarrierList) {
		var ret []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner
		return ret
	}
	return o.CarrierList
}

// GetCarrierListOk returns a tuple with the CarrierList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) GetCarrierListOk() ([]FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner, bool) {
	if o == nil || IsNil(o.CarrierList) {
		return nil, false
	}
	return o.CarrierList, true
}

// HasCarrierList returns a boolean if a field has been set.
func (o *FreightResponseFreightEstimateResponseDistributionInner) HasCarrierList() bool {
	if o != nil && !IsNil(o.CarrierList) {
		return true
	}

	return false
}

// SetCarrierList gets a reference to the given []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner and assigns it to the CarrierList field.
func (o *FreightResponseFreightEstimateResponseDistributionInner) SetCarrierList(v []FreightResponseFreightEstimateResponseDistributionInnerCarrierListInner) {
	o.CarrierList = v
}

func (o FreightResponseFreightEstimateResponseDistributionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreightResponseFreightEstimateResponseDistributionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShipFromBranchNumber) {
		toSerialize["shipFromBranchNumber"] = o.ShipFromBranchNumber
	}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	if !IsNil(o.ShipVia) {
		toSerialize["shipVia"] = o.ShipVia
	}
	if !IsNil(o.FreightRate) {
		toSerialize["freightRate"] = o.FreightRate
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["totalWeight"] = o.TotalWeight
	}
	if !IsNil(o.TransitDays) {
		toSerialize["transitDays"] = o.TransitDays
	}
	if !IsNil(o.CarrierList) {
		toSerialize["carrierList"] = o.CarrierList
	}
	return toSerialize, nil
}

type NullableFreightResponseFreightEstimateResponseDistributionInner struct {
	value *FreightResponseFreightEstimateResponseDistributionInner
	isSet bool
}

func (v NullableFreightResponseFreightEstimateResponseDistributionInner) Get() *FreightResponseFreightEstimateResponseDistributionInner {
	return v.value
}

func (v *NullableFreightResponseFreightEstimateResponseDistributionInner) Set(val *FreightResponseFreightEstimateResponseDistributionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFreightResponseFreightEstimateResponseDistributionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFreightResponseFreightEstimateResponseDistributionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreightResponseFreightEstimateResponseDistributionInner(val *FreightResponseFreightEstimateResponseDistributionInner) *NullableFreightResponseFreightEstimateResponseDistributionInner {
	return &NullableFreightResponseFreightEstimateResponseDistributionInner{value: val, isSet: true}
}

func (v NullableFreightResponseFreightEstimateResponseDistributionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreightResponseFreightEstimateResponseDistributionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


