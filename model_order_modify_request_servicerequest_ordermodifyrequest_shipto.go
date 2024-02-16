/*
XI SDK Resellers

For Resellers. Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi_sdk_resellers

import (
	"encoding/json"
)

// checks if the OrderModifyRequestServicerequestOrdermodifyrequestShipto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderModifyRequestServicerequestOrdermodifyrequestShipto{}

// OrderModifyRequestServicerequestOrdermodifyrequestShipto struct for OrderModifyRequestServicerequestOrdermodifyrequestShipto
type OrderModifyRequestServicerequestOrdermodifyrequestShipto struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"name,omitempty"`
	Addressline *string `json:"addressline,omitempty"`
	City *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Postalcode *string `json:"postalcode,omitempty"`
	Countrycode *string `json:"countrycode,omitempty"`
}

// NewOrderModifyRequestServicerequestOrdermodifyrequestShipto instantiates a new OrderModifyRequestServicerequestOrdermodifyrequestShipto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderModifyRequestServicerequestOrdermodifyrequestShipto() *OrderModifyRequestServicerequestOrdermodifyrequestShipto {
	this := OrderModifyRequestServicerequestOrdermodifyrequestShipto{}
	return &this
}

// NewOrderModifyRequestServicerequestOrdermodifyrequestShiptoWithDefaults instantiates a new OrderModifyRequestServicerequestOrdermodifyrequestShipto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderModifyRequestServicerequestOrdermodifyrequestShiptoWithDefaults() *OrderModifyRequestServicerequestOrdermodifyrequestShipto {
	this := OrderModifyRequestServicerequestOrdermodifyrequestShipto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetName(v string) {
	o.Name = &v
}

// GetAddressline returns the Addressline field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetAddressline() string {
	if o == nil || IsNil(o.Addressline) {
		var ret string
		return ret
	}
	return *o.Addressline
}

// GetAddresslineOk returns a tuple with the Addressline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetAddresslineOk() (*string, bool) {
	if o == nil || IsNil(o.Addressline) {
		return nil, false
	}
	return o.Addressline, true
}

// HasAddressline returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasAddressline() bool {
	if o != nil && !IsNil(o.Addressline) {
		return true
	}

	return false
}

// SetAddressline gets a reference to the given string and assigns it to the Addressline field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetAddressline(v string) {
	o.Addressline = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetState(v string) {
	o.State = &v
}

// GetPostalcode returns the Postalcode field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetPostalcode() string {
	if o == nil || IsNil(o.Postalcode) {
		var ret string
		return ret
	}
	return *o.Postalcode
}

// GetPostalcodeOk returns a tuple with the Postalcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetPostalcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Postalcode) {
		return nil, false
	}
	return o.Postalcode, true
}

// HasPostalcode returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasPostalcode() bool {
	if o != nil && !IsNil(o.Postalcode) {
		return true
	}

	return false
}

// SetPostalcode gets a reference to the given string and assigns it to the Postalcode field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetPostalcode(v string) {
	o.Postalcode = &v
}

// GetCountrycode returns the Countrycode field value if set, zero value otherwise.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetCountrycode() string {
	if o == nil || IsNil(o.Countrycode) {
		var ret string
		return ret
	}
	return *o.Countrycode
}

// GetCountrycodeOk returns a tuple with the Countrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) GetCountrycodeOk() (*string, bool) {
	if o == nil || IsNil(o.Countrycode) {
		return nil, false
	}
	return o.Countrycode, true
}

// HasCountrycode returns a boolean if a field has been set.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) HasCountrycode() bool {
	if o != nil && !IsNil(o.Countrycode) {
		return true
	}

	return false
}

// SetCountrycode gets a reference to the given string and assigns it to the Countrycode field.
func (o *OrderModifyRequestServicerequestOrdermodifyrequestShipto) SetCountrycode(v string) {
	o.Countrycode = &v
}

func (o OrderModifyRequestServicerequestOrdermodifyrequestShipto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderModifyRequestServicerequestOrdermodifyrequestShipto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Addressline) {
		toSerialize["addressline"] = o.Addressline
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Postalcode) {
		toSerialize["postalcode"] = o.Postalcode
	}
	if !IsNil(o.Countrycode) {
		toSerialize["countrycode"] = o.Countrycode
	}
	return toSerialize, nil
}

type NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto struct {
	value *OrderModifyRequestServicerequestOrdermodifyrequestShipto
	isSet bool
}

func (v NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto) Get() *OrderModifyRequestServicerequestOrdermodifyrequestShipto {
	return v.value
}

func (v *NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto) Set(val *OrderModifyRequestServicerequestOrdermodifyrequestShipto) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderModifyRequestServicerequestOrdermodifyrequestShipto(val *OrderModifyRequestServicerequestOrdermodifyrequestShipto) *NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto {
	return &NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto{value: val, isSet: true}
}

func (v NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderModifyRequestServicerequestOrdermodifyrequestShipto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


