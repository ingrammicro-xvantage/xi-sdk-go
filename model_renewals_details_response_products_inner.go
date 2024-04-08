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

// checks if the RenewalsDetailsResponseProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RenewalsDetailsResponseProductsInner{}

// RenewalsDetailsResponseProductsInner struct for RenewalsDetailsResponseProductsInner
type RenewalsDetailsResponseProductsInner struct {
	// Unique Ingram Micro line number.
	IngramLineNumber *string `json:"ingramLineNumber,omitempty"`
	// The description of the product.
	ProductDescription *string `json:"productDescription,omitempty"`
	// The vendor's part number for the line item.
	VendorPartNumber *string `json:"vendorPartNumber,omitempty"`
	// Unique IngramMicro part number.
	IngramPartNumber *string `json:"ingramPartNumber,omitempty"`
	// The manufacturer's part number for the line item.
	ManufacturerPartNumber *string `json:"manufacturerPartNumber,omitempty"`
	// The quantity of the line item.
	Quantity *string `json:"quantity,omitempty"`
	// The unit price of the line item.
	UnitPrice *float32 `json:"unitPrice,omitempty"`
	// Is the line item consolidated? Yes or No.
	IsConsolidated *string `json:"isConsolidated,omitempty"`
}

// NewRenewalsDetailsResponseProductsInner instantiates a new RenewalsDetailsResponseProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRenewalsDetailsResponseProductsInner() *RenewalsDetailsResponseProductsInner {
	this := RenewalsDetailsResponseProductsInner{}
	return &this
}

// NewRenewalsDetailsResponseProductsInnerWithDefaults instantiates a new RenewalsDetailsResponseProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRenewalsDetailsResponseProductsInnerWithDefaults() *RenewalsDetailsResponseProductsInner {
	this := RenewalsDetailsResponseProductsInner{}
	return &this
}

// GetIngramLineNumber returns the IngramLineNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetIngramLineNumber() string {
	if o == nil || IsNil(o.IngramLineNumber) {
		var ret string
		return ret
	}
	return *o.IngramLineNumber
}

// GetIngramLineNumberOk returns a tuple with the IngramLineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetIngramLineNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramLineNumber) {
		return nil, false
	}
	return o.IngramLineNumber, true
}

// HasIngramLineNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasIngramLineNumber() bool {
	if o != nil && !IsNil(o.IngramLineNumber) {
		return true
	}

	return false
}

// SetIngramLineNumber gets a reference to the given string and assigns it to the IngramLineNumber field.
func (o *RenewalsDetailsResponseProductsInner) SetIngramLineNumber(v string) {
	o.IngramLineNumber = &v
}

// GetProductDescription returns the ProductDescription field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetProductDescription() string {
	if o == nil || IsNil(o.ProductDescription) {
		var ret string
		return ret
	}
	return *o.ProductDescription
}

// GetProductDescriptionOk returns a tuple with the ProductDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetProductDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductDescription) {
		return nil, false
	}
	return o.ProductDescription, true
}

// HasProductDescription returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasProductDescription() bool {
	if o != nil && !IsNil(o.ProductDescription) {
		return true
	}

	return false
}

// SetProductDescription gets a reference to the given string and assigns it to the ProductDescription field.
func (o *RenewalsDetailsResponseProductsInner) SetProductDescription(v string) {
	o.ProductDescription = &v
}

// GetVendorPartNumber returns the VendorPartNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetVendorPartNumber() string {
	if o == nil || IsNil(o.VendorPartNumber) {
		var ret string
		return ret
	}
	return *o.VendorPartNumber
}

// GetVendorPartNumberOk returns a tuple with the VendorPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetVendorPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VendorPartNumber) {
		return nil, false
	}
	return o.VendorPartNumber, true
}

// HasVendorPartNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasVendorPartNumber() bool {
	if o != nil && !IsNil(o.VendorPartNumber) {
		return true
	}

	return false
}

// SetVendorPartNumber gets a reference to the given string and assigns it to the VendorPartNumber field.
func (o *RenewalsDetailsResponseProductsInner) SetVendorPartNumber(v string) {
	o.VendorPartNumber = &v
}

// GetIngramPartNumber returns the IngramPartNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetIngramPartNumber() string {
	if o == nil || IsNil(o.IngramPartNumber) {
		var ret string
		return ret
	}
	return *o.IngramPartNumber
}

// GetIngramPartNumberOk returns a tuple with the IngramPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetIngramPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IngramPartNumber) {
		return nil, false
	}
	return o.IngramPartNumber, true
}

// HasIngramPartNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasIngramPartNumber() bool {
	if o != nil && !IsNil(o.IngramPartNumber) {
		return true
	}

	return false
}

// SetIngramPartNumber gets a reference to the given string and assigns it to the IngramPartNumber field.
func (o *RenewalsDetailsResponseProductsInner) SetIngramPartNumber(v string) {
	o.IngramPartNumber = &v
}

// GetManufacturerPartNumber returns the ManufacturerPartNumber field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetManufacturerPartNumber() string {
	if o == nil || IsNil(o.ManufacturerPartNumber) {
		var ret string
		return ret
	}
	return *o.ManufacturerPartNumber
}

// GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetManufacturerPartNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ManufacturerPartNumber) {
		return nil, false
	}
	return o.ManufacturerPartNumber, true
}

// HasManufacturerPartNumber returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasManufacturerPartNumber() bool {
	if o != nil && !IsNil(o.ManufacturerPartNumber) {
		return true
	}

	return false
}

// SetManufacturerPartNumber gets a reference to the given string and assigns it to the ManufacturerPartNumber field.
func (o *RenewalsDetailsResponseProductsInner) SetManufacturerPartNumber(v string) {
	o.ManufacturerPartNumber = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *RenewalsDetailsResponseProductsInner) SetQuantity(v string) {
	o.Quantity = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetUnitPrice() float32 {
	if o == nil || IsNil(o.UnitPrice) {
		var ret float32
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetUnitPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given float32 and assigns it to the UnitPrice field.
func (o *RenewalsDetailsResponseProductsInner) SetUnitPrice(v float32) {
	o.UnitPrice = &v
}

// GetIsConsolidated returns the IsConsolidated field value if set, zero value otherwise.
func (o *RenewalsDetailsResponseProductsInner) GetIsConsolidated() string {
	if o == nil || IsNil(o.IsConsolidated) {
		var ret string
		return ret
	}
	return *o.IsConsolidated
}

// GetIsConsolidatedOk returns a tuple with the IsConsolidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RenewalsDetailsResponseProductsInner) GetIsConsolidatedOk() (*string, bool) {
	if o == nil || IsNil(o.IsConsolidated) {
		return nil, false
	}
	return o.IsConsolidated, true
}

// HasIsConsolidated returns a boolean if a field has been set.
func (o *RenewalsDetailsResponseProductsInner) HasIsConsolidated() bool {
	if o != nil && !IsNil(o.IsConsolidated) {
		return true
	}

	return false
}

// SetIsConsolidated gets a reference to the given string and assigns it to the IsConsolidated field.
func (o *RenewalsDetailsResponseProductsInner) SetIsConsolidated(v string) {
	o.IsConsolidated = &v
}

func (o RenewalsDetailsResponseProductsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RenewalsDetailsResponseProductsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IngramLineNumber) {
		toSerialize["ingramLineNumber"] = o.IngramLineNumber
	}
	if !IsNil(o.ProductDescription) {
		toSerialize["productDescription"] = o.ProductDescription
	}
	if !IsNil(o.VendorPartNumber) {
		toSerialize["vendorPartNumber"] = o.VendorPartNumber
	}
	if !IsNil(o.IngramPartNumber) {
		toSerialize["ingramPartNumber"] = o.IngramPartNumber
	}
	if !IsNil(o.ManufacturerPartNumber) {
		toSerialize["manufacturerPartNumber"] = o.ManufacturerPartNumber
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	if !IsNil(o.IsConsolidated) {
		toSerialize["isConsolidated"] = o.IsConsolidated
	}
	return toSerialize, nil
}

type NullableRenewalsDetailsResponseProductsInner struct {
	value *RenewalsDetailsResponseProductsInner
	isSet bool
}

func (v NullableRenewalsDetailsResponseProductsInner) Get() *RenewalsDetailsResponseProductsInner {
	return v.value
}

func (v *NullableRenewalsDetailsResponseProductsInner) Set(val *RenewalsDetailsResponseProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRenewalsDetailsResponseProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRenewalsDetailsResponseProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRenewalsDetailsResponseProductsInner(val *RenewalsDetailsResponseProductsInner) *NullableRenewalsDetailsResponseProductsInner {
	return &NullableRenewalsDetailsResponseProductsInner{value: val, isSet: true}
}

func (v NullableRenewalsDetailsResponseProductsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRenewalsDetailsResponseProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


