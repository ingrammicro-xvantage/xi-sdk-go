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

// checks if the ProductDetailResponseAdditionalInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDetailResponseAdditionalInformation{}

// ProductDetailResponseAdditionalInformation Additional Information related to the product.
type ProductDetailResponseAdditionalInformation struct {
	// Weight information related to the product.
	ProductWeight []ProductDetailResponseAdditionalInformationProductWeightInner `json:"productWeight,omitempty"`
	// Example : true or false
	IsBulkFreight *bool `json:"isBulkFreight,omitempty"`
	// Example : '5.2 Inches'
	Height *string `json:"height,omitempty"`
	// Example : '13 inches'
	Width *string `json:"width,omitempty"`
	// Example : '20.4 inches'
	Length *string `json:"length,omitempty"`
	// Example : '10 lb'
	NetWeight *string `json:"netWeight,omitempty"`
	// Example : 'Unit value'
	DimensionUnit *string `json:"dimensionUnit,omitempty"`
}

// NewProductDetailResponseAdditionalInformation instantiates a new ProductDetailResponseAdditionalInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDetailResponseAdditionalInformation() *ProductDetailResponseAdditionalInformation {
	this := ProductDetailResponseAdditionalInformation{}
	return &this
}

// NewProductDetailResponseAdditionalInformationWithDefaults instantiates a new ProductDetailResponseAdditionalInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDetailResponseAdditionalInformationWithDefaults() *ProductDetailResponseAdditionalInformation {
	this := ProductDetailResponseAdditionalInformation{}
	return &this
}

// GetProductWeight returns the ProductWeight field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetProductWeight() []ProductDetailResponseAdditionalInformationProductWeightInner {
	if o == nil || IsNil(o.ProductWeight) {
		var ret []ProductDetailResponseAdditionalInformationProductWeightInner
		return ret
	}
	return o.ProductWeight
}

// GetProductWeightOk returns a tuple with the ProductWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetProductWeightOk() ([]ProductDetailResponseAdditionalInformationProductWeightInner, bool) {
	if o == nil || IsNil(o.ProductWeight) {
		return nil, false
	}
	return o.ProductWeight, true
}

// HasProductWeight returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasProductWeight() bool {
	if o != nil && !IsNil(o.ProductWeight) {
		return true
	}

	return false
}

// SetProductWeight gets a reference to the given []ProductDetailResponseAdditionalInformationProductWeightInner and assigns it to the ProductWeight field.
func (o *ProductDetailResponseAdditionalInformation) SetProductWeight(v []ProductDetailResponseAdditionalInformationProductWeightInner) {
	o.ProductWeight = v
}

// GetIsBulkFreight returns the IsBulkFreight field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetIsBulkFreight() bool {
	if o == nil || IsNil(o.IsBulkFreight) {
		var ret bool
		return ret
	}
	return *o.IsBulkFreight
}

// GetIsBulkFreightOk returns a tuple with the IsBulkFreight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetIsBulkFreightOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBulkFreight) {
		return nil, false
	}
	return o.IsBulkFreight, true
}

// HasIsBulkFreight returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasIsBulkFreight() bool {
	if o != nil && !IsNil(o.IsBulkFreight) {
		return true
	}

	return false
}

// SetIsBulkFreight gets a reference to the given bool and assigns it to the IsBulkFreight field.
func (o *ProductDetailResponseAdditionalInformation) SetIsBulkFreight(v bool) {
	o.IsBulkFreight = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetHeight() string {
	if o == nil || IsNil(o.Height) {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetHeightOk() (*string, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *ProductDetailResponseAdditionalInformation) SetHeight(v string) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *ProductDetailResponseAdditionalInformation) SetWidth(v string) {
	o.Width = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetLength() string {
	if o == nil || IsNil(o.Length) {
		var ret string
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetLengthOk() (*string, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given string and assigns it to the Length field.
func (o *ProductDetailResponseAdditionalInformation) SetLength(v string) {
	o.Length = &v
}

// GetNetWeight returns the NetWeight field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetNetWeight() string {
	if o == nil || IsNil(o.NetWeight) {
		var ret string
		return ret
	}
	return *o.NetWeight
}

// GetNetWeightOk returns a tuple with the NetWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetNetWeightOk() (*string, bool) {
	if o == nil || IsNil(o.NetWeight) {
		return nil, false
	}
	return o.NetWeight, true
}

// HasNetWeight returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasNetWeight() bool {
	if o != nil && !IsNil(o.NetWeight) {
		return true
	}

	return false
}

// SetNetWeight gets a reference to the given string and assigns it to the NetWeight field.
func (o *ProductDetailResponseAdditionalInformation) SetNetWeight(v string) {
	o.NetWeight = &v
}

// GetDimensionUnit returns the DimensionUnit field value if set, zero value otherwise.
func (o *ProductDetailResponseAdditionalInformation) GetDimensionUnit() string {
	if o == nil || IsNil(o.DimensionUnit) {
		var ret string
		return ret
	}
	return *o.DimensionUnit
}

// GetDimensionUnitOk returns a tuple with the DimensionUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDetailResponseAdditionalInformation) GetDimensionUnitOk() (*string, bool) {
	if o == nil || IsNil(o.DimensionUnit) {
		return nil, false
	}
	return o.DimensionUnit, true
}

// HasDimensionUnit returns a boolean if a field has been set.
func (o *ProductDetailResponseAdditionalInformation) HasDimensionUnit() bool {
	if o != nil && !IsNil(o.DimensionUnit) {
		return true
	}

	return false
}

// SetDimensionUnit gets a reference to the given string and assigns it to the DimensionUnit field.
func (o *ProductDetailResponseAdditionalInformation) SetDimensionUnit(v string) {
	o.DimensionUnit = &v
}

func (o ProductDetailResponseAdditionalInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDetailResponseAdditionalInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductWeight) {
		toSerialize["productWeight"] = o.ProductWeight
	}
	if !IsNil(o.IsBulkFreight) {
		toSerialize["isBulkFreight"] = o.IsBulkFreight
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.NetWeight) {
		toSerialize["netWeight"] = o.NetWeight
	}
	if !IsNil(o.DimensionUnit) {
		toSerialize["dimensionUnit"] = o.DimensionUnit
	}
	return toSerialize, nil
}

type NullableProductDetailResponseAdditionalInformation struct {
	value *ProductDetailResponseAdditionalInformation
	isSet bool
}

func (v NullableProductDetailResponseAdditionalInformation) Get() *ProductDetailResponseAdditionalInformation {
	return v.value
}

func (v *NullableProductDetailResponseAdditionalInformation) Set(val *ProductDetailResponseAdditionalInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDetailResponseAdditionalInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDetailResponseAdditionalInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDetailResponseAdditionalInformation(val *ProductDetailResponseAdditionalInformation) *NullableProductDetailResponseAdditionalInformation {
	return &NullableProductDetailResponseAdditionalInformation{value: val, isSet: true}
}

func (v NullableProductDetailResponseAdditionalInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDetailResponseAdditionalInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


