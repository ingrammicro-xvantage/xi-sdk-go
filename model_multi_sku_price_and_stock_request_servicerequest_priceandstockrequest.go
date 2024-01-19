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

// checks if the MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest{}

// MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest struct for MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest
type MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest struct {
	// True/false to show the availability of individual warehouses
	Showwarehouseavailability *string `json:"showwarehouseavailability,omitempty"`
	// Y/N to show extra availability flag
	Extravailabilityflag *string `json:"extravailabilityflag,omitempty"`
	// Flag to indicate if the price and stock information is required for all Ingram Micro systems. If it is set to true, the price and stock details will be returned from all Ingram Micro systems and if false, the price and stock will have returned from the system where the reseller number is set up in.
	Includeallsystems *bool `json:"includeallsystems,omitempty"`
	Item *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem `json:"item,omitempty"`
}

// NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest instantiates a new MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest {
	this := MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest{}
	return &this
}

// NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestWithDefaults instantiates a new MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiSKUPriceAndStockRequestServicerequestPriceandstockrequestWithDefaults() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest {
	this := MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest{}
	return &this
}

// GetShowwarehouseavailability returns the Showwarehouseavailability field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetShowwarehouseavailability() string {
	if o == nil || IsNil(o.Showwarehouseavailability) {
		var ret string
		return ret
	}
	return *o.Showwarehouseavailability
}

// GetShowwarehouseavailabilityOk returns a tuple with the Showwarehouseavailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetShowwarehouseavailabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Showwarehouseavailability) {
		return nil, false
	}
	return o.Showwarehouseavailability, true
}

// HasShowwarehouseavailability returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasShowwarehouseavailability() bool {
	if o != nil && !IsNil(o.Showwarehouseavailability) {
		return true
	}

	return false
}

// SetShowwarehouseavailability gets a reference to the given string and assigns it to the Showwarehouseavailability field.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetShowwarehouseavailability(v string) {
	o.Showwarehouseavailability = &v
}

// GetExtravailabilityflag returns the Extravailabilityflag field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetExtravailabilityflag() string {
	if o == nil || IsNil(o.Extravailabilityflag) {
		var ret string
		return ret
	}
	return *o.Extravailabilityflag
}

// GetExtravailabilityflagOk returns a tuple with the Extravailabilityflag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetExtravailabilityflagOk() (*string, bool) {
	if o == nil || IsNil(o.Extravailabilityflag) {
		return nil, false
	}
	return o.Extravailabilityflag, true
}

// HasExtravailabilityflag returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasExtravailabilityflag() bool {
	if o != nil && !IsNil(o.Extravailabilityflag) {
		return true
	}

	return false
}

// SetExtravailabilityflag gets a reference to the given string and assigns it to the Extravailabilityflag field.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetExtravailabilityflag(v string) {
	o.Extravailabilityflag = &v
}

// GetIncludeallsystems returns the Includeallsystems field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetIncludeallsystems() bool {
	if o == nil || IsNil(o.Includeallsystems) {
		var ret bool
		return ret
	}
	return *o.Includeallsystems
}

// GetIncludeallsystemsOk returns a tuple with the Includeallsystems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetIncludeallsystemsOk() (*bool, bool) {
	if o == nil || IsNil(o.Includeallsystems) {
		return nil, false
	}
	return o.Includeallsystems, true
}

// HasIncludeallsystems returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasIncludeallsystems() bool {
	if o != nil && !IsNil(o.Includeallsystems) {
		return true
	}

	return false
}

// SetIncludeallsystems gets a reference to the given bool and assigns it to the Includeallsystems field.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetIncludeallsystems(v bool) {
	o.Includeallsystems = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetItem() MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem {
	if o == nil || IsNil(o.Item) {
		var ret MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) GetItemOk() (*MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem and assigns it to the Item field.
func (o *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) SetItem(v MultiSKUPriceAndStockRequestServicerequestPriceandstockrequestItem) {
	o.Item = &v
}

func (o MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Showwarehouseavailability) {
		toSerialize["showwarehouseavailability"] = o.Showwarehouseavailability
	}
	if !IsNil(o.Extravailabilityflag) {
		toSerialize["extravailabilityflag"] = o.Extravailabilityflag
	}
	if !IsNil(o.Includeallsystems) {
		toSerialize["includeallsystems"] = o.Includeallsystems
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	return toSerialize, nil
}

type NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest struct {
	value *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest
	isSet bool
}

func (v NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) Get() *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest {
	return v.value
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) Set(val *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest(val *MultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) *NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest {
	return &NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest{value: val, isSet: true}
}

func (v NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiSKUPriceAndStockRequestServicerequestPriceandstockrequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


