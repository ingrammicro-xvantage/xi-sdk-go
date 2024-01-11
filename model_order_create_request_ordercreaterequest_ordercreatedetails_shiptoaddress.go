/*
Reseller API Documentation

For Resellers. <br> Who are looking to Innovate with Ingram Micro's API SolutionsAutomate your eCommerce with our offering of APIs and Webhooks to create a seamless experience for your customers.

API version: 6.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package xi-sdk-resellers-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress{}

// OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress struct for OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress
type OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress struct {
	// Customer contact name
	Attention *string `json:"attention,omitempty"`
	// Company Name or person to deliver. *If there isn’t an attention line please add the company name on address line 1.   UPS and FedEx will create surcharges if address line 1 contains a physical address.
	Addressline1 string `json:"addressline1"`
	// Street address for delivery
	Addressline2 string `json:"addressline2"`
	// Continuation of address line 2
	Addressline3 *string `json:"addressline3,omitempty"`
	// Ship to city
	City string `json:"city"`
	// Ship to State or Region
	State string `json:"state"`
	// Ship to Zip code or Postal code
	Postalcode string `json:"postalcode"`
	// Ship to country
	Countrycode *string `json:"countrycode,omitempty"`
}

type _OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress

// NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress(addressline1 string, addressline2 string, city string, state string, postalcode string) *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress {
	this := OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress{}
	this.Addressline1 = addressline1
	this.Addressline2 = addressline2
	this.City = city
	this.State = state
	this.Postalcode = postalcode
	return &this
}

// NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddressWithDefaults instantiates a new OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddressWithDefaults() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress {
	this := OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress{}
	return &this
}

// GetAttention returns the Attention field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAttention() string {
	if o == nil || IsNil(o.Attention) {
		var ret string
		return ret
	}
	return *o.Attention
}

// GetAttentionOk returns a tuple with the Attention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAttentionOk() (*string, bool) {
	if o == nil || IsNil(o.Attention) {
		return nil, false
	}
	return o.Attention, true
}

// HasAttention returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) HasAttention() bool {
	if o != nil && !IsNil(o.Attention) {
		return true
	}

	return false
}

// SetAttention gets a reference to the given string and assigns it to the Attention field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAttention(v string) {
	o.Attention = &v
}

// GetAddressline1 returns the Addressline1 field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addressline1
}

// GetAddressline1Ok returns a tuple with the Addressline1 field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addressline1, true
}

// SetAddressline1 sets field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAddressline1(v string) {
	o.Addressline1 = v
}

// GetAddressline2 returns the Addressline2 field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Addressline2
}

// GetAddressline2Ok returns a tuple with the Addressline2 field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addressline2, true
}

// SetAddressline2 sets field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAddressline2(v string) {
	o.Addressline2 = v
}

// GetAddressline3 returns the Addressline3 field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline3() string {
	if o == nil || IsNil(o.Addressline3) {
		var ret string
		return ret
	}
	return *o.Addressline3
}

// GetAddressline3Ok returns a tuple with the Addressline3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetAddressline3Ok() (*string, bool) {
	if o == nil || IsNil(o.Addressline3) {
		return nil, false
	}
	return o.Addressline3, true
}

// HasAddressline3 returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) HasAddressline3() bool {
	if o != nil && !IsNil(o.Addressline3) {
		return true
	}

	return false
}

// SetAddressline3 gets a reference to the given string and assigns it to the Addressline3 field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetAddressline3(v string) {
	o.Addressline3 = &v
}

// GetCity returns the City field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetState(v string) {
	o.State = v
}

// GetPostalcode returns the Postalcode field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetPostalcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Postalcode
}

// GetPostalcodeOk returns a tuple with the Postalcode field value
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetPostalcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Postalcode, true
}

// SetPostalcode sets field value
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetPostalcode(v string) {
	o.Postalcode = v
}

// GetCountrycode returns the Countrycode field value if set, zero value otherwise.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCountrycode() string {
	if o == nil || IsNil(o.Countrycode) {
		var ret string
		return ret
	}
	return *o.Countrycode
}

// GetCountrycodeOk returns a tuple with the Countrycode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) GetCountrycodeOk() (*string, bool) {
	if o == nil || IsNil(o.Countrycode) {
		return nil, false
	}
	return o.Countrycode, true
}

// HasCountrycode returns a boolean if a field has been set.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) HasCountrycode() bool {
	if o != nil && !IsNil(o.Countrycode) {
		return true
	}

	return false
}

// SetCountrycode gets a reference to the given string and assigns it to the Countrycode field.
func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) SetCountrycode(v string) {
	o.Countrycode = &v
}

func (o OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attention) {
		toSerialize["attention"] = o.Attention
	}
	toSerialize["addressline1"] = o.Addressline1
	toSerialize["addressline2"] = o.Addressline2
	if !IsNil(o.Addressline3) {
		toSerialize["addressline3"] = o.Addressline3
	}
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["postalcode"] = o.Postalcode
	if !IsNil(o.Countrycode) {
		toSerialize["countrycode"] = o.Countrycode
	}
	return toSerialize, nil
}

func (o *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addressline1",
		"addressline2",
		"city",
		"state",
		"postalcode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress := _OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress)

	if err != nil {
		return err
	}

	*o = OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress(varOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress)

	return err
}

type NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress struct {
	value *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress
	isSet bool
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) Get() *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress {
	return v.value
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) Set(val *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress(val *OrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress {
	return &NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress{value: val, isSet: true}
}

func (v NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderCreateRequestOrdercreaterequestOrdercreatedetailsShiptoaddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


